package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new ElectionMembership
func (s *shrikeServiceServer) CreateElectionMembership(ctx context.Context, req *v1.CreateElectionMembershipRequest) (*v1.CreateElectionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ElectionMembership Manager
	m := models.NewElectionMembershipManager(s.db)

	// Get a list of electionMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateElectionMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get electionMembership by id.
func (s *shrikeServiceServer) GetElectionMembership(ctx context.Context, req *v1.GetElectionMembershipRequest) (*v1.GetElectionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ElectionMembership Manager
	m := models.NewElectionMembershipManager(s.db)

	// Get a list of electionMemberships given filters, ordering, and limit rules.
	electionMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetElectionMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(electionMembership),
	}, nil

}

// Read all ElectionMembership
func (s *shrikeServiceServer) ListElectionMembership(ctx context.Context, req *v1.ListElectionMembershipRequest) (*v1.ListElectionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a ElectionMembership Manager
	m := models.NewElectionMembershipManager(s.db)

	// Get a list of electionMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListElectionMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update ElectionMembership
func (s *shrikeServiceServer) UpdateElectionMembership(ctx context.Context, req *v1.UpdateElectionMembershipRequest) (*v1.UpdateElectionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ElectionMembership Manager
	m := models.NewElectionMembershipManager(s.db)

	// Get a list of electionMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateElectionMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete electionMembership
func (s *shrikeServiceServer) DeleteElectionMembership(ctx context.Context, req *v1.DeleteElectionMembershipRequest) (*v1.DeleteElectionMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a ElectionMembership Manager
	m := models.NewElectionMembershipManager(s.db)

	// Get a list of electionMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteElectionMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
