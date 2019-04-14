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

// Election is a type for election db element.
type Election struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// ElectionManager manages queries returning a election or list of elections.
// It is configured with a db field to contain the db driver.
type ElectionManager struct {
	db *sql.DB
}

// NewElectionManager creates a election manager
func NewElectionManager(db *sql.DB) *ElectionManager {
	return &ElectionManager{db: db}
}

// CRUD Methods for the ElectionManager.

// CreateElection creates a election.
func (m *ElectionManager) CreateElection(ctx context.Context, item *v1.CreateElection) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO election (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Election-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Election-> "+err.Error())
	}
	return &id, nil
}

// GetElection gets a single election from the database by ID.
func (m *ElectionManager) GetElection(ctx context.Context, id string) (*Election, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Election by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM election WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Election-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Election-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Election with ID='%s' is not found", id))
	}

	// scan Election data into protobuf model
	var election Election

	if err := rows.Scan(&election.ID, &election.CreatedAt, &election.UpdatedAt, &election.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Election row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Election rows with ID='%s'",
			id))
	}
	return &election, nil
}

// ListElection returns a slice of all elections meeting the filter criteria.
func (m *ElectionManager) ListElection(ctx context.Context, filters []*v1.ElectionFilterRule, orderings []*v1.ElectionOrdering, limit int64) ([]*Election, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Election Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildElectionListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Election-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Election{}
	for rows.Next() {
		election := new(Election)
		if err := rows.Scan(&election.ID, &election.CreatedAt, &election.UpdatedAt, &election.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Election row-> "+err.Error())
		}
		list = append(list, election)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Election-> "+err.Error())
	}
	return list, nil
}

// UpdateElection runs an update query on the provided db and returns the rows affected as an int64.
func (m *ElectionManager) UpdateElection(ctx context.Context, item *v1.Election) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE election SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Election-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Election with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteElection creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ElectionManager) DeleteElection(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM election WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Election-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Election with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToElectionProto accepts a election struct and returns a protobuf election struct.
func convertToElectionProto(c *Election) *v1.Election {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Election{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a election.
func (*ElectionManager) GetProtoList(l []*Election) []*v1.Election {
	list := []*v1.Election{}
	for _, v := range l {
		list = append(list, convertToElectionProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a election.
func (*ElectionManager) GetProto(c *Election) *v1.Election {
	return convertToElectionProto(c)
}

// BuildElectionListQuery takes a filter and ordering object for a election.
// and returns an SQL string
func BuildElectionListQuery(filters []*v1.ElectionFilterRule, orderings []*v1.ElectionOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM election"
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
func (m *ElectionManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
