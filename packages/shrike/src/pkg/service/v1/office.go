package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Office
func (s *shrikeServiceServer) CreateOffice(ctx context.Context, req *v1.CreateOfficeRequest) (*v1.CreateOfficeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Office Manager
	m := models.NewOfficeManager(s.db)

	// Get a list of offices given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateOfficeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get office by id.
func (s *shrikeServiceServer) GetOffice(ctx context.Context, req *v1.GetOfficeRequest) (*v1.GetOfficeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Office Manager
	m := models.NewOfficeManager(s.db)

	// Get a list of offices given filters, ordering, and limit rules.
	office, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetOfficeResponse{
		Api:  apiVersion,
		Item: m.GetProto(office),
	}, nil

}

// Read all Office
func (s *shrikeServiceServer) ListOffice(ctx context.Context, req *v1.ListOfficeRequest) (*v1.ListOfficeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Office Manager
	m := models.NewOfficeManager(s.db)

	// Get a list of offices given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListOfficeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Office
func (s *shrikeServiceServer) UpdateOffice(ctx context.Context, req *v1.UpdateOfficeRequest) (*v1.UpdateOfficeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Office Manager
	m := models.NewOfficeManager(s.db)

	// Get a list of offices given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateOfficeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete office
func (s *shrikeServiceServer) DeleteOffice(ctx context.Context, req *v1.DeleteOfficeRequest) (*v1.DeleteOfficeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Office Manager
	m := models.NewOfficeManager(s.db)

	// Get a list of offices given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteOfficeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
