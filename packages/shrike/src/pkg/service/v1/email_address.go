package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new EmailAddress
func (s *shrikeServiceServer) CreateEmailAddress(ctx context.Context, req *v1.CreateEmailAddressRequest) (*v1.CreateEmailAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a EmailAddress Manager
	m := models.NewEmailAddressManager(s.db)

	// Get a list of emailAddresss given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateEmailAddressResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get emailAddress by id.
func (s *shrikeServiceServer) GetEmailAddress(ctx context.Context, req *v1.GetEmailAddressRequest) (*v1.GetEmailAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a EmailAddress Manager
	m := models.NewEmailAddressManager(s.db)

	// Get a list of emailAddresss given filters, ordering, and limit rules.
	emailAddress, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetEmailAddressResponse{
		Api:  apiVersion,
		Item: m.GetProto(emailAddress),
	}, nil

}

// Read all EmailAddress
func (s *shrikeServiceServer) ListEmailAddress(ctx context.Context, req *v1.ListEmailAddressRequest) (*v1.ListEmailAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a EmailAddress Manager
	m := models.NewEmailAddressManager(s.db)

	// Get a list of emailAddresss given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListEmailAddressResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update EmailAddress
func (s *shrikeServiceServer) UpdateEmailAddress(ctx context.Context, req *v1.UpdateEmailAddressRequest) (*v1.UpdateEmailAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a EmailAddress Manager
	m := models.NewEmailAddressManager(s.db)

	// Get a list of emailAddresss given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateEmailAddressResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete emailAddress
func (s *shrikeServiceServer) DeleteEmailAddress(ctx context.Context, req *v1.DeleteEmailAddressRequest) (*v1.DeleteEmailAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a EmailAddress Manager
	m := models.NewEmailAddressManager(s.db)

	// Get a list of emailAddresss given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteEmailAddressResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
