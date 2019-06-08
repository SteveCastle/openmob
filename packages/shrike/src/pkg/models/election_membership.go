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

// ElectionMembership is a type for election_membership db element.
type ElectionMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Election  uuid.UUID
}

// ElectionMembershipManager manages queries returning a electionMembership or list of electionMemberships.
// It is configured with a db field to contain the db driver.
type ElectionMembershipManager struct {
	db *sql.DB
}

// NewElectionMembershipManager creates a electionMembership manager
func NewElectionMembershipManager(db *sql.DB) *ElectionMembershipManager {
	return &ElectionMembershipManager{db: db}
}

// CRUD Methods for the ElectionMembershipManager.

// Create creates a electionMembership.
func (m *ElectionMembershipManager) Create(ctx context.Context, item *v1.CreateElectionMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO election_membership (cause, election) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.Election).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ElectionMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ElectionMembership-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single electionMembership from the database by ID.
func (m *ElectionMembershipManager) Get(ctx context.Context, id string) (*ElectionMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ElectionMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, election FROM election_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ElectionMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ElectionMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%s' is not found", id))
	}

	// scan ElectionMembership data into protobuf model
	var electionMembership ElectionMembership

	if err := rows.Scan(&electionMembership.ID, &electionMembership.CreatedAt, &electionMembership.UpdatedAt, &electionMembership.Cause, &electionMembership.Election); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ElectionMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ElectionMembership rows with ID='%s'",
			id))
	}
	return &electionMembership, nil
}

// List returns a slice of all electionMemberships meeting the filter criteria.
func (m *ElectionMembershipManager) List(ctx context.Context, filters []*v1.ElectionMembershipFilterRule, orderings []*v1.ElectionMembershipOrdering, limit int64) ([]*ElectionMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ElectionMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildElectionMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ElectionMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ElectionMembership{}
	for rows.Next() {
		electionMembership := new(ElectionMembership)
		if err := rows.Scan(&electionMembership.ID, &electionMembership.CreatedAt, &electionMembership.UpdatedAt, &electionMembership.Cause, &electionMembership.Election); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ElectionMembership row-> "+err.Error())
		}
		list = append(list, electionMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ElectionMembership-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *ElectionMembershipManager) Update(ctx context.Context, item *v1.ElectionMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE election_membership SET cause=$2, election=$3 WHERE id=$1",
		item.ID, item.Cause, item.Election)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ElectionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ElectionMembershipManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM election_membership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ElectionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ElectionMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToElectionMembershipProto accepts a electionMembership struct and returns a protobuf electionMembership struct.
func convertToElectionMembershipProto(c *ElectionMembership) *v1.ElectionMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ElectionMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		Election:  c.Election.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a electionMembership.
func (*ElectionMembershipManager) GetProtoList(l []*ElectionMembership) []*v1.ElectionMembership {
	list := []*v1.ElectionMembership{}
	for _, v := range l {
		list = append(list, convertToElectionMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a electionMembership.
func (*ElectionMembershipManager) GetProto(c *ElectionMembership) *v1.ElectionMembership {
	return convertToElectionMembershipProto(c)
}

// BuildElectionMembershipListQuery takes a filter and ordering object for a electionMembership.
// and returns an SQL string
func BuildElectionMembershipListQuery(filters []*v1.ElectionMembershipFilterRule, orderings []*v1.ElectionMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, election FROM election_membership"
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
func (m *ElectionMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
