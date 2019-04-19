package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Election
func (s *shrikeServiceServer) CreateElection(ctx context.Context, req *v1.CreateElectionRequest) (*v1.CreateElectionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Election Manager
	m := models.NewElectionManager(s.db)

	// Get a list of elections given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateElectionResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get election by id.
func (s *shrikeServiceServer) GetElection(ctx context.Context, req *v1.GetElectionRequest) (*v1.GetElectionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Election Manager
	m := models.NewElectionManager(s.db)

	// Get a list of elections given filters, ordering, and limit rules.
	election, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetElectionResponse{
		Api:  apiVersion,
		Item: m.GetProto(election),
	}, nil

}

// Read all Election
func (s *shrikeServiceServer) ListElection(ctx context.Context, req *v1.ListElectionRequest) (*v1.ListElectionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Election Manager
	m := models.NewElectionManager(s.db)

	// Get a list of elections given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListElectionResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Election
func (s *shrikeServiceServer) UpdateElection(ctx context.Context, req *v1.UpdateElectionRequest) (*v1.UpdateElectionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Election Manager
	m := models.NewElectionManager(s.db)

	// Get a list of elections given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateElectionResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete election
func (s *shrikeServiceServer) DeleteElection(ctx context.Context, req *v1.DeleteElectionRequest) (*v1.DeleteElectionResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Election Manager
	m := models.NewElectionManager(s.db)

	// Get a list of elections given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteElectionResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
