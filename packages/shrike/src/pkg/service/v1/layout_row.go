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

// Create new LayoutRow
func (s *shrikeServiceServer) CreateLayoutRow(ctx context.Context, req *v1.CreateLayoutRowRequest) (*v1.CreateLayoutRowResponse, error) {
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
	// insert LayoutRow entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layout_row (layout, container) VALUES($1, $2)  RETURNING id;",
		req.Item.Layout, req.Item.Container).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LayoutRow-> "+err.Error())
	}

	// get ID of creates LayoutRow
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LayoutRow-> "+err.Error())
	}

	return &v1.CreateLayoutRowResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get layout_row by id.
func (s *shrikeServiceServer) GetLayoutRow(ctx context.Context, req *v1.GetLayoutRowRequest) (*v1.GetLayoutRowResponse, error) {
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

	// query LayoutRow by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout, container FROM layout_row WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutRow-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutRow-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%s' is not found",
			req.ID))
	}

	// scan LayoutRow data into protobuf model
	var layoutrow v1.LayoutRow
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&layoutrow.ID, &createdAt, &updatedAt, &layoutrow.Layout, &layoutrow.Container); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutRow row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	layoutrow.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	layoutrow.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutRow rows with ID='%s'",
			req.ID))
	}

	return &v1.GetLayoutRowResponse{
		Api:  apiVersion,
		Item: &layoutrow,
	}, nil

}

// Read all LayoutRow
func (s *shrikeServiceServer) ListLayoutRow(ctx context.Context, req *v1.ListLayoutRowRequest) (*v1.ListLayoutRowResponse, error) {
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

	// Generate SQL to select all columns in LayoutRow Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, layout, container FROM layout_row"
	querySQL := queries.BuildLayoutRowFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutRow-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.LayoutRow{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		layoutrow := new(v1.LayoutRow)
		if err := rows.Scan(&layoutrow.ID, &createdAt, &updatedAt, &layoutrow.Layout, &layoutrow.Container); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutRow row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		layoutrow.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		layoutrow.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}

		list = append(list, layoutrow)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutRow-> "+err.Error())
	}

	return &v1.ListLayoutRowResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update LayoutRow
func (s *shrikeServiceServer) UpdateLayoutRow(ctx context.Context, req *v1.UpdateLayoutRowRequest) (*v1.UpdateLayoutRowResponse, error) {
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

	// update layout_row
	res, err := c.ExecContext(ctx, "UPDATE layout_row SET layout=$2, container=$3 WHERE id=$1",
		req.Item.ID, req.Item.Layout, req.Item.Container)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutRow-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateLayoutRowResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete layout_row
func (s *shrikeServiceServer) DeleteLayoutRow(ctx context.Context, req *v1.DeleteLayoutRowRequest) (*v1.DeleteLayoutRowResponse, error) {
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

	// delete layout_row
	res, err := c.ExecContext(ctx, "DELETE FROM layout_row WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LayoutRow-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteLayoutRowResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
