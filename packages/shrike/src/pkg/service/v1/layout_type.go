package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new LayoutType
func (s *shrikeServiceServer) CreateLayoutType(ctx context.Context, req *v1.CreateLayoutTypeRequest) (*v1.CreateLayoutTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutType Manager
	m := models.NewLayoutTypeManager(s.db)

	// Get a list of layoutTypes given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateLayoutTypeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get layoutType by id.
func (s *shrikeServiceServer) GetLayoutType(ctx context.Context, req *v1.GetLayoutTypeRequest) (*v1.GetLayoutTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutType Manager
	m := models.NewLayoutTypeManager(s.db)

	// Get a list of layoutTypes given filters, ordering, and limit rules.
	layoutType, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetLayoutTypeResponse{
		Api:  apiVersion,
		Item: m.GetProto(layoutType),
	}, nil

}

// Read all LayoutType
func (s *shrikeServiceServer) ListLayoutType(ctx context.Context, req *v1.ListLayoutTypeRequest) (*v1.ListLayoutTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a LayoutType Manager
	m := models.NewLayoutTypeManager(s.db)

	// Get a list of layoutTypes given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListLayoutTypeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update LayoutType
func (s *shrikeServiceServer) UpdateLayoutType(ctx context.Context, req *v1.UpdateLayoutTypeRequest) (*v1.UpdateLayoutTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutType Manager
	m := models.NewLayoutTypeManager(s.db)

	// Get a list of layoutTypes given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateLayoutTypeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete layoutType
func (s *shrikeServiceServer) DeleteLayoutType(ctx context.Context, req *v1.DeleteLayoutTypeRequest) (*v1.DeleteLayoutTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutType Manager
	m := models.NewLayoutTypeManager(s.db)

	// Get a list of layoutTypes given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteLayoutTypeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
