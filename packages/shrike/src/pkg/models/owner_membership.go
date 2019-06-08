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

// OwnerMembership is a type for owner_membership db element.
type OwnerMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Account   uuid.UUID
}

// OwnerMembershipManager manages queries returning a ownerMembership or list of ownerMemberships.
// It is configured with a db field to contain the db driver.
type OwnerMembershipManager struct {
	db *sql.DB
}

// NewOwnerMembershipManager creates a ownerMembership manager
func NewOwnerMembershipManager(db *sql.DB) *OwnerMembershipManager {
	return &OwnerMembershipManager{db: db}
}

// CRUD Methods for the OwnerMembershipManager.

// Create creates a ownerMembership.
func (m *OwnerMembershipManager) Create(ctx context.Context, item *v1.CreateOwnerMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO owner_membership (cause, account) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.Account).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into OwnerMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created OwnerMembership-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single ownerMembership from the database by ID.
func (m *OwnerMembershipManager) Get(ctx context.Context, id string) (*OwnerMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query OwnerMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, account FROM owner_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from OwnerMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from OwnerMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%s' is not found", id))
	}

	// scan OwnerMembership data into protobuf model
	var ownerMembership OwnerMembership

	if err := rows.Scan(&ownerMembership.ID, &ownerMembership.CreatedAt, &ownerMembership.UpdatedAt, &ownerMembership.Cause, &ownerMembership.Account); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from OwnerMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple OwnerMembership rows with ID='%s'",
			id))
	}
	return &ownerMembership, nil
}

// List returns a slice of all ownerMemberships meeting the filter criteria.
func (m *OwnerMembershipManager) List(ctx context.Context, filters []*v1.OwnerMembershipFilterRule, orderings []*v1.OwnerMembershipOrdering, limit int64) ([]*OwnerMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in OwnerMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildOwnerMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from OwnerMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*OwnerMembership{}
	for rows.Next() {
		ownerMembership := new(OwnerMembership)
		if err := rows.Scan(&ownerMembership.ID, &ownerMembership.CreatedAt, &ownerMembership.UpdatedAt, &ownerMembership.Cause, &ownerMembership.Account); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from OwnerMembership row-> "+err.Error())
		}
		list = append(list, ownerMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from OwnerMembership-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *OwnerMembershipManager) Update(ctx context.Context, item *v1.OwnerMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE owner_membership SET cause=$2, account=$3 WHERE id=$1",
		item.ID, item.Cause, item.Account)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update OwnerMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *OwnerMembershipManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM owner_membership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete OwnerMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("OwnerMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToOwnerMembershipProto accepts a ownerMembership struct and returns a protobuf ownerMembership struct.
func convertToOwnerMembershipProto(c *OwnerMembership) *v1.OwnerMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.OwnerMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		Account:   c.Account.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a ownerMembership.
func (*OwnerMembershipManager) GetProtoList(l []*OwnerMembership) []*v1.OwnerMembership {
	list := []*v1.OwnerMembership{}
	for _, v := range l {
		list = append(list, convertToOwnerMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a ownerMembership.
func (*OwnerMembershipManager) GetProto(c *OwnerMembership) *v1.OwnerMembership {
	return convertToOwnerMembershipProto(c)
}

// BuildOwnerMembershipListQuery takes a filter and ordering object for a ownerMembership.
// and returns an SQL string
func BuildOwnerMembershipListQuery(filters []*v1.OwnerMembershipFilterRule, orderings []*v1.OwnerMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, account FROM owner_membership"
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
func (m *OwnerMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
