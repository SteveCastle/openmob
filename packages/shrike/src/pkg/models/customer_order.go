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

// CustomerOrder is a type for customer_order db element.
type CustomerOrder struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CustomerCart uuid.UUID
}

// CustomerOrderManager manages queries returning a customerOrder or list of customerOrders.
// It is configured with a db field to contain the db driver.
type CustomerOrderManager struct {
	db *sql.DB
}

// NewCustomerOrderManager creates a customerOrder manager
func NewCustomerOrderManager(db *sql.DB) *CustomerOrderManager {
	return &CustomerOrderManager{db: db}
}

// CRUD Methods for the CustomerOrderManager.

// Create creates a customerOrder.
func (m *CustomerOrderManager) Create(ctx context.Context, item *v1.CreateCustomerOrder) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO customer_order (customer_cart) VALUES($1)  RETURNING id;",
		item.CustomerCart).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into CustomerOrder-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created CustomerOrder-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single customerOrder from the database by ID.
func (m *CustomerOrderManager) Get(ctx context.Context, id string) (*CustomerOrder, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query CustomerOrder by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_cart FROM customer_order WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerOrder-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerOrder-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerOrder with ID='%s' is not found", id))
	}

	// scan CustomerOrder data into protobuf model
	var customerOrder CustomerOrder

	if err := rows.Scan(&customerOrder.ID, &customerOrder.CreatedAt, &customerOrder.UpdatedAt, &customerOrder.CustomerCart); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerOrder row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple CustomerOrder rows with ID='%s'",
			id))
	}
	return &customerOrder, nil
}

// List returns a slice of all customerOrders meeting the filter criteria.
func (m *CustomerOrderManager) List(ctx context.Context, filters []*v1.CustomerOrderFilterRule, orderings []*v1.CustomerOrderOrdering, limit int64) ([]*CustomerOrder, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in CustomerOrder Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildCustomerOrderListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from CustomerOrder-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*CustomerOrder{}
	for rows.Next() {
		customerOrder := new(CustomerOrder)
		if err := rows.Scan(&customerOrder.ID, &customerOrder.CreatedAt, &customerOrder.UpdatedAt, &customerOrder.CustomerCart); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from CustomerOrder row-> "+err.Error())
		}
		list = append(list, customerOrder)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from CustomerOrder-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *CustomerOrderManager) Update(ctx context.Context, item *v1.CustomerOrder) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE customer_order SET customer_cart=$2 WHERE id=$1",
		item.ID, item.CustomerCart)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update CustomerOrder-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerOrder with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *CustomerOrderManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM customer_order WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete CustomerOrder-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("CustomerOrder with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToCustomerOrderProto accepts a customerOrder struct and returns a protobuf customerOrder struct.
func convertToCustomerOrderProto(c *CustomerOrder) *v1.CustomerOrder {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.CustomerOrder{
		ID:           c.ID.String(),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		CustomerCart: c.CustomerCart.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a customerOrder.
func (*CustomerOrderManager) GetProtoList(l []*CustomerOrder) []*v1.CustomerOrder {
	list := []*v1.CustomerOrder{}
	for _, v := range l {
		list = append(list, convertToCustomerOrderProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a customerOrder.
func (*CustomerOrderManager) GetProto(c *CustomerOrder) *v1.CustomerOrder {
	return convertToCustomerOrderProto(c)
}

// BuildCustomerOrderListQuery takes a filter and ordering object for a customerOrder.
// and returns an SQL string
func BuildCustomerOrderListQuery(filters []*v1.CustomerOrderFilterRule, orderings []*v1.CustomerOrderOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, customer_cart FROM customer_order"
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
func (m *CustomerOrderManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
