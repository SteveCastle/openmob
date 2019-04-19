package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new PhoneNumber
func (s *shrikeServiceServer) CreatePhoneNumber(ctx context.Context, req *v1.CreatePhoneNumberRequest) (*v1.CreatePhoneNumberResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PhoneNumber Manager
	m := models.NewPhoneNumberManager(s.db)

	// Get a list of phoneNumbers given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePhoneNumberResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get phoneNumber by id.
func (s *shrikeServiceServer) GetPhoneNumber(ctx context.Context, req *v1.GetPhoneNumberRequest) (*v1.GetPhoneNumberResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PhoneNumber Manager
	m := models.NewPhoneNumberManager(s.db)

	// Get a list of phoneNumbers given filters, ordering, and limit rules.
	phoneNumber, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetPhoneNumberResponse{
		Api:  apiVersion,
		Item: m.GetProto(phoneNumber),
	}, nil

}

// Read all PhoneNumber
func (s *shrikeServiceServer) ListPhoneNumber(ctx context.Context, req *v1.ListPhoneNumberRequest) (*v1.ListPhoneNumberResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a PhoneNumber Manager
	m := models.NewPhoneNumberManager(s.db)

	// Get a list of phoneNumbers given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListPhoneNumberResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update PhoneNumber
func (s *shrikeServiceServer) UpdatePhoneNumber(ctx context.Context, req *v1.UpdatePhoneNumberRequest) (*v1.UpdatePhoneNumberResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PhoneNumber Manager
	m := models.NewPhoneNumberManager(s.db)

	// Get a list of phoneNumbers given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePhoneNumberResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete phoneNumber
func (s *shrikeServiceServer) DeletePhoneNumber(ctx context.Context, req *v1.DeletePhoneNumberRequest) (*v1.DeletePhoneNumberResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a PhoneNumber Manager
	m := models.NewPhoneNumberManager(s.db)

	// Get a list of phoneNumbers given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePhoneNumberResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
