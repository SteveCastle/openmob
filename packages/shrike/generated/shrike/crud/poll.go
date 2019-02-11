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

// NewShrikeServiceServer creates Poll service
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
func (s *shrikeServiceServer) CreatePoll(ctx context.Context, req *v1.CreatePollRequest) (*v1.CreatePollResponse, error) {
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
	// insert Poll entity data
	err = c.QueryRowContext(ctx, "INSERT INTO poll (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Poll-> "+err.Error())
	}

	// get ID of creates Poll
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Poll-> "+err.Error())
	}

	return &v1.CreatePollResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get poll by id.
func (s *shrikeServiceServer) GetPoll(ctx context.Context, req *v1.GetPollRequest) (*v1.GetPollResponse, error) {
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

	// query Poll by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM poll WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Poll-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Poll-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%d' is not found",
			req.Id))
	}

	// get Poll data
	var td v1.Poll
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Poll row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Poll rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPollResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListPoll(ctx context.Context, req *v1.ListPollRequest) (*v1.ListPollResponse, error) {
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

	// get Poll list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM poll")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Poll-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Poll{}
	for rows.Next() {
		td := new(v1.Poll)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Poll row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Poll-> "+err.Error())
	}

	return &v1.ListPollResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdatePoll(ctx context.Context, req *v1.UpdatePollRequest) (*v1.UpdatePollResponse, error) {
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

	// update poll
	res, err := c.ExecContext(ctx, "UPDATE poll SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Poll-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePollResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete poll
func (s *shrikeServiceServer) DeletePoll(ctx context.Context, req *v1.DeletePollRequest) (*v1.DeletePollResponse, error) {
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

	// delete poll
	res, err := c.ExecContext(ctx, "DELETE FROM poll WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Poll-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePollResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
