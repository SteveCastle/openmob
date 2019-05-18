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

// PollMembership is a type for poll_membership db element.
type PollMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Petition  uuid.UUID
}

// PollMembershipManager manages queries returning a pollMembership or list of pollMemberships.
// It is configured with a db field to contain the db driver.
type PollMembershipManager struct {
	db *sql.DB
}

// NewPollMembershipManager creates a pollMembership manager
func NewPollMembershipManager(db *sql.DB) *PollMembershipManager {
	return &PollMembershipManager{db: db}
}

// CRUD Methods for the PollMembershipManager.

// Create creates a pollMembership.
func (m *PollMembershipManager) Create(ctx context.Context, item *v1.CreatePollMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO poll_membership (cause, petition) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.Petition).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PollMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PollMembership-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single pollMembership from the database by ID.
func (m *PollMembershipManager) Get(ctx context.Context, id string) (*PollMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query PollMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, petition FROM poll_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PollMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%s' is not found", id))
	}

	// scan PollMembership data into protobuf model
	var pollMembership PollMembership

	if err := rows.Scan(&pollMembership.ID, &pollMembership.CreatedAt, &pollMembership.UpdatedAt, &pollMembership.Cause, &pollMembership.Petition); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PollMembership rows with ID='%s'",
			id))
	}
	return &pollMembership, nil
}

// List returns a slice of all pollMemberships meeting the filter criteria.
func (m *PollMembershipManager) List(ctx context.Context, filters []*v1.PollMembershipFilterRule, orderings []*v1.PollMembershipOrdering, limit int64) ([]*PollMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in PollMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPollMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*PollMembership{}
	for rows.Next() {
		pollMembership := new(PollMembership)
		if err := rows.Scan(&pollMembership.ID, &pollMembership.CreatedAt, &pollMembership.UpdatedAt, &pollMembership.Cause, &pollMembership.Petition); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollMembership row-> "+err.Error())
		}
		list = append(list, pollMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PollMembership-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *PollMembershipManager) Update(ctx context.Context, item *v1.PollMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE poll_membership SET cause=$2, petition=$3 WHERE id=$1",
		item.ID, item.Cause, item.Petition)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PollMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PollMembershipManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM pollMembership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PollMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPollMembershipProto accepts a pollMembership struct and returns a protobuf pollMembership struct.
func convertToPollMembershipProto(c *PollMembership) *v1.PollMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.PollMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		Petition:  c.Petition.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a pollMembership.
func (*PollMembershipManager) GetProtoList(l []*PollMembership) []*v1.PollMembership {
	list := []*v1.PollMembership{}
	for _, v := range l {
		list = append(list, convertToPollMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a pollMembership.
func (*PollMembershipManager) GetProto(c *PollMembership) *v1.PollMembership {
	return convertToPollMembershipProto(c)
}

// BuildPollMembershipListQuery takes a filter and ordering object for a pollMembership.
// and returns an SQL string
func BuildPollMembershipListQuery(filters []*v1.PollMembershipFilterRule, orderings []*v1.PollMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, petition FROM poll_membership"
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
func (m *PollMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
