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

// Component is a type for component db element.
type Component struct {
	ID                      uuid.UUID
	CreatedAt               time.Time
	UpdatedAt               time.Time
	ComponentType           uuid.UUID
	ComponentImplementation uuid.UUID
	LayoutColumn            uuid.NullUUID
	Weight                  sql.NullInt64
}

// ComponentManager manages queries returning a component or list of components.
// It is configured with a db field to contain the db driver.
type ComponentManager struct {
	db *sql.DB
}

// NewComponentManager creates a component manager
func NewComponentManager(db *sql.DB) *ComponentManager {
	return &ComponentManager{db: db}
}

// CRUD Methods for the ComponentManager.

// CreateComponent creates a component.
func (m *ComponentManager) CreateComponent(ctx context.Context, item *v1.CreateComponent) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO component (component_type, component_implementation, layout_column, weight) VALUES($1, $2, $3, $4)  RETURNING id;",
		item.ComponentType, item.ComponentImplementation, item.LayoutColumn, item.Weight).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Component-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Component-> "+err.Error())
	}
	return &id, nil
}

// GetComponent gets a single component from the database by ID.
func (m *ComponentManager) GetComponent(ctx context.Context, id string) (*Component, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Component by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, component_type, component_implementation, layout_column, weight FROM component WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Component-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Component-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%s' is not found", id))
	}

	// scan Component data into protobuf model
	var component Component

	if err := rows.Scan(&component.ID, &component.CreatedAt, &component.UpdatedAt, &component.ComponentType, &component.ComponentImplementation, &component.LayoutColumn, &component.Weight); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Component row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Component rows with ID='%s'",
			id))
	}
	return &component, nil
}

// ListComponent returns a slice of all components meeting the filter criteria.
func (m *ComponentManager) ListComponent(ctx context.Context, filters []*v1.ComponentFilterRule, orderings []*v1.ComponentOrdering, limit int64) ([]*Component, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Component Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildComponentListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Component-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Component{}
	for rows.Next() {
		component := new(Component)
		if err := rows.Scan(&component.ID, &component.CreatedAt, &component.UpdatedAt, &component.ComponentType, &component.ComponentImplementation, &component.LayoutColumn, &component.Weight); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Component row-> "+err.Error())
		}
		list = append(list, component)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Component-> "+err.Error())
	}
	return list, nil
}

// UpdateComponent runs an update query on the provided db and returns the rows affected as an int64.
func (m *ComponentManager) UpdateComponent(ctx context.Context, item *v1.Component) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE component SET component_type=$2, component_implementation=$3, layout_column=$4, weight=$5 WHERE id=$1",
		item.ID, item.ComponentType, item.ComponentImplementation, item.LayoutColumn, item.Weight)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Component-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteComponent creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ComponentManager) DeleteComponent(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM component WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Component-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Component with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToComponentProto accepts a component struct and returns a protobuf component struct.
func convertToComponentProto(c *Component) *v1.Component {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Component{
		ID:                      c.ID.String(),
		CreatedAt:               createdAt,
		UpdatedAt:               updatedAt,
		ComponentType:           c.ComponentType.String(),
		ComponentImplementation: c.ComponentImplementation.String(),
		LayoutColumn:            *safeNullUUID(c.LayoutColumn),
		Weight:                  *safeNullInt64(c.Weight),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a component.
func (*ComponentManager) GetProtoList(l []*Component) []*v1.Component {
	list := []*v1.Component{}
	for _, v := range l {
		list = append(list, convertToComponentProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a component.
func (*ComponentManager) GetProto(c *Component) *v1.Component {
	return convertToComponentProto(c)
}

// BuildComponentListQuery takes a filter and ordering object for a component.
// and returns an SQL string
func BuildComponentListQuery(filters []*v1.ComponentFilterRule, orderings []*v1.ComponentOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, component_type, component_implementation, layout_column, weight FROM component"
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
func (m *ComponentManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
