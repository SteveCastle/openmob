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

// Follower is a type for follower db element.
type Follower struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Contact   uuid.UUID
	Cause     uuid.UUID
}

// FollowerManager manages queries returning a follower or list of followers.
// It is configured with a db field to contain the db driver.
type FollowerManager struct {
	db *sql.DB
}

// NewFollowerManager creates a follower manager
func NewFollowerManager(db *sql.DB) *FollowerManager {
	return &FollowerManager{db: db}
}

// CRUD Methods for the FollowerManager.

// CreateFollower creates a follower.
func (m *FollowerManager) CreateFollower(ctx context.Context, item *v1.CreateFollower) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO follower (contact, cause) VALUES($1, $2)  RETURNING id;",
		item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Follower-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Follower-> "+err.Error())
	}
	return &id, nil
}

// GetFollower gets a single follower from the database by ID.
func (m *FollowerManager) GetFollower(ctx context.Context, id string) (*Follower, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Follower by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, contact, cause FROM follower WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Follower-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Follower-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Follower with ID='%s' is not found", id))
	}

	// scan Follower data into protobuf model
	var follower Follower

	if err := rows.Scan(&follower.ID, &follower.CreatedAt, &follower.UpdatedAt, &follower.Contact, &follower.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Follower row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Follower rows with ID='%s'",
			id))
	}
	return &follower, nil
}

// ListFollower returns a slice of all followers meeting the filter criteria.
func (m *FollowerManager) ListFollower(ctx context.Context, filters []*v1.FollowerFilterRule, orderings []*v1.FollowerOrdering, limit int64) ([]*Follower, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Follower Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildFollowerListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Follower-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Follower{}
	for rows.Next() {
		follower := new(Follower)
		if err := rows.Scan(&follower.ID, &follower.CreatedAt, &follower.UpdatedAt, &follower.Contact, &follower.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Follower row-> "+err.Error())
		}
		list = append(list, follower)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Follower-> "+err.Error())
	}
	return list, nil
}

// UpdateFollower runs an update query on the provided db and returns the rows affected as an int64.
func (m *FollowerManager) UpdateFollower(ctx context.Context, item *v1.Follower) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE follower SET contact=$2, cause=$3 WHERE id=$1",
		item.ID, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Follower-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Follower with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteFollower creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *FollowerManager) DeleteFollower(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM follower WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Follower-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Follower with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToFollowerProto accepts a follower struct and returns a protobuf follower struct.
func convertToFollowerProto(c *Follower) *v1.Follower {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Follower{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Contact:   c.Contact.String(),
		Cause:     c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a follower.
func (*FollowerManager) GetProtoList(l []*Follower) []*v1.Follower {
	list := []*v1.Follower{}
	for _, v := range l {
		list = append(list, convertToFollowerProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a follower.
func (*FollowerManager) GetProto(c *Follower) *v1.Follower {
	return convertToFollowerProto(c)
}

// BuildFollowerListQuery takes a filter and ordering object for a follower.
// and returns an SQL string
func BuildFollowerListQuery(filters []*v1.FollowerFilterRule, orderings []*v1.FollowerOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, contact, cause FROM follower"
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
func (m *FollowerManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
