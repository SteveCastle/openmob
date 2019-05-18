package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new ExperimentVariant
func (s *shrikeServiceServer) CreateExperimentVariant(ctx context.Context, req *v1.CreateExperimentVariantRequest) (*v1.CreateExperimentVariantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ExperimentVariant Manager
	m := models.NewExperimentVariantManager(s.db)

	// Get a list of experimentVariants given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateExperimentVariantResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get experimentVariant by id.
func (s *shrikeServiceServer) GetExperimentVariant(ctx context.Context, req *v1.GetExperimentVariantRequest) (*v1.GetExperimentVariantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ExperimentVariant Manager
	m := models.NewExperimentVariantManager(s.db)

	// Get a list of experimentVariants given filters, ordering, and limit rules.
	experimentVariant, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetExperimentVariantResponse{
		Api:  apiVersion,
		Item: m.GetProto(experimentVariant),
	}, nil

}

// Read all ExperimentVariant
func (s *shrikeServiceServer) ListExperimentVariant(ctx context.Context, req *v1.ListExperimentVariantRequest) (*v1.ListExperimentVariantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ExperimentVariant Manager
	m := models.NewExperimentVariantManager(s.db)

	// Get a list of experimentVariants given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListExperimentVariantResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ExperimentVariant
func (s *shrikeServiceServer) UpdateExperimentVariant(ctx context.Context, req *v1.UpdateExperimentVariantRequest) (*v1.UpdateExperimentVariantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ExperimentVariant Manager
	m := models.NewExperimentVariantManager(s.db)

	// Get a list of experimentVariants given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateExperimentVariantResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete experimentVariant
func (s *shrikeServiceServer) DeleteExperimentVariant(ctx context.Context, req *v1.DeleteExperimentVariantRequest) (*v1.DeleteExperimentVariantResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ExperimentVariant Manager
	m := models.NewExperimentVariantManager(s.db)

	// Get a list of experimentVariants given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteExperimentVariantResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
