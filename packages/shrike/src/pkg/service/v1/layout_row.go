package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new LayoutRow
func (s *shrikeServiceServer) CreateLayoutRow(ctx context.Context, req *v1.CreateLayoutRowRequest) (*v1.CreateLayoutRowResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutRow Manager
	m := models.NewLayoutRowManager(s.db)

	// Get a list of layoutRows given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateLayoutRowResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get layoutRow by id.
func (s *shrikeServiceServer) GetLayoutRow(ctx context.Context, req *v1.GetLayoutRowRequest) (*v1.GetLayoutRowResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutRow Manager
	m := models.NewLayoutRowManager(s.db)

	// Get a list of layoutRows given filters, ordering, and limit rules.
	layoutRow, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetLayoutRowResponse{
		Api:  apiVersion,
		Item: m.GetProto(layoutRow),
	}, nil

}

// Read all LayoutRow
func (s *shrikeServiceServer) ListLayoutRow(ctx context.Context, req *v1.ListLayoutRowRequest) (*v1.ListLayoutRowResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a LayoutRow Manager
	m := models.NewLayoutRowManager(s.db)

	// Get a list of layoutRows given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListLayoutRowResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update LayoutRow
func (s *shrikeServiceServer) UpdateLayoutRow(ctx context.Context, req *v1.UpdateLayoutRowRequest) (*v1.UpdateLayoutRowResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutRow Manager
	m := models.NewLayoutRowManager(s.db)

	// Get a list of layoutRows given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateLayoutRowResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete layoutRow
func (s *shrikeServiceServer) DeleteLayoutRow(ctx context.Context, req *v1.DeleteLayoutRowRequest) (*v1.DeleteLayoutRowResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LayoutRow Manager
	m := models.NewLayoutRowManager(s.db)

	// Get a list of layoutRows given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteLayoutRowResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
