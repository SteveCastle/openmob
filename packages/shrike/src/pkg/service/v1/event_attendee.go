package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new EventAttendee
func (s *shrikeServiceServer) CreateEventAttendee(ctx context.Context, req *v1.CreateEventAttendeeRequest) (*v1.CreateEventAttendeeResponse, error) {
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
	// insert EventAttendee entity data
	err = c.QueryRowContext(ctx, "INSERT INTO event_attendee (live_event, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		req.Item.LiveEvent, req.Item.Contact, req.Item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into EventAttendee-> "+err.Error())
	}

	// get ID of creates EventAttendee
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created EventAttendee-> "+err.Error())
	}

	return &v1.CreateEventAttendeeResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get event_attendee by id.
func (s *shrikeServiceServer) GetEventAttendee(ctx context.Context, req *v1.GetEventAttendeeRequest) (*v1.GetEventAttendeeResponse, error) {
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

	// query EventAttendee by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, live_event, contact, cause FROM event_attendee WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EventAttendee-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from EventAttendee-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%d' is not found",
			req.ID))
	}

	// scan EventAttendee data into protobuf model
	var eventattendee v1.EventAttendee
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&eventattendee.ID, &createdAt, &updatedAt, &eventattendee.LiveEvent, &eventattendee.Contact, &eventattendee.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from EventAttendee row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	eventattendee.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	eventattendee.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple EventAttendee rows with ID='%d'",
			req.ID))
	}

	return &v1.GetEventAttendeeResponse{
		Api:  apiVersion,
		Item: &eventattendee,
	}, nil

}

// Read all EventAttendee
func (s *shrikeServiceServer) ListEventAttendee(ctx context.Context, req *v1.ListEventAttendeeRequest) (*v1.ListEventAttendeeResponse, error) {
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

	// get EventAttendee list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, live_event, contact, cause FROM event_attendee")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EventAttendee-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.EventAttendee{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		eventattendee := new(v1.EventAttendee)
		if err := rows.Scan(&eventattendee.ID, &createdAt, &updatedAt, &eventattendee.LiveEvent, &eventattendee.Contact, &eventattendee.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from EventAttendee row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		eventattendee.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		eventattendee.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, eventattendee)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from EventAttendee-> "+err.Error())
	}

	return &v1.ListEventAttendeeResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update EventAttendee
func (s *shrikeServiceServer) UpdateEventAttendee(ctx context.Context, req *v1.UpdateEventAttendeeRequest) (*v1.UpdateEventAttendeeResponse, error) {
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

	// update event_attendee
	res, err := c.ExecContext(ctx, "UPDATE event_attendee SET live_event=$2, contact=$3, cause=$4 WHERE id=$1",
		req.Item.ID, req.Item.LiveEvent, req.Item.Contact, req.Item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update EventAttendee-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateEventAttendeeResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete event_attendee
func (s *shrikeServiceServer) DeleteEventAttendee(ctx context.Context, req *v1.DeleteEventAttendeeRequest) (*v1.DeleteEventAttendeeResponse, error) {
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

	// delete event_attendee
	res, err := c.ExecContext(ctx, "DELETE FROM event_attendee WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete EventAttendee-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteEventAttendeeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
