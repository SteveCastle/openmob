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

// LayoutColumn is a type for layout_column db element.
type LayoutColumn struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	LayoutRow uuid.UUID
	Width     int64
	Weight    sql.NullInt64
}

// LayoutColumnManager manages queries returning a layoutColumn or list of layoutColumns.
// It is configured with a db field to contain the db driver.
type LayoutColumnManager struct {
	db *sql.DB
}

// NewLayoutColumnManager creates a layoutColumn manager
func NewLayoutColumnManager(db *sql.DB) *LayoutColumnManager {
	return &LayoutColumnManager{db: db}
}

// CRUD Methods for the LayoutColumnManager.

// CreateLayoutColumn creates a layoutColumn.
func (m *LayoutColumnManager) CreateLayoutColumn(ctx context.Context, item *v1.CreateLayoutColumn) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO layout_column (layout_row, width, weight) VALUES($1, $2, $3)  RETURNING id;",
		item.LayoutRow, item.Width, item.Weight).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LayoutColumn-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LayoutColumn-> "+err.Error())
	}
	return &id, nil
}

// GetLayoutColumn gets a single layoutColumn from the database by ID.
func (m *LayoutColumnManager) GetLayoutColumn(ctx context.Context, id string) (*LayoutColumn, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query LayoutColumn by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout_row, width, weight FROM layout_column WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutColumn-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutColumn-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutColumn with ID='%s' is not found", id))
	}

	// scan LayoutColumn data into protobuf model
	var layoutColumn LayoutColumn

	if err := rows.Scan(&layoutColumn.ID, &layoutColumn.CreatedAt, &layoutColumn.UpdatedAt, &layoutColumn.LayoutRow, &layoutColumn.Width, &layoutColumn.Weight); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutColumn row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutColumn rows with ID='%s'",
			id))
	}
	return &layoutColumn, nil
}

// ListLayoutColumn returns a slice of all layoutColumns meeting the filter criteria.
func (m *LayoutColumnManager) ListLayoutColumn(ctx context.Context, filters []*v1.LayoutColumnFilterRule, orderings []*v1.LayoutColumnOrdering, limit int64) ([]*LayoutColumn, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in LayoutColumn Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildLayoutColumnListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutColumn-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*LayoutColumn{}
	for rows.Next() {
		layoutColumn := new(LayoutColumn)
		if err := rows.Scan(&layoutColumn.ID, &layoutColumn.CreatedAt, &layoutColumn.UpdatedAt, &layoutColumn.LayoutRow, &layoutColumn.Width, &layoutColumn.Weight); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutColumn row-> "+err.Error())
		}
		list = append(list, layoutColumn)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutColumn-> "+err.Error())
	}
	return list, nil
}

// UpdateLayoutColumn runs an update query on the provided db and returns the rows affected as an int64.
func (m *LayoutColumnManager) UpdateLayoutColumn(ctx context.Context, item *v1.LayoutColumn) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE layout_column SET layout_row=$2, width=$3, weight=$4 WHERE id=$1",
		item.ID, item.LayoutRow, item.Width, item.Weight)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutColumn-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutColumn with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteLayoutColumn creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *LayoutColumnManager) DeleteLayoutColumn(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM layoutColumn WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LayoutColumn-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutColumn with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToLayoutColumnProto accepts a layoutColumn struct and returns a protobuf layoutColumn struct.
func convertToLayoutColumnProto(c *LayoutColumn) *v1.LayoutColumn {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.LayoutColumn{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		LayoutRow: c.LayoutRow.String(),
		Width:     c.Width,
		Weight:    *safeNullInt64(c.Weight),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a layoutColumn.
func (*LayoutColumnManager) GetProtoList(l []*LayoutColumn) []*v1.LayoutColumn {
	list := []*v1.LayoutColumn{}
	for _, v := range l {
		list = append(list, convertToLayoutColumnProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a layoutColumn.
func (*LayoutColumnManager) GetProto(c *LayoutColumn) *v1.LayoutColumn {
	return convertToLayoutColumnProto(c)
}

// BuildLayoutColumnListQuery takes a filter and ordering object for a layoutColumn.
// and returns an SQL string
func BuildLayoutColumnListQuery(filters []*v1.LayoutColumnFilterRule, orderings []*v1.LayoutColumnOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, layout_row, width, weight FROM layout_column"
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
func (m *LayoutColumnManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
