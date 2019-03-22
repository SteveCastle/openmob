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

// Create new Activity
func (s *shrikeServiceServer) CreateActivity(ctx context.Context, req *v1.CreateActivityRequest) (*v1.CreateActivityResponse, error) {
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
	// insert Activity entity data
	err = c.QueryRowContext(ctx, "INSERT INTO activity (title, activity_type, contact, cause) VALUES($1, $2, $3, $4)  RETURNING id;",
		req.Item.Title, req.Item.ActivityType, req.Item.Contact, req.Item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Activity-> "+err.Error())
	}

	// get ID of creates Activity
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Activity-> "+err.Error())
	}

	return &v1.CreateActivityResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get activity by id.
func (s *shrikeServiceServer) GetActivity(ctx context.Context, req *v1.GetActivityRequest) (*v1.GetActivityResponse, error) {
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

	// query Activity by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, activity_type, contact, cause FROM activity WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Activity-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Activity-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Activity with ID='%s' is not found",
			req.ID))
	}

	// scan Activity data into protobuf model
	var activity v1.Activity
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&activity.ID, &createdAt, &updatedAt, &activity.Title, &activity.ActivityType, &activity.Contact, &activity.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Activity row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	activity.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	activity.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Activity rows with ID='%s'",
			req.ID))
	}

	return &v1.GetActivityResponse{
		Api:  apiVersion,
		Item: &activity,
	}, nil

}

// Read all Activity
func (s *shrikeServiceServer) ListActivity(ctx context.Context, req *v1.ListActivityRequest) (*v1.ListActivityResponse, error) {
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

	// Generate SQL to select all columns in Activity Table
	// Then generate filtering and ordering sql and finally run query.

	baseSQL := "SELECT id, created_at, updated_at, title, activity_type, contact, cause FROM activity"
	querySQL := queries.BuildActivityFilters(req.Filters, req.Ordering, req.Limit)
	SQL := fmt.Sprintf("%s %s", baseSQL, querySQL)
	rows, err := c.QueryContext(ctx, SQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Activity-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.Activity{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		activity := new(v1.Activity)
		if err := rows.Scan(&activity.ID, &createdAt, &updatedAt, &activity.Title, &activity.ActivityType, &activity.Contact, &activity.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Activity row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		activity.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		activity.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}

		list = append(list, activity)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Activity-> "+err.Error())
	}

	return &v1.ListActivityResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Activity
func (s *shrikeServiceServer) UpdateActivity(ctx context.Context, req *v1.UpdateActivityRequest) (*v1.UpdateActivityResponse, error) {
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

	// update activity
	res, err := c.ExecContext(ctx, "UPDATE activity SET title=$2, activity_type=$3, contact=$4, cause=$5 WHERE id=$1",
		req.Item.ID, req.Item.Title, req.Item.ActivityType, req.Item.Contact, req.Item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Activity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Activity with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateActivityResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete activity
func (s *shrikeServiceServer) DeleteActivity(ctx context.Context, req *v1.DeleteActivityRequest) (*v1.DeleteActivityResponse, error) {
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

	// delete activity
	res, err := c.ExecContext(ctx, "DELETE FROM activity WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Activity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Activity with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteActivityResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
