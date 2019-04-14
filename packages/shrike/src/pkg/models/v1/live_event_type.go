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

// LiveEventType is a type for live_event_type db element.
type LiveEventType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// LiveEventTypeManager manages queries returning a liveEventType or list of liveEventTypes.
// It is configured with a db field to contain the db driver.
type LiveEventTypeManager struct {
	db *sql.DB
}

// NewLiveEventTypeManager creates a liveEventType manager
func NewLiveEventTypeManager(db *sql.DB) *LiveEventTypeManager {
	return &LiveEventTypeManager{db: db}
}

// CRUD Methods for the LiveEventTypeManager.

// CreateLiveEventType creates a liveEventType.
func (m *LiveEventTypeManager) CreateLiveEventType(ctx context.Context, item *v1.CreateLiveEventType) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO live_event_type (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEventType-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEventType-> "+err.Error())
	}
	return &id, nil
}

// GetLiveEventType gets a single liveEventType from the database by ID.
func (m *LiveEventTypeManager) GetLiveEventType(ctx context.Context, id string) (*LiveEventType, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query LiveEventType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM live_event_type WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%s' is not found", id))
	}

	// scan LiveEventType data into protobuf model
	var liveEventType LiveEventType

	if err := rows.Scan(&liveEventType.ID, &liveEventType.CreatedAt, &liveEventType.UpdatedAt, &liveEventType.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEventType rows with ID='%s'",
			id))
	}
	return &liveEventType, nil
}

// ListLiveEventType returns a slice of all liveEventTypes meeting the filter criteria.
func (m *LiveEventTypeManager) ListLiveEventType(ctx context.Context, filters []*v1.LiveEventTypeFilterRule, orderings []*v1.LiveEventTypeOrdering, limit int64) ([]*LiveEventType, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in LiveEventType Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildLiveEventTypeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventType-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*LiveEventType{}
	for rows.Next() {
		liveEventType := new(LiveEventType)
		if err := rows.Scan(&liveEventType.ID, &liveEventType.CreatedAt, &liveEventType.UpdatedAt, &liveEventType.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventType row-> "+err.Error())
		}
		list = append(list, liveEventType)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventType-> "+err.Error())
	}
	return list, nil
}

// UpdateLiveEventType runs an update query on the provided db and returns the rows affected as an int64.
func (m *LiveEventTypeManager) UpdateLiveEventType(ctx context.Context, item *v1.LiveEventType) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE live_event_type SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEventType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteLiveEventType creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *LiveEventTypeManager) DeleteLiveEventType(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM liveEventType WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEventType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventType with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToLiveEventTypeProto accepts a liveEventType struct and returns a protobuf liveEventType struct.
func convertToLiveEventTypeProto(c *LiveEventType) *v1.LiveEventType {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.LiveEventType{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a liveEventType.
func (*LiveEventTypeManager) GetProtoList(l []*LiveEventType) []*v1.LiveEventType {
	list := []*v1.LiveEventType{}
	for _, v := range l {
		list = append(list, convertToLiveEventTypeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a liveEventType.
func (*LiveEventTypeManager) GetProto(c *LiveEventType) *v1.LiveEventType {
	return convertToLiveEventTypeProto(c)
}

// BuildLiveEventTypeListQuery takes a filter and ordering object for a liveEventType.
// and returns an SQL string
func BuildLiveEventTypeListQuery(filters []*v1.LiveEventTypeFilterRule, orderings []*v1.LiveEventTypeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM live_event_type"
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
func (m *LiveEventTypeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
