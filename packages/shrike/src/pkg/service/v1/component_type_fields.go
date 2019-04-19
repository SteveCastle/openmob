package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new ComponentTypeFields
func (s *shrikeServiceServer) CreateComponentTypeFields(ctx context.Context, req *v1.CreateComponentTypeFieldsRequest) (*v1.CreateComponentTypeFieldsResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentTypeFields Manager
	m := models.NewComponentTypeFieldsManager(s.db)

	// Get a list of componentTypeFieldss given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateComponentTypeFieldsResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get componentTypeFields by id.
func (s *shrikeServiceServer) GetComponentTypeFields(ctx context.Context, req *v1.GetComponentTypeFieldsRequest) (*v1.GetComponentTypeFieldsResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentTypeFields Manager
	m := models.NewComponentTypeFieldsManager(s.db)

	// Get a list of componentTypeFieldss given filters, ordering, and limit rules.
	componentTypeFields, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetComponentTypeFieldsResponse{
		Api:  apiVersion,
		Item: m.GetProto(componentTypeFields),
	}, nil

}

// Read all ComponentTypeFields
func (s *shrikeServiceServer) ListComponentTypeFields(ctx context.Context, req *v1.ListComponentTypeFieldsRequest) (*v1.ListComponentTypeFieldsResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ComponentTypeFields Manager
	m := models.NewComponentTypeFieldsManager(s.db)

	// Get a list of componentTypeFieldss given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListComponentTypeFieldsResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ComponentTypeFields
func (s *shrikeServiceServer) UpdateComponentTypeFields(ctx context.Context, req *v1.UpdateComponentTypeFieldsRequest) (*v1.UpdateComponentTypeFieldsResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentTypeFields Manager
	m := models.NewComponentTypeFieldsManager(s.db)

	// Get a list of componentTypeFieldss given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateComponentTypeFieldsResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete componentTypeFields
func (s *shrikeServiceServer) DeleteComponentTypeFields(ctx context.Context, req *v1.DeleteComponentTypeFieldsRequest) (*v1.DeleteComponentTypeFieldsResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ComponentTypeFields Manager
	m := models.NewComponentTypeFieldsManager(s.db)

	// Get a list of componentTypeFieldss given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteComponentTypeFieldsResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
