package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new ActivityType
func (s *shrikeServiceServer) CreateActivityType(ctx context.Context, req *v1.CreateActivityTypeRequest) (*v1.CreateActivityTypeResponse, error) {
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
	// insert ActivityType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO activity_type (id, created_at, updated_at, title) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ActivityType-> "+err.Error())
	}

	// get ID of creates ActivityType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ActivityType-> "+err.Error())
	}

	return &v1.CreateActivityTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get activity_type by id.
func (s *shrikeServiceServer) GetActivityType(ctx context.Context, req *v1.GetActivityTypeRequest) (*v1.GetActivityTypeResponse, error) {
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

	// query ActivityType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM activity_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ActivityType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ActivityType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ActivityType with ID='%d' is not found",
			req.ID))
	}

	// get ActivityType data
	var activitytype v1.ActivityType
	if err := rows.Scan( &activitytype.ID,  &activitytype.CreatedAt,  &activitytype.UpdatedAt,  &activitytype.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ActivityType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ActivityType rows with ID='%d'",
			req.ID))
	}

	return &v1.GetActivityTypeResponse{
		Api:  apiVersion,
		Item: &activitytype,
	}, nil

}

// Read all ActivityType
func (s *shrikeServiceServer) ListActivityType(ctx context.Context, req *v1.ListActivityTypeRequest) (*v1.ListActivityTypeResponse, error) {
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

	// get ActivityType list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM activity_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ActivityType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.ActivityType{}
	for rows.Next() {
		activitytype := new(v1.ActivityType)
		if err := rows.Scan( &activitytype.ID,  &activitytype.CreatedAt,  &activitytype.UpdatedAt,  &activitytype.Title, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ActivityType row-> "+err.Error())
		}
		list = append(list, activitytype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ActivityType-> "+err.Error())
	}

	return &v1.ListActivityTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update ActivityType
func (s *shrikeServiceServer) UpdateActivityType(ctx context.Context, req *v1.UpdateActivityTypeRequest) (*v1.UpdateActivityTypeResponse, error) {
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

	// update activity_type
	res, err := c.ExecContext(ctx, "UPDATE activity_type SET $1 ,$2 ,$3 ,$4  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ActivityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ActivityType with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateActivityTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete activity_type
func (s *shrikeServiceServer) DeleteActivityType(ctx context.Context, req *v1.DeleteActivityTypeRequest) (*v1.DeleteActivityTypeResponse, error) {
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

	// delete activity_type
	res, err := c.ExecContext(ctx, "DELETE FROM activity_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ActivityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ActivityType with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteActivityTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
