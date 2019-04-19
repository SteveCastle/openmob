package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new PollItem
func (s *shrikeServiceServer) CreatePollItem(ctx context.Context, req *v1.CreatePollItemRequest) (*v1.CreatePollItemResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollItem Manager
	m := models.NewPollItemManager(s.db)

	// Get a list of pollItems given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePollItemResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get pollItem by id.
func (s *shrikeServiceServer) GetPollItem(ctx context.Context, req *v1.GetPollItemRequest) (*v1.GetPollItemResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollItem Manager
	m := models.NewPollItemManager(s.db)

	// Get a list of pollItems given filters, ordering, and limit rules.
	pollItem, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPollItemResponse{
		Api:  apiVersion,
		Item: m.GetProto(pollItem),
	}, nil

}

// Read all PollItem
func (s *shrikeServiceServer) ListPollItem(ctx context.Context, req *v1.ListPollItemRequest) (*v1.ListPollItemResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a PollItem Manager
	m := models.NewPollItemManager(s.db)

	// Get a list of pollItems given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPollItemResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update PollItem
func (s *shrikeServiceServer) UpdatePollItem(ctx context.Context, req *v1.UpdatePollItemRequest) (*v1.UpdatePollItemResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollItem Manager
	m := models.NewPollItemManager(s.db)

	// Get a list of pollItems given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePollItemResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete pollItem
func (s *shrikeServiceServer) DeletePollItem(ctx context.Context, req *v1.DeletePollItemRequest) (*v1.DeletePollItemResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollItem Manager
	m := models.NewPollItemManager(s.db)

	// Get a list of pollItems given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePollItemResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
