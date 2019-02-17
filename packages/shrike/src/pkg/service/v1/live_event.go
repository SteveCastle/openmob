package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


// Create new LiveEvent
func (s *shrikeServiceServer) CreateLiveEvent(ctx context.Context, req *v1.CreateLiveEventRequest) (*v1.CreateLiveEventResponse, error) {
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
	// insert LiveEvent entity data
	err = c.QueryRowContext(ctx, "INSERT INTO live_event (id, created_at, updated_at, title, live_event_type, ) VALUES($1, $2, $3, $4, $5, )  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Title,  req.Item.LiveEventType, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEvent-> "+err.Error())
	}

	// get ID of creates LiveEvent
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEvent-> "+err.Error())
	}

	return &v1.CreateLiveEventResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get live_event by id.
func (s *shrikeServiceServer) GetLiveEvent(ctx context.Context, req *v1.GetLiveEventRequest) (*v1.GetLiveEventResponse, error) {
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

	// query LiveEvent by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, live_event_type,  FROM live_event WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEvent-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEvent-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%d' is not found",
			req.ID))
	}

	// get LiveEvent data
	var liveevent v1.LiveEvent
	if err := rows.Scan( &liveevent.ID,  &liveevent.CreatedAt,  &liveevent.UpdatedAt,  &liveevent.Title,  &liveevent.LiveEventType, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEvent row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEvent rows with ID='%d'",
			req.ID))
	}

	return &v1.GetLiveEventResponse{
		Api:  apiVersion,
		Item: &liveevent,
	}, nil

}

// Read all LiveEvent
func (s *shrikeServiceServer) ListLiveEvent(ctx context.Context, req *v1.ListLiveEventRequest) (*v1.ListLiveEventResponse, error) {
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

	// get LiveEvent list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, live_event_type,  FROM live_event")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEvent-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.LiveEvent{}
	for rows.Next() {
		liveevent := new(v1.LiveEvent)
		if err := rows.Scan( &liveevent.ID,  &liveevent.CreatedAt,  &liveevent.UpdatedAt,  &liveevent.Title,  &liveevent.LiveEventType, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEvent row-> "+err.Error())
		}
		list = append(list, liveevent)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEvent-> "+err.Error())
	}

	return &v1.ListLiveEventResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update LiveEvent
func (s *shrikeServiceServer) UpdateLiveEvent(ctx context.Context, req *v1.UpdateLiveEventRequest) (*v1.UpdateLiveEventResponse, error) {
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

	// update live_event
	res, err := c.ExecContext(ctx, "UPDATE live_event SET $1, $2, $3, $4, $5,  WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Title,req.Item.LiveEventType, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEvent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateLiveEventResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete live_event
func (s *shrikeServiceServer) DeleteLiveEvent(ctx context.Context, req *v1.DeleteLiveEventRequest) (*v1.DeleteLiveEventResponse, error) {
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

	// delete live_event
	res, err := c.ExecContext(ctx, "DELETE FROM live_event WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEvent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteLiveEventResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
