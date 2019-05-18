package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new LandingPage
func (s *shrikeServiceServer) CreateLandingPage(ctx context.Context, req *v1.CreateLandingPageRequest) (*v1.CreateLandingPageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LandingPage Manager
	m := models.NewLandingPageManager(s.db)

	// Get a list of landingPages given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateLandingPageResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get landingPage by id.
func (s *shrikeServiceServer) GetLandingPage(ctx context.Context, req *v1.GetLandingPageRequest) (*v1.GetLandingPageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LandingPage Manager
	m := models.NewLandingPageManager(s.db)

	// Get a list of landingPages given filters, ordering, and limit rules.
	landingPage, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetLandingPageResponse{
		Api:  apiVersion,
		Item: m.GetProto(landingPage),
	}, nil

}

// Read all LandingPage
func (s *shrikeServiceServer) ListLandingPage(ctx context.Context, req *v1.ListLandingPageRequest) (*v1.ListLandingPageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a LandingPage Manager
	m := models.NewLandingPageManager(s.db)

	// Get a list of landingPages given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListLandingPageResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update LandingPage
func (s *shrikeServiceServer) UpdateLandingPage(ctx context.Context, req *v1.UpdateLandingPageRequest) (*v1.UpdateLandingPageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LandingPage Manager
	m := models.NewLandingPageManager(s.db)

	// Get a list of landingPages given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateLandingPageResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete landingPage
func (s *shrikeServiceServer) DeleteLandingPage(ctx context.Context, req *v1.DeleteLandingPageRequest) (*v1.DeleteLandingPageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a LandingPage Manager
	m := models.NewLandingPageManager(s.db)

	// Get a list of landingPages given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteLandingPageResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
