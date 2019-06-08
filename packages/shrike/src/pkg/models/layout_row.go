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

// LayoutRow is a type for layout_row db element.
type LayoutRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Layout    uuid.UUID
	Container sql.NullBool
	Weight    sql.NullInt64
}

// LayoutRowManager manages queries returning a layoutRow or list of layoutRows.
// It is configured with a db field to contain the db driver.
type LayoutRowManager struct {
	db *sql.DB
}

// NewLayoutRowManager creates a layoutRow manager
func NewLayoutRowManager(db *sql.DB) *LayoutRowManager {
	return &LayoutRowManager{db: db}
}

// CRUD Methods for the LayoutRowManager.

// Create creates a layoutRow.
func (m *LayoutRowManager) Create(ctx context.Context, item *v1.CreateLayoutRow) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO layout_row (layout, container, weight) VALUES($1, $2, $3)  RETURNING id;",
		item.Layout, item.Container, item.Weight).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LayoutRow-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LayoutRow-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single layoutRow from the database by ID.
func (m *LayoutRowManager) Get(ctx context.Context, id string) (*LayoutRow, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query LayoutRow by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, layout, container, weight FROM layout_row WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutRow-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutRow-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%s' is not found", id))
	}

	// scan LayoutRow data into protobuf model
	var layoutRow LayoutRow

	if err := rows.Scan(&layoutRow.ID, &layoutRow.CreatedAt, &layoutRow.UpdatedAt, &layoutRow.Layout, &layoutRow.Container, &layoutRow.Weight); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutRow row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutRow rows with ID='%s'",
			id))
	}
	return &layoutRow, nil
}

// List returns a slice of all layoutRows meeting the filter criteria.
func (m *LayoutRowManager) List(ctx context.Context, filters []*v1.LayoutRowFilterRule, orderings []*v1.LayoutRowOrdering, limit int64) ([]*LayoutRow, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in LayoutRow Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildLayoutRowListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutRow-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*LayoutRow{}
	for rows.Next() {
		layoutRow := new(LayoutRow)
		if err := rows.Scan(&layoutRow.ID, &layoutRow.CreatedAt, &layoutRow.UpdatedAt, &layoutRow.Layout, &layoutRow.Container, &layoutRow.Weight); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutRow row-> "+err.Error())
		}
		list = append(list, layoutRow)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutRow-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *LayoutRowManager) Update(ctx context.Context, item *v1.LayoutRow) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE layout_row SET layout=$2, container=$3, weight=$4 WHERE id=$1",
		item.ID, item.Layout, item.Container, item.Weight)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutRow-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *LayoutRowManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM layout_row WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LayoutRow-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutRow with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToLayoutRowProto accepts a layoutRow struct and returns a protobuf layoutRow struct.
func convertToLayoutRowProto(c *LayoutRow) *v1.LayoutRow {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.LayoutRow{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Layout:    c.Layout.String(),
		Container: *safeNullBool(c.Container),
		Weight:    *safeNullInt64(c.Weight),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a layoutRow.
func (*LayoutRowManager) GetProtoList(l []*LayoutRow) []*v1.LayoutRow {
	list := []*v1.LayoutRow{}
	for _, v := range l {
		list = append(list, convertToLayoutRowProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a layoutRow.
func (*LayoutRowManager) GetProto(c *LayoutRow) *v1.LayoutRow {
	return convertToLayoutRowProto(c)
}

// BuildLayoutRowListQuery takes a filter and ordering object for a layoutRow.
// and returns an SQL string
func BuildLayoutRowListQuery(filters []*v1.LayoutRowFilterRule, orderings []*v1.LayoutRowOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, layout, container, weight FROM layout_row"
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
func (m *LayoutRowManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
