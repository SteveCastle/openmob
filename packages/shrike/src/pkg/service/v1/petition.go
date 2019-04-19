package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Petition
func (s *shrikeServiceServer) CreatePetition(ctx context.Context, req *v1.CreatePetitionRequest) (*v1.CreatePetitionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Petition Manager
	m := models.NewPetitionManager(s.db)

	// Get a list of petitions given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePetitionResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get petition by id.
func (s *shrikeServiceServer) GetPetition(ctx context.Context, req *v1.GetPetitionRequest) (*v1.GetPetitionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Petition Manager
	m := models.NewPetitionManager(s.db)

	// Get a list of petitions given filters, ordering, and limit rules.
	petition, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPetitionResponse{
		Api:  apiVersion,
		Item: m.GetProto(petition),
	}, nil

}

// Read all Petition
func (s *shrikeServiceServer) ListPetition(ctx context.Context, req *v1.ListPetitionRequest) (*v1.ListPetitionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Petition Manager
	m := models.NewPetitionManager(s.db)

	// Get a list of petitions given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPetitionResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Petition
func (s *shrikeServiceServer) UpdatePetition(ctx context.Context, req *v1.UpdatePetitionRequest) (*v1.UpdatePetitionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Petition Manager
	m := models.NewPetitionManager(s.db)

	// Get a list of petitions given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePetitionResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete petition
func (s *shrikeServiceServer) DeletePetition(ctx context.Context, req *v1.DeletePetitionRequest) (*v1.DeletePetitionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Petition Manager
	m := models.NewPetitionManager(s.db)

	// Get a list of petitions given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePetitionResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
