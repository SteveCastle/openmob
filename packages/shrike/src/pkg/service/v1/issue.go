package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Issue
func (s *shrikeServiceServer) CreateIssue(ctx context.Context, req *v1.CreateIssueRequest) (*v1.CreateIssueResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Issue Manager
	m := models.NewIssueManager(s.db)

	// Get a list of issues given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateIssueResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get issue by id.
func (s *shrikeServiceServer) GetIssue(ctx context.Context, req *v1.GetIssueRequest) (*v1.GetIssueResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Issue Manager
	m := models.NewIssueManager(s.db)

	// Get a list of issues given filters, ordering, and limit rules.
	issue, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetIssueResponse{
		Api:  apiVersion,
		Item: m.GetProto(issue),
	}, nil

}

// Read all Issue
func (s *shrikeServiceServer) ListIssue(ctx context.Context, req *v1.ListIssueRequest) (*v1.ListIssueResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Issue Manager
	m := models.NewIssueManager(s.db)

	// Get a list of issues given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListIssueResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Issue
func (s *shrikeServiceServer) UpdateIssue(ctx context.Context, req *v1.UpdateIssueRequest) (*v1.UpdateIssueResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Issue Manager
	m := models.NewIssueManager(s.db)

	// Get a list of issues given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateIssueResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete issue
func (s *shrikeServiceServer) DeleteIssue(ctx context.Context, req *v1.DeleteIssueRequest) (*v1.DeleteIssueResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Issue Manager
	m := models.NewIssueManager(s.db)

	// Get a list of issues given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteIssueResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
