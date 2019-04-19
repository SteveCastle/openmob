package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Delivery
func (s *shrikeServiceServer) CreateDelivery(ctx context.Context, req *v1.CreateDeliveryRequest) (*v1.CreateDeliveryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Delivery Manager
	m := models.NewDeliveryManager(s.db)

	// Get a list of deliverys given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDeliveryResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get delivery by id.
func (s *shrikeServiceServer) GetDelivery(ctx context.Context, req *v1.GetDeliveryRequest) (*v1.GetDeliveryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Delivery Manager
	m := models.NewDeliveryManager(s.db)

	// Get a list of deliverys given filters, ordering, and limit rules.
	delivery, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetDeliveryResponse{
		Api:  apiVersion,
		Item: m.GetProto(delivery),
	}, nil

}

// Read all Delivery
func (s *shrikeServiceServer) ListDelivery(ctx context.Context, req *v1.ListDeliveryRequest) (*v1.ListDeliveryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Delivery Manager
	m := models.NewDeliveryManager(s.db)

	// Get a list of deliverys given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListDeliveryResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Delivery
func (s *shrikeServiceServer) UpdateDelivery(ctx context.Context, req *v1.UpdateDeliveryRequest) (*v1.UpdateDeliveryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Delivery Manager
	m := models.NewDeliveryManager(s.db)

	// Get a list of deliverys given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDeliveryResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete delivery
func (s *shrikeServiceServer) DeleteDelivery(ctx context.Context, req *v1.DeleteDeliveryRequest) (*v1.DeleteDeliveryResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Delivery Manager
	m := models.NewDeliveryManager(s.db)

	// Get a list of deliverys given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteDeliveryResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
