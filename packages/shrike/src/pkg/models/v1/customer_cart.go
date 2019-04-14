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

// CustomerCart is a type for customer_cart db element.
type CustomerCart struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CustomerCartManager manages queries returning a customerCart or list of customerCarts.
// It is configured with a db field to contain the db driver.
type CustomerCartManager struct {
	db *sql.DB
}

// NewCustomerCartManager creates a customerCart manager
func NewCustomerCartManager(db *sql.DB) *CustomerCartManager {
	return &CustomerCartManager{db: db}
}

// CRUD Methods for the CustomerCartManager.

// CreateCustomerCart creates a customerCart.
func (m *CustomerCartManager) CreateCustomerCart(ctx context.Context, item *v1.CreateCustomerCart) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO customer_cart () VALUES()  RETURNING id;").Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into CustomerCart-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created CustomerCart-> "+err.Error())
	}
	return &id, nil
}

// GetCustomerCart gets a single customerCart from the database by ID.
func (m *CustomerCartManager) GetCustomerCart(ctx context.Context, id string) (*CustomerCart, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query CustomerCart by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM customer_cart WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerCart-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerCart-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerCart with ID='%s' is not found", id))
	}

	// scan CustomerCart data into protobuf model
	var customerCart CustomerCart

	if err := rows.Scan(&customerCart.ID, &customerCart.CreatedAt, &customerCart.UpdatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerCart row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple CustomerCart rows with ID='%s'",
			id))
	}
	return &customerCart, nil
}

// ListCustomerCart returns a slice of all customerCarts meeting the filter criteria.
func (m *CustomerCartManager) ListCustomerCart(ctx context.Context, filters []*v1.CustomerCartFilterRule, orderings []*v1.CustomerCartOrdering, limit int64) ([]*CustomerCart, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in CustomerCart Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildCustomerCartListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerCart-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*CustomerCart{}
	for rows.Next() {
		customerCart := new(CustomerCart)
		if err := rows.Scan(&customerCart.ID, &customerCart.CreatedAt, &customerCart.UpdatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerCart row-> "+err.Error())
		}
		list = append(list, customerCart)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerCart-> "+err.Error())
	}
	return list, nil
}

// UpdateCustomerCart runs an update query on the provided db and returns the rows affected as an int64.
func (m *CustomerCartManager) UpdateCustomerCart(ctx context.Context, item *v1.CustomerCart) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE customer_cart SET  WHERE id=$1",
		item.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update CustomerCart-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerCart with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteCustomerCart creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *CustomerCartManager) DeleteCustomerCart(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM customerCart WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete CustomerCart-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerCart with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToCustomerCartProto accepts a customerCart struct and returns a protobuf customerCart struct.
func convertToCustomerCartProto(c *CustomerCart) *v1.CustomerCart {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.CustomerCart{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a customerCart.
func (*CustomerCartManager) GetProtoList(l []*CustomerCart) []*v1.CustomerCart {
	list := []*v1.CustomerCart{}
	for _, v := range l {
		list = append(list, convertToCustomerCartProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a customerCart.
func (*CustomerCartManager) GetProto(c *CustomerCart) *v1.CustomerCart {
	return convertToCustomerCartProto(c)
}

// BuildCustomerCartListQuery takes a filter and ordering object for a customerCart.
// and returns an SQL string
func BuildCustomerCartListQuery(filters []*v1.CustomerCartFilterRule, orderings []*v1.CustomerCartOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at FROM customer_cart"
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
func (m *CustomerCartManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
