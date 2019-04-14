package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new ComponentType
func (s *shrikeServiceServer) CreateComponentType(ctx context.Context, req *v1.CreateComponentTypeRequest) (*v1.CreateComponentTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentType Manager
	m := models.NewComponentTypeManager(s.db)

	// Get a list of componentTypes given filters, ordering, and limit rules.
	id, err := m.CreateComponentType(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateComponentTypeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get componentType by id.
func (s *shrikeServiceServer) GetComponentType(ctx context.Context, req *v1.GetComponentTypeRequest) (*v1.GetComponentTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentType Manager
	m := models.NewComponentTypeManager(s.db)

	// Get a list of componentTypes given filters, ordering, and limit rules.
	componentType, err := m.GetComponentType(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetComponentTypeResponse{
		Api:  apiVersion,
		Item: m.GetProto(componentType),
	}, nil

}

// Read all ComponentType
func (s *shrikeServiceServer) ListComponentType(ctx context.Context, req *v1.ListComponentTypeRequest) (*v1.ListComponentTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ComponentType Manager
	m := models.NewComponentTypeManager(s.db)

	// Get a list of componentTypes given filters, ordering, and limit rules.
	list, err := m.ListComponentType(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListComponentTypeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ComponentType
func (s *shrikeServiceServer) UpdateComponentType(ctx context.Context, req *v1.UpdateComponentTypeRequest) (*v1.UpdateComponentTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentType Manager
	m := models.NewComponentTypeManager(s.db)

	// Get a list of componentTypes given filters, ordering, and limit rules.
	rows, err := m.UpdateComponentType(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateComponentTypeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete componentType
func (s *shrikeServiceServer) DeleteComponentType(ctx context.Context, req *v1.DeleteComponentTypeRequest) (*v1.DeleteComponentTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentType Manager
	m := models.NewComponentTypeManager(s.db)

	// Get a list of componentTypes given filters, ordering, and limit rules.
	rows, err := m.DeleteComponentType(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteComponentTypeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
