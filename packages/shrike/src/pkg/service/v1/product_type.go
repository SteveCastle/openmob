package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new ProductType
func (s *shrikeServiceServer) CreateProductType(ctx context.Context, req *v1.CreateProductTypeRequest) (*v1.CreateProductTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ProductType Manager
	m := models.NewProductTypeManager(s.db)

	// Get a list of productTypes given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateProductTypeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get productType by id.
func (s *shrikeServiceServer) GetProductType(ctx context.Context, req *v1.GetProductTypeRequest) (*v1.GetProductTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ProductType Manager
	m := models.NewProductTypeManager(s.db)

	// Get a list of productTypes given filters, ordering, and limit rules.
	productType, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetProductTypeResponse{
		Api:  apiVersion,
		Item: m.GetProto(productType),
	}, nil

}

// Read all ProductType
func (s *shrikeServiceServer) ListProductType(ctx context.Context, req *v1.ListProductTypeRequest) (*v1.ListProductTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ProductType Manager
	m := models.NewProductTypeManager(s.db)

	// Get a list of productTypes given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListProductTypeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ProductType
func (s *shrikeServiceServer) UpdateProductType(ctx context.Context, req *v1.UpdateProductTypeRequest) (*v1.UpdateProductTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ProductType Manager
	m := models.NewProductTypeManager(s.db)

	// Get a list of productTypes given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateProductTypeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete productType
func (s *shrikeServiceServer) DeleteProductType(ctx context.Context, req *v1.DeleteProductTypeRequest) (*v1.DeleteProductTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ProductType Manager
	m := models.NewProductTypeManager(s.db)

	// Get a list of productTypes given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteProductTypeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
