package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new BoycottMembership
func (s *shrikeServiceServer) CreateBoycottMembership(ctx context.Context, req *v1.CreateBoycottMembershipRequest) (*v1.CreateBoycottMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a BoycottMembership Manager
	m := models.NewBoycottMembershipManager(s.db)

	// Get a list of boycottMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateBoycottMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get boycottMembership by id.
func (s *shrikeServiceServer) GetBoycottMembership(ctx context.Context, req *v1.GetBoycottMembershipRequest) (*v1.GetBoycottMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a BoycottMembership Manager
	m := models.NewBoycottMembershipManager(s.db)

	// Get a list of boycottMemberships given filters, ordering, and limit rules.
	boycottMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetBoycottMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(boycottMembership),
	}, nil

}

// Read all BoycottMembership
func (s *shrikeServiceServer) ListBoycottMembership(ctx context.Context, req *v1.ListBoycottMembershipRequest) (*v1.ListBoycottMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a BoycottMembership Manager
	m := models.NewBoycottMembershipManager(s.db)

	// Get a list of boycottMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListBoycottMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update BoycottMembership
func (s *shrikeServiceServer) UpdateBoycottMembership(ctx context.Context, req *v1.UpdateBoycottMembershipRequest) (*v1.UpdateBoycottMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a BoycottMembership Manager
	m := models.NewBoycottMembershipManager(s.db)

	// Get a list of boycottMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateBoycottMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete boycottMembership
func (s *shrikeServiceServer) DeleteBoycottMembership(ctx context.Context, req *v1.DeleteBoycottMembershipRequest) (*v1.DeleteBoycottMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a BoycottMembership Manager
	m := models.NewBoycottMembershipManager(s.db)

	// Get a list of boycottMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteBoycottMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
