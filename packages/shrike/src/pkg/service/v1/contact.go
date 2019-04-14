package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Contact
func (s *shrikeServiceServer) CreateContact(ctx context.Context, req *v1.CreateContactRequest) (*v1.CreateContactResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Contact Manager
	m := models.NewContactManager(s.db)

	// Get a list of contacts given filters, ordering, and limit rules.
	id, err := m.CreateContact(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateContactResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get contact by id.
func (s *shrikeServiceServer) GetContact(ctx context.Context, req *v1.GetContactRequest) (*v1.GetContactResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Contact Manager
	m := models.NewContactManager(s.db)

	// Get a list of contacts given filters, ordering, and limit rules.
	contact, err := m.GetContact(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetContactResponse{
		Api:  apiVersion,
		Item: m.GetProto(contact),
	}, nil

}

// Read all Contact
func (s *shrikeServiceServer) ListContact(ctx context.Context, req *v1.ListContactRequest) (*v1.ListContactResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Contact Manager
	m := models.NewContactManager(s.db)

	// Get a list of contacts given filters, ordering, and limit rules.
	list, err := m.ListContact(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListContactResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Contact
func (s *shrikeServiceServer) UpdateContact(ctx context.Context, req *v1.UpdateContactRequest) (*v1.UpdateContactResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Contact Manager
	m := models.NewContactManager(s.db)

	// Get a list of contacts given filters, ordering, and limit rules.
	rows, err := m.UpdateContact(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateContactResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete contact
func (s *shrikeServiceServer) DeleteContact(ctx context.Context, req *v1.DeleteContactRequest) (*v1.DeleteContactResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Contact Manager
	m := models.NewContactManager(s.db)

	// Get a list of contacts given filters, ordering, and limit rules.
	rows, err := m.DeleteContact(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteContactResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
