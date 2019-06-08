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

// EventAttendee is a type for event_attendee db element.
type EventAttendee struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	LiveEvent uuid.UUID
	Contact   uuid.UUID
	Cause     uuid.UUID
}

// EventAttendeeManager manages queries returning a eventAttendee or list of eventAttendees.
// It is configured with a db field to contain the db driver.
type EventAttendeeManager struct {
	db *sql.DB
}

// NewEventAttendeeManager creates a eventAttendee manager
func NewEventAttendeeManager(db *sql.DB) *EventAttendeeManager {
	return &EventAttendeeManager{db: db}
}

// CRUD Methods for the EventAttendeeManager.

// Create creates a eventAttendee.
func (m *EventAttendeeManager) Create(ctx context.Context, item *v1.CreateEventAttendee) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO event_attendee (live_event, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		item.LiveEvent, item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into EventAttendee-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created EventAttendee-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single eventAttendee from the database by ID.
func (m *EventAttendeeManager) Get(ctx context.Context, id string) (*EventAttendee, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query EventAttendee by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, live_event, contact, cause FROM event_attendee WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EventAttendee-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from EventAttendee-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%s' is not found", id))
	}

	// scan EventAttendee data into protobuf model
	var eventAttendee EventAttendee

	if err := rows.Scan(&eventAttendee.ID, &eventAttendee.CreatedAt, &eventAttendee.UpdatedAt, &eventAttendee.LiveEvent, &eventAttendee.Contact, &eventAttendee.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from EventAttendee row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple EventAttendee rows with ID='%s'",
			id))
	}
	return &eventAttendee, nil
}

// List returns a slice of all eventAttendees meeting the filter criteria.
func (m *EventAttendeeManager) List(ctx context.Context, filters []*v1.EventAttendeeFilterRule, orderings []*v1.EventAttendeeOrdering, limit int64) ([]*EventAttendee, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in EventAttendee Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildEventAttendeeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from EventAttendee-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*EventAttendee{}
	for rows.Next() {
		eventAttendee := new(EventAttendee)
		if err := rows.Scan(&eventAttendee.ID, &eventAttendee.CreatedAt, &eventAttendee.UpdatedAt, &eventAttendee.LiveEvent, &eventAttendee.Contact, &eventAttendee.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from EventAttendee row-> "+err.Error())
		}
		list = append(list, eventAttendee)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from EventAttendee-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *EventAttendeeManager) Update(ctx context.Context, item *v1.EventAttendee) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE event_attendee SET live_event=$2, contact=$3, cause=$4 WHERE id=$1",
		item.ID, item.LiveEvent, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update EventAttendee-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *EventAttendeeManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM event_attendee WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete EventAttendee-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("EventAttendee with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToEventAttendeeProto accepts a eventAttendee struct and returns a protobuf eventAttendee struct.
func convertToEventAttendeeProto(c *EventAttendee) *v1.EventAttendee {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.EventAttendee{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		LiveEvent: c.LiveEvent.String(),
		Contact:   c.Contact.String(),
		Cause:     c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a eventAttendee.
func (*EventAttendeeManager) GetProtoList(l []*EventAttendee) []*v1.EventAttendee {
	list := []*v1.EventAttendee{}
	for _, v := range l {
		list = append(list, convertToEventAttendeeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a eventAttendee.
func (*EventAttendeeManager) GetProto(c *EventAttendee) *v1.EventAttendee {
	return convertToEventAttendeeProto(c)
}

// BuildEventAttendeeListQuery takes a filter and ordering object for a eventAttendee.
// and returns an SQL string
func BuildEventAttendeeListQuery(filters []*v1.EventAttendeeFilterRule, orderings []*v1.EventAttendeeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, live_event, contact, cause FROM event_attendee"
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
func (m *EventAttendeeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
