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

// NewShrikeServiceServer creates LayoutColumn service
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
func (s *shrikeServiceServer) CreateLayoutColumn(ctx context.Context, req *v1.CreateLayoutColumnRequest) (*v1.CreateLayoutColumnResponse, error) {
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
	// insert LayoutColumn entity data
	err = c.QueryRowContext(ctx, "INSERT INTO layoutcolumn (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LayoutColumn-> "+err.Error())
	}

	// get ID of creates LayoutColumn
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LayoutColumn-> "+err.Error())
	}

	return &v1.CreateLayoutColumnResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get layoutcolumn by id.
func (s *shrikeServiceServer) GetLayoutColumn(ctx context.Context, req *v1.GetLayoutColumnRequest) (*v1.GetLayoutColumnResponse, error) {
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

	// query LayoutColumn by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM layoutcolumn WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutColumn-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutColumn-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutColumn with ID='%d' is not found",
			req.Id))
	}

	// get LayoutColumn data
	var td v1.LayoutColumn
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutColumn row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutColumn rows with ID='%d'",
			req.Id))
	}

	return &v1.GetLayoutColumnResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListLayoutColumn(ctx context.Context, req *v1.ListLayoutColumnRequest) (*v1.ListLayoutColumnResponse, error) {
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

	// get LayoutColumn list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM LayoutColumn")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutColumn-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LayoutColumn{}
	for rows.Next() {
		td := new(v1.LayoutColumn)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutColumn row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutColumn-> "+err.Error())
	}

	return &v1.ListLayoutColumnResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateLayoutColumn(ctx context.Context, req *v1.UpdateLayoutColumnRequest) (*v1.UpdateLayoutColumnResponse, error) {
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

	// update layoutcolumn
	res, err := c.ExecContext(ctx, "UPDATE layoutcolumn SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update layoutcolumn-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("layoutcolumn with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateLayoutColumnResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete layoutcolumn
func (s *shrikeServiceServer) DeleteLayoutColumn(ctx context.Context, req *v1.DeleteLayoutColumnRequest) (*v1.DeleteLayoutColumnResponse, error) {
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

	// delete layoutcolumn
	res, err := c.ExecContext(ctx, "DELETE FROM layoutcolumn WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete layoutcolumn-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("layoutcolumn with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteLayoutColumnResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
