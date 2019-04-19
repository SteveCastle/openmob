package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new ContactMembership
func (s *shrikeServiceServer) CreateContactMembership(ctx context.Context, req *v1.CreateContactMembershipRequest) (*v1.CreateContactMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ContactMembership Manager
	m := models.NewContactMembershipManager(s.db)

	// Get a list of contactMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateContactMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get contactMembership by id.
func (s *shrikeServiceServer) GetContactMembership(ctx context.Context, req *v1.GetContactMembershipRequest) (*v1.GetContactMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ContactMembership Manager
	m := models.NewContactMembershipManager(s.db)

	// Get a list of contactMemberships given filters, ordering, and limit rules.
	contactMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetContactMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(contactMembership),
	}, nil

}

// Read all ContactMembership
func (s *shrikeServiceServer) ListContactMembership(ctx context.Context, req *v1.ListContactMembershipRequest) (*v1.ListContactMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ContactMembership Manager
	m := models.NewContactMembershipManager(s.db)

	// Get a list of contactMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListContactMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ContactMembership
func (s *shrikeServiceServer) UpdateContactMembership(ctx context.Context, req *v1.UpdateContactMembershipRequest) (*v1.UpdateContactMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ContactMembership Manager
	m := models.NewContactMembershipManager(s.db)

	// Get a list of contactMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateContactMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete contactMembership
func (s *shrikeServiceServer) DeleteContactMembership(ctx context.Context, req *v1.DeleteContactMembershipRequest) (*v1.DeleteContactMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ContactMembership Manager
	m := models.NewContactMembershipManager(s.db)

	// Get a list of contactMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteContactMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
