package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Product
func (s *shrikeServiceServer) CreateProduct(ctx context.Context, req *v1.CreateProductRequest) (*v1.CreateProductResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Product Manager
	m := models.NewProductManager(s.db)

	// Get a list of products given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateProductResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get product by id.
func (s *shrikeServiceServer) GetProduct(ctx context.Context, req *v1.GetProductRequest) (*v1.GetProductResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Product Manager
	m := models.NewProductManager(s.db)

	// Get a list of products given filters, ordering, and limit rules.
	product, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetProductResponse{
		Api:  apiVersion,
		Item: m.GetProto(product),
	}, nil

}

// Read all Product
func (s *shrikeServiceServer) ListProduct(ctx context.Context, req *v1.ListProductRequest) (*v1.ListProductResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Product Manager
	m := models.NewProductManager(s.db)

	// Get a list of products given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListProductResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Product
func (s *shrikeServiceServer) UpdateProduct(ctx context.Context, req *v1.UpdateProductRequest) (*v1.UpdateProductResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Product Manager
	m := models.NewProductManager(s.db)

	// Get a list of products given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateProductResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete product
func (s *shrikeServiceServer) DeleteProduct(ctx context.Context, req *v1.DeleteProductRequest) (*v1.DeleteProductResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Product Manager
	m := models.NewProductManager(s.db)

	// Get a list of products given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteProductResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
