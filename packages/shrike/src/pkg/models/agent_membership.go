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

// AgentMembership is a type for agent_membership db element.
type AgentMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Agent     uuid.UUID
}

// AgentMembershipManager manages queries returning a agentMembership or list of agentMemberships.
// It is configured with a db field to contain the db driver.
type AgentMembershipManager struct {
	db *sql.DB
}

// NewAgentMembershipManager creates a agentMembership manager
func NewAgentMembershipManager(db *sql.DB) *AgentMembershipManager {
	return &AgentMembershipManager{db: db}
}

// CRUD Methods for the AgentMembershipManager.

// Create creates a agentMembership.
func (m *AgentMembershipManager) Create(ctx context.Context, item *v1.CreateAgentMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO agent_membership (cause, agent) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.Agent).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into AgentMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created AgentMembership-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single agentMembership from the database by ID.
func (m *AgentMembershipManager) Get(ctx context.Context, id string) (*AgentMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query AgentMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, agent FROM agent_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from AgentMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from AgentMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("AgentMembership with ID='%s' is not found", id))
	}

	// scan AgentMembership data into protobuf model
	var agentMembership AgentMembership

	if err := rows.Scan(&agentMembership.ID, &agentMembership.CreatedAt, &agentMembership.UpdatedAt, &agentMembership.Cause, &agentMembership.Agent); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from AgentMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple AgentMembership rows with ID='%s'",
			id))
	}
	return &agentMembership, nil
}

// List returns a slice of all agentMemberships meeting the filter criteria.
func (m *AgentMembershipManager) List(ctx context.Context, filters []*v1.AgentMembershipFilterRule, orderings []*v1.AgentMembershipOrdering, limit int64) ([]*AgentMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in AgentMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildAgentMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from AgentMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*AgentMembership{}
	for rows.Next() {
		agentMembership := new(AgentMembership)
		if err := rows.Scan(&agentMembership.ID, &agentMembership.CreatedAt, &agentMembership.UpdatedAt, &agentMembership.Cause, &agentMembership.Agent); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from AgentMembership row-> "+err.Error())
		}
		list = append(list, agentMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from AgentMembership-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *AgentMembershipManager) Update(ctx context.Context, item *v1.AgentMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE agent_membership SET cause=$2, agent=$3 WHERE id=$1",
		item.ID, item.Cause, item.Agent)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update AgentMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("AgentMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *AgentMembershipManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM agent_membership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete AgentMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("AgentMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToAgentMembershipProto accepts a agentMembership struct and returns a protobuf agentMembership struct.
func convertToAgentMembershipProto(c *AgentMembership) *v1.AgentMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.AgentMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		Agent:     c.Agent.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a agentMembership.
func (*AgentMembershipManager) GetProtoList(l []*AgentMembership) []*v1.AgentMembership {
	list := []*v1.AgentMembership{}
	for _, v := range l {
		list = append(list, convertToAgentMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a agentMembership.
func (*AgentMembershipManager) GetProto(c *AgentMembership) *v1.AgentMembership {
	return convertToAgentMembershipProto(c)
}

// BuildAgentMembershipListQuery takes a filter and ordering object for a agentMembership.
// and returns an SQL string
func BuildAgentMembershipListQuery(filters []*v1.AgentMembershipFilterRule, orderings []*v1.AgentMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, agent FROM agent_membership"
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
func (m *AgentMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
