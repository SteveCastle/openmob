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

// Create new Follower
func (s *shrikeServiceServer) CreateFollower(ctx context.Context, req *v1.CreateFollowerRequest) (*v1.CreateFollowerResponse, error) {
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
	// insert Follower entity data
	err = c.QueryRowContext(ctx, "INSERT INTO follower (contact, cause) VALUES($1, $2)  RETURNING id;",
		req.Item.Contact, req.Item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Follower-> "+err.Error())
	}

	// get ID of creates Follower
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Follower-> "+err.Error())
	}

	return &v1.CreateFollowerResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get follower by id.
func (s *shrikeServiceServer) GetFollower(ctx context.Context, req *v1.GetFollowerRequest) (*v1.GetFollowerResponse, error) {
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

	// query Follower by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, contact, cause FROM follower WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Follower-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Follower-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Follower with ID='%s' is not found",
			req.ID))
	}

	// scan Follower data into protobuf model
	var follower v1.Follower
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&follower.ID, &createdAt, &updatedAt, &follower.Contact, &follower.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Follower row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	follower.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	follower.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Follower rows with ID='%s'",
			req.ID))
	}

	return &v1.GetFollowerResponse{
		Api:  apiVersion,
		Item: &follower,
	}, nil

}

// Read all Follower
func (s *shrikeServiceServer) ListFollower(ctx context.Context, req *v1.ListFollowerRequest) (*v1.ListFollowerResponse, error) {
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

	// Generate SQL to select all columns in Follower Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, contact, cause FROM follower"
	querySQL := queries.BuildFollowerFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Follower-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Follower{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		follower := new(v1.Follower)
		if err := rows.Scan(&follower.ID, &createdAt, &updatedAt, &follower.Contact, &follower.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Follower row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		follower.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		follower.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, follower)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Follower-> "+err.Error())
	}

	return &v1.ListFollowerResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Follower
func (s *shrikeServiceServer) UpdateFollower(ctx context.Context, req *v1.UpdateFollowerRequest) (*v1.UpdateFollowerResponse, error) {
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

	// update follower
	res, err := c.ExecContext(ctx, "UPDATE follower SET contact=$2, cause=$3 WHERE id=$1",
		req.Item.ID, req.Item.Contact, req.Item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Follower-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Follower with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateFollowerResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete follower
func (s *shrikeServiceServer) DeleteFollower(ctx context.Context, req *v1.DeleteFollowerRequest) (*v1.DeleteFollowerResponse, error) {
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

	// delete follower
	res, err := c.ExecContext(ctx, "DELETE FROM follower WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Follower-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Follower with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteFollowerResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
