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

// Payment is a type for payment db element.
type Payment struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CustomerOrder uuid.UUID
}

// PaymentManager manages queries returning a payment or list of payments.
// It is configured with a db field to contain the db driver.
type PaymentManager struct {
	db *sql.DB
}

// NewPaymentManager creates a payment manager
func NewPaymentManager(db *sql.DB) *PaymentManager {
	return &PaymentManager{db: db}
}

// CRUD Methods for the PaymentManager.

// CreatePayment creates a payment.
func (m *PaymentManager) CreatePayment(ctx context.Context, item *v1.CreatePayment) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO payment (customer_order) VALUES($1)  RETURNING id;",
		item.CustomerOrder).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Payment-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Payment-> "+err.Error())
	}
	return &id, nil
}

// GetPayment gets a single payment from the database by ID.
func (m *PaymentManager) GetPayment(ctx context.Context, id string) (*Payment, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Payment by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order FROM payment WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Payment-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Payment-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%s' is not found", id))
	}

	// scan Payment data into protobuf model
	var payment Payment

	if err := rows.Scan(&payment.ID, &payment.CreatedAt, &payment.UpdatedAt, &payment.CustomerOrder); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Payment row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Payment rows with ID='%s'",
			id))
	}
	return &payment, nil
}

// ListPayment returns a slice of all payments meeting the filter criteria.
func (m *PaymentManager) ListPayment(ctx context.Context, filters []*v1.PaymentFilterRule, orderings []*v1.PaymentOrdering, limit int64) ([]*Payment, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Payment Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPaymentListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Payment-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Payment{}
	for rows.Next() {
		payment := new(Payment)
		if err := rows.Scan(&payment.ID, &payment.CreatedAt, &payment.UpdatedAt, &payment.CustomerOrder); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Payment row-> "+err.Error())
		}
		list = append(list, payment)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Payment-> "+err.Error())
	}
	return list, nil
}

// UpdatePayment runs an update query on the provided db and returns the rows affected as an int64.
func (m *PaymentManager) UpdatePayment(ctx context.Context, item *v1.Payment) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE payment SET customer_order=$2 WHERE id=$1",
		item.ID, item.CustomerOrder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Payment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeletePayment creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PaymentManager) DeletePayment(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM payment WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Payment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Payment with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPaymentProto accepts a payment struct and returns a protobuf payment struct.
func convertToPaymentProto(c *Payment) *v1.Payment {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Payment{
		ID:            c.ID.String(),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		CustomerOrder: c.CustomerOrder.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a payment.
func (*PaymentManager) GetProtoList(l []*Payment) []*v1.Payment {
	list := []*v1.Payment{}
	for _, v := range l {
		list = append(list, convertToPaymentProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a payment.
func (*PaymentManager) GetProto(c *Payment) *v1.Payment {
	return convertToPaymentProto(c)
}

// BuildPaymentListQuery takes a filter and ordering object for a payment.
// and returns an SQL string
func BuildPaymentListQuery(filters []*v1.PaymentFilterRule, orderings []*v1.PaymentOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, customer_order FROM payment"
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
func (m *PaymentManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
