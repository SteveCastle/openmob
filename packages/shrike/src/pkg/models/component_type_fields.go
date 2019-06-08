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

// ComponentTypeFields is a type for component_type_fields db element.
type ComponentTypeFields struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ComponentType uuid.UUID
	FieldType     uuid.UUID
	Weight        sql.NullInt64
	Required      sql.NullBool
}

// ComponentTypeFieldsManager manages queries returning a componentTypeFields or list of componentTypeFieldss.
// It is configured with a db field to contain the db driver.
type ComponentTypeFieldsManager struct {
	db *sql.DB
}

// NewComponentTypeFieldsManager creates a componentTypeFields manager
func NewComponentTypeFieldsManager(db *sql.DB) *ComponentTypeFieldsManager {
	return &ComponentTypeFieldsManager{db: db}
}

// CRUD Methods for the ComponentTypeFieldsManager.

// Create creates a componentTypeFields.
func (m *ComponentTypeFieldsManager) Create(ctx context.Context, item *v1.CreateComponentTypeFields) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO component_type_fields (component_type, field_type, weight, required) VALUES($1, $2, $3, $4)  RETURNING id;",
		item.ComponentType, item.FieldType, item.Weight, item.Required).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ComponentTypeFields-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ComponentTypeFields-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single componentTypeFields from the database by ID.
func (m *ComponentTypeFieldsManager) Get(ctx context.Context, id string) (*ComponentTypeFields, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ComponentTypeFields by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, component_type, field_type, weight, required FROM component_type_fields WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentTypeFields-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentTypeFields-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentTypeFields with ID='%s' is not found", id))
	}

	// scan ComponentTypeFields data into protobuf model
	var componentTypeFields ComponentTypeFields

	if err := rows.Scan(&componentTypeFields.ID, &componentTypeFields.CreatedAt, &componentTypeFields.UpdatedAt, &componentTypeFields.ComponentType, &componentTypeFields.FieldType, &componentTypeFields.Weight, &componentTypeFields.Required); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentTypeFields row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ComponentTypeFields rows with ID='%s'",
			id))
	}
	return &componentTypeFields, nil
}

// List returns a slice of all componentTypeFieldss meeting the filter criteria.
func (m *ComponentTypeFieldsManager) List(ctx context.Context, filters []*v1.ComponentTypeFieldsFilterRule, orderings []*v1.ComponentTypeFieldsOrdering, limit int64) ([]*ComponentTypeFields, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ComponentTypeFields Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildComponentTypeFieldsListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentTypeFields-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ComponentTypeFields{}
	for rows.Next() {
		componentTypeFields := new(ComponentTypeFields)
		if err := rows.Scan(&componentTypeFields.ID, &componentTypeFields.CreatedAt, &componentTypeFields.UpdatedAt, &componentTypeFields.ComponentType, &componentTypeFields.FieldType, &componentTypeFields.Weight, &componentTypeFields.Required); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentTypeFields row-> "+err.Error())
		}
		list = append(list, componentTypeFields)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentTypeFields-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *ComponentTypeFieldsManager) Update(ctx context.Context, item *v1.ComponentTypeFields) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE component_type_fields SET component_type=$2, field_type=$3, weight=$4, required=$5 WHERE id=$1",
		item.ID, item.ComponentType, item.FieldType, item.Weight, item.Required)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ComponentTypeFields-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentTypeFields with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ComponentTypeFieldsManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM component_type_fields WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ComponentTypeFields-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentTypeFields with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToComponentTypeFieldsProto accepts a componentTypeFields struct and returns a protobuf componentTypeFields struct.
func convertToComponentTypeFieldsProto(c *ComponentTypeFields) *v1.ComponentTypeFields {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ComponentTypeFields{
		ID:            c.ID.String(),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		ComponentType: c.ComponentType.String(),
		FieldType:     c.FieldType.String(),
		Weight:        *safeNullInt64(c.Weight),
		Required:      *safeNullBool(c.Required),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a componentTypeFields.
func (*ComponentTypeFieldsManager) GetProtoList(l []*ComponentTypeFields) []*v1.ComponentTypeFields {
	list := []*v1.ComponentTypeFields{}
	for _, v := range l {
		list = append(list, convertToComponentTypeFieldsProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a componentTypeFields.
func (*ComponentTypeFieldsManager) GetProto(c *ComponentTypeFields) *v1.ComponentTypeFields {
	return convertToComponentTypeFieldsProto(c)
}

// BuildComponentTypeFieldsListQuery takes a filter and ordering object for a componentTypeFields.
// and returns an SQL string
func BuildComponentTypeFieldsListQuery(filters []*v1.ComponentTypeFieldsFilterRule, orderings []*v1.ComponentTypeFieldsOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, component_type, field_type, weight, required FROM component_type_fields"
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
func (m *ComponentTypeFieldsManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
