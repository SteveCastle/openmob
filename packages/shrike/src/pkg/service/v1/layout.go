package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Layout
func (s *shrikeServiceServer) CreateLayout(ctx context.Context, req *v1.CreateLayoutRequest) (*v1.CreateLayoutResponse, error) {
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
	// insert Layout entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layout (id, created_at, updated_at, layout_type) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.LayoutType, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Layout-> "+err.Error())
	}

	// get ID of creates Layout
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Layout-> "+err.Error())
	}

	return &v1.CreateLayoutResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get layout by id.
func (s *shrikeServiceServer) GetLayout(ctx context.Context, req *v1.GetLayoutRequest) (*v1.GetLayoutResponse, error) {
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

	// query Layout by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout_type FROM layout WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Layout-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Layout-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%d' is not found",
			req.ID))
	}

	// get Layout data
	var layout v1.Layout
	if err := rows.Scan( &layout.ID,  &layout.CreatedAt,  &layout.UpdatedAt,  &layout.LayoutType, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Layout row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Layout rows with ID='%d'",
			req.ID))
	}

	return &v1.GetLayoutResponse{
		Api:  apiVersion,
		Item: &layout,
	}, nil

}

// Read all Layout
func (s *shrikeServiceServer) ListLayout(ctx context.Context, req *v1.ListLayoutRequest) (*v1.ListLayoutResponse, error) {
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

	// get Layout list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout_type FROM layout")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Layout-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Layout{}
	for rows.Next() {
		layout := new(v1.Layout)
		if err := rows.Scan( &layout.ID,  &layout.CreatedAt,  &layout.UpdatedAt,  &layout.LayoutType, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Layout row-> "+err.Error())
		}
		list = append(list, layout)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Layout-> "+err.Error())
	}

	return &v1.ListLayoutResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Layout
func (s *shrikeServiceServer) UpdateLayout(ctx context.Context, req *v1.UpdateLayoutRequest) (*v1.UpdateLayoutResponse, error) {
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

	// update layout
	res, err := c.ExecContext(ctx, "UPDATE layout SET id=$1, created_at=$2, updated_at=$3, layout_type=$4 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.LayoutType, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Layout-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateLayoutResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete layout
func (s *shrikeServiceServer) DeleteLayout(ctx context.Context, req *v1.DeleteLayoutRequest) (*v1.DeleteLayoutResponse, error) {
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

	// delete layout
	res, err := c.ExecContext(ctx, "DELETE FROM layout WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Layout-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteLayoutResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
