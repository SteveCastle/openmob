package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new ProductMembership
func (s *shrikeServiceServer) CreateProductMembership(ctx context.Context, req *v1.CreateProductMembershipRequest) (*v1.CreateProductMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ProductMembership Manager
	m := models.NewProductMembershipManager(s.db)

	// Get a list of productMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateProductMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get productMembership by id.
func (s *shrikeServiceServer) GetProductMembership(ctx context.Context, req *v1.GetProductMembershipRequest) (*v1.GetProductMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ProductMembership Manager
	m := models.NewProductMembershipManager(s.db)

	// Get a list of productMemberships given filters, ordering, and limit rules.
	productMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetProductMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(productMembership),
	}, nil

}

// Read all ProductMembership
func (s *shrikeServiceServer) ListProductMembership(ctx context.Context, req *v1.ListProductMembershipRequest) (*v1.ListProductMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ProductMembership Manager
	m := models.NewProductMembershipManager(s.db)

	// Get a list of productMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListProductMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ProductMembership
func (s *shrikeServiceServer) UpdateProductMembership(ctx context.Context, req *v1.UpdateProductMembershipRequest) (*v1.UpdateProductMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ProductMembership Manager
	m := models.NewProductMembershipManager(s.db)

	// Get a list of productMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateProductMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete productMembership
func (s *shrikeServiceServer) DeleteProductMembership(ctx context.Context, req *v1.DeleteProductMembershipRequest) (*v1.DeleteProductMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ProductMembership Manager
	m := models.NewProductMembershipManager(s.db)

	// Get a list of productMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteProductMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
