package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new LiveEvent
func (s *shrikeServiceServer) CreateLiveEvent(ctx context.Context, req *v1.CreateLiveEventRequest) (*v1.CreateLiveEventResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEvent Manager
	m := models.NewLiveEventManager(s.db)

	// Get a list of liveEvents given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateLiveEventResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get liveEvent by id.
func (s *shrikeServiceServer) GetLiveEvent(ctx context.Context, req *v1.GetLiveEventRequest) (*v1.GetLiveEventResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEvent Manager
	m := models.NewLiveEventManager(s.db)

	// Get a list of liveEvents given filters, ordering, and limit rules.
	liveEvent, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetLiveEventResponse{
		Api:  apiVersion,
		Item: m.GetProto(liveEvent),
	}, nil

}

// Read all LiveEvent
func (s *shrikeServiceServer) ListLiveEvent(ctx context.Context, req *v1.ListLiveEventRequest) (*v1.ListLiveEventResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a LiveEvent Manager
	m := models.NewLiveEventManager(s.db)

	// Get a list of liveEvents given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListLiveEventResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update LiveEvent
func (s *shrikeServiceServer) UpdateLiveEvent(ctx context.Context, req *v1.UpdateLiveEventRequest) (*v1.UpdateLiveEventResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEvent Manager
	m := models.NewLiveEventManager(s.db)

	// Get a list of liveEvents given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateLiveEventResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete liveEvent
func (s *shrikeServiceServer) DeleteLiveEvent(ctx context.Context, req *v1.DeleteLiveEventRequest) (*v1.DeleteLiveEventResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEvent Manager
	m := models.NewLiveEventManager(s.db)

	// Get a list of liveEvents given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteLiveEventResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
