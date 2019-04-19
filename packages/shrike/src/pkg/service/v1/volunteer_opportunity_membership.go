package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new VolunteerOpportunityMembership
func (s *shrikeServiceServer) CreateVolunteerOpportunityMembership(ctx context.Context, req *v1.CreateVolunteerOpportunityMembershipRequest) (*v1.CreateVolunteerOpportunityMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunityMembership Manager
	m := models.NewVolunteerOpportunityMembershipManager(s.db)

	// Get a list of volunteerOpportunityMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateVolunteerOpportunityMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get volunteerOpportunityMembership by id.
func (s *shrikeServiceServer) GetVolunteerOpportunityMembership(ctx context.Context, req *v1.GetVolunteerOpportunityMembershipRequest) (*v1.GetVolunteerOpportunityMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunityMembership Manager
	m := models.NewVolunteerOpportunityMembershipManager(s.db)

	// Get a list of volunteerOpportunityMemberships given filters, ordering, and limit rules.
	volunteerOpportunityMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetVolunteerOpportunityMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(volunteerOpportunityMembership),
	}, nil

}

// Read all VolunteerOpportunityMembership
func (s *shrikeServiceServer) ListVolunteerOpportunityMembership(ctx context.Context, req *v1.ListVolunteerOpportunityMembershipRequest) (*v1.ListVolunteerOpportunityMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a VolunteerOpportunityMembership Manager
	m := models.NewVolunteerOpportunityMembershipManager(s.db)

	// Get a list of volunteerOpportunityMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListVolunteerOpportunityMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update VolunteerOpportunityMembership
func (s *shrikeServiceServer) UpdateVolunteerOpportunityMembership(ctx context.Context, req *v1.UpdateVolunteerOpportunityMembershipRequest) (*v1.UpdateVolunteerOpportunityMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunityMembership Manager
	m := models.NewVolunteerOpportunityMembershipManager(s.db)

	// Get a list of volunteerOpportunityMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateVolunteerOpportunityMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete volunteerOpportunityMembership
func (s *shrikeServiceServer) DeleteVolunteerOpportunityMembership(ctx context.Context, req *v1.DeleteVolunteerOpportunityMembershipRequest) (*v1.DeleteVolunteerOpportunityMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunityMembership Manager
	m := models.NewVolunteerOpportunityMembershipManager(s.db)

	// Get a list of volunteerOpportunityMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteVolunteerOpportunityMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
