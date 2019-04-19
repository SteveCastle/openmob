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

// PollItem is a type for poll_item db element.
type PollItem struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Poll      uuid.UUID
}

// PollItemManager manages queries returning a pollItem or list of pollItems.
// It is configured with a db field to contain the db driver.
type PollItemManager struct {
	db *sql.DB
}

// NewPollItemManager creates a pollItem manager
func NewPollItemManager(db *sql.DB) *PollItemManager {
	return &PollItemManager{db: db}
}

// CRUD Methods for the PollItemManager.

// Create creates a pollItem.
func (m *PollItemManager) Create(ctx context.Context, item *v1.CreatePollItem) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO poll_item (title, poll) VALUES($1, $2)  RETURNING id;",
		item.Title, item.Poll).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PollItem-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PollItem-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single pollItem from the database by ID.
func (m *PollItemManager) Get(ctx context.Context, id string) (*PollItem, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query PollItem by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, poll FROM poll_item WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollItem-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PollItem-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollItem with ID='%s' is not found", id))
	}

	// scan PollItem data into protobuf model
	var pollItem PollItem

	if err := rows.Scan(&pollItem.ID, &pollItem.CreatedAt, &pollItem.UpdatedAt, &pollItem.Title, &pollItem.Poll); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollItem row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PollItem rows with ID='%s'",
			id))
	}
	return &pollItem, nil
}

// List returns a slice of all pollItems meeting the filter criteria.
func (m *PollItemManager) List(ctx context.Context, filters []*v1.PollItemFilterRule, orderings []*v1.PollItemOrdering, limit int64) ([]*PollItem, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in PollItem Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPollItemListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollItem-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*PollItem{}
	for rows.Next() {
		pollItem := new(PollItem)
		if err := rows.Scan(&pollItem.ID, &pollItem.CreatedAt, &pollItem.UpdatedAt, &pollItem.Title, &pollItem.Poll); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollItem row-> "+err.Error())
		}
		list = append(list, pollItem)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PollItem-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *PollItemManager) Update(ctx context.Context, item *v1.PollItem) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE poll_item SET title=$2, poll=$3 WHERE id=$1",
		item.ID, item.Title, item.Poll)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PollItem-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollItem with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PollItemManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM pollItem WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PollItem-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollItem with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPollItemProto accepts a pollItem struct and returns a protobuf pollItem struct.
func convertToPollItemProto(c *PollItem) *v1.PollItem {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.PollItem{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		Poll:      c.Poll.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a pollItem.
func (*PollItemManager) GetProtoList(l []*PollItem) []*v1.PollItem {
	list := []*v1.PollItem{}
	for _, v := range l {
		list = append(list, convertToPollItemProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a pollItem.
func (*PollItemManager) GetProto(c *PollItem) *v1.PollItem {
	return convertToPollItemProto(c)
}

// BuildPollItemListQuery takes a filter and ordering object for a pollItem.
// and returns an SQL string
func BuildPollItemListQuery(filters []*v1.PollItemFilterRule, orderings []*v1.PollItemOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, poll FROM poll_item"
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
func (m *PollItemManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
