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

// NewShrikeServiceServer creates Boycott service
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
func (s *shrikeServiceServer) CreateBoycott(ctx context.Context, req *v1.CreateBoycottRequest) (*v1.CreateBoycottResponse, error) {
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
	// insert Boycott entity data
	err = c.QueryRowContext(ctx, "INSERT INTO boycott (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Boycott-> "+err.Error())
	}

	// get ID of creates Boycott
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Boycott-> "+err.Error())
	}

	return &v1.CreateBoycottResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get boycott by id.
func (s *shrikeServiceServer) GetBoycott(ctx context.Context, req *v1.GetBoycottRequest) (*v1.GetBoycottResponse, error) {
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

	// query Boycott by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM boycott WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Boycott-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Boycott-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%d' is not found",
			req.Id))
	}

	// get Boycott data
	var td v1.Boycott
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Boycott row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Boycott rows with ID='%d'",
			req.Id))
	}

	return &v1.GetBoycottResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListBoycott(ctx context.Context, req *v1.ListBoycottRequest) (*v1.ListBoycottResponse, error) {
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

	// get Boycott list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM boycott")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Boycott-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Boycott{}
	for rows.Next() {
		td := new(v1.Boycott)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Boycott row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Boycott-> "+err.Error())
	}

	return &v1.ListBoycottResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateBoycott(ctx context.Context, req *v1.UpdateBoycottRequest) (*v1.UpdateBoycottResponse, error) {
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

	// update boycott
	res, err := c.ExecContext(ctx, "UPDATE boycott SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Boycott-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateBoycottResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete boycott
func (s *shrikeServiceServer) DeleteBoycott(ctx context.Context, req *v1.DeleteBoycottRequest) (*v1.DeleteBoycottResponse, error) {
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

	// delete boycott
	res, err := c.ExecContext(ctx, "DELETE FROM boycott WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Boycott-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteBoycottResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}