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

// NewShrikeServiceServer creates FieldType service
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

// Create new todo task
func (s *shrikeServiceServer) CreateFieldType(ctx context.Context, req *v1.CreateFieldTypeRequest) (*v1.CreateFieldTypeResponse, error) {
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
	// insert FieldType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO field_type (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into FieldType-> "+err.Error())
	}

	// get ID of creates FieldType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created FieldType-> "+err.Error())
	}

	return &v1.CreateFieldTypeResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get field_type by id.
func (s *shrikeServiceServer) GetFieldType(ctx context.Context, req *v1.GetFieldTypeRequest) (*v1.GetFieldTypeResponse, error) {
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

	// query FieldType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM field_type WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from FieldType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from FieldType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%d' is not found",
			req.Id))
	}

	// get FieldType data
	var td v1.FieldType
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from FieldType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple FieldType rows with ID='%d'",
			req.Id))
	}

	return &v1.GetFieldTypeResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListFieldType(ctx context.Context, req *v1.ListFieldTypeRequest) (*v1.ListFieldTypeResponse, error) {
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

	// get FieldType list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM field_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from FieldType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.FieldType{}
	for rows.Next() {
		td := new(v1.FieldType)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from FieldType row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from FieldType-> "+err.Error())
	}

	return &v1.ListFieldTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateFieldType(ctx context.Context, req *v1.UpdateFieldTypeRequest) (*v1.UpdateFieldTypeResponse, error) {
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

	// update field_type
	res, err := c.ExecContext(ctx, "UPDATE field_type SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update FieldType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateFieldTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete field_type
func (s *shrikeServiceServer) DeleteFieldType(ctx context.Context, req *v1.DeleteFieldTypeRequest) (*v1.DeleteFieldTypeResponse, error) {
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

	// delete field_type
	res, err := c.ExecContext(ctx, "DELETE FROM field_type WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete FieldType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteFieldTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
