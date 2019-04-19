package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new PollMembership
func (s *shrikeServiceServer) CreatePollMembership(ctx context.Context, req *v1.CreatePollMembershipRequest) (*v1.CreatePollMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollMembership Manager
	m := models.NewPollMembershipManager(s.db)

	// Get a list of pollMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePollMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get pollMembership by id.
func (s *shrikeServiceServer) GetPollMembership(ctx context.Context, req *v1.GetPollMembershipRequest) (*v1.GetPollMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollMembership Manager
	m := models.NewPollMembershipManager(s.db)

	// Get a list of pollMemberships given filters, ordering, and limit rules.
	pollMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPollMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(pollMembership),
	}, nil

}

// Read all PollMembership
func (s *shrikeServiceServer) ListPollMembership(ctx context.Context, req *v1.ListPollMembershipRequest) (*v1.ListPollMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a PollMembership Manager
	m := models.NewPollMembershipManager(s.db)

	// Get a list of pollMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPollMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update PollMembership
func (s *shrikeServiceServer) UpdatePollMembership(ctx context.Context, req *v1.UpdatePollMembershipRequest) (*v1.UpdatePollMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollMembership Manager
	m := models.NewPollMembershipManager(s.db)

	// Get a list of pollMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePollMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete pollMembership
func (s *shrikeServiceServer) DeletePollMembership(ctx context.Context, req *v1.DeletePollMembershipRequest) (*v1.DeletePollMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PollMembership Manager
	m := models.NewPollMembershipManager(s.db)

	// Get a list of pollMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePollMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
