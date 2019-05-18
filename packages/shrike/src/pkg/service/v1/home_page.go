package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new HomePage
func (s *shrikeServiceServer) CreateHomePage(ctx context.Context, req *v1.CreateHomePageRequest) (*v1.CreateHomePageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a HomePage Manager
	m := models.NewHomePageManager(s.db)

	// Get a list of homePages given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateHomePageResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get homePage by id.
func (s *shrikeServiceServer) GetHomePage(ctx context.Context, req *v1.GetHomePageRequest) (*v1.GetHomePageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a HomePage Manager
	m := models.NewHomePageManager(s.db)

	// Get a list of homePages given filters, ordering, and limit rules.
	homePage, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetHomePageResponse{
		Api:  apiVersion,
		Item: m.GetProto(homePage),
	}, nil

}

// Read all HomePage
func (s *shrikeServiceServer) ListHomePage(ctx context.Context, req *v1.ListHomePageRequest) (*v1.ListHomePageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a HomePage Manager
	m := models.NewHomePageManager(s.db)

	// Get a list of homePages given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListHomePageResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update HomePage
func (s *shrikeServiceServer) UpdateHomePage(ctx context.Context, req *v1.UpdateHomePageRequest) (*v1.UpdateHomePageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a HomePage Manager
	m := models.NewHomePageManager(s.db)

	// Get a list of homePages given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateHomePageResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete homePage
func (s *shrikeServiceServer) DeleteHomePage(ctx context.Context, req *v1.DeleteHomePageRequest) (*v1.DeleteHomePageResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a HomePage Manager
	m := models.NewHomePageManager(s.db)

	// Get a list of homePages given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteHomePageResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
