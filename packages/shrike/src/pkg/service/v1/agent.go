package v1

import (
	"context"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new Agent
func (s *shrikeServiceServer) CreateAgent(ctx context.Context, req *v1.CreateAgentRequest) (*v1.CreateAgentResponse, error) {
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
	var id int64
	// insert Agent entity data
	err = c.QueryRowContext(ctx, "INSERT INTO agent (id, created_at, updated_at, account) VALUES($1, $2, $3, $4)  RETURNING id;",
		 req.Item.ID,  req.Item.CreatedAt,  req.Item.UpdatedAt,  req.Item.Account, ).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Agent-> "+err.Error())
	}

	// get ID of creates Agent
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Agent-> "+err.Error())
	}

	return &v1.CreateAgentResponse{
		Api: apiVersion,
		ID:  id,
	}, nil
}

// Get agent by id.
func (s *shrikeServiceServer) GetAgent(ctx context.Context, req *v1.GetAgentRequest) (*v1.GetAgentResponse, error) {
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

	// query Agent by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, account FROM agent WHERE id=$1",
		req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Agent-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Agent-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Agent with ID='%d' is not found",
			req.ID))
	}

	// get Agent data
	var agent v1.Agent
	if err := rows.Scan( &agent.ID,  &agent.CreatedAt,  &agent.UpdatedAt,  &agent.Account, ); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Agent row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Agent rows with ID='%d'",
			req.ID))
	}

	return &v1.GetAgentResponse{
		Api:  apiVersion,
		Item: &agent,
	}, nil

}

// Read all Agent
func (s *shrikeServiceServer) ListAgent(ctx context.Context, req *v1.ListAgentRequest) (*v1.ListAgentResponse, error) {
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

	// get Agent list
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, account FROM agent")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Agent-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.Agent{}
	for rows.Next() {
		agent := new(v1.Agent)
		if err := rows.Scan( &agent.ID,  &agent.CreatedAt,  &agent.UpdatedAt,  &agent.Account, ); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Agent row-> "+err.Error())
		}
		list = append(list, agent)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Agent-> "+err.Error())
	}

	return &v1.ListAgentResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update Agent
func (s *shrikeServiceServer) UpdateAgent(ctx context.Context, req *v1.UpdateAgentRequest) (*v1.UpdateAgentResponse, error) {
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

	// update agent
	res, err := c.ExecContext(ctx, "UPDATE agent SET id=$1, created_at=$2, updated_at=$3, account=$4 WHERE id=$1",
		req.Item.ID,req.Item.CreatedAt,req.Item.UpdatedAt,req.Item.Account, )
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Agent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Agent with ID='%d' is not found",
			req.Item.ID))
	}

	return &v1.UpdateAgentResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete agent
func (s *shrikeServiceServer) DeleteAgent(ctx context.Context, req *v1.DeleteAgentRequest) (*v1.DeleteAgentResponse, error) {
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

	// delete agent
	res, err := c.ExecContext(ctx, "DELETE FROM agent WHERE id=$1", req.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Agent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Agent with ID='%d' is not found",
			req.ID))
	}

	return &v1.DeleteAgentResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
