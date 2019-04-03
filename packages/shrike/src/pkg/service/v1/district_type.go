package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new DistrictType
func (s *shrikeServiceServer) CreateDistrictType(ctx context.Context, req *v1.CreateDistrictTypeRequest) (*v1.CreateDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// insert DistrictType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO district_type (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into DistrictType-> "+err.Error())
	}

	// get ID of creates DistrictType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created DistrictType-> "+err.Error())
	}

	return &v1.CreateDistrictTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get district_type by id.
func (s *shrikeServiceServer) GetDistrictType(ctx context.Context, req *v1.GetDistrictTypeRequest) (*v1.GetDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query DistrictType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM district_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DistrictType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from DistrictType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DistrictType with ID='%s' is not found",
			req.ID))
	}

	// scan DistrictType data into protobuf model
	var districttype v1.DistrictType
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&districttype.ID, &createdAt, &updatedAt, &districttype.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from DistrictType row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		districttype.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		districttype.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple DistrictType rows with ID='%s'",
			req.ID))
	}

	return &v1.GetDistrictTypeResponse{
		Api:  apiVersion,
		Item: &districttype,
	}, nil

}

// Read all DistrictType
func (s *shrikeServiceServer) ListDistrictType(ctx context.Context, req *v1.ListDistrictTypeRequest) (*v1.ListDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in DistrictType Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildDistrictTypeListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DistrictType-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.DistrictType{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		districttype := new(v1.DistrictType)
		if err := rows.Scan(&districttype.ID, &createdAt, &updatedAt, &districttype.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from DistrictType row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			districttype.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			districttype.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, districttype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from DistrictType-> "+err.Error())
	}

	return &v1.ListDistrictTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update DistrictType
func (s *shrikeServiceServer) UpdateDistrictType(ctx context.Context, req *v1.UpdateDistrictTypeRequest) (*v1.UpdateDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// update district_type
	res, err := c.ExecContext(ctx, "UPDATE district_type SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update DistrictType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DistrictType with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateDistrictTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete district_type
func (s *shrikeServiceServer) DeleteDistrictType(ctx context.Context, req *v1.DeleteDistrictTypeRequest) (*v1.DeleteDistrictTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// delete district_type
	res, err := c.ExecContext(ctx, "DELETE FROM district_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete DistrictType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DistrictType with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteDistrictTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
