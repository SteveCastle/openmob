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

// LiveEvent is a type for live_event db element.
type LiveEvent struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Title         string
	LiveEventType uuid.UUID
}

// LiveEventManager manages queries returning a liveEvent or list of liveEvents.
// It is configured with a db field to contain the db driver.
type LiveEventManager struct {
	db *sql.DB
}

// NewLiveEventManager creates a liveEvent manager
func NewLiveEventManager(db *sql.DB) *LiveEventManager {
	return &LiveEventManager{db: db}
}

// CRUD Methods for the LiveEventManager.

// CreateLiveEvent creates a liveEvent.
func (m *LiveEventManager) CreateLiveEvent(ctx context.Context, item *v1.CreateLiveEvent) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO live_event (title, live_event_type) VALUES($1, $2)  RETURNING id;",
		item.Title, item.LiveEventType).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEvent-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEvent-> "+err.Error())
	}
	return &id, nil
}

// GetLiveEvent gets a single liveEvent from the database by ID.
func (m *LiveEventManager) GetLiveEvent(ctx context.Context, id string) (*LiveEvent, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query LiveEvent by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, live_event_type FROM live_event WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEvent-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEvent-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%s' is not found", id))
	}

	// scan LiveEvent data into protobuf model
	var liveEvent LiveEvent

	if err := rows.Scan(&liveEvent.ID, &liveEvent.CreatedAt, &liveEvent.UpdatedAt, &liveEvent.Title, &liveEvent.LiveEventType); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEvent row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEvent rows with ID='%s'",
			id))
	}
	return &liveEvent, nil
}

// ListLiveEvent returns a slice of all liveEvents meeting the filter criteria.
func (m *LiveEventManager) ListLiveEvent(ctx context.Context, filters []*v1.LiveEventFilterRule, orderings []*v1.LiveEventOrdering, limit int64) ([]*LiveEvent, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in LiveEvent Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildLiveEventListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEvent-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*LiveEvent{}
	for rows.Next() {
		liveEvent := new(LiveEvent)
		if err := rows.Scan(&liveEvent.ID, &liveEvent.CreatedAt, &liveEvent.UpdatedAt, &liveEvent.Title, &liveEvent.LiveEventType); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEvent row-> "+err.Error())
		}
		list = append(list, liveEvent)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEvent-> "+err.Error())
	}
	return list, nil
}

// UpdateLiveEvent runs an update query on the provided db and returns the rows affected as an int64.
func (m *LiveEventManager) UpdateLiveEvent(ctx context.Context, item *v1.LiveEvent) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE live_event SET title=$2, live_event_type=$3 WHERE id=$1",
		item.ID, item.Title, item.LiveEventType)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEvent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteLiveEvent creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *LiveEventManager) DeleteLiveEvent(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM liveEvent WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEvent-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEvent with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToLiveEventProto accepts a liveEvent struct and returns a protobuf liveEvent struct.
func convertToLiveEventProto(c *LiveEvent) *v1.LiveEvent {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.LiveEvent{
		ID:            c.ID.String(),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		Title:         c.Title,
		LiveEventType: c.LiveEventType.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a liveEvent.
func (*LiveEventManager) GetProtoList(l []*LiveEvent) []*v1.LiveEvent {
	list := []*v1.LiveEvent{}
	for _, v := range l {
		list = append(list, convertToLiveEventProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a liveEvent.
func (*LiveEventManager) GetProto(c *LiveEvent) *v1.LiveEvent {
	return convertToLiveEventProto(c)
}

// BuildLiveEventListQuery takes a filter and ordering object for a liveEvent.
// and returns an SQL string
func BuildLiveEventListQuery(filters []*v1.LiveEventFilterRule, orderings []*v1.LiveEventOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, live_event_type FROM live_event"
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
func (m *LiveEventManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
