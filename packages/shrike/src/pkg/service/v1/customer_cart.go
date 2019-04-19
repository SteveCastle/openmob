package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new CustomerCart
func (s *shrikeServiceServer) CreateCustomerCart(ctx context.Context, req *v1.CreateCustomerCartRequest) (*v1.CreateCustomerCartResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a CustomerCart Manager
	m := models.NewCustomerCartManager(s.db)

	// Get a list of customerCarts given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateCustomerCartResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get customerCart by id.
func (s *shrikeServiceServer) GetCustomerCart(ctx context.Context, req *v1.GetCustomerCartRequest) (*v1.GetCustomerCartResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a CustomerCart Manager
	m := models.NewCustomerCartManager(s.db)

	// Get a list of customerCarts given filters, ordering, and limit rules.
	customerCart, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetCustomerCartResponse{
		Api:  apiVersion,
		Item: m.GetProto(customerCart),
	}, nil

}

// Read all CustomerCart
func (s *shrikeServiceServer) ListCustomerCart(ctx context.Context, req *v1.ListCustomerCartRequest) (*v1.ListCustomerCartResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a CustomerCart Manager
	m := models.NewCustomerCartManager(s.db)

	// Get a list of customerCarts given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListCustomerCartResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update CustomerCart
func (s *shrikeServiceServer) UpdateCustomerCart(ctx context.Context, req *v1.UpdateCustomerCartRequest) (*v1.UpdateCustomerCartResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a CustomerCart Manager
	m := models.NewCustomerCartManager(s.db)

	// Get a list of customerCarts given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateCustomerCartResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete customerCart
func (s *shrikeServiceServer) DeleteCustomerCart(ctx context.Context, req *v1.DeleteCustomerCartRequest) (*v1.DeleteCustomerCartResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a CustomerCart Manager
	m := models.NewCustomerCartManager(s.db)

	// Get a list of customerCarts given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteCustomerCartResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
