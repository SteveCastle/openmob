package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new VolunteerOpportunity
func (s *shrikeServiceServer) CreateVolunteerOpportunity(ctx context.Context, req *v1.CreateVolunteerOpportunityRequest) (*v1.CreateVolunteerOpportunityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunity Manager
	m := models.NewVolunteerOpportunityManager(s.db)

	// Get a list of volunteerOpportunitys given filters, ordering, and limit rules.
	id, err := m.CreateVolunteerOpportunity(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateVolunteerOpportunityResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get volunteerOpportunity by id.
func (s *shrikeServiceServer) GetVolunteerOpportunity(ctx context.Context, req *v1.GetVolunteerOpportunityRequest) (*v1.GetVolunteerOpportunityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunity Manager
	m := models.NewVolunteerOpportunityManager(s.db)

	// Get a list of volunteerOpportunitys given filters, ordering, and limit rules.
	volunteerOpportunity, err := m.GetVolunteerOpportunity(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetVolunteerOpportunityResponse{
		Api:  apiVersion,
		Item: m.GetProto(volunteerOpportunity),
	}, nil

}

// Read all VolunteerOpportunity
func (s *shrikeServiceServer) ListVolunteerOpportunity(ctx context.Context, req *v1.ListVolunteerOpportunityRequest) (*v1.ListVolunteerOpportunityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a VolunteerOpportunity Manager
	m := models.NewVolunteerOpportunityManager(s.db)

	// Get a list of volunteerOpportunitys given filters, ordering, and limit rules.
	list, err := m.ListVolunteerOpportunity(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListVolunteerOpportunityResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update VolunteerOpportunity
func (s *shrikeServiceServer) UpdateVolunteerOpportunity(ctx context.Context, req *v1.UpdateVolunteerOpportunityRequest) (*v1.UpdateVolunteerOpportunityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunity Manager
	m := models.NewVolunteerOpportunityManager(s.db)

	// Get a list of volunteerOpportunitys given filters, ordering, and limit rules.
	rows, err := m.UpdateVolunteerOpportunity(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateVolunteerOpportunityResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete volunteerOpportunity
func (s *shrikeServiceServer) DeleteVolunteerOpportunity(ctx context.Context, req *v1.DeleteVolunteerOpportunityRequest) (*v1.DeleteVolunteerOpportunityResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunity Manager
	m := models.NewVolunteerOpportunityManager(s.db)

	// Get a list of volunteerOpportunitys given filters, ordering, and limit rules.
	rows, err := m.DeleteVolunteerOpportunity(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteVolunteerOpportunityResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
