package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new LayoutType
func (s *shrikeServiceServer) CreateLayoutType(ctx context.Context, req *v1.CreateLayoutTypeRequest) (*v1.CreateLayoutTypeResponse, error) {
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
	// insert LayoutType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layout_type (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LayoutType-> "+err.Error())
	}

	// get ID of creates LayoutType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LayoutType-> "+err.Error())
	}

	return &v1.CreateLayoutTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get layout_type by id.
func (s *shrikeServiceServer) GetLayoutType(ctx context.Context, req *v1.GetLayoutTypeRequest) (*v1.GetLayoutTypeResponse, error) {
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

	// query LayoutType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM layout_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%s' is not found",
			req.ID))
	}

	// scan LayoutType data into protobuf model
	var layouttype v1.LayoutType
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&layouttype.ID, &createdAt, &updatedAt, &layouttype.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutType row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	layouttype.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	layouttype.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutType rows with ID='%s'",
			req.ID))
	}

	return &v1.GetLayoutTypeResponse{
		Api:  apiVersion,
		Item: &layouttype,
	}, nil

}

// Read all LayoutType
func (s *shrikeServiceServer) ListLayoutType(ctx context.Context, req *v1.ListLayoutTypeRequest) (*v1.ListLayoutTypeResponse, error) {
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

	// Generate SQL to select all columns in LayoutType Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title FROM layout_type"
	querySQL := queries.BuildLayoutTypeFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutType-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.LayoutType{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		layouttype := new(v1.LayoutType)
		if err := rows.Scan(&layouttype.ID, &createdAt, &updatedAt, &layouttype.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutType row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		layouttype.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		layouttype.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}

		list = append(list, layouttype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutType-> "+err.Error())
	}

	return &v1.ListLayoutTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update LayoutType
func (s *shrikeServiceServer) UpdateLayoutType(ctx context.Context, req *v1.UpdateLayoutTypeRequest) (*v1.UpdateLayoutTypeResponse, error) {
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

	// update layout_type
	res, err := c.ExecContext(ctx, "UPDATE layout_type SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateLayoutTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete layout_type
func (s *shrikeServiceServer) DeleteLayoutType(ctx context.Context, req *v1.DeleteLayoutTypeRequest) (*v1.DeleteLayoutTypeResponse, error) {
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

	// delete layout_type
	res, err := c.ExecContext(ctx, "DELETE FROM layout_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LayoutType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteLayoutTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
