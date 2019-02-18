package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new ComponentType
func (s *shrikeServiceServer) CreateComponentType(ctx context.Context, req *v1.CreateComponentTypeRequest) (*v1.CreateComponentTypeResponse, error) {
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
	// insert ComponentType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO component_type (id, created_at, updated_at, title) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ComponentType-> "+err.Error())
	}

	// get ID of creates ComponentType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ComponentType-> "+err.Error())
	}

	return &v1.CreateComponentTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get component_type by id.
func (s *shrikeServiceServer) GetComponentType(ctx context.Context, req *v1.GetComponentTypeRequest) (*v1.GetComponentTypeResponse, error) {
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

	// query ComponentType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM component_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentType with ID='%d' is not found",
			req.ID))
	}

	// get ComponentType data
	var componenttype v1.ComponentType
	if err := rows.Scan( &componenttype.ID,  &componenttype.CreatedAt,  &componenttype.UpdatedAt,  &componenttype.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ComponentType rows with ID='%d'",
			req.ID))
	}

	return &v1.GetComponentTypeResponse{
		Api:  apiVersion,
		Item: &componenttype,
	}, nil

}

// Read all ComponentType
func (s *shrikeServiceServer) ListComponentType(ctx context.Context, req *v1.ListComponentTypeRequest) (*v1.ListComponentTypeResponse, error) {
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

	// get ComponentType list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM component_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.ComponentType{}
	for rows.Next() {
		componenttype := new(v1.ComponentType)
		if err := rows.Scan( &componenttype.ID,  &componenttype.CreatedAt,  &componenttype.UpdatedAt,  &componenttype.Title, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentType row-> "+err.Error())
		}
		list = append(list, componenttype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentType-> "+err.Error())
	}

	return &v1.ListComponentTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ComponentType
func (s *shrikeServiceServer) UpdateComponentType(ctx context.Context, req *v1.UpdateComponentTypeRequest) (*v1.UpdateComponentTypeResponse, error) {
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

	// update component_type
	res, err := c.ExecContext(ctx, "UPDATE component_type SET id=$1, created_at=$2, updated_at=$3, title=$4 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ComponentType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentType with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateComponentTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete component_type
func (s *shrikeServiceServer) DeleteComponentType(ctx context.Context, req *v1.DeleteComponentTypeRequest) (*v1.DeleteComponentTypeResponse, error) {
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

	// delete component_type
	res, err := c.ExecContext(ctx, "DELETE FROM component_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ComponentType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentType with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteComponentTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}