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

// Territory is a type for territory db element.
type Territory struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// TerritoryManager manages queries returning a territory or list of territorys.
// It is configured with a db field to contain the db driver.
type TerritoryManager struct {
	db *sql.DB
}

// NewTerritoryManager creates a territory manager
func NewTerritoryManager(db *sql.DB) *TerritoryManager {
	return &TerritoryManager{db: db}
}

// CRUD Methods for the TerritoryManager.

// Create creates a territory.
func (m *TerritoryManager) Create(ctx context.Context, item *v1.CreateTerritory) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO territory (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Territory-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Territory-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single territory from the database by ID.
func (m *TerritoryManager) Get(ctx context.Context, id string) (*Territory, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Territory by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM territory WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Territory-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Territory-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%s' is not found", id))
	}

	// scan Territory data into protobuf model
	var territory Territory

	if err := rows.Scan(&territory.ID, &territory.CreatedAt, &territory.UpdatedAt, &territory.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Territory row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Territory rows with ID='%s'",
			id))
	}
	return &territory, nil
}

// List returns a slice of all territorys meeting the filter criteria.
func (m *TerritoryManager) List(ctx context.Context, filters []*v1.TerritoryFilterRule, orderings []*v1.TerritoryOrdering, limit int64) ([]*Territory, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Territory Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildTerritoryListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Territory-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Territory{}
	for rows.Next() {
		territory := new(Territory)
		if err := rows.Scan(&territory.ID, &territory.CreatedAt, &territory.UpdatedAt, &territory.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Territory row-> "+err.Error())
		}
		list = append(list, territory)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Territory-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *TerritoryManager) Update(ctx context.Context, item *v1.Territory) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE territory SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Territory-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *TerritoryManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM territory WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Territory-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Territory with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToTerritoryProto accepts a territory struct and returns a protobuf territory struct.
func convertToTerritoryProto(c *Territory) *v1.Territory {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Territory{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a territory.
func (*TerritoryManager) GetProtoList(l []*Territory) []*v1.Territory {
	list := []*v1.Territory{}
	for _, v := range l {
		list = append(list, convertToTerritoryProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a territory.
func (*TerritoryManager) GetProto(c *Territory) *v1.Territory {
	return convertToTerritoryProto(c)
}

// BuildTerritoryListQuery takes a filter and ordering object for a territory.
// and returns an SQL string
func BuildTerritoryListQuery(filters []*v1.TerritoryFilterRule, orderings []*v1.TerritoryOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM territory"
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
func (m *TerritoryManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
