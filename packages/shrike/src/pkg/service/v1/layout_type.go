package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

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
	var id int64
	// insert LayoutType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layout_type (id, created_at, updated_at, title) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title, ).Scan(&id)
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%d' is not found",
			req.ID))
	}

	// get LayoutType data
	var layouttype v1.LayoutType
	if err := rows.Scan( &layouttype.ID,  &layouttype.CreatedAt,  &layouttype.UpdatedAt,  &layouttype.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutType rows with ID='%d'",
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

	// get LayoutType list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM layout_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LayoutType{}
	for rows.Next() {
		layouttype := new(v1.LayoutType)
		if err := rows.Scan( &layouttype.ID,  &layouttype.CreatedAt,  &layouttype.UpdatedAt,  &layouttype.Title, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutType row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE layout_type SET id=$1, created_at=$2, updated_at=$3, title=$4 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%d' is not found",
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteLayoutTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}