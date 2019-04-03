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

// Create new ACL
func (s *shrikeServiceServer) CreateACL(ctx context.Context, req *v1.CreateACLRequest) (*v1.CreateACLResponse, error) {
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
	// insert ACL entity data
	err = c.QueryRowContext(ctx, "INSERT INTO acl () VALUES()  RETURNING id;").Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ACL-> "+err.Error())
	}

	// get ID of creates ACL
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ACL-> "+err.Error())
	}

	return &v1.CreateACLResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get acl by id.
func (s *shrikeServiceServer) GetACL(ctx context.Context, req *v1.GetACLRequest) (*v1.GetACLResponse, error) {
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

	// query ACL by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM acl WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ACL-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ACL-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ACL with ID='%s' is not found",
			req.ID))
	}

	// scan ACL data into protobuf model
	var acl v1.ACL
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	if err := rows.Scan(&acl.ID, &createdAt, &updatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ACL row-> "+err.Error())
	}

	// Convert pq.NullTime from database into proto timestamp.
	if createdAt.Valid {
		acl.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}
	if updatedAt.Valid {
		acl.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ACL rows with ID='%s'",
			req.ID))
	}

	return &v1.GetACLResponse{
		Api:  apiVersion,
		Item: &acl,
	}, nil

}

// Read all ACL
func (s *shrikeServiceServer) ListACL(ctx context.Context, req *v1.ListACLRequest) (*v1.ListACLResponse, error) {
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

	// Generate SQL to select all columns in ACL Table
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildACLListQuery(req.Filters, req.Ordering, req.Limit)
	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ACL-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.ACL{}
	var createdAt pq.NullTime
	var updatedAt pq.NullTime

	for rows.Next() {
		acl := new(v1.ACL)
		if err := rows.Scan(&acl.ID, &createdAt, &updatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ACL row-> "+err.Error())
		}
		// Convert pq.NullTime from database into proto timestamp.
		if createdAt.Valid {
			acl.CreatedAt, err = ptypes.TimestampProto(createdAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}
		if updatedAt.Valid {
			acl.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time)
			if err != nil {
				return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
			}
		}

		list = append(list, acl)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ACL-> "+err.Error())
	}

	return &v1.ListACLResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ACL
func (s *shrikeServiceServer) UpdateACL(ctx context.Context, req *v1.UpdateACLRequest) (*v1.UpdateACLResponse, error) {
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

	// update acl
	res, err := c.ExecContext(ctx, "UPDATE acl SET  WHERE id=$1",
		req.Item.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ACL-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ACL with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateACLResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete acl
func (s *shrikeServiceServer) DeleteACL(ctx context.Context, req *v1.DeleteACLRequest) (*v1.DeleteACLResponse, error) {
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

	// delete acl
	res, err := c.ExecContext(ctx, "DELETE FROM acl WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ACL-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ACL with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteACLResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
