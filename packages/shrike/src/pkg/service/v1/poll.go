package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Poll
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
	var id string
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
		ID:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM poll WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Poll-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Poll-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%s' is not found",
			req.ID))
	}

	// scan Poll data into protobuf model
	var poll v1.Poll
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&poll.ID, &createdAt, &updatedAt, &poll.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Poll row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	poll.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	poll.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Poll rows with ID='%s'",
			req.ID))
	}

	return &v1.GetPollResponse{
		Api:  apiVersion,
		Item: &poll,
	}, nil

}

// Read all Poll
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

	// Generate SQL to select all columns in Poll Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title FROM poll"
	querySQL := queries.BuildPollFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Poll-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Poll{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		poll := new(v1.Poll)
		if err := rows.Scan(&poll.ID, &createdAt, &updatedAt, &poll.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Poll row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		poll.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		poll.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, poll)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Poll-> "+err.Error())
	}

	return &v1.ListPollResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Poll
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
	res, err := c.ExecContext(ctx, "UPDATE poll SET title=$2 WHERE id=$1",
		req.Item.ID, req.Item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Poll-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%s' is not found",
			req.Item.ID))
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
	res, err := c.ExecContext(ctx, "DELETE FROM poll WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Poll-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePollResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
