package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new CustomerOrder
func (s *shrikeServiceServer) CreateCustomerOrder(ctx context.Context, req *v1.CreateCustomerOrderRequest) (*v1.CreateCustomerOrderResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a CustomerOrder Manager
	m := models.NewCustomerOrderManager(s.db)

	// Get a list of customerOrders given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateCustomerOrderResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get customerOrder by id.
func (s *shrikeServiceServer) GetCustomerOrder(ctx context.Context, req *v1.GetCustomerOrderRequest) (*v1.GetCustomerOrderResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a CustomerOrder Manager
	m := models.NewCustomerOrderManager(s.db)

	// Get a list of customerOrders given filters, ordering, and limit rules.
	customerOrder, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetCustomerOrderResponse{
		Api:  apiVersion,
		Item: m.GetProto(customerOrder),
	}, nil

}

// Read all CustomerOrder
func (s *shrikeServiceServer) ListCustomerOrder(ctx context.Context, req *v1.ListCustomerOrderRequest) (*v1.ListCustomerOrderResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a CustomerOrder Manager
	m := models.NewCustomerOrderManager(s.db)

	// Get a list of customerOrders given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListCustomerOrderResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update CustomerOrder
func (s *shrikeServiceServer) UpdateCustomerOrder(ctx context.Context, req *v1.UpdateCustomerOrderRequest) (*v1.UpdateCustomerOrderResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a CustomerOrder Manager
	m := models.NewCustomerOrderManager(s.db)

	// Get a list of customerOrders given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateCustomerOrderResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete customerOrder
func (s *shrikeServiceServer) DeleteCustomerOrder(ctx context.Context, req *v1.DeleteCustomerOrderRequest) (*v1.DeleteCustomerOrderResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a CustomerOrder Manager
	m := models.NewCustomerOrderManager(s.db)

	// Get a list of customerOrders given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteCustomerOrderResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
