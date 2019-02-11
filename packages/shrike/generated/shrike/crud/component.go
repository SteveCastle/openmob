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

// NewShrikeServiceServer creates Component service
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
func (s *shrikeServiceServer) CreateComponent(ctx context.Context, req *v1.CreateComponentRequest) (*v1.CreateComponentResponse, error) {
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
	// insert Component entity data
	err = c.QueryRowContext(ctx, "INSERT INTO component (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Component-> "+err.Error())
	}

	// get ID of creates Component
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Component-> "+err.Error())
	}

	return &v1.CreateComponentResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get component by id.
func (s *shrikeServiceServer) GetComponent(ctx context.Context, req *v1.GetComponentRequest) (*v1.GetComponentResponse, error) {
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

	// query Component by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM component WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Component-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Component-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%d' is not found",
			req.Id))
	}

	// get Component data
	var td v1.Component
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Component row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Component rows with ID='%d'",
			req.Id))
	}

	return &v1.GetComponentResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListComponent(ctx context.Context, req *v1.ListComponentRequest) (*v1.ListComponentResponse, error) {
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

	// get Component list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM component")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Component-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Component{}
	for rows.Next() {
		td := new(v1.Component)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Component row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Component-> "+err.Error())
	}

	return &v1.ListComponentResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateComponent(ctx context.Context, req *v1.UpdateComponentRequest) (*v1.UpdateComponentResponse, error) {
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

	// update component
	res, err := c.ExecContext(ctx, "UPDATE component SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Component-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateComponentResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete component
func (s *shrikeServiceServer) DeleteComponent(ctx context.Context, req *v1.DeleteComponentRequest) (*v1.DeleteComponentResponse, error) {
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

	// delete component
	res, err := c.ExecContext(ctx, "DELETE FROM component WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Component-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteComponentResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
