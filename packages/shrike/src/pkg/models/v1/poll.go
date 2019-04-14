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

// Poll is a type for poll db element.
type Poll struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// PollManager manages queries returning a poll or list of polls.
// It is configured with a db field to contain the db driver.
type PollManager struct {
	db *sql.DB
}

// NewPollManager creates a poll manager
func NewPollManager(db *sql.DB) *PollManager {
	return &PollManager{db: db}
}

// CRUD Methods for the PollManager.

// CreatePoll creates a poll.
func (m *PollManager) CreatePoll(ctx context.Context, item *v1.CreatePoll) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO poll (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Poll-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Poll-> "+err.Error())
	}
	return &id, nil
}

// GetPoll gets a single poll from the database by ID.
func (m *PollManager) GetPoll(ctx context.Context, id string) (*Poll, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Poll by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM poll WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Poll-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Poll-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%s' is not found", id))
	}

	// scan Poll data into protobuf model
	var poll Poll

	if err := rows.Scan(&poll.ID, &poll.CreatedAt, &poll.UpdatedAt, &poll.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Poll row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Poll rows with ID='%s'",
			id))
	}
	return &poll, nil
}

// ListPoll returns a slice of all polls meeting the filter criteria.
func (m *PollManager) ListPoll(ctx context.Context, filters []*v1.PollFilterRule, orderings []*v1.PollOrdering, limit int64) ([]*Poll, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Poll Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPollListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Poll-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Poll{}
	for rows.Next() {
		poll := new(Poll)
		if err := rows.Scan(&poll.ID, &poll.CreatedAt, &poll.UpdatedAt, &poll.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Poll row-> "+err.Error())
		}
		list = append(list, poll)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Poll-> "+err.Error())
	}
	return list, nil
}

// UpdatePoll runs an update query on the provided db and returns the rows affected as an int64.
func (m *PollManager) UpdatePoll(ctx context.Context, item *v1.Poll) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE poll SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Poll-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeletePoll creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PollManager) DeletePoll(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM poll WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Poll-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Poll with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPollProto accepts a poll struct and returns a protobuf poll struct.
func convertToPollProto(c *Poll) *v1.Poll {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Poll{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a poll.
func (*PollManager) GetProtoList(l []*Poll) []*v1.Poll {
	list := []*v1.Poll{}
	for _, v := range l {
		list = append(list, convertToPollProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a poll.
func (*PollManager) GetProto(c *Poll) *v1.Poll {
	return convertToPollProto(c)
}

// BuildPollListQuery takes a filter and ordering object for a poll.
// and returns an SQL string
func BuildPollListQuery(filters []*v1.PollFilterRule, orderings []*v1.PollOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM poll"
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
func (m *PollManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
