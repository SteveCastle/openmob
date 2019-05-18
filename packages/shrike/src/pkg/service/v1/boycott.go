package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Boycott
func (s *shrikeServiceServer) CreateBoycott(ctx context.Context, req *v1.CreateBoycottRequest) (*v1.CreateBoycottResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Boycott Manager
	m := models.NewBoycottManager(s.db)

	// Get a list of boycotts given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateBoycottResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get boycott by id.
func (s *shrikeServiceServer) GetBoycott(ctx context.Context, req *v1.GetBoycottRequest) (*v1.GetBoycottResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Boycott Manager
	m := models.NewBoycottManager(s.db)

	// Get a list of boycotts given filters, ordering, and limit rules.
	boycott, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetBoycottResponse{
		Api:  apiVersion,
		Item: m.GetProto(boycott),
	}, nil

}

// Read all Boycott
func (s *shrikeServiceServer) ListBoycott(ctx context.Context, req *v1.ListBoycottRequest) (*v1.ListBoycottResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Boycott Manager
	m := models.NewBoycottManager(s.db)

	// Get a list of boycotts given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListBoycottResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Boycott
func (s *shrikeServiceServer) UpdateBoycott(ctx context.Context, req *v1.UpdateBoycottRequest) (*v1.UpdateBoycottResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Boycott Manager
	m := models.NewBoycottManager(s.db)

	// Get a list of boycotts given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateBoycottResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete boycott
func (s *shrikeServiceServer) DeleteBoycott(ctx context.Context, req *v1.DeleteBoycottRequest) (*v1.DeleteBoycottResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Boycott Manager
	m := models.NewBoycottManager(s.db)

	// Get a list of boycotts given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteBoycottResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
