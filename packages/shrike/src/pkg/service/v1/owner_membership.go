package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new OwnerMembership
func (s *shrikeServiceServer) CreateOwnerMembership(ctx context.Context, req *v1.CreateOwnerMembershipRequest) (*v1.CreateOwnerMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a OwnerMembership Manager
	m := models.NewOwnerMembershipManager(s.db)

	// Get a list of ownerMemberships given filters, ordering, and limit rules.
	id, err := m.CreateOwnerMembership(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateOwnerMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get ownerMembership by id.
func (s *shrikeServiceServer) GetOwnerMembership(ctx context.Context, req *v1.GetOwnerMembershipRequest) (*v1.GetOwnerMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a OwnerMembership Manager
	m := models.NewOwnerMembershipManager(s.db)

	// Get a list of ownerMemberships given filters, ordering, and limit rules.
	ownerMembership, err := m.GetOwnerMembership(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetOwnerMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(ownerMembership),
	}, nil

}

// Read all OwnerMembership
func (s *shrikeServiceServer) ListOwnerMembership(ctx context.Context, req *v1.ListOwnerMembershipRequest) (*v1.ListOwnerMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a OwnerMembership Manager
	m := models.NewOwnerMembershipManager(s.db)

	// Get a list of ownerMemberships given filters, ordering, and limit rules.
	list, err := m.ListOwnerMembership(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListOwnerMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update OwnerMembership
func (s *shrikeServiceServer) UpdateOwnerMembership(ctx context.Context, req *v1.UpdateOwnerMembershipRequest) (*v1.UpdateOwnerMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a OwnerMembership Manager
	m := models.NewOwnerMembershipManager(s.db)

	// Get a list of ownerMemberships given filters, ordering, and limit rules.
	rows, err := m.UpdateOwnerMembership(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateOwnerMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete ownerMembership
func (s *shrikeServiceServer) DeleteOwnerMembership(ctx context.Context, req *v1.DeleteOwnerMembershipRequest) (*v1.DeleteOwnerMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a OwnerMembership Manager
	m := models.NewOwnerMembershipManager(s.db)

	// Get a list of ownerMemberships given filters, ordering, and limit rules.
	rows, err := m.DeleteOwnerMembership(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteOwnerMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
