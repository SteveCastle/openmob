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

// MailingAddress is a type for mailing_address db element.
type MailingAddress struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	StreetAddress string
	City          string
	State         string
	ZipCode       string
}

// MailingAddressManager manages queries returning a mailingAddress or list of mailingAddresss.
// It is configured with a db field to contain the db driver.
type MailingAddressManager struct {
	db *sql.DB
}

// NewMailingAddressManager creates a mailingAddress manager
func NewMailingAddressManager(db *sql.DB) *MailingAddressManager {
	return &MailingAddressManager{db: db}
}

// CRUD Methods for the MailingAddressManager.

// Create creates a mailingAddress.
func (m *MailingAddressManager) Create(ctx context.Context, item *v1.CreateMailingAddress) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO mailing_address (street_address, city, state, zip_code) VALUES($1, $2, $3, $4)  RETURNING id;",
		item.StreetAddress, item.City, item.State, item.ZipCode).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into MailingAddress-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created MailingAddress-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single mailingAddress from the database by ID.
func (m *MailingAddressManager) Get(ctx context.Context, id string) (*MailingAddress, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query MailingAddress by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, street_address, city, state, zip_code FROM mailing_address WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from MailingAddress-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from MailingAddress-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("MailingAddress with ID='%s' is not found", id))
	}

	// scan MailingAddress data into protobuf model
	var mailingAddress MailingAddress

	if err := rows.Scan(&mailingAddress.ID, &mailingAddress.CreatedAt, &mailingAddress.UpdatedAt, &mailingAddress.StreetAddress, &mailingAddress.City, &mailingAddress.State, &mailingAddress.ZipCode); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from MailingAddress row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple MailingAddress rows with ID='%s'",
			id))
	}
	return &mailingAddress, nil
}

// List returns a slice of all mailingAddresss meeting the filter criteria.
func (m *MailingAddressManager) List(ctx context.Context, filters []*v1.MailingAddressFilterRule, orderings []*v1.MailingAddressOrdering, limit int64) ([]*MailingAddress, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in MailingAddress Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildMailingAddressListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from MailingAddress-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*MailingAddress{}
	for rows.Next() {
		mailingAddress := new(MailingAddress)
		if err := rows.Scan(&mailingAddress.ID, &mailingAddress.CreatedAt, &mailingAddress.UpdatedAt, &mailingAddress.StreetAddress, &mailingAddress.City, &mailingAddress.State, &mailingAddress.ZipCode); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from MailingAddress row-> "+err.Error())
		}
		list = append(list, mailingAddress)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from MailingAddress-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *MailingAddressManager) Update(ctx context.Context, item *v1.MailingAddress) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE mailing_address SET street_address=$2, city=$3, state=$4, zip_code=$5 WHERE id=$1",
		item.ID, item.StreetAddress, item.City, item.State, item.ZipCode)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update MailingAddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("MailingAddress with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *MailingAddressManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM mailing_address WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete MailingAddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("MailingAddress with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToMailingAddressProto accepts a mailingAddress struct and returns a protobuf mailingAddress struct.
func convertToMailingAddressProto(c *MailingAddress) *v1.MailingAddress {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.MailingAddress{
		ID:            c.ID.String(),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		StreetAddress: c.StreetAddress,
		City:          c.City,
		State:         c.State,
		ZipCode:       c.ZipCode,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a mailingAddress.
func (*MailingAddressManager) GetProtoList(l []*MailingAddress) []*v1.MailingAddress {
	list := []*v1.MailingAddress{}
	for _, v := range l {
		list = append(list, convertToMailingAddressProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a mailingAddress.
func (*MailingAddressManager) GetProto(c *MailingAddress) *v1.MailingAddress {
	return convertToMailingAddressProto(c)
}

// BuildMailingAddressListQuery takes a filter and ordering object for a mailingAddress.
// and returns an SQL string
func BuildMailingAddressListQuery(filters []*v1.MailingAddressFilterRule, orderings []*v1.MailingAddressOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, street_address, city, state, zip_code FROM mailing_address"
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
func (m *MailingAddressManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
