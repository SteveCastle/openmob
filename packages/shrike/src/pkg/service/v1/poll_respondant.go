package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new PollRespondant
func (s *shrikeServiceServer) CreatePollRespondant(ctx context.Context, req *v1.CreatePollRespondantRequest) (*v1.CreatePollRespondantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollRespondant Manager
	m := models.NewPollRespondantManager(s.db)

	// Get a list of pollRespondants given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePollRespondantResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get pollRespondant by id.
func (s *shrikeServiceServer) GetPollRespondant(ctx context.Context, req *v1.GetPollRespondantRequest) (*v1.GetPollRespondantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollRespondant Manager
	m := models.NewPollRespondantManager(s.db)

	// Get a list of pollRespondants given filters, ordering, and limit rules.
	pollRespondant, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPollRespondantResponse{
		Api:  apiVersion,
		Item: m.GetProto(pollRespondant),
	}, nil

}

// Read all PollRespondant
func (s *shrikeServiceServer) ListPollRespondant(ctx context.Context, req *v1.ListPollRespondantRequest) (*v1.ListPollRespondantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a PollRespondant Manager
	m := models.NewPollRespondantManager(s.db)

	// Get a list of pollRespondants given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPollRespondantResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update PollRespondant
func (s *shrikeServiceServer) UpdatePollRespondant(ctx context.Context, req *v1.UpdatePollRespondantRequest) (*v1.UpdatePollRespondantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollRespondant Manager
	m := models.NewPollRespondantManager(s.db)

	// Get a list of pollRespondants given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePollRespondantResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete pollRespondant
func (s *shrikeServiceServer) DeletePollRespondant(ctx context.Context, req *v1.DeletePollRespondantRequest) (*v1.DeletePollRespondantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollRespondant Manager
	m := models.NewPollRespondantManager(s.db)

	// Get a list of pollRespondants given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePollRespondantResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
