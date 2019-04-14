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

// Activity is a type for activity db element.
type Activity struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Title        string
	ActivityType uuid.UUID
	Contact      uuid.UUID
	Cause        uuid.UUID
}

// ActivityManager manages queries returning a activity or list of activitys.
// It is configured with a db field to contain the db driver.
type ActivityManager struct {
	db *sql.DB
}

// NewActivityManager creates a activity manager
func NewActivityManager(db *sql.DB) *ActivityManager {
	return &ActivityManager{db: db}
}

// CRUD Methods for the ActivityManager.

// CreateActivity creates a activity.
func (m *ActivityManager) CreateActivity(ctx context.Context, item *v1.CreateActivity) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO activity (title, activity_type, contact, cause) VALUES($1, $2, $3, $4)  RETURNING id;",
		item.Title, item.ActivityType, item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Activity-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Activity-> "+err.Error())
	}
	return &id, nil
}

// GetActivity gets a single activity from the database by ID.
func (m *ActivityManager) GetActivity(ctx context.Context, id string) (*Activity, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Activity by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, activity_type, contact, cause FROM activity WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Activity-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Activity-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Activity with ID='%s' is not found", id))
	}

	// scan Activity data into protobuf model
	var activity Activity

	if err := rows.Scan(&activity.ID, &activity.CreatedAt, &activity.UpdatedAt, &activity.Title, &activity.ActivityType, &activity.Contact, &activity.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Activity row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Activity rows with ID='%s'",
			id))
	}
	return &activity, nil
}

// ListActivity returns a slice of all activitys meeting the filter criteria.
func (m *ActivityManager) ListActivity(ctx context.Context, filters []*v1.ActivityFilterRule, orderings []*v1.ActivityOrdering, limit int64) ([]*Activity, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Activity Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildActivityListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Activity-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Activity{}
	for rows.Next() {
		activity := new(Activity)
		if err := rows.Scan(&activity.ID, &activity.CreatedAt, &activity.UpdatedAt, &activity.Title, &activity.ActivityType, &activity.Contact, &activity.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Activity row-> "+err.Error())
		}
		list = append(list, activity)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Activity-> "+err.Error())
	}
	return list, nil
}

// UpdateActivity runs an update query on the provided db and returns the rows affected as an int64.
func (m *ActivityManager) UpdateActivity(ctx context.Context, item *v1.Activity) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE activity SET title=$2, activity_type=$3, contact=$4, cause=$5 WHERE id=$1",
		item.ID, item.Title, item.ActivityType, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Activity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Activity with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteActivity creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ActivityManager) DeleteActivity(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM activity WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Activity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Activity with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToActivityProto accepts a activity struct and returns a protobuf activity struct.
func convertToActivityProto(c *Activity) *v1.Activity {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Activity{
		ID:           c.ID.String(),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		Title:        c.Title,
		ActivityType: c.ActivityType.String(),
		Contact:      c.Contact.String(),
		Cause:        c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a activity.
func (*ActivityManager) GetProtoList(l []*Activity) []*v1.Activity {
	list := []*v1.Activity{}
	for _, v := range l {
		list = append(list, convertToActivityProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a activity.
func (*ActivityManager) GetProto(c *Activity) *v1.Activity {
	return convertToActivityProto(c)
}

// BuildActivityListQuery takes a filter and ordering object for a activity.
// and returns an SQL string
func BuildActivityListQuery(filters []*v1.ActivityFilterRule, orderings []*v1.ActivityOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, activity_type, contact, cause FROM activity"
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
func (m *ActivityManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
