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

// PhoneNumber is a type for phone_number db element.
type PhoneNumber struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PhoneNumber string
}

// PhoneNumberManager manages queries returning a phoneNumber or list of phoneNumbers.
// It is configured with a db field to contain the db driver.
type PhoneNumberManager struct {
	db *sql.DB
}

// NewPhoneNumberManager creates a phoneNumber manager
func NewPhoneNumberManager(db *sql.DB) *PhoneNumberManager {
	return &PhoneNumberManager{db: db}
}

// CRUD Methods for the PhoneNumberManager.

// Create creates a phoneNumber.
func (m *PhoneNumberManager) Create(ctx context.Context, item *v1.CreatePhoneNumber) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO phone_number (phone_number) VALUES($1)  RETURNING id;",
		item.PhoneNumber).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PhoneNumber-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PhoneNumber-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single phoneNumber from the database by ID.
func (m *PhoneNumberManager) Get(ctx context.Context, id string) (*PhoneNumber, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query PhoneNumber by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, phone_number FROM phone_number WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PhoneNumber-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PhoneNumber-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%s' is not found", id))
	}

	// scan PhoneNumber data into protobuf model
	var phoneNumber PhoneNumber

	if err := rows.Scan(&phoneNumber.ID, &phoneNumber.CreatedAt, &phoneNumber.UpdatedAt, &phoneNumber.PhoneNumber); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PhoneNumber row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PhoneNumber rows with ID='%s'",
			id))
	}
	return &phoneNumber, nil
}

// List returns a slice of all phoneNumbers meeting the filter criteria.
func (m *PhoneNumberManager) List(ctx context.Context, filters []*v1.PhoneNumberFilterRule, orderings []*v1.PhoneNumberOrdering, limit int64) ([]*PhoneNumber, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in PhoneNumber Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPhoneNumberListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PhoneNumber-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*PhoneNumber{}
	for rows.Next() {
		phoneNumber := new(PhoneNumber)
		if err := rows.Scan(&phoneNumber.ID, &phoneNumber.CreatedAt, &phoneNumber.UpdatedAt, &phoneNumber.PhoneNumber); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PhoneNumber row-> "+err.Error())
		}
		list = append(list, phoneNumber)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PhoneNumber-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *PhoneNumberManager) Update(ctx context.Context, item *v1.PhoneNumber) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE phone_number SET phone_number=$2 WHERE id=$1",
		item.ID, item.PhoneNumber)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PhoneNumber-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PhoneNumberManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM phone_number WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PhoneNumber-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PhoneNumber with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPhoneNumberProto accepts a phoneNumber struct and returns a protobuf phoneNumber struct.
func convertToPhoneNumberProto(c *PhoneNumber) *v1.PhoneNumber {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.PhoneNumber{
		ID:          c.ID.String(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		PhoneNumber: c.PhoneNumber,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a phoneNumber.
func (*PhoneNumberManager) GetProtoList(l []*PhoneNumber) []*v1.PhoneNumber {
	list := []*v1.PhoneNumber{}
	for _, v := range l {
		list = append(list, convertToPhoneNumberProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a phoneNumber.
func (*PhoneNumberManager) GetProto(c *PhoneNumber) *v1.PhoneNumber {
	return convertToPhoneNumberProto(c)
}

// BuildPhoneNumberListQuery takes a filter and ordering object for a phoneNumber.
// and returns an SQL string
func BuildPhoneNumberListQuery(filters []*v1.PhoneNumberFilterRule, orderings []*v1.PhoneNumberOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, phone_number FROM phone_number"
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
func (m *PhoneNumberManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
