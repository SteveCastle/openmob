package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

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
	var id int64
	// insert LayoutRow entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layout_row (id, created_at, updated_at, layout) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Layout, ).Scan(&id)
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout FROM layout_row WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutRow-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutRow-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%d' is not found",
			req.ID))
	}

	// get LayoutRow data
	var layoutrow v1.LayoutRow
	if err := rows.Scan( &layoutrow.ID,  &layoutrow.CreatedAt,  &layoutrow.UpdatedAt,  &layoutrow.Layout, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutRow row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutRow rows with ID='%d'",
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

	// get LayoutRow list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout FROM layout_row")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutRow-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LayoutRow{}
	for rows.Next() {
		layoutrow := new(v1.LayoutRow)
		if err := rows.Scan( &layoutrow.ID,  &layoutrow.CreatedAt,  &layoutrow.UpdatedAt,  &layoutrow.Layout, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutRow row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE layout_row SET id=$1, created_at=$2, updated_at=$3, layout=$4 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Layout, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutRow-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%d' is not found",
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteLayoutRowResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}