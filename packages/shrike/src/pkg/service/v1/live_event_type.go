package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new LiveEventType
func (s *shrikeServiceServer) CreateLiveEventType(ctx context.Context, req *v1.CreateLiveEventTypeRequest) (*v1.CreateLiveEventTypeResponse, error) {
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
	// insert LiveEventType entity data
	err = c.QueryRowContext(ctx, "INSERT INTO live_event_type (id, created_at, updated_at, title) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEventType-> "+err.Error())
	}

	// get ID of creates LiveEventType
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEventType-> "+err.Error())
	}

	return &v1.CreateLiveEventTypeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get live_event_type by id.
func (s *shrikeServiceServer) GetLiveEventType(ctx context.Context, req *v1.GetLiveEventTypeRequest) (*v1.GetLiveEventTypeResponse, error) {
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

	// query LiveEventType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM live_event_type WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%d' is not found",
			req.ID))
	}

	// get LiveEventType data
	var liveeventtype v1.LiveEventType
	if err := rows.Scan( &liveeventtype.ID,  &liveeventtype.CreatedAt,  &liveeventtype.UpdatedAt,  &liveeventtype.Title, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEventType rows with ID='%d'",
			req.ID))
	}

	return &v1.GetLiveEventTypeResponse{
		Api:  apiVersion,
		Item: &liveeventtype,
	}, nil

}

// Read all LiveEventType
func (s *shrikeServiceServer) ListLiveEventType(ctx context.Context, req *v1.ListLiveEventTypeRequest) (*v1.ListLiveEventTypeResponse, error) {
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

	// get LiveEventType list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM live_event_type")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventType-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LiveEventType{}
	for rows.Next() {
		liveeventtype := new(v1.LiveEventType)
		if err := rows.Scan( &liveeventtype.ID,  &liveeventtype.CreatedAt,  &liveeventtype.UpdatedAt,  &liveeventtype.Title, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventType row-> "+err.Error())
		}
		list = append(list, liveeventtype)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventType-> "+err.Error())
	}

	return &v1.ListLiveEventTypeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update LiveEventType
func (s *shrikeServiceServer) UpdateLiveEventType(ctx context.Context, req *v1.UpdateLiveEventTypeRequest) (*v1.UpdateLiveEventTypeResponse, error) {
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

	// update live_event_type
	res, err := c.ExecContext(ctx, "UPDATE live_event_type SET id=$1, created_at=$2, updated_at=$3, title=$4 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEventType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateLiveEventTypeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete live_event_type
func (s *shrikeServiceServer) DeleteLiveEventType(ctx context.Context, req *v1.DeleteLiveEventTypeRequest) (*v1.DeleteLiveEventTypeResponse, error) {
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

	// delete live_event_type
	res, err := c.ExecContext(ctx, "DELETE FROM live_event_type WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEventType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteLiveEventTypeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}