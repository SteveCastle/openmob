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

// Purchaser is a type for purchaser db element.
type Purchaser struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CustomerOrder uuid.UUID
	Contact       uuid.UUID
	Cause         uuid.UUID
}

// PurchaserManager manages queries returning a purchaser or list of purchasers.
// It is configured with a db field to contain the db driver.
type PurchaserManager struct {
	db *sql.DB
}

// NewPurchaserManager creates a purchaser manager
func NewPurchaserManager(db *sql.DB) *PurchaserManager {
	return &PurchaserManager{db: db}
}

// CRUD Methods for the PurchaserManager.

// CreatePurchaser creates a purchaser.
func (m *PurchaserManager) CreatePurchaser(ctx context.Context, item *v1.CreatePurchaser) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO purchaser (customer_order, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		item.CustomerOrder, item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Purchaser-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Purchaser-> "+err.Error())
	}
	return &id, nil
}

// GetPurchaser gets a single purchaser from the database by ID.
func (m *PurchaserManager) GetPurchaser(ctx context.Context, id string) (*Purchaser, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Purchaser by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order, contact, cause FROM purchaser WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Purchaser-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Purchaser-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%s' is not found", id))
	}

	// scan Purchaser data into protobuf model
	var purchaser Purchaser

	if err := rows.Scan(&purchaser.ID, &purchaser.CreatedAt, &purchaser.UpdatedAt, &purchaser.CustomerOrder, &purchaser.Contact, &purchaser.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Purchaser row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Purchaser rows with ID='%s'",
			id))
	}
	return &purchaser, nil
}

// ListPurchaser returns a slice of all purchasers meeting the filter criteria.
func (m *PurchaserManager) ListPurchaser(ctx context.Context, filters []*v1.PurchaserFilterRule, orderings []*v1.PurchaserOrdering, limit int64) ([]*Purchaser, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Purchaser Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPurchaserListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Purchaser-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Purchaser{}
	for rows.Next() {
		purchaser := new(Purchaser)
		if err := rows.Scan(&purchaser.ID, &purchaser.CreatedAt, &purchaser.UpdatedAt, &purchaser.CustomerOrder, &purchaser.Contact, &purchaser.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Purchaser row-> "+err.Error())
		}
		list = append(list, purchaser)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Purchaser-> "+err.Error())
	}
	return list, nil
}

// UpdatePurchaser runs an update query on the provided db and returns the rows affected as an int64.
func (m *PurchaserManager) UpdatePurchaser(ctx context.Context, item *v1.Purchaser) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE purchaser SET customer_order=$2, contact=$3, cause=$4 WHERE id=$1",
		item.ID, item.CustomerOrder, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Purchaser-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeletePurchaser creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PurchaserManager) DeletePurchaser(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM purchaser WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Purchaser-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Purchaser with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPurchaserProto accepts a purchaser struct and returns a protobuf purchaser struct.
func convertToPurchaserProto(c *Purchaser) *v1.Purchaser {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Purchaser{
		ID:            c.ID.String(),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		CustomerOrder: c.CustomerOrder.String(),
		Contact:       c.Contact.String(),
		Cause:         c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a purchaser.
func (*PurchaserManager) GetProtoList(l []*Purchaser) []*v1.Purchaser {
	list := []*v1.Purchaser{}
	for _, v := range l {
		list = append(list, convertToPurchaserProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a purchaser.
func (*PurchaserManager) GetProto(c *Purchaser) *v1.Purchaser {
	return convertToPurchaserProto(c)
}

// BuildPurchaserListQuery takes a filter and ordering object for a purchaser.
// and returns an SQL string
func BuildPurchaserListQuery(filters []*v1.PurchaserFilterRule, orderings []*v1.PurchaserOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, customer_order, contact, cause FROM purchaser"
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
func (m *PurchaserManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
