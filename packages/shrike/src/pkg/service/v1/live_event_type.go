package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new LiveEventType
func (s *shrikeServiceServer) CreateLiveEventType(ctx context.Context, req *v1.CreateLiveEventTypeRequest) (*v1.CreateLiveEventTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEventType Manager
	m := models.NewLiveEventTypeManager(s.db)

	// Get a list of liveEventTypes given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateLiveEventTypeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get liveEventType by id.
func (s *shrikeServiceServer) GetLiveEventType(ctx context.Context, req *v1.GetLiveEventTypeRequest) (*v1.GetLiveEventTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEventType Manager
	m := models.NewLiveEventTypeManager(s.db)

	// Get a list of liveEventTypes given filters, ordering, and limit rules.
	liveEventType, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetLiveEventTypeResponse{
		Api:  apiVersion,
		Item: m.GetProto(liveEventType),
	}, nil

}

// Read all LiveEventType
func (s *shrikeServiceServer) ListLiveEventType(ctx context.Context, req *v1.ListLiveEventTypeRequest) (*v1.ListLiveEventTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a LiveEventType Manager
	m := models.NewLiveEventTypeManager(s.db)

	// Get a list of liveEventTypes given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListLiveEventTypeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update LiveEventType
func (s *shrikeServiceServer) UpdateLiveEventType(ctx context.Context, req *v1.UpdateLiveEventTypeRequest) (*v1.UpdateLiveEventTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEventType Manager
	m := models.NewLiveEventTypeManager(s.db)

	// Get a list of liveEventTypes given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateLiveEventTypeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete liveEventType
func (s *shrikeServiceServer) DeleteLiveEventType(ctx context.Context, req *v1.DeleteLiveEventTypeRequest) (*v1.DeleteLiveEventTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LiveEventType Manager
	m := models.NewLiveEventTypeManager(s.db)

	// Get a list of liveEventTypes given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteLiveEventTypeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
