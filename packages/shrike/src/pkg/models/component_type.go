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

// ComponentType is a type for component_type db element.
type ComponentType struct {
	ID                      uuid.UUID
	CreatedAt               time.Time
	UpdatedAt               time.Time
	Title                   string
	ComponentImplementation uuid.UUID
}

// ComponentTypeManager manages queries returning a componentType or list of componentTypes.
// It is configured with a db field to contain the db driver.
type ComponentTypeManager struct {
	db *sql.DB
}

// NewComponentTypeManager creates a componentType manager
func NewComponentTypeManager(db *sql.DB) *ComponentTypeManager {
	return &ComponentTypeManager{db: db}
}

// CRUD Methods for the ComponentTypeManager.

// Create creates a componentType.
func (m *ComponentTypeManager) Create(ctx context.Context, item *v1.CreateComponentType) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO component_type (title, component_implementation) VALUES($1, $2)  RETURNING id;",
		item.Title, item.ComponentImplementation).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ComponentType-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ComponentType-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single componentType from the database by ID.
func (m *ComponentTypeManager) Get(ctx context.Context, id string) (*ComponentType, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ComponentType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, component_implementation FROM component_type WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentType with ID='%s' is not found", id))
	}

	// scan ComponentType data into protobuf model
	var componentType ComponentType

	if err := rows.Scan(&componentType.ID, &componentType.CreatedAt, &componentType.UpdatedAt, &componentType.Title, &componentType.ComponentImplementation); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ComponentType rows with ID='%s'",
			id))
	}
	return &componentType, nil
}

// List returns a slice of all componentTypes meeting the filter criteria.
func (m *ComponentTypeManager) List(ctx context.Context, filters []*v1.ComponentTypeFilterRule, orderings []*v1.ComponentTypeOrdering, limit int64) ([]*ComponentType, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ComponentType Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildComponentTypeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentType-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ComponentType{}
	for rows.Next() {
		componentType := new(ComponentType)
		if err := rows.Scan(&componentType.ID, &componentType.CreatedAt, &componentType.UpdatedAt, &componentType.Title, &componentType.ComponentImplementation); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentType row-> "+err.Error())
		}
		list = append(list, componentType)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentType-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *ComponentTypeManager) Update(ctx context.Context, item *v1.ComponentType) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE component_type SET title=$2, component_implementation=$3 WHERE id=$1",
		item.ID, item.Title, item.ComponentImplementation)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ComponentType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentType with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ComponentTypeManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM component_type WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ComponentType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentType with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToComponentTypeProto accepts a componentType struct and returns a protobuf componentType struct.
func convertToComponentTypeProto(c *ComponentType) *v1.ComponentType {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ComponentType{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		ComponentImplementation: c.ComponentImplementation.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a componentType.
func (*ComponentTypeManager) GetProtoList(l []*ComponentType) []*v1.ComponentType {
	list := []*v1.ComponentType{}
	for _, v := range l {
		list = append(list, convertToComponentTypeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a componentType.
func (*ComponentTypeManager) GetProto(c *ComponentType) *v1.ComponentType {
	return convertToComponentTypeProto(c)
}

// BuildComponentTypeListQuery takes a filter and ordering object for a componentType.
// and returns an SQL string
func BuildComponentTypeListQuery(filters []*v1.ComponentTypeFilterRule, orderings []*v1.ComponentTypeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, component_implementation FROM component_type"
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
func (m *ComponentTypeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
