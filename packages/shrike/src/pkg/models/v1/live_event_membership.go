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

// LiveEventMembership is a type for live_event_membership db element.
type LiveEventMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	LiveEvent uuid.UUID
}

// LiveEventMembershipManager manages queries returning a liveEventMembership or list of liveEventMemberships.
// It is configured with a db field to contain the db driver.
type LiveEventMembershipManager struct {
	db *sql.DB
}

// NewLiveEventMembershipManager creates a liveEventMembership manager
func NewLiveEventMembershipManager(db *sql.DB) *LiveEventMembershipManager {
	return &LiveEventMembershipManager{db: db}
}

// CRUD Methods for the LiveEventMembershipManager.

// CreateLiveEventMembership creates a liveEventMembership.
func (m *LiveEventMembershipManager) CreateLiveEventMembership(ctx context.Context, item *v1.CreateLiveEventMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO live_event_membership (cause, live_event) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.LiveEvent).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LiveEventMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LiveEventMembership-> "+err.Error())
	}
	return &id, nil
}

// GetLiveEventMembership gets a single liveEventMembership from the database by ID.
func (m *LiveEventMembershipManager) GetLiveEventMembership(ctx context.Context, id string) (*LiveEventMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query LiveEventMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, live_event FROM live_event_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%s' is not found", id))
	}

	// scan LiveEventMembership data into protobuf model
	var liveEventMembership LiveEventMembership

	if err := rows.Scan(&liveEventMembership.ID, &liveEventMembership.CreatedAt, &liveEventMembership.UpdatedAt, &liveEventMembership.Cause, &liveEventMembership.LiveEvent); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LiveEventMembership rows with ID='%s'",
			id))
	}
	return &liveEventMembership, nil
}

// ListLiveEventMembership returns a slice of all liveEventMemberships meeting the filter criteria.
func (m *LiveEventMembershipManager) ListLiveEventMembership(ctx context.Context, filters []*v1.LiveEventMembershipFilterRule, orderings []*v1.LiveEventMembershipOrdering, limit int64) ([]*LiveEventMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in LiveEventMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildLiveEventMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LiveEventMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*LiveEventMembership{}
	for rows.Next() {
		liveEventMembership := new(LiveEventMembership)
		if err := rows.Scan(&liveEventMembership.ID, &liveEventMembership.CreatedAt, &liveEventMembership.UpdatedAt, &liveEventMembership.Cause, &liveEventMembership.LiveEvent); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LiveEventMembership row-> "+err.Error())
		}
		list = append(list, liveEventMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LiveEventMembership-> "+err.Error())
	}
	return list, nil
}

// UpdateLiveEventMembership runs an update query on the provided db and returns the rows affected as an int64.
func (m *LiveEventMembershipManager) UpdateLiveEventMembership(ctx context.Context, item *v1.LiveEventMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE live_event_membership SET cause=$2, live_event=$3 WHERE id=$1",
		item.ID, item.Cause, item.LiveEvent)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LiveEventMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteLiveEventMembership creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *LiveEventMembershipManager) DeleteLiveEventMembership(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM liveEventMembership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LiveEventMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LiveEventMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToLiveEventMembershipProto accepts a liveEventMembership struct and returns a protobuf liveEventMembership struct.
func convertToLiveEventMembershipProto(c *LiveEventMembership) *v1.LiveEventMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.LiveEventMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		LiveEvent: c.LiveEvent.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a liveEventMembership.
func (*LiveEventMembershipManager) GetProtoList(l []*LiveEventMembership) []*v1.LiveEventMembership {
	list := []*v1.LiveEventMembership{}
	for _, v := range l {
		list = append(list, convertToLiveEventMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a liveEventMembership.
func (*LiveEventMembershipManager) GetProto(c *LiveEventMembership) *v1.LiveEventMembership {
	return convertToLiveEventMembershipProto(c)
}

// BuildLiveEventMembershipListQuery takes a filter and ordering object for a liveEventMembership.
// and returns an SQL string
func BuildLiveEventMembershipListQuery(filters []*v1.LiveEventMembershipFilterRule, orderings []*v1.LiveEventMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, live_event FROM live_event_membership"
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
func (m *LiveEventMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
