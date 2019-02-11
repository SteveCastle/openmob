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

// NewShrikeServiceServer creates ComponentImplementation service
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
func (s *shrikeServiceServer) CreateComponentImplementation(ctx context.Context, req *v1.CreateComponentImplementationRequest) (*v1.CreateComponentImplementationResponse, error) {
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
	// insert ComponentImplementation entity data
	err = c.QueryRowContext(ctx, "INSERT INTO component_implementation (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ComponentImplementation-> "+err.Error())
	}

	// get ID of creates ComponentImplementation
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ComponentImplementation-> "+err.Error())
	}

	return &v1.CreateComponentImplementationResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get component_implementation by id.
func (s *shrikeServiceServer) GetComponentImplementation(ctx context.Context, req *v1.GetComponentImplementationRequest) (*v1.GetComponentImplementationResponse, error) {
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

	// query ComponentImplementation by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM component_implementation WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentImplementation-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentImplementation-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%d' is not found",
			req.Id))
	}

	// get ComponentImplementation data
	var td v1.ComponentImplementation
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentImplementation row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ComponentImplementation rows with ID='%d'",
			req.Id))
	}

	return &v1.GetComponentImplementationResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListComponentImplementation(ctx context.Context, req *v1.ListComponentImplementationRequest) (*v1.ListComponentImplementationResponse, error) {
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

	// get ComponentImplementation list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM component_implementation")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentImplementation-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.ComponentImplementation{}
	for rows.Next() {
		td := new(v1.ComponentImplementation)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentImplementation row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentImplementation-> "+err.Error())
	}

	return &v1.ListComponentImplementationResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateComponentImplementation(ctx context.Context, req *v1.UpdateComponentImplementationRequest) (*v1.UpdateComponentImplementationResponse, error) {
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

	// update component_implementation
	res, err := c.ExecContext(ctx, "UPDATE component_implementation SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ComponentImplementation-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateComponentImplementationResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete component_implementation
func (s *shrikeServiceServer) DeleteComponentImplementation(ctx context.Context, req *v1.DeleteComponentImplementationRequest) (*v1.DeleteComponentImplementationResponse, error) {
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

	// delete component_implementation
	res, err := c.ExecContext(ctx, "DELETE FROM component_implementation WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ComponentImplementation-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteComponentImplementationResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
