package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Payment
func (s *shrikeServiceServer) CreatePayment(ctx context.Context, req *v1.CreatePaymentRequest) (*v1.CreatePaymentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Payment Manager
	m := models.NewPaymentManager(s.db)

	// Get a list of payments given filters, ordering, and limit rules.
	id, err := m.CreatePayment(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePaymentResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get payment by id.
func (s *shrikeServiceServer) GetPayment(ctx context.Context, req *v1.GetPaymentRequest) (*v1.GetPaymentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Payment Manager
	m := models.NewPaymentManager(s.db)

	// Get a list of payments given filters, ordering, and limit rules.
	payment, err := m.GetPayment(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPaymentResponse{
		Api:  apiVersion,
		Item: m.GetProto(payment),
	}, nil

}

// Read all Payment
func (s *shrikeServiceServer) ListPayment(ctx context.Context, req *v1.ListPaymentRequest) (*v1.ListPaymentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Payment Manager
	m := models.NewPaymentManager(s.db)

	// Get a list of payments given filters, ordering, and limit rules.
	list, err := m.ListPayment(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPaymentResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Payment
func (s *shrikeServiceServer) UpdatePayment(ctx context.Context, req *v1.UpdatePaymentRequest) (*v1.UpdatePaymentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Payment Manager
	m := models.NewPaymentManager(s.db)

	// Get a list of payments given filters, ordering, and limit rules.
	rows, err := m.UpdatePayment(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePaymentResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete payment
func (s *shrikeServiceServer) DeletePayment(ctx context.Context, req *v1.DeletePaymentRequest) (*v1.DeletePaymentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Payment Manager
	m := models.NewPaymentManager(s.db)

	// Get a list of payments given filters, ordering, and limit rules.
	rows, err := m.DeletePayment(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePaymentResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
