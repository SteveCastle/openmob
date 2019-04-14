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

// BoycottMembership is a type for boycott_membership db element.
type BoycottMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Boycott   uuid.UUID
}

// BoycottMembershipManager manages queries returning a boycottMembership or list of boycottMemberships.
// It is configured with a db field to contain the db driver.
type BoycottMembershipManager struct {
	db *sql.DB
}

// NewBoycottMembershipManager creates a boycottMembership manager
func NewBoycottMembershipManager(db *sql.DB) *BoycottMembershipManager {
	return &BoycottMembershipManager{db: db}
}

// CRUD Methods for the BoycottMembershipManager.

// CreateBoycottMembership creates a boycottMembership.
func (m *BoycottMembershipManager) CreateBoycottMembership(ctx context.Context, item *v1.CreateBoycottMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO boycott_membership (cause, boycott) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.Boycott).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into BoycottMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created BoycottMembership-> "+err.Error())
	}
	return &id, nil
}

// GetBoycottMembership gets a single boycottMembership from the database by ID.
func (m *BoycottMembershipManager) GetBoycottMembership(ctx context.Context, id string) (*BoycottMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query BoycottMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, boycott FROM boycott_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from BoycottMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from BoycottMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%s' is not found", id))
	}

	// scan BoycottMembership data into protobuf model
	var boycottMembership BoycottMembership

	if err := rows.Scan(&boycottMembership.ID, &boycottMembership.CreatedAt, &boycottMembership.UpdatedAt, &boycottMembership.Cause, &boycottMembership.Boycott); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from BoycottMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple BoycottMembership rows with ID='%s'",
			id))
	}
	return &boycottMembership, nil
}

// ListBoycottMembership returns a slice of all boycottMemberships meeting the filter criteria.
func (m *BoycottMembershipManager) ListBoycottMembership(ctx context.Context, filters []*v1.BoycottMembershipFilterRule, orderings []*v1.BoycottMembershipOrdering, limit int64) ([]*BoycottMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in BoycottMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildBoycottMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from BoycottMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*BoycottMembership{}
	for rows.Next() {
		boycottMembership := new(BoycottMembership)
		if err := rows.Scan(&boycottMembership.ID, &boycottMembership.CreatedAt, &boycottMembership.UpdatedAt, &boycottMembership.Cause, &boycottMembership.Boycott); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from BoycottMembership row-> "+err.Error())
		}
		list = append(list, boycottMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from BoycottMembership-> "+err.Error())
	}
	return list, nil
}

// UpdateBoycottMembership runs an update query on the provided db and returns the rows affected as an int64.
func (m *BoycottMembershipManager) UpdateBoycottMembership(ctx context.Context, item *v1.BoycottMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE boycott_membership SET cause=$2, boycott=$3 WHERE id=$1",
		item.ID, item.Cause, item.Boycott)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update BoycottMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteBoycottMembership creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *BoycottMembershipManager) DeleteBoycottMembership(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM boycottMembership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete BoycottMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BoycottMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToBoycottMembershipProto accepts a boycottMembership struct and returns a protobuf boycottMembership struct.
func convertToBoycottMembershipProto(c *BoycottMembership) *v1.BoycottMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.BoycottMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		Boycott:   c.Boycott.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a boycottMembership.
func (*BoycottMembershipManager) GetProtoList(l []*BoycottMembership) []*v1.BoycottMembership {
	list := []*v1.BoycottMembership{}
	for _, v := range l {
		list = append(list, convertToBoycottMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a boycottMembership.
func (*BoycottMembershipManager) GetProto(c *BoycottMembership) *v1.BoycottMembership {
	return convertToBoycottMembershipProto(c)
}

// BuildBoycottMembershipListQuery takes a filter and ordering object for a boycottMembership.
// and returns an SQL string
func BuildBoycottMembershipListQuery(filters []*v1.BoycottMembershipFilterRule, orderings []*v1.BoycottMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, boycott FROM boycott_membership"
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
func (m *BoycottMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
