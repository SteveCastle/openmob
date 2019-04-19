package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Component
func (s *shrikeServiceServer) CreateComponent(ctx context.Context, req *v1.CreateComponentRequest) (*v1.CreateComponentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Component Manager
	m := models.NewComponentManager(s.db)

	// Get a list of components given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateComponentResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get component by id.
func (s *shrikeServiceServer) GetComponent(ctx context.Context, req *v1.GetComponentRequest) (*v1.GetComponentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Component Manager
	m := models.NewComponentManager(s.db)

	// Get a list of components given filters, ordering, and limit rules.
	component, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetComponentResponse{
		Api:  apiVersion,
		Item: m.GetProto(component),
	}, nil

}

// Read all Component
func (s *shrikeServiceServer) ListComponent(ctx context.Context, req *v1.ListComponentRequest) (*v1.ListComponentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Component Manager
	m := models.NewComponentManager(s.db)

	// Get a list of components given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListComponentResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Component
func (s *shrikeServiceServer) UpdateComponent(ctx context.Context, req *v1.UpdateComponentRequest) (*v1.UpdateComponentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Component Manager
	m := models.NewComponentManager(s.db)

	// Get a list of components given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateComponentResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete component
func (s *shrikeServiceServer) DeleteComponent(ctx context.Context, req *v1.DeleteComponentRequest) (*v1.DeleteComponentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Component Manager
	m := models.NewComponentManager(s.db)

	// Get a list of components given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteComponentResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
