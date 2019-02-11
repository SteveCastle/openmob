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

// NewShrikeServiceServer creates Issue service
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
func (s *shrikeServiceServer) CreateIssue(ctx context.Context, req *v1.CreateIssueRequest) (*v1.CreateIssueResponse, error) {
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
	// insert Issue entity data
	err = c.QueryRowContext(ctx, "INSERT INTO issue (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Issue-> "+err.Error())
	}

	// get ID of creates Issue
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Issue-> "+err.Error())
	}

	return &v1.CreateIssueResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get issue by id.
func (s *shrikeServiceServer) GetIssue(ctx context.Context, req *v1.GetIssueRequest) (*v1.GetIssueResponse, error) {
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

	// query Issue by ID
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM issue WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Issue-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Issue-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%d' is not found",
			req.Id))
	}

	// get Issue data
	var td v1.Issue
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Issue row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Issue rows with ID='%d'",
			req.Id))
	}

	return &v1.GetIssueResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
func (s *shrikeServiceServer) ListIssue(ctx context.Context, req *v1.ListIssueRequest) (*v1.ListIssueResponse, error) {
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

	// get Issue list
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM issue")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Issue-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Issue{}
	for rows.Next() {
		td := new(v1.Issue)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Issue row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Issue-> "+err.Error())
	}

	return &v1.ListIssueResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
func (s *shrikeServiceServer) UpdateIssue(ctx context.Context, req *v1.UpdateIssueRequest) (*v1.UpdateIssueResponse, error) {
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

	// update issue
	res, err := c.ExecContext(ctx, "UPDATE issue SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Issue-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateIssueResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete issue
func (s *shrikeServiceServer) DeleteIssue(ctx context.Context, req *v1.DeleteIssueRequest) (*v1.DeleteIssueResponse, error) {
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

	// delete issue
	res, err := c.ExecContext(ctx, "DELETE FROM issue WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Issue-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteIssueResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
