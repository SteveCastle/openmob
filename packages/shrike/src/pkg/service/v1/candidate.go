package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Candidate
func (s *shrikeServiceServer) CreateCandidate(ctx context.Context, req *v1.CreateCandidateRequest) (*v1.CreateCandidateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Candidate Manager
	m := models.NewCandidateManager(s.db)

	// Get a list of candidates given filters, ordering, and limit rules.
	id, err := m.CreateCandidate(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateCandidateResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get candidate by id.
func (s *shrikeServiceServer) GetCandidate(ctx context.Context, req *v1.GetCandidateRequest) (*v1.GetCandidateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Candidate Manager
	m := models.NewCandidateManager(s.db)

	// Get a list of candidates given filters, ordering, and limit rules.
	candidate, err := m.GetCandidate(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetCandidateResponse{
		Api:  apiVersion,
		Item: m.GetProto(candidate),
	}, nil

}

// Read all Candidate
func (s *shrikeServiceServer) ListCandidate(ctx context.Context, req *v1.ListCandidateRequest) (*v1.ListCandidateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Candidate Manager
	m := models.NewCandidateManager(s.db)

	// Get a list of candidates given filters, ordering, and limit rules.
	list, err := m.ListCandidate(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListCandidateResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Candidate
func (s *shrikeServiceServer) UpdateCandidate(ctx context.Context, req *v1.UpdateCandidateRequest) (*v1.UpdateCandidateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Candidate Manager
	m := models.NewCandidateManager(s.db)

	// Get a list of candidates given filters, ordering, and limit rules.
	rows, err := m.UpdateCandidate(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateCandidateResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete candidate
func (s *shrikeServiceServer) DeleteCandidate(ctx context.Context, req *v1.DeleteCandidateRequest) (*v1.DeleteCandidateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Candidate Manager
	m := models.NewCandidateManager(s.db)

	// Get a list of candidates given filters, ordering, and limit rules.
	rows, err := m.DeleteCandidate(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteCandidateResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
