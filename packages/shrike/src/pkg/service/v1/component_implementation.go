package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new ComponentImplementation
func (s *shrikeServiceServer) CreateComponentImplementation(ctx context.Context, req *v1.CreateComponentImplementationRequest) (*v1.CreateComponentImplementationResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentImplementation Manager
	m := models.NewComponentImplementationManager(s.db)

	// Get a list of componentImplementations given filters, ordering, and limit rules.
	id, err := m.CreateComponentImplementation(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateComponentImplementationResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get componentImplementation by id.
func (s *shrikeServiceServer) GetComponentImplementation(ctx context.Context, req *v1.GetComponentImplementationRequest) (*v1.GetComponentImplementationResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentImplementation Manager
	m := models.NewComponentImplementationManager(s.db)

	// Get a list of componentImplementations given filters, ordering, and limit rules.
	componentImplementation, err := m.GetComponentImplementation(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetComponentImplementationResponse{
		Api:  apiVersion,
		Item: m.GetProto(componentImplementation),
	}, nil

}

// Read all ComponentImplementation
func (s *shrikeServiceServer) ListComponentImplementation(ctx context.Context, req *v1.ListComponentImplementationRequest) (*v1.ListComponentImplementationResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ComponentImplementation Manager
	m := models.NewComponentImplementationManager(s.db)

	// Get a list of componentImplementations given filters, ordering, and limit rules.
	list, err := m.ListComponentImplementation(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListComponentImplementationResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ComponentImplementation
func (s *shrikeServiceServer) UpdateComponentImplementation(ctx context.Context, req *v1.UpdateComponentImplementationRequest) (*v1.UpdateComponentImplementationResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentImplementation Manager
	m := models.NewComponentImplementationManager(s.db)

	// Get a list of componentImplementations given filters, ordering, and limit rules.
	rows, err := m.UpdateComponentImplementation(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateComponentImplementationResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete componentImplementation
func (s *shrikeServiceServer) DeleteComponentImplementation(ctx context.Context, req *v1.DeleteComponentImplementationRequest) (*v1.DeleteComponentImplementationResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentImplementation Manager
	m := models.NewComponentImplementationManager(s.db)

	// Get a list of componentImplementations given filters, ordering, and limit rules.
	rows, err := m.DeleteComponentImplementation(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteComponentImplementationResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
