package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Layout
func (s *shrikeServiceServer) CreateLayout(ctx context.Context, req *v1.CreateLayoutRequest) (*v1.CreateLayoutResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Layout Manager
	m := models.NewLayoutManager(s.db)

	// Get a list of layouts given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateLayoutResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get layout by id.
func (s *shrikeServiceServer) GetLayout(ctx context.Context, req *v1.GetLayoutRequest) (*v1.GetLayoutResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Layout Manager
	m := models.NewLayoutManager(s.db)

	// Get a list of layouts given filters, ordering, and limit rules.
	layout, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetLayoutResponse{
		Api:  apiVersion,
		Item: m.GetProto(layout),
	}, nil

}

// Read all Layout
func (s *shrikeServiceServer) ListLayout(ctx context.Context, req *v1.ListLayoutRequest) (*v1.ListLayoutResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Layout Manager
	m := models.NewLayoutManager(s.db)

	// Get a list of layouts given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListLayoutResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Layout
func (s *shrikeServiceServer) UpdateLayout(ctx context.Context, req *v1.UpdateLayoutRequest) (*v1.UpdateLayoutResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Layout Manager
	m := models.NewLayoutManager(s.db)

	// Get a list of layouts given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateLayoutResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete layout
func (s *shrikeServiceServer) DeleteLayout(ctx context.Context, req *v1.DeleteLayoutRequest) (*v1.DeleteLayoutResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Layout Manager
	m := models.NewLayoutManager(s.db)

	// Get a list of layouts given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteLayoutResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
