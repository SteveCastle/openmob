package v1

import (
	"context"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/models"
)

// Create new Agent
func (s *shrikeServiceServer) CreateAgent(ctx context.Context, req *v1.CreateAgentRequest) (*v1.CreateAgentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Agent Manager
	m := models.NewAgentManager(s.db)

	// Get a list of agents given filters, ordering, and limit rules.
	id, err := m.Create(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &v1.CreateAgentResponse{
		Api: apiVersion,
		ID:  *id,
	}, nil
}

// Get agent by id.
func (s *shrikeServiceServer) GetAgent(ctx context.Context, req *v1.GetAgentRequest) (*v1.GetAgentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Agent Manager
	m := models.NewAgentManager(s.db)

	// Get a list of agents given filters, ordering, and limit rules.
	agent, err := m.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetAgentResponse{
		Api:  apiVersion,
		Item: m.GetProto(agent),
	}, nil

}

// Read all Agent
func (s *shrikeServiceServer) ListAgent(ctx context.Context, req *v1.ListAgentRequest) (*v1.ListAgentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Create a Agent Manager
	m := models.NewAgentManager(s.db)

	// Get a list of agents given filters, ordering, and limit rules.
	list, err := m.List(ctx, req.Filters, req.Ordering, req.Limit)
	if err != nil {
		return nil, err
	}

	return &v1.ListAgentResponse{
		Api:   apiVersion,
		Items: m.GetProtoList(list),
	}, nil
}

// Update Agent
func (s *shrikeServiceServer) UpdateAgent(ctx context.Context, req *v1.UpdateAgentRequest) (*v1.UpdateAgentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Agent Manager
	m := models.NewAgentManager(s.db)

	// Get a list of agents given filters, ordering, and limit rules.
	rows, err := m.Update(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateAgentResponse{
		Api:     apiVersion,
		Updated: *rows,
	}, nil
}

// Delete agent
func (s *shrikeServiceServer) DeleteAgent(ctx context.Context, req *v1.DeleteAgentRequest) (*v1.DeleteAgentResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Create a Agent Manager
	m := models.NewAgentManager(s.db)

	// Get a list of agents given filters, ordering, and limit rules.
	rows, err := m.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteAgentResponse{
		Api:     apiVersion,
		Deleted: *rows,
	}, nil
}
