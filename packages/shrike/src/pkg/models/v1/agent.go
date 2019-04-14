package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/structs"
	uuid "github.com/gofrs/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Agent is a type for agent db element.
type Agent struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Account   uuid.UUID
}

// AgentManager manages queries returning a agent or list of agents.
// It is configured with a db field to contain the db driver.
type AgentManager struct {
	db *sql.DB
}

// NewAgentManager creates a agent manager
func NewAgentManager(db *sql.DB) *AgentManager {
	return &AgentManager{db: db}
}

// CRUD Methods for the AgentManager.

// CreateAgent creates a agent.
func (m *AgentManager) CreateAgent(ctx context.Context, item *v1.CreateAgent) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO agent (account) VALUES($1)  RETURNING id;",
		item.Account).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Agent-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Agent-> "+err.Error())
	}
	return &id, nil
}

// GetAgent gets a single agent from the database by ID.
func (m *AgentManager) GetAgent(ctx context.Context, id string) (*Agent, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Agent by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, account FROM agent WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Agent-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Agent-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Agent with ID='%s' is not found", id))
	}

	// scan Agent data into protobuf model
	var agent Agent

	if err := rows.Scan(&agent.ID, &agent.CreatedAt, &agent.UpdatedAt, &agent.Account); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Agent row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Agent rows with ID='%s'",
			id))
	}
	return &agent, nil
}

// ListAgent returns a slice of all agents meeting the filter criteria.
func (m *AgentManager) ListAgent(ctx context.Context, filters []*v1.AgentFilterRule, orderings []*v1.AgentOrdering, limit int64) ([]*Agent, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Agent Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildAgentListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Agent-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Agent{}
	for rows.Next() {
		agent := new(Agent)
		if err := rows.Scan(&agent.ID, &agent.CreatedAt, &agent.UpdatedAt, &agent.Account); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Agent row-> "+err.Error())
		}
		list = append(list, agent)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Agent-> "+err.Error())
	}
	return list, nil
}

// UpdateAgent runs an update query on the provided db and returns the rows affected as an int64.
func (m *AgentManager) UpdateAgent(ctx context.Context, item *v1.Agent) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE agent SET account=$2 WHERE id=$1",
		item.ID, item.Account)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Agent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Agent with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteAgent creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *AgentManager) DeleteAgent(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM agent WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Agent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Agent with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToAgentProto accepts a agent struct and returns a protobuf agent struct.
func convertToAgentProto(c *Agent) *v1.Agent {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Agent{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Account:   c.Account.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a agent.
func (*AgentManager) GetProtoList(l []*Agent) []*v1.Agent {
	list := []*v1.Agent{}
	for _, v := range l {
		list = append(list, convertToAgentProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a agent.
func (*AgentManager) GetProto(c *Agent) *v1.Agent {
	return convertToAgentProto(c)
}

// BuildAgentListQuery takes a filter and ordering object for a agent.
// and returns an SQL string
func BuildAgentListQuery(filters []*v1.AgentFilterRule, orderings []*v1.AgentOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, account FROM agent"
	// Range over the provided rules and create where clauses.
	for i, r := range filters {
		if i == 0 {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "WHERE")
		} else {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, r.LogicalOperator)
		}
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			if f.IsExported() {
				baseSQL = fmt.Sprintf("%s %s %s '%s'", baseSQL, ToSnakeCase(f.Name()), Comparison[r.Rule.String()], f.Value())
			}
		}
	}
	// Range over ordering rules and create ORDER BY clauses.
	for _, r := range orderings {
		fmt.Println(r.Direction)
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "ORDER BY")
			if f.IsExported() {
				baseSQL = fmt.Sprintf("%s %s %s", baseSQL, ToSnakeCase(f.Name()), SQLDirections[r.Direction.String()])
			}
		}

	}
	baseSQL = fmt.Sprintf("%s LIMIT %d;", baseSQL, limit)
	fmt.Printf("List SQL Executed: %v\n", baseSQL)
	return baseSQL
}

// connect returns SQL database connection from the pool.
func (m *AgentManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
