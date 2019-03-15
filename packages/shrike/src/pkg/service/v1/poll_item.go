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

// Create new PollItem
func (s *shrikeServiceServer) CreatePollItem(ctx context.Context, req *v1.CreatePollItemRequest) (*v1.CreatePollItemResponse, error) {
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
	// insert PollItem entity data
	err = c.QueryRowContext(ctx, "INSERT INTO poll_item (title, poll) VALUES($1, $2)  RETURNING id;",
		req.Item.Title, req.Item.Poll).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PollItem-> "+err.Error())
	}

	// get ID of creates PollItem
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PollItem-> "+err.Error())
	}

	return &v1.CreatePollItemResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get poll_item by id.
func (s *shrikeServiceServer) GetPollItem(ctx context.Context, req *v1.GetPollItemRequest) (*v1.GetPollItemResponse, error) {
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

	// query PollItem by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, poll FROM poll_item WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollItem-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PollItem-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollItem with ID='%s' is not found",
			req.ID))
	}

	// scan PollItem data into protobuf model
	var pollitem v1.PollItem
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&pollitem.ID, &createdAt, &updatedAt, &pollitem.Title, &pollitem.Poll); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollItem row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	pollitem.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	pollitem.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PollItem rows with ID='%s'",
			req.ID))
	}

	return &v1.GetPollItemResponse{
		Api:  apiVersion,
		Item: &pollitem,
	}, nil

}

// Read all PollItem
func (s *shrikeServiceServer) ListPollItem(ctx context.Context, req *v1.ListPollItemRequest) (*v1.ListPollItemResponse, error) {
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

	// Generate SQL to select all columns in PollItem Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title, poll FROM poll_item"
	querySQL := queries.BuildPollItemFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollItem-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.PollItem{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		pollitem := new(v1.PollItem)
		if err := rows.Scan(&pollitem.ID, &createdAt, &updatedAt, &pollitem.Title, &pollitem.Poll); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollItem row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		pollitem.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		pollitem.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, pollitem)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PollItem-> "+err.Error())
	}

	return &v1.ListPollItemResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update PollItem
func (s *shrikeServiceServer) UpdatePollItem(ctx context.Context, req *v1.UpdatePollItemRequest) (*v1.UpdatePollItemResponse, error) {
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

	// update poll_item
	res, err := c.ExecContext(ctx, "UPDATE poll_item SET title=$2, poll=$3 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.Poll)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PollItem-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollItem with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdatePollItemResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete poll_item
func (s *shrikeServiceServer) DeletePollItem(ctx context.Context, req *v1.DeletePollItemRequest) (*v1.DeletePollItemResponse, error) {
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

	// delete poll_item
	res, err := c.ExecContext(ctx, "DELETE FROM poll_item WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PollItem-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollItem with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeletePollItemResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
