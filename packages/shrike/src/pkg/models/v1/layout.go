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

// Layout is a type for layout db element.
type Layout struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	LayoutType uuid.NullUUID
}

// LayoutManager manages queries returning a layout or list of layouts.
// It is configured with a db field to contain the db driver.
type LayoutManager struct {
	db *sql.DB
}

// NewLayoutManager creates a layout manager
func NewLayoutManager(db *sql.DB) *LayoutManager {
	return &LayoutManager{db: db}
}

// CRUD Methods for the LayoutManager.

// CreateLayout creates a layout.
func (m *LayoutManager) CreateLayout(ctx context.Context, item *v1.CreateLayout) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO layout (layout_type) VALUES($1)  RETURNING id;",
		item.LayoutType).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Layout-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Layout-> "+err.Error())
	}
	return &id, nil
}

// GetLayout gets a single layout from the database by ID.
func (m *LayoutManager) GetLayout(ctx context.Context, id string) (*Layout, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Layout by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout_type FROM layout WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Layout-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Layout-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%s' is not found", id))
	}

	// scan Layout data into protobuf model
	var layout Layout

	if err := rows.Scan(&layout.ID, &layout.CreatedAt, &layout.UpdatedAt, &layout.LayoutType); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Layout row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Layout rows with ID='%s'",
			id))
	}
	return &layout, nil
}

// ListLayout returns a slice of all layouts meeting the filter criteria.
func (m *LayoutManager) ListLayout(ctx context.Context, filters []*v1.LayoutFilterRule, orderings []*v1.LayoutOrdering, limit int64) ([]*Layout, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Layout Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildLayoutListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Layout-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Layout{}
	for rows.Next() {
		layout := new(Layout)
		if err := rows.Scan(&layout.ID, &layout.CreatedAt, &layout.UpdatedAt, &layout.LayoutType); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Layout row-> "+err.Error())
		}
		list = append(list, layout)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Layout-> "+err.Error())
	}
	return list, nil
}

// UpdateLayout runs an update query on the provided db and returns the rows affected as an int64.
func (m *LayoutManager) UpdateLayout(ctx context.Context, item *v1.Layout) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE layout SET layout_type=$2 WHERE id=$1",
		item.ID, item.LayoutType)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Layout-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteLayout creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *LayoutManager) DeleteLayout(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM layout WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Layout-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Layout with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToLayoutProto accepts a layout struct and returns a protobuf layout struct.
func convertToLayoutProto(c *Layout) *v1.Layout {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Layout{
		ID:         c.ID.String(),
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		LayoutType: *safeNullUUID(c.LayoutType),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a layout.
func (*LayoutManager) GetProtoList(l []*Layout) []*v1.Layout {
	list := []*v1.Layout{}
	for _, v := range l {
		list = append(list, convertToLayoutProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a layout.
func (*LayoutManager) GetProto(c *Layout) *v1.Layout {
	return convertToLayoutProto(c)
}

// BuildLayoutListQuery takes a filter and ordering object for a layout.
// and returns an SQL string
func BuildLayoutListQuery(filters []*v1.LayoutFilterRule, orderings []*v1.LayoutOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, layout_type FROM layout"
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
func (m *LayoutManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
