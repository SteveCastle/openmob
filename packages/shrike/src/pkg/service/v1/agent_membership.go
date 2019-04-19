package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models/v1"
)

// Create new AgentMembership
func (s *shrikeServiceServer) CreateAgentMembership(ctx context.Context, req *v1.CreateAgentMembershipRequest) (*v1.CreateAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a AgentMembership Manager
	m := models.NewAgentMembershipManager(s.db)

	// Get a list of agentMemberships given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateAgentMembershipResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get agentMembership by id.
func (s *shrikeServiceServer) GetAgentMembership(ctx context.Context, req *v1.GetAgentMembershipRequest) (*v1.GetAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a AgentMembership Manager
	m := models.NewAgentMembershipManager(s.db)

	// Get a list of agentMemberships given filters, ordering, and limit rules.
	agentMembership, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetAgentMembershipResponse{
		Api:  apiVersion,
		Item: m.GetProto(agentMembership),
	}, nil

}

// Read all AgentMembership
func (s *shrikeServiceServer) ListAgentMembership(ctx context.Context, req *v1.ListAgentMembershipRequest) (*v1.ListAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a AgentMembership Manager
	m := models.NewAgentMembershipManager(s.db)

	// Get a list of agentMemberships given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListAgentMembershipResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update AgentMembership
func (s *shrikeServiceServer) UpdateAgentMembership(ctx context.Context, req *v1.UpdateAgentMembershipRequest) (*v1.UpdateAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a AgentMembership Manager
	m := models.NewAgentMembershipManager(s.db)

	// Get a list of agentMemberships given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateAgentMembershipResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete agentMembership
func (s *shrikeServiceServer) DeleteAgentMembership(ctx context.Context, req *v1.DeleteAgentMembershipRequest) (*v1.DeleteAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a AgentMembership Manager
	m := models.NewAgentMembershipManager(s.db)

	// Get a list of agentMemberships given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteAgentMembershipResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
