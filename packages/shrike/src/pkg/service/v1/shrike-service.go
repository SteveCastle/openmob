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

// NewShrikeServiceServer creates Cause service
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
func (s *shrikeServiceServer) CreateCause(ctx context.Context, req *v1.CreateCauseRequest) (*v1.CreateCauseResponse, error) {
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
	// insert Cause entity data
	err = c.QueryRowContext(ctx, "INSERT INTO cause (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Cause-> "+err.Error())
	}

	// get ID of creates Cause
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Cause-> "+err.Error())
	}

	return &v1.CreateCauseResponse{
		Id: id,
	}, nil
}

// Get cause by id.
func (s *shrikeServiceServer) GetCause(ctx context.Context, req *v1.GetCauseRequest) (*v1.GetCauseResponse, error) {
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

	// query Cause by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM cause WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cause-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Cause-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%d' is not found",
			req.Id))
	}

	// get Cause data
	var td v1.Cause
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cause row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Cause rows with ID='%d'",
			req.Id))
	}

	return &v1.GetCauseResponse{
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListCause(ctx context.Context, req *v1.ListCauseRequest) (*v1.ListCauseResponse, error) {
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

	// get Cause list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM Cause")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cause-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Cause{}
	for rows.Next() {
		td := new(v1.Cause)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cause row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Cause-> "+err.Error())
	}

	return &v1.ListCauseResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}
