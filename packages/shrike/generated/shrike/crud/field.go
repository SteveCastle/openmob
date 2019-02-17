package v1

import (
	"context"
	"database/sql"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// shrikeServiceServer is implementation of v1.ShrikeServiceServer proto interface
type shrikeServiceServer struct {
	db *sql.DB
}

// NewShrikeServiceServer creates Field service
func NewShrikeServiceServer(db *sql.DB) v1.ShrikeServiceServer {
	return &shrikeServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *shrikeServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *shrikeServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Create new Field
func (s *shrikeServiceServer) CreateField(ctx context.Context, req *v1.CreateFieldRequest) (*v1.CreateFieldResponse, error) {
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
	// insert Field entity data
	err = c.QueryRowContext(ctx, "INSERT INTO field (id, created_at, updated_at, field_type, component, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemFieldType  req.ItemComponent ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Field-> "+err.Error())
	}

	// get ID of creates Field
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Field-> "+err.Error())
	}

	return &v1.CreateFieldResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get field by id.
func (s *shrikeServiceServer) GetField(ctx context.Context, req *v1.GetFieldRequest) (*v1.GetFieldResponse, error) {
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

	// query Field by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, field_type, component,  FROM field WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Field-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Field-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%d' is not found",
			req.Id))
	}

	// get Field data
	var field v1.Field
	if err := rows.Scan( &field.ID,  &field.CreatedAt,  &field.UpdatedAt,  &field.FieldType,  &field.Component, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Field row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Field rows with ID='%d'",
			req.Id))
	}

	return &v1.GetFieldResponse{
		Api:  apiVersion,
		Item: &field,
	}, nil

}

// Read all Field
func (s *shrikeServiceServer) ListField(ctx context.Context, req *v1.ListFieldRequest) (*v1.ListFieldResponse, error) {
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

	// get Field list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, field_type, component,  FROM field")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Field-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Field{}
	for rows.Next() {
		field := new(v1.Field)
		if err := rows.Scan( &field.ID,  &field.CreatedAt,  &field.UpdatedAt,  &field.FieldType,  &field.Component, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Field row-> "+err.Error())
		}
		list = append(list, field)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Field-> "+err.Error())
	}

	return &v1.ListFieldResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Field
func (s *shrikeServiceServer) UpdateField(ctx context.Context, req *v1.UpdateFieldRequest) (*v1.UpdateFieldResponse, error) {
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

	// update field
	res, err := c.ExecContext(ctx, "UPDATE field SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Field-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateFieldResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete field
func (s *shrikeServiceServer) DeleteField(ctx context.Context, req *v1.DeleteFieldRequest) (*v1.DeleteFieldResponse, error) {
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

	// delete field
	res, err := c.ExecContext(ctx, "DELETE FROM field WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Field-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteFieldResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
