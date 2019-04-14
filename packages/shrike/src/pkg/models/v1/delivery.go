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

// Delivery is a type for delivery db element.
type Delivery struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// DeliveryManager manages queries returning a delivery or list of deliverys.
// It is configured with a db field to contain the db driver.
type DeliveryManager struct {
	db *sql.DB
}

// NewDeliveryManager creates a delivery manager
func NewDeliveryManager(db *sql.DB) *DeliveryManager {
	return &DeliveryManager{db: db}
}

// CRUD Methods for the DeliveryManager.

// CreateDelivery creates a delivery.
func (m *DeliveryManager) CreateDelivery(ctx context.Context, item *v1.CreateDelivery) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO delivery () VALUES()  RETURNING id;").Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Delivery-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Delivery-> "+err.Error())
	}
	return &id, nil
}

// GetDelivery gets a single delivery from the database by ID.
func (m *DeliveryManager) GetDelivery(ctx context.Context, id string) (*Delivery, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Delivery by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM delivery WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Delivery-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Delivery-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%s' is not found", id))
	}

	// scan Delivery data into protobuf model
	var delivery Delivery

	if err := rows.Scan(&delivery.ID, &delivery.CreatedAt, &delivery.UpdatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Delivery row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Delivery rows with ID='%s'",
			id))
	}
	return &delivery, nil
}

// ListDelivery returns a slice of all deliverys meeting the filter criteria.
func (m *DeliveryManager) ListDelivery(ctx context.Context, filters []*v1.DeliveryFilterRule, orderings []*v1.DeliveryOrdering, limit int64) ([]*Delivery, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Delivery Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildDeliveryListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Delivery-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Delivery{}
	for rows.Next() {
		delivery := new(Delivery)
		if err := rows.Scan(&delivery.ID, &delivery.CreatedAt, &delivery.UpdatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Delivery row-> "+err.Error())
		}
		list = append(list, delivery)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Delivery-> "+err.Error())
	}
	return list, nil
}

// UpdateDelivery runs an update query on the provided db and returns the rows affected as an int64.
func (m *DeliveryManager) UpdateDelivery(ctx context.Context, item *v1.Delivery) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE delivery SET  WHERE id=$1",
		item.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Delivery-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteDelivery creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *DeliveryManager) DeleteDelivery(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM delivery WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Delivery-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Delivery with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToDeliveryProto accepts a delivery struct and returns a protobuf delivery struct.
func convertToDeliveryProto(c *Delivery) *v1.Delivery {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Delivery{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a delivery.
func (*DeliveryManager) GetProtoList(l []*Delivery) []*v1.Delivery {
	list := []*v1.Delivery{}
	for _, v := range l {
		list = append(list, convertToDeliveryProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a delivery.
func (*DeliveryManager) GetProto(c *Delivery) *v1.Delivery {
	return convertToDeliveryProto(c)
}

// BuildDeliveryListQuery takes a filter and ordering object for a delivery.
// and returns an SQL string
func BuildDeliveryListQuery(filters []*v1.DeliveryFilterRule, orderings []*v1.DeliveryOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at FROM delivery"
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
func (m *DeliveryManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
