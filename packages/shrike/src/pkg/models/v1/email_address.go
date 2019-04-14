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

// EmailAddress is a type for email_address db element.
type EmailAddress struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Address   string
}

// EmailAddressManager manages queries returning a emailAddress or list of emailAddresss.
// It is configured with a db field to contain the db driver.
type EmailAddressManager struct {
	db *sql.DB
}

// NewEmailAddressManager creates a emailAddress manager
func NewEmailAddressManager(db *sql.DB) *EmailAddressManager {
	return &EmailAddressManager{db: db}
}

// CRUD Methods for the EmailAddressManager.

// CreateEmailAddress creates a emailAddress.
func (m *EmailAddressManager) CreateEmailAddress(ctx context.Context, item *v1.CreateEmailAddress) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO email_address (address) VALUES($1)  RETURNING id;",
		item.Address).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into EmailAddress-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created EmailAddress-> "+err.Error())
	}
	return &id, nil
}

// GetEmailAddress gets a single emailAddress from the database by ID.
func (m *EmailAddressManager) GetEmailAddress(ctx context.Context, id string) (*EmailAddress, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query EmailAddress by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, address FROM email_address WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EmailAddress-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from EmailAddress-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EmailAddress with ID='%s' is not found", id))
	}

	// scan EmailAddress data into protobuf model
	var emailAddress EmailAddress

	if err := rows.Scan(&emailAddress.ID, &emailAddress.CreatedAt, &emailAddress.UpdatedAt, &emailAddress.Address); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from EmailAddress row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple EmailAddress rows with ID='%s'",
			id))
	}
	return &emailAddress, nil
}

// ListEmailAddress returns a slice of all emailAddresss meeting the filter criteria.
func (m *EmailAddressManager) ListEmailAddress(ctx context.Context, filters []*v1.EmailAddressFilterRule, orderings []*v1.EmailAddressOrdering, limit int64) ([]*EmailAddress, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in EmailAddress Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildEmailAddressListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EmailAddress-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*EmailAddress{}
	for rows.Next() {
		emailAddress := new(EmailAddress)
		if err := rows.Scan(&emailAddress.ID, &emailAddress.CreatedAt, &emailAddress.UpdatedAt, &emailAddress.Address); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from EmailAddress row-> "+err.Error())
		}
		list = append(list, emailAddress)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from EmailAddress-> "+err.Error())
	}
	return list, nil
}

// UpdateEmailAddress runs an update query on the provided db and returns the rows affected as an int64.
func (m *EmailAddressManager) UpdateEmailAddress(ctx context.Context, item *v1.EmailAddress) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE email_address SET address=$2 WHERE id=$1",
		item.ID, item.Address)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update EmailAddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EmailAddress with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteEmailAddress creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *EmailAddressManager) DeleteEmailAddress(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM emailAddress WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete EmailAddress-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EmailAddress with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToEmailAddressProto accepts a emailAddress struct and returns a protobuf emailAddress struct.
func convertToEmailAddressProto(c *EmailAddress) *v1.EmailAddress {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.EmailAddress{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Address:   c.Address,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a emailAddress.
func (*EmailAddressManager) GetProtoList(l []*EmailAddress) []*v1.EmailAddress {
	list := []*v1.EmailAddress{}
	for _, v := range l {
		list = append(list, convertToEmailAddressProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a emailAddress.
func (*EmailAddressManager) GetProto(c *EmailAddress) *v1.EmailAddress {
	return convertToEmailAddressProto(c)
}

// BuildEmailAddressListQuery takes a filter and ordering object for a emailAddress.
// and returns an SQL string
func BuildEmailAddressListQuery(filters []*v1.EmailAddressFilterRule, orderings []*v1.EmailAddressOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, address FROM email_address"
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
func (m *EmailAddressManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
