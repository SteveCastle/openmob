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

// Account is a type for account db element.
type Account struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
}

// AccountManager manages queries returning a account or list of accounts.
// It is configured with a db field to contain the db driver.
type AccountManager struct {
	db *sql.DB
}

// NewAccountManager creates a account manager
func NewAccountManager(db *sql.DB) *AccountManager {
	return &AccountManager{db: db}
}

// CRUD Methods for the AccountManager.

// Create creates a account.
func (m *AccountManager) Create(ctx context.Context, item *v1.CreateAccount) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO account (username) VALUES($1)  RETURNING id;",
		item.Username).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Account-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Account-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single account from the database by ID.
func (m *AccountManager) Get(ctx context.Context, id string) (*Account, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Account by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, username FROM account WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Account-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Account-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%s' is not found", id))
	}

	// scan Account data into protobuf model
	var account Account

	if err := rows.Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt, &account.Username); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Account row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Account rows with ID='%s'",
			id))
	}
	return &account, nil
}

// List returns a slice of all accounts meeting the filter criteria.
func (m *AccountManager) List(ctx context.Context, filters []*v1.AccountFilterRule, orderings []*v1.AccountOrdering, limit int64) ([]*Account, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Account Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildAccountListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Account-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Account{}
	for rows.Next() {
		account := new(Account)
		if err := rows.Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt, &account.Username); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Account row-> "+err.Error())
		}
		list = append(list, account)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Account-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *AccountManager) Update(ctx context.Context, item *v1.Account) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE account SET username=$2 WHERE id=$1",
		item.ID, item.Username)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Account-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *AccountManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM account WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Account-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToAccountProto accepts a account struct and returns a protobuf account struct.
func convertToAccountProto(c *Account) *v1.Account {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Account{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Username:  c.Username,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a account.
func (*AccountManager) GetProtoList(l []*Account) []*v1.Account {
	list := []*v1.Account{}
	for _, v := range l {
		list = append(list, convertToAccountProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a account.
func (*AccountManager) GetProto(c *Account) *v1.Account {
	return convertToAccountProto(c)
}

// BuildAccountListQuery takes a filter and ordering object for a account.
// and returns an SQL string
func BuildAccountListQuery(filters []*v1.AccountFilterRule, orderings []*v1.AccountOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, username FROM account"
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
func (m *AccountManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
