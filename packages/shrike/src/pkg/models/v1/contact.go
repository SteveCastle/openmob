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

// Contact is a type for contact db element.
type Contact struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FirstName   sql.NullString
	MiddleName  sql.NullString
	LastName    sql.NullString
	Email       sql.NullString
	PhoneNumber sql.NullString
}

// ContactManager manages queries returning a contact or list of contacts.
// It is configured with a db field to contain the db driver.
type ContactManager struct {
	db *sql.DB
}

// NewContactManager creates a contact manager
func NewContactManager(db *sql.DB) *ContactManager {
	return &ContactManager{db: db}
}

// CRUD Methods for the ContactManager.

// CreateContact creates a contact.
func (m *ContactManager) CreateContact(ctx context.Context, item *v1.CreateContact) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO contact (first_name, middle_name, last_name, email, phone_number) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		item.FirstName, item.MiddleName, item.LastName, item.Email, item.PhoneNumber).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Contact-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Contact-> "+err.Error())
	}
	return &id, nil
}

// GetContact gets a single contact from the database by ID.
func (m *ContactManager) GetContact(ctx context.Context, id string) (*Contact, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Contact by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, first_name, middle_name, last_name, email, phone_number FROM contact WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Contact-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Contact-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%s' is not found", id))
	}

	// scan Contact data into protobuf model
	var contact Contact

	if err := rows.Scan(&contact.ID, &contact.CreatedAt, &contact.UpdatedAt, &contact.FirstName, &contact.MiddleName, &contact.LastName, &contact.Email, &contact.PhoneNumber); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Contact row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Contact rows with ID='%s'",
			id))
	}
	return &contact, nil
}

// ListContact returns a slice of all contacts meeting the filter criteria.
func (m *ContactManager) ListContact(ctx context.Context, filters []*v1.ContactFilterRule, orderings []*v1.ContactOrdering, limit int64) ([]*Contact, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Contact Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildContactListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Contact-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Contact{}
	for rows.Next() {
		contact := new(Contact)
		if err := rows.Scan(&contact.ID, &contact.CreatedAt, &contact.UpdatedAt, &contact.FirstName, &contact.MiddleName, &contact.LastName, &contact.Email, &contact.PhoneNumber); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Contact row-> "+err.Error())
		}
		list = append(list, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Contact-> "+err.Error())
	}
	return list, nil
}

// UpdateContact runs an update query on the provided db and returns the rows affected as an int64.
func (m *ContactManager) UpdateContact(ctx context.Context, item *v1.Contact) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE contact SET first_name=$2, middle_name=$3, last_name=$4, email=$5, phone_number=$6 WHERE id=$1",
		item.ID, item.FirstName, item.MiddleName, item.LastName, item.Email, item.PhoneNumber)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Contact-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteContact creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ContactManager) DeleteContact(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM contact WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Contact-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Contact with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToContactProto accepts a contact struct and returns a protobuf contact struct.
func convertToContactProto(c *Contact) *v1.Contact {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Contact{
		ID:          c.ID.String(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		FirstName:   *safeNullString(c.FirstName),
		MiddleName:  *safeNullString(c.MiddleName),
		LastName:    *safeNullString(c.LastName),
		Email:       *safeNullString(c.Email),
		PhoneNumber: *safeNullString(c.PhoneNumber),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a contact.
func (*ContactManager) GetProtoList(l []*Contact) []*v1.Contact {
	list := []*v1.Contact{}
	for _, v := range l {
		list = append(list, convertToContactProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a contact.
func (*ContactManager) GetProto(c *Contact) *v1.Contact {
	return convertToContactProto(c)
}

// BuildContactListQuery takes a filter and ordering object for a contact.
// and returns an SQL string
func BuildContactListQuery(filters []*v1.ContactFilterRule, orderings []*v1.ContactOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, first_name, middle_name, last_name, email, phone_number FROM contact"
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
func (m *ContactManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
