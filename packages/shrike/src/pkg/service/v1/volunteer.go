package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Volunteer
func (s *shrikeServiceServer) CreateVolunteer(ctx context.Context, req *v1.CreateVolunteerRequest) (*v1.CreateVolunteerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Volunteer Manager
	m := models.NewVolunteerManager(s.db)

	// Get a list of volunteers given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateVolunteerResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get volunteer by id.
func (s *shrikeServiceServer) GetVolunteer(ctx context.Context, req *v1.GetVolunteerRequest) (*v1.GetVolunteerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Volunteer Manager
	m := models.NewVolunteerManager(s.db)

	// Get a list of volunteers given filters, ordering, and limit rules.
	volunteer, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetVolunteerResponse{
		Api:  apiVersion,
		Item: m.GetProto(volunteer),
	}, nil

}

// Read all Volunteer
func (s *shrikeServiceServer) ListVolunteer(ctx context.Context, req *v1.ListVolunteerRequest) (*v1.ListVolunteerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Volunteer Manager
	m := models.NewVolunteerManager(s.db)

	// Get a list of volunteers given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListVolunteerResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Volunteer
func (s *shrikeServiceServer) UpdateVolunteer(ctx context.Context, req *v1.UpdateVolunteerRequest) (*v1.UpdateVolunteerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Volunteer Manager
	m := models.NewVolunteerManager(s.db)

	// Get a list of volunteers given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateVolunteerResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete volunteer
func (s *shrikeServiceServer) DeleteVolunteer(ctx context.Context, req *v1.DeleteVolunteerRequest) (*v1.DeleteVolunteerResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Volunteer Manager
	m := models.NewVolunteerManager(s.db)

	// Get a list of volunteers given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteVolunteerResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
