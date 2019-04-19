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

// Boycott is a type for boycott db element.
type Boycott struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// BoycottManager manages queries returning a boycott or list of boycotts.
// It is configured with a db field to contain the db driver.
type BoycottManager struct {
	db *sql.DB
}

// NewBoycottManager creates a boycott manager
func NewBoycottManager(db *sql.DB) *BoycottManager {
	return &BoycottManager{db: db}
}

// CRUD Methods for the BoycottManager.

// Create creates a boycott.
func (m *BoycottManager) Create(ctx context.Context, item *v1.CreateBoycott) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO boycott (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Boycott-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Boycott-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single boycott from the database by ID.
func (m *BoycottManager) Get(ctx context.Context, id string) (*Boycott, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Boycott by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM boycott WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Boycott-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Boycott-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%s' is not found", id))
	}

	// scan Boycott data into protobuf model
	var boycott Boycott

	if err := rows.Scan(&boycott.ID, &boycott.CreatedAt, &boycott.UpdatedAt, &boycott.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Boycott row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Boycott rows with ID='%s'",
			id))
	}
	return &boycott, nil
}

// List returns a slice of all boycotts meeting the filter criteria.
func (m *BoycottManager) List(ctx context.Context, filters []*v1.BoycottFilterRule, orderings []*v1.BoycottOrdering, limit int64) ([]*Boycott, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Boycott Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildBoycottListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Boycott-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Boycott{}
	for rows.Next() {
		boycott := new(Boycott)
		if err := rows.Scan(&boycott.ID, &boycott.CreatedAt, &boycott.UpdatedAt, &boycott.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Boycott row-> "+err.Error())
		}
		list = append(list, boycott)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Boycott-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *BoycottManager) Update(ctx context.Context, item *v1.Boycott) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE boycott SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Boycott-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *BoycottManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM boycott WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Boycott-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Boycott with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToBoycottProto accepts a boycott struct and returns a protobuf boycott struct.
func convertToBoycottProto(c *Boycott) *v1.Boycott {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Boycott{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a boycott.
func (*BoycottManager) GetProtoList(l []*Boycott) []*v1.Boycott {
	list := []*v1.Boycott{}
	for _, v := range l {
		list = append(list, convertToBoycottProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a boycott.
func (*BoycottManager) GetProto(c *Boycott) *v1.Boycott {
	return convertToBoycottProto(c)
}

// BuildBoycottListQuery takes a filter and ordering object for a boycott.
// and returns an SQL string
func BuildBoycottListQuery(filters []*v1.BoycottFilterRule, orderings []*v1.BoycottOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM boycott"
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
func (m *BoycottManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
