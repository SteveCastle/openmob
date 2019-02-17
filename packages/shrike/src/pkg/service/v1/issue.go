package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


// Create new Issue
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
	err = c.QueryRowContext(ctx, "INSERT INTO issue (id, created_at, updated_at, title, election, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title,  req.Item.Election, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Issue-> "+err.Error())
	}

	// get ID of creates Issue
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Issue-> "+err.Error())
	}

	return &v1.CreateIssueResponse{
		Api: apiVersion,
		ID:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, election,  FROM issue WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Issue-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Issue-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%d' is not found",
			req.ID))
	}

	// get Issue data
	var issue v1.Issue
	if err := rows.Scan( &issue.ID,  &issue.CreatedAt,  &issue.UpdatedAt,  &issue.Title,  &issue.Election, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Issue row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Issue rows with ID='%d'",
			req.ID))
	}

	return &v1.GetIssueResponse{
		Api:  apiVersion,
		Item: &issue,
	}, nil

}

// Read all Issue
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, election,  FROM issue")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Issue-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Issue{}
	for rows.Next() {
		issue := new(v1.Issue)
		if err := rows.Scan( &issue.ID,  &issue.CreatedAt,  &issue.UpdatedAt,  &issue.Title,  &issue.Election, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Issue row-> "+err.Error())
		}
		list = append(list, issue)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Issue-> "+err.Error())
	}

	return &v1.ListIssueResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Issue
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
	res, err := c.ExecContext(ctx, "UPDATE issue SET $1, $2, $3, $4, $5,  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title,req.Item.Election, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Issue-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%d' is not found",
			req.Item.ID))
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
	res, err := c.ExecContext(ctx, "DELETE FROM issue WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Issue-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteIssueResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
