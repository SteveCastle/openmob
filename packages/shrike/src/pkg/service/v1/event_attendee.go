package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new EventAttendee
func (s *shrikeServiceServer) CreateEventAttendee(ctx context.Context, req *v1.CreateEventAttendeeRequest) (*v1.CreateEventAttendeeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a EventAttendee Manager
	m := models.NewEventAttendeeManager(s.db)

	// Get a list of eventAttendees given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateEventAttendeeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get eventAttendee by id.
func (s *shrikeServiceServer) GetEventAttendee(ctx context.Context, req *v1.GetEventAttendeeRequest) (*v1.GetEventAttendeeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a EventAttendee Manager
	m := models.NewEventAttendeeManager(s.db)

	// Get a list of eventAttendees given filters, ordering, and limit rules.
	eventAttendee, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetEventAttendeeResponse{
		Api:  apiVersion,
		Item: m.GetProto(eventAttendee),
	}, nil

}

// Read all EventAttendee
func (s *shrikeServiceServer) ListEventAttendee(ctx context.Context, req *v1.ListEventAttendeeRequest) (*v1.ListEventAttendeeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a EventAttendee Manager
	m := models.NewEventAttendeeManager(s.db)

	// Get a list of eventAttendees given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListEventAttendeeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update EventAttendee
func (s *shrikeServiceServer) UpdateEventAttendee(ctx context.Context, req *v1.UpdateEventAttendeeRequest) (*v1.UpdateEventAttendeeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a EventAttendee Manager
	m := models.NewEventAttendeeManager(s.db)

	// Get a list of eventAttendees given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateEventAttendeeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete eventAttendee
func (s *shrikeServiceServer) DeleteEventAttendee(ctx context.Context, req *v1.DeleteEventAttendeeRequest) (*v1.DeleteEventAttendeeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a EventAttendee Manager
	m := models.NewEventAttendeeManager(s.db)

	// Get a list of eventAttendees given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteEventAttendeeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
