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

// NewShrikeServiceServer creates EventAttendee service
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
	err = c.QueryRowContext(ctx, "INSERT INTO event_attendee ( id  created_at  updated_at  live_event  contact  cause ) VALUES( $1 $2 $3 $4 $5 $6)  RETURNING id;",
		 req.ItemID  req.ItemCreatedAt  req.ItemUpdatedAt  req.ItemLiveEvent  req.ItemContact  req.ItemCause ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into EventAttendee-> "+err.Error())
	}

	// get ID of creates EventAttendee
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created EventAttendee-> "+err.Error())
	}

	return &v1.CreateEventAttendeeResponse{
		Api: apiVersion,
		Id:  id,
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM event_attendee WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EventAttendee-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from EventAttendee-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%d' is not found",
			req.Id))
	}

	// get EventAttendee data
	var eventattendee v1.EventAttendee
	if err := rows.Scan(&eventattendee.Id, &eventattendee.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from EventAttendee row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple EventAttendee rows with ID='%d'",
			req.Id))
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM event_attendee")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EventAttendee-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.EventAttendee{}
	for rows.Next() {
		eventattendee := new(v1.EventAttendee)
		if err := rows.Scan(&eventattendee.Id, &eventattendee.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from EventAttendee row-> "+err.Error())
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
	res, err := c.ExecContext(ctx, "UPDATE event_attendee SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update EventAttendee-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%d' is not found",
			req.Item.Id))
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
	res, err := c.ExecContext(ctx, "DELETE FROM event_attendee WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete EventAttendee-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteEventAttendeeResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
