package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new VolunteerOpportunityType
func (s *shrikeServiceServer) CreateVolunteerOpportunityType(ctx context.Context, req *v1.CreateVolunteerOpportunityTypeRequest) (*v1.CreateVolunteerOpportunityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunityType Manager
	m := models.NewVolunteerOpportunityTypeManager(s.db)

	// Get a list of volunteerOpportunityTypes given filters, ordering, and limit rules.
	id, err := m.CreateVolunteerOpportunityType(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateVolunteerOpportunityTypeResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get volunteerOpportunityType by id.
func (s *shrikeServiceServer) GetVolunteerOpportunityType(ctx context.Context, req *v1.GetVolunteerOpportunityTypeRequest) (*v1.GetVolunteerOpportunityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunityType Manager
	m := models.NewVolunteerOpportunityTypeManager(s.db)

	// Get a list of volunteerOpportunityTypes given filters, ordering, and limit rules.
	volunteerOpportunityType, err := m.GetVolunteerOpportunityType(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetVolunteerOpportunityTypeResponse{
		Api:  apiVersion,
		Item: m.GetProto(volunteerOpportunityType),
	}, nil

}

// Read all VolunteerOpportunityType
func (s *shrikeServiceServer) ListVolunteerOpportunityType(ctx context.Context, req *v1.ListVolunteerOpportunityTypeRequest) (*v1.ListVolunteerOpportunityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a VolunteerOpportunityType Manager
	m := models.NewVolunteerOpportunityTypeManager(s.db)

	// Get a list of volunteerOpportunityTypes given filters, ordering, and limit rules.
	list, err := m.ListVolunteerOpportunityType(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListVolunteerOpportunityTypeResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update VolunteerOpportunityType
func (s *shrikeServiceServer) UpdateVolunteerOpportunityType(ctx context.Context, req *v1.UpdateVolunteerOpportunityTypeRequest) (*v1.UpdateVolunteerOpportunityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunityType Manager
	m := models.NewVolunteerOpportunityTypeManager(s.db)

	// Get a list of volunteerOpportunityTypes given filters, ordering, and limit rules.
	rows, err := m.UpdateVolunteerOpportunityType(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateVolunteerOpportunityTypeResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete volunteerOpportunityType
func (s *shrikeServiceServer) DeleteVolunteerOpportunityType(ctx context.Context, req *v1.DeleteVolunteerOpportunityTypeRequest) (*v1.DeleteVolunteerOpportunityTypeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a VolunteerOpportunityType Manager
	m := models.NewVolunteerOpportunityTypeManager(s.db)

	// Get a list of volunteerOpportunityTypes given filters, ordering, and limit rules.
	rows, err := m.DeleteVolunteerOpportunityType(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteVolunteerOpportunityTypeResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
