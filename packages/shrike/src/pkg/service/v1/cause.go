package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Cause
func (s *shrikeServiceServer) CreateCause(ctx context.Context, req *v1.CreateCauseRequest) (*v1.CreateCauseResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Cause Manager
	m := models.NewCauseManager(s.db)

	// Get a list of causes given filters, ordering, and limit rules.
	id, err := m.CreateCause(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateCauseResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get cause by id.
func (s *shrikeServiceServer) GetCause(ctx context.Context, req *v1.GetCauseRequest) (*v1.GetCauseResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Cause Manager
	m := models.NewCauseManager(s.db)

	// Get a list of causes given filters, ordering, and limit rules.
	cause, err := m.GetCause(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetCauseResponse{
		Api:  apiVersion,
		Item: m.GetProto(cause),
	}, nil

}

// Read all Cause
func (s *shrikeServiceServer) ListCause(ctx context.Context, req *v1.ListCauseRequest) (*v1.ListCauseResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Cause Manager
	m := models.NewCauseManager(s.db)

	// Get a list of causes given filters, ordering, and limit rules.
	list, err := m.ListCause(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListCauseResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Cause
func (s *shrikeServiceServer) UpdateCause(ctx context.Context, req *v1.UpdateCauseRequest) (*v1.UpdateCauseResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Cause Manager
	m := models.NewCauseManager(s.db)

	// Get a list of causes given filters, ordering, and limit rules.
	rows, err := m.UpdateCause(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateCauseResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete cause
func (s *shrikeServiceServer) DeleteCause(ctx context.Context, req *v1.DeleteCauseRequest) (*v1.DeleteCauseResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Cause Manager
	m := models.NewCauseManager(s.db)

	// Get a list of causes given filters, ordering, and limit rules.
	rows, err := m.DeleteCause(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteCauseResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
