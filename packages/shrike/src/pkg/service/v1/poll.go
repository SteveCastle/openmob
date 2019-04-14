package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Poll
func (s *shrikeServiceServer) CreatePoll(ctx context.Context, req *v1.CreatePollRequest) (*v1.CreatePollResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Poll Manager
	m := models.NewPollManager(s.db)

	// Get a list of polls given filters, ordering, and limit rules.
	id, err := m.CreatePoll(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePollResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get poll by id.
func (s *shrikeServiceServer) GetPoll(ctx context.Context, req *v1.GetPollRequest) (*v1.GetPollResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Poll Manager
	m := models.NewPollManager(s.db)

	// Get a list of polls given filters, ordering, and limit rules.
	poll, err := m.GetPoll(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPollResponse{
		Api:  apiVersion,
		Item: m.GetProto(poll),
	}, nil

}

// Read all Poll
func (s *shrikeServiceServer) ListPoll(ctx context.Context, req *v1.ListPollRequest) (*v1.ListPollResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Poll Manager
	m := models.NewPollManager(s.db)

	// Get a list of polls given filters, ordering, and limit rules.
	list, err := m.ListPoll(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPollResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Poll
func (s *shrikeServiceServer) UpdatePoll(ctx context.Context, req *v1.UpdatePollRequest) (*v1.UpdatePollResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Poll Manager
	m := models.NewPollManager(s.db)

	// Get a list of polls given filters, ordering, and limit rules.
	rows, err := m.UpdatePoll(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePollResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete poll
func (s *shrikeServiceServer) DeletePoll(ctx context.Context, req *v1.DeletePollRequest) (*v1.DeletePollResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Poll Manager
	m := models.NewPollManager(s.db)

	// Get a list of polls given filters, ordering, and limit rules.
	rows, err := m.DeletePoll(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePollResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
