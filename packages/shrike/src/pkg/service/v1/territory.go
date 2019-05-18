package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Territory
func (s *shrikeServiceServer) CreateTerritory(ctx context.Context, req *v1.CreateTerritoryRequest) (*v1.CreateTerritoryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Territory Manager
	m := models.NewTerritoryManager(s.db)

	// Get a list of territorys given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateTerritoryResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get territory by id.
func (s *shrikeServiceServer) GetTerritory(ctx context.Context, req *v1.GetTerritoryRequest) (*v1.GetTerritoryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Territory Manager
	m := models.NewTerritoryManager(s.db)

	// Get a list of territorys given filters, ordering, and limit rules.
	territory, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetTerritoryResponse{
		Api:  apiVersion,
		Item: m.GetProto(territory),
	}, nil

}

// Read all Territory
func (s *shrikeServiceServer) ListTerritory(ctx context.Context, req *v1.ListTerritoryRequest) (*v1.ListTerritoryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Territory Manager
	m := models.NewTerritoryManager(s.db)

	// Get a list of territorys given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListTerritoryResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Territory
func (s *shrikeServiceServer) UpdateTerritory(ctx context.Context, req *v1.UpdateTerritoryRequest) (*v1.UpdateTerritoryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Territory Manager
	m := models.NewTerritoryManager(s.db)

	// Get a list of territorys given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateTerritoryResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete territory
func (s *shrikeServiceServer) DeleteTerritory(ctx context.Context, req *v1.DeleteTerritoryRequest) (*v1.DeleteTerritoryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Territory Manager
	m := models.NewTerritoryManager(s.db)

	// Get a list of territorys given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteTerritoryResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
