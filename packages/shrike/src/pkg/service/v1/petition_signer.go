package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new PetitionSigner
func (s *shrikeServiceServer) CreatePetitionSigner(ctx context.Context, req *v1.CreatePetitionSignerRequest) (*v1.CreatePetitionSignerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PetitionSigner Manager
	m := models.NewPetitionSignerManager(s.db)

	// Get a list of petitionSigners given filters, ordering, and limit rules.
	id, err := m.CreatePetitionSigner(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePetitionSignerResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get petitionSigner by id.
func (s *shrikeServiceServer) GetPetitionSigner(ctx context.Context, req *v1.GetPetitionSignerRequest) (*v1.GetPetitionSignerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PetitionSigner Manager
	m := models.NewPetitionSignerManager(s.db)

	// Get a list of petitionSigners given filters, ordering, and limit rules.
	petitionSigner, err := m.GetPetitionSigner(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPetitionSignerResponse{
		Api:  apiVersion,
		Item: m.GetProto(petitionSigner),
	}, nil

}

// Read all PetitionSigner
func (s *shrikeServiceServer) ListPetitionSigner(ctx context.Context, req *v1.ListPetitionSignerRequest) (*v1.ListPetitionSignerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a PetitionSigner Manager
	m := models.NewPetitionSignerManager(s.db)

	// Get a list of petitionSigners given filters, ordering, and limit rules.
	list, err := m.ListPetitionSigner(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPetitionSignerResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update PetitionSigner
func (s *shrikeServiceServer) UpdatePetitionSigner(ctx context.Context, req *v1.UpdatePetitionSignerRequest) (*v1.UpdatePetitionSignerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PetitionSigner Manager
	m := models.NewPetitionSignerManager(s.db)

	// Get a list of petitionSigners given filters, ordering, and limit rules.
	rows, err := m.UpdatePetitionSigner(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePetitionSignerResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete petitionSigner
func (s *shrikeServiceServer) DeletePetitionSigner(ctx context.Context, req *v1.DeletePetitionSignerRequest) (*v1.DeletePetitionSignerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PetitionSigner Manager
	m := models.NewPetitionSignerManager(s.db)

	// Get a list of petitionSigners given filters, ordering, and limit rules.
	rows, err := m.DeletePetitionSigner(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePetitionSignerResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
