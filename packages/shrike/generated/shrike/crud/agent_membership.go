package v1

import (
	"context"
	"database/sql"
	"fmt"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// shrikeServiceServer is implementation of v1.ShrikeServiceServer proto interface
type shrikeServiceServer struct {
	db *sql.DB
}

// NewShrikeServiceServer creates AgentMembership service
func NewShrikeServiceServer(db *sql.DB) v1.ShrikeServiceServer {
	return &shrikeServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *shrikeServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *shrikeServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Create new todo task
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
	var id int64
	// insert AgentMembership entity data
	err = c.QueryRowContext(ctx, "INSERT INTO agentmembership (title) VALUES($1)  RETURNING id;",
		req.Item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into AgentMembership-> "+err.Error())
	}

	// get ID of creates AgentMembership
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created AgentMembership-> "+err.Error())
	}

	return &v1.CreateAgentMembershipResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Get agentmembership by id.
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
	rows, err := c.QueryContext(ctx, "SELECT id, title FROM agentmembership WHERE id=$1",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from AgentMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from AgentMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("AgentMembership with ID='%d' is not found",
			req.Id))
	}

	// get AgentMembership data
	var td v1.AgentMembership
	if err := rows.Scan(&td.Id, &td.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from AgentMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple AgentMembership rows with ID='%d'",
			req.Id))
	}

	return &v1.GetAgentMembershipResponse{
		Api:  apiVersion,
		Item: &td,
	}, nil

}

// Read all todo tasks
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
	rows, err := c.QueryContext(ctx, "SELECT id,title FROM AgentMembership")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from AgentMembership-> "+err.Error())
	}
	defer rows.Close()

	list := []*v1.AgentMembership{}
	for rows.Next() {
		td := new(v1.AgentMembership)
		if err := rows.Scan(&td.Id, &td.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from AgentMembership row-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from AgentMembership-> "+err.Error())
	}

	return &v1.ListAgentMembershipResponse{
		Api:   apiVersion,
		Items: list,
	}, nil
}

// Update todo task
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

	// update agentmembership
	res, err := c.ExecContext(ctx, "UPDATE agentmembership SET title=$1 WHERE id=$2",
		req.Item.Title, req.Item.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update agentmembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("agentmembership with ID='%d' is not found",
			req.Item.Id))
	}

	return &v1.UpdateAgentMembershipResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete agentmembership
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

	// delete agentmembership
	res, err := c.ExecContext(ctx, "DELETE FROM agentmembership WHERE id=$1", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete agentmembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("agentmembership with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteAgentMembershipResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}
