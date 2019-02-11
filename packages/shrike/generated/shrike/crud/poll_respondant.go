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

// NewShrikeServiceServer creates PollRespondant service
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
func (s *shrikeServiceServer) CreatePollRespondant(ctx context.Context, req *v1.CreatePollRespondantRequest) (*v1.CreatePollRespondantResponse, error) {
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
	// insert PollRespondant entity data
	err = c.QueryRowContext(ctx, "INSERT INTO poll_respondant (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PollRespondant-> "+err.Error())
	}

	// get ID of creates PollRespondant
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PollRespondant-> "+err.Error())
	}

	return &v1.CreatePollRespondantResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get poll_respondant by id.
func (s *shrikeServiceServer) GetPollRespondant(ctx context.Context, req *v1.GetPollRespondantRequest) (*v1.GetPollRespondantResponse, error) {
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

	// query PollRespondant by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM poll_respondant WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollRespondant-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PollRespondant-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%d' is not found",
			req.Id))
	}

	// get PollRespondant data
	var td v1.PollRespondant
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollRespondant row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PollRespondant rows with ID='%d'",
			req.Id))
	}

	return &v1.GetPollRespondantResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListPollRespondant(ctx context.Context, req *v1.ListPollRespondantRequest) (*v1.ListPollRespondantResponse, error) {
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

	// get PollRespondant list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM poll_respondant")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollRespondant-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.PollRespondant{}
	for rows.Next() {
		td := new(v1.PollRespondant)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollRespondant row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PollRespondant-> "+err.Error())
	}

	return &v1.ListPollRespondantResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdatePollRespondant(ctx context.Context, req *v1.UpdatePollRespondantRequest) (*v1.UpdatePollRespondantResponse, error) {
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

	// update poll_respondant
	res, err := c.ExecContext(ctx, "UPDATE poll_respondant SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PollRespondant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdatePollRespondantResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete poll_respondant
func (s *shrikeServiceServer) DeletePollRespondant(ctx context.Context, req *v1.DeletePollRespondantRequest) (*v1.DeletePollRespondantResponse, error) {
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

	// delete poll_respondant
	res, err := c.ExecContext(ctx, "DELETE FROM poll_respondant WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PollRespondant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeletePollRespondantResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
