package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new AgentMembership
func (s *shrikeServiceServer) CreateAgentMembership(ctx context.Context, req *v1.CreateAgentMembershipRequest) (*v1.CreateAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// insert AgentMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO agent_membership (cause, agent) VALUES($1, $2)  RETURNING id;",
		req.Item.Cause, req.Item.Agent).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into AgentMembership-> "+err.Error())
	}

	// get ID of creates AgentMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created AgentMembership-> "+err.Error())
	}

	return &v1.CreateAgentMembershipResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get agent_membership by id.
func (s *shrikeServiceServer) GetAgentMembership(ctx context.Context, req *v1.GetAgentMembershipRequest) (*v1.GetAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query AgentMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, agent FROM agent_membership WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from AgentMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from AgentMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("AgentMembership with ID='%s' is not found",
			req.ID))
	}

	// scan AgentMembership data into protobuf model
	var agentmembership v1.AgentMembership
	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&agentmembership.ID, &createdAt, &updatedAt, &agentmembership.Cause, &agentmembership.Agent); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from AgentMembership row-> "+err.Error())
	}

	// Convert time.Time from database into proto timestamp.
	agentmembership.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
	}
	agentmembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple AgentMembership rows with ID='%s'",
			req.ID))
	}

	return &v1.GetAgentMembershipResponse{
		Api:  apiVersion,
		Item: &agentmembership,
	}, nil

}

// Read all AgentMembership
func (s *shrikeServiceServer) ListAgentMembership(ctx context.Context, req *v1.ListAgentMembershipRequest) (*v1.ListAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// get AgentMembership list
	queries.BuildAgentMembershipFilters(req.Filters, req.Ordering, req.Limit)
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, agent FROM agent_membership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from AgentMembership-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
	list := []*v1.AgentMembership{}
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		agentmembership := new(v1.AgentMembership)
		if err := rows.Scan(&agentmembership.ID, &createdAt, &updatedAt, &agentmembership.Cause, &agentmembership.Agent); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from AgentMembership row-> "+err.Error())
		}
		// Convert time.Time from database into proto timestamp.
		agentmembership.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		agentmembership.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updatedAt field has invalid format-> "+err.Error())
		}

		list = append(list, agentmembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from AgentMembership-> "+err.Error())
	}

	return &v1.ListAgentMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update AgentMembership
func (s *shrikeServiceServer) UpdateAgentMembership(ctx context.Context, req *v1.UpdateAgentMembershipRequest) (*v1.UpdateAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// update agent_membership
	res, err := c.ExecContext(ctx, "UPDATE agent_membership SET cause=$2, agent=$3 WHERE id=$1",
		req.Item.ID, req.Item.Cause, req.Item.Agent)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update AgentMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("AgentMembership with ID='%s' is not found",
			req.Item.ID))
	}

	return &v1.UpdateAgentMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete agent_membership
func (s *shrikeServiceServer) DeleteAgentMembership(ctx context.Context, req *v1.DeleteAgentMembershipRequest) (*v1.DeleteAgentMembershipResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// delete agent_membership
	res, err := c.ExecContext(ctx, "DELETE FROM agent_membership WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete AgentMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("AgentMembership with ID='%s' is not found",
			req.ID))
	}

	return &v1.DeleteAgentMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
