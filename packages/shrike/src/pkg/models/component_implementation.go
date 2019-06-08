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

// ComponentImplementation is a type for component_implementation db element.
type ComponentImplementation struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Title         string
	Path          string
	ComponentType uuid.UUID
}

// ComponentImplementationManager manages queries returning a componentImplementation or list of componentImplementations.
// It is configured with a db field to contain the db driver.
type ComponentImplementationManager struct {
	db *sql.DB
}

// NewComponentImplementationManager creates a componentImplementation manager
func NewComponentImplementationManager(db *sql.DB) *ComponentImplementationManager {
	return &ComponentImplementationManager{db: db}
}

// CRUD Methods for the ComponentImplementationManager.

// Create creates a componentImplementation.
func (m *ComponentImplementationManager) Create(ctx context.Context, item *v1.CreateComponentImplementation) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO component_implementation (title, path, component_type) VALUES($1, $2, $3)  RETURNING id;",
		item.Title, item.Path, item.ComponentType).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ComponentImplementation-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ComponentImplementation-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single componentImplementation from the database by ID.
func (m *ComponentImplementationManager) Get(ctx context.Context, id string) (*ComponentImplementation, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ComponentImplementation by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, path, component_type FROM component_implementation WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentImplementation-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentImplementation-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%s' is not found", id))
	}

	// scan ComponentImplementation data into protobuf model
	var componentImplementation ComponentImplementation

	if err := rows.Scan(&componentImplementation.ID, &componentImplementation.CreatedAt, &componentImplementation.UpdatedAt, &componentImplementation.Title, &componentImplementation.Path, &componentImplementation.ComponentType); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentImplementation row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ComponentImplementation rows with ID='%s'",
			id))
	}
	return &componentImplementation, nil
}

// List returns a slice of all componentImplementations meeting the filter criteria.
func (m *ComponentImplementationManager) List(ctx context.Context, filters []*v1.ComponentImplementationFilterRule, orderings []*v1.ComponentImplementationOrdering, limit int64) ([]*ComponentImplementation, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ComponentImplementation Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildComponentImplementationListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ComponentImplementation-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ComponentImplementation{}
	for rows.Next() {
		componentImplementation := new(ComponentImplementation)
		if err := rows.Scan(&componentImplementation.ID, &componentImplementation.CreatedAt, &componentImplementation.UpdatedAt, &componentImplementation.Title, &componentImplementation.Path, &componentImplementation.ComponentType); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ComponentImplementation row-> "+err.Error())
		}
		list = append(list, componentImplementation)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ComponentImplementation-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *ComponentImplementationManager) Update(ctx context.Context, item *v1.ComponentImplementation) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE component_implementation SET title=$2, path=$3, component_type=$4 WHERE id=$1",
		item.ID, item.Title, item.Path, item.ComponentType)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ComponentImplementation-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ComponentImplementationManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM component_implementation WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ComponentImplementation-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ComponentImplementation with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToComponentImplementationProto accepts a componentImplementation struct and returns a protobuf componentImplementation struct.
func convertToComponentImplementationProto(c *ComponentImplementation) *v1.ComponentImplementation {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ComponentImplementation{
		ID:            c.ID.String(),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		Title:         c.Title,
		Path:          c.Path,
		ComponentType: c.ComponentType.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a componentImplementation.
func (*ComponentImplementationManager) GetProtoList(l []*ComponentImplementation) []*v1.ComponentImplementation {
	list := []*v1.ComponentImplementation{}
	for _, v := range l {
		list = append(list, convertToComponentImplementationProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a componentImplementation.
func (*ComponentImplementationManager) GetProto(c *ComponentImplementation) *v1.ComponentImplementation {
	return convertToComponentImplementationProto(c)
}

// BuildComponentImplementationListQuery takes a filter and ordering object for a componentImplementation.
// and returns an SQL string
func BuildComponentImplementationListQuery(filters []*v1.ComponentImplementationFilterRule, orderings []*v1.ComponentImplementationOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, path, component_type FROM component_implementation"
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
func (m *ComponentImplementationManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
