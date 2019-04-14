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

// ContactMembership is a type for contact_membership db element.
type ContactMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Contact   uuid.UUID
}

// ContactMembershipManager manages queries returning a contactMembership or list of contactMemberships.
// It is configured with a db field to contain the db driver.
type ContactMembershipManager struct {
	db *sql.DB
}

// NewContactMembershipManager creates a contactMembership manager
func NewContactMembershipManager(db *sql.DB) *ContactMembershipManager {
	return &ContactMembershipManager{db: db}
}

// CRUD Methods for the ContactMembershipManager.

// CreateContactMembership creates a contactMembership.
func (m *ContactMembershipManager) CreateContactMembership(ctx context.Context, item *v1.CreateContactMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO contact_membership (cause, contact) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.Contact).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ContactMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ContactMembership-> "+err.Error())
	}
	return &id, nil
}

// GetContactMembership gets a single contactMembership from the database by ID.
func (m *ContactMembershipManager) GetContactMembership(ctx context.Context, id string) (*ContactMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ContactMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, contact FROM contact_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ContactMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ContactMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%s' is not found", id))
	}

	// scan ContactMembership data into protobuf model
	var contactMembership ContactMembership

	if err := rows.Scan(&contactMembership.ID, &contactMembership.CreatedAt, &contactMembership.UpdatedAt, &contactMembership.Cause, &contactMembership.Contact); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ContactMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ContactMembership rows with ID='%s'",
			id))
	}
	return &contactMembership, nil
}

// ListContactMembership returns a slice of all contactMemberships meeting the filter criteria.
func (m *ContactMembershipManager) ListContactMembership(ctx context.Context, filters []*v1.ContactMembershipFilterRule, orderings []*v1.ContactMembershipOrdering, limit int64) ([]*ContactMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ContactMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildContactMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ContactMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ContactMembership{}
	for rows.Next() {
		contactMembership := new(ContactMembership)
		if err := rows.Scan(&contactMembership.ID, &contactMembership.CreatedAt, &contactMembership.UpdatedAt, &contactMembership.Cause, &contactMembership.Contact); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ContactMembership row-> "+err.Error())
		}
		list = append(list, contactMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ContactMembership-> "+err.Error())
	}
	return list, nil
}

// UpdateContactMembership runs an update query on the provided db and returns the rows affected as an int64.
func (m *ContactMembershipManager) UpdateContactMembership(ctx context.Context, item *v1.ContactMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE contact_membership SET cause=$2, contact=$3 WHERE id=$1",
		item.ID, item.Cause, item.Contact)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ContactMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteContactMembership creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ContactMembershipManager) DeleteContactMembership(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM contactMembership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ContactMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ContactMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToContactMembershipProto accepts a contactMembership struct and returns a protobuf contactMembership struct.
func convertToContactMembershipProto(c *ContactMembership) *v1.ContactMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ContactMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		Contact:   c.Contact.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a contactMembership.
func (*ContactMembershipManager) GetProtoList(l []*ContactMembership) []*v1.ContactMembership {
	list := []*v1.ContactMembership{}
	for _, v := range l {
		list = append(list, convertToContactMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a contactMembership.
func (*ContactMembershipManager) GetProto(c *ContactMembership) *v1.ContactMembership {
	return convertToContactMembershipProto(c)
}

// BuildContactMembershipListQuery takes a filter and ordering object for a contactMembership.
// and returns an SQL string
func BuildContactMembershipListQuery(filters []*v1.ContactMembershipFilterRule, orderings []*v1.ContactMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, contact FROM contact_membership"
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
func (m *ContactMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
