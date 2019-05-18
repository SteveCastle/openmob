package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Purchaser
func (s *shrikeServiceServer) CreatePurchaser(ctx context.Context, req *v1.CreatePurchaserRequest) (*v1.CreatePurchaserResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Purchaser Manager
	m := models.NewPurchaserManager(s.db)

	// Get a list of purchasers given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePurchaserResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get purchaser by id.
func (s *shrikeServiceServer) GetPurchaser(ctx context.Context, req *v1.GetPurchaserRequest) (*v1.GetPurchaserResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Purchaser Manager
	m := models.NewPurchaserManager(s.db)

	// Get a list of purchasers given filters, ordering, and limit rules.
	purchaser, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPurchaserResponse{
		Api:  apiVersion,
		Item: m.GetProto(purchaser),
	}, nil

}

// Read all Purchaser
func (s *shrikeServiceServer) ListPurchaser(ctx context.Context, req *v1.ListPurchaserRequest) (*v1.ListPurchaserResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Purchaser Manager
	m := models.NewPurchaserManager(s.db)

	// Get a list of purchasers given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPurchaserResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Purchaser
func (s *shrikeServiceServer) UpdatePurchaser(ctx context.Context, req *v1.UpdatePurchaserRequest) (*v1.UpdatePurchaserResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Purchaser Manager
	m := models.NewPurchaserManager(s.db)

	// Get a list of purchasers given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePurchaserResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete purchaser
func (s *shrikeServiceServer) DeletePurchaser(ctx context.Context, req *v1.DeletePurchaserRequest) (*v1.DeletePurchaserResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Purchaser Manager
	m := models.NewPurchaserManager(s.db)

	// Get a list of purchasers given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePurchaserResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
