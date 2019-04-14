package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new Voter
func (s *shrikeServiceServer) CreateVoter(ctx context.Context, req *v1.CreateVoterRequest) (*v1.CreateVoterResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Voter Manager
	m := models.NewVoterManager(s.db)

	// Get a list of voters given filters, ordering, and limit rules.
	id, err := m.CreateVoter(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateVoterResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get voter by id.
func (s *shrikeServiceServer) GetVoter(ctx context.Context, req *v1.GetVoterRequest) (*v1.GetVoterResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Voter Manager
	m := models.NewVoterManager(s.db)

	// Get a list of voters given filters, ordering, and limit rules.
	voter, err := m.GetVoter(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetVoterResponse{
		Api:  apiVersion,
		Item: m.GetProto(voter),
	}, nil

}

// Read all Voter
func (s *shrikeServiceServer) ListVoter(ctx context.Context, req *v1.ListVoterRequest) (*v1.ListVoterResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Voter Manager
	m := models.NewVoterManager(s.db)

	// Get a list of voters given filters, ordering, and limit rules.
	list, err := m.ListVoter(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListVoterResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Voter
func (s *shrikeServiceServer) UpdateVoter(ctx context.Context, req *v1.UpdateVoterRequest) (*v1.UpdateVoterResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Voter Manager
	m := models.NewVoterManager(s.db)

	// Get a list of voters given filters, ordering, and limit rules.
	rows, err := m.UpdateVoter(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateVoterResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete voter
func (s *shrikeServiceServer) DeleteVoter(ctx context.Context, req *v1.DeleteVoterRequest) (*v1.DeleteVoterResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Voter Manager
	m := models.NewVoterManager(s.db)

	// Get a list of voters given filters, ordering, and limit rules.
	rows, err := m.DeleteVoter(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteVoterResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
