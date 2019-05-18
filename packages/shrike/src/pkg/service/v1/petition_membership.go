package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new PetitionMembership
func (s *shrikeServiceServer) CreatePetitionMembership(ctx context.Context, req *v1.CreatePetitionMembershipRequest) (*v1.CreatePetitionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PetitionMembership Manager
	m := models.NewPetitionMembershipManager(s.db)

	// Get a list of petitionMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePetitionMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get petitionMembership by id.
func (s *shrikeServiceServer) GetPetitionMembership(ctx context.Context, req *v1.GetPetitionMembershipRequest) (*v1.GetPetitionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PetitionMembership Manager
	m := models.NewPetitionMembershipManager(s.db)

	// Get a list of petitionMemberships given filters, ordering, and limit rules.
	petitionMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPetitionMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(petitionMembership),
	}, nil

}

// Read all PetitionMembership
func (s *shrikeServiceServer) ListPetitionMembership(ctx context.Context, req *v1.ListPetitionMembershipRequest) (*v1.ListPetitionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a PetitionMembership Manager
	m := models.NewPetitionMembershipManager(s.db)

	// Get a list of petitionMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPetitionMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update PetitionMembership
func (s *shrikeServiceServer) UpdatePetitionMembership(ctx context.Context, req *v1.UpdatePetitionMembershipRequest) (*v1.UpdatePetitionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PetitionMembership Manager
	m := models.NewPetitionMembershipManager(s.db)

	// Get a list of petitionMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePetitionMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete petitionMembership
func (s *shrikeServiceServer) DeletePetitionMembership(ctx context.Context, req *v1.DeletePetitionMembershipRequest) (*v1.DeletePetitionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PetitionMembership Manager
	m := models.NewPetitionMembershipManager(s.db)

	// Get a list of petitionMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePetitionMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
