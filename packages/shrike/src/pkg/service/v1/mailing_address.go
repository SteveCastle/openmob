package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new MailingAddress
func (s *shrikeServiceServer) CreateMailingAddress(ctx context.Context, req *v1.CreateMailingAddressRequest) (*v1.CreateMailingAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a MailingAddress Manager
	m := models.NewMailingAddressManager(s.db)

	// Get a list of mailingAddresss given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateMailingAddressResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get mailingAddress by id.
func (s *shrikeServiceServer) GetMailingAddress(ctx context.Context, req *v1.GetMailingAddressRequest) (*v1.GetMailingAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a MailingAddress Manager
	m := models.NewMailingAddressManager(s.db)

	// Get a list of mailingAddresss given filters, ordering, and limit rules.
	mailingAddress, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetMailingAddressResponse{
		Api:  apiVersion,
		Item: m.GetProto(mailingAddress),
	}, nil

}

// Read all MailingAddress
func (s *shrikeServiceServer) ListMailingAddress(ctx context.Context, req *v1.ListMailingAddressRequest) (*v1.ListMailingAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a MailingAddress Manager
	m := models.NewMailingAddressManager(s.db)

	// Get a list of mailingAddresss given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListMailingAddressResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update MailingAddress
func (s *shrikeServiceServer) UpdateMailingAddress(ctx context.Context, req *v1.UpdateMailingAddressRequest) (*v1.UpdateMailingAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a MailingAddress Manager
	m := models.NewMailingAddressManager(s.db)

	// Get a list of mailingAddresss given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateMailingAddressResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete mailingAddress
func (s *shrikeServiceServer) DeleteMailingAddress(ctx context.Context, req *v1.DeleteMailingAddressRequest) (*v1.DeleteMailingAddressResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a MailingAddress Manager
	m := models.NewMailingAddressManager(s.db)

	// Get a list of mailingAddresss given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteMailingAddressResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
