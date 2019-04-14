package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Experiment
func (s *shrikeServiceServer) CreateExperiment(ctx context.Context, req *v1.CreateExperimentRequest) (*v1.CreateExperimentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Experiment Manager
	m := models.NewExperimentManager(s.db)

	// Get a list of experiments given filters, ordering, and limit rules.
	id, err := m.CreateExperiment(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateExperimentResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get experiment by id.
func (s *shrikeServiceServer) GetExperiment(ctx context.Context, req *v1.GetExperimentRequest) (*v1.GetExperimentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Experiment Manager
	m := models.NewExperimentManager(s.db)

	// Get a list of experiments given filters, ordering, and limit rules.
	experiment, err := m.GetExperiment(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetExperimentResponse{
		Api:  apiVersion,
		Item: m.GetProto(experiment),
	}, nil

}

// Read all Experiment
func (s *shrikeServiceServer) ListExperiment(ctx context.Context, req *v1.ListExperimentRequest) (*v1.ListExperimentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Experiment Manager
	m := models.NewExperimentManager(s.db)

	// Get a list of experiments given filters, ordering, and limit rules.
	list, err := m.ListExperiment(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListExperimentResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Experiment
func (s *shrikeServiceServer) UpdateExperiment(ctx context.Context, req *v1.UpdateExperimentRequest) (*v1.UpdateExperimentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Experiment Manager
	m := models.NewExperimentManager(s.db)

	// Get a list of experiments given filters, ordering, and limit rules.
	rows, err := m.UpdateExperiment(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateExperimentResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete experiment
func (s *shrikeServiceServer) DeleteExperiment(ctx context.Context, req *v1.DeleteExperimentRequest) (*v1.DeleteExperimentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Experiment Manager
	m := models.NewExperimentManager(s.db)

	// Get a list of experiments given filters, ordering, and limit rules.
	rows, err := m.DeleteExperiment(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteExperimentResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
