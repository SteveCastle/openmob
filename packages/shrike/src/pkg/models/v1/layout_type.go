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

// LayoutType is a type for layout_type db element.
type LayoutType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// LayoutTypeManager manages queries returning a layoutType or list of layoutTypes.
// It is configured with a db field to contain the db driver.
type LayoutTypeManager struct {
	db *sql.DB
}

// NewLayoutTypeManager creates a layoutType manager
func NewLayoutTypeManager(db *sql.DB) *LayoutTypeManager {
	return &LayoutTypeManager{db: db}
}

// CRUD Methods for the LayoutTypeManager.

// CreateLayoutType creates a layoutType.
func (m *LayoutTypeManager) CreateLayoutType(ctx context.Context, item *v1.CreateLayoutType) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO layout_type (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LayoutType-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LayoutType-> "+err.Error())
	}
	return &id, nil
}

// GetLayoutType gets a single layoutType from the database by ID.
func (m *LayoutTypeManager) GetLayoutType(ctx context.Context, id string) (*LayoutType, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query LayoutType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM layout_type WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%s' is not found", id))
	}

	// scan LayoutType data into protobuf model
	var layoutType LayoutType

	if err := rows.Scan(&layoutType.ID, &layoutType.CreatedAt, &layoutType.UpdatedAt, &layoutType.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LayoutType rows with ID='%s'",
			id))
	}
	return &layoutType, nil
}

// ListLayoutType returns a slice of all layoutTypes meeting the filter criteria.
func (m *LayoutTypeManager) ListLayoutType(ctx context.Context, filters []*v1.LayoutTypeFilterRule, orderings []*v1.LayoutTypeOrdering, limit int64) ([]*LayoutType, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in LayoutType Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildLayoutTypeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LayoutType-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*LayoutType{}
	for rows.Next() {
		layoutType := new(LayoutType)
		if err := rows.Scan(&layoutType.ID, &layoutType.CreatedAt, &layoutType.UpdatedAt, &layoutType.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LayoutType row-> "+err.Error())
		}
		list = append(list, layoutType)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LayoutType-> "+err.Error())
	}
	return list, nil
}

// UpdateLayoutType runs an update query on the provided db and returns the rows affected as an int64.
func (m *LayoutTypeManager) UpdateLayoutType(ctx context.Context, item *v1.LayoutType) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE layout_type SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LayoutType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteLayoutType creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *LayoutTypeManager) DeleteLayoutType(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM layoutType WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LayoutType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LayoutType with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToLayoutTypeProto accepts a layoutType struct and returns a protobuf layoutType struct.
func convertToLayoutTypeProto(c *LayoutType) *v1.LayoutType {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.LayoutType{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a layoutType.
func (*LayoutTypeManager) GetProtoList(l []*LayoutType) []*v1.LayoutType {
	list := []*v1.LayoutType{}
	for _, v := range l {
		list = append(list, convertToLayoutTypeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a layoutType.
func (*LayoutTypeManager) GetProto(c *LayoutType) *v1.LayoutType {
	return convertToLayoutTypeProto(c)
}

// BuildLayoutTypeListQuery takes a filter and ordering object for a layoutType.
// and returns an SQL string
func BuildLayoutTypeListQuery(filters []*v1.LayoutTypeFilterRule, orderings []*v1.LayoutTypeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM layout_type"
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
func (m *LayoutTypeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
