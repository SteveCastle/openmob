package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new LayoutColumn
func (s *shrikeServiceServer) CreateLayoutColumn(ctx context.Context, req *v1.CreateLayoutColumnRequest) (*v1.CreateLayoutColumnResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutColumn Manager
	m := models.NewLayoutColumnManager(s.db)

	// Get a list of layoutColumns given filters, ordering, and limit rules.
	id, err := m.CreateLayoutColumn(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateLayoutColumnResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get layoutColumn by id.
func (s *shrikeServiceServer) GetLayoutColumn(ctx context.Context, req *v1.GetLayoutColumnRequest) (*v1.GetLayoutColumnResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutColumn Manager
	m := models.NewLayoutColumnManager(s.db)

	// Get a list of layoutColumns given filters, ordering, and limit rules.
	layoutColumn, err := m.GetLayoutColumn(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetLayoutColumnResponse{
		Api:  apiVersion,
		Item: m.GetProto(layoutColumn),
	}, nil

}

// Read all LayoutColumn
func (s *shrikeServiceServer) ListLayoutColumn(ctx context.Context, req *v1.ListLayoutColumnRequest) (*v1.ListLayoutColumnResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a LayoutColumn Manager
	m := models.NewLayoutColumnManager(s.db)

	// Get a list of layoutColumns given filters, ordering, and limit rules.
	list, err := m.ListLayoutColumn(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListLayoutColumnResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update LayoutColumn
func (s *shrikeServiceServer) UpdateLayoutColumn(ctx context.Context, req *v1.UpdateLayoutColumnRequest) (*v1.UpdateLayoutColumnResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutColumn Manager
	m := models.NewLayoutColumnManager(s.db)

	// Get a list of layoutColumns given filters, ordering, and limit rules.
	rows, err := m.UpdateLayoutColumn(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateLayoutColumnResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete layoutColumn
func (s *shrikeServiceServer) DeleteLayoutColumn(ctx context.Context, req *v1.DeleteLayoutColumnRequest) (*v1.DeleteLayoutColumnResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutColumn Manager
	m := models.NewLayoutColumnManager(s.db)

	// Get a list of layoutColumns given filters, ordering, and limit rules.
	rows, err := m.DeleteLayoutColumn(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteLayoutColumnResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
