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

// ActivityType is a type for activity_type db element.
type ActivityType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// ActivityTypeManager manages queries returning a activityType or list of activityTypes.
// It is configured with a db field to contain the db driver.
type ActivityTypeManager struct {
	db *sql.DB
}

// NewActivityTypeManager creates a activityType manager
func NewActivityTypeManager(db *sql.DB) *ActivityTypeManager {
	return &ActivityTypeManager{db: db}
}

// CRUD Methods for the ActivityTypeManager.

// CreateActivityType creates a activityType.
func (m *ActivityTypeManager) CreateActivityType(ctx context.Context, item *v1.CreateActivityType) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO activity_type (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ActivityType-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ActivityType-> "+err.Error())
	}
	return &id, nil
}

// GetActivityType gets a single activityType from the database by ID.
func (m *ActivityTypeManager) GetActivityType(ctx context.Context, id string) (*ActivityType, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ActivityType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM activity_type WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ActivityType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ActivityType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ActivityType with ID='%s' is not found", id))
	}

	// scan ActivityType data into protobuf model
	var activityType ActivityType

	if err := rows.Scan(&activityType.ID, &activityType.CreatedAt, &activityType.UpdatedAt, &activityType.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ActivityType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ActivityType rows with ID='%s'",
			id))
	}
	return &activityType, nil
}

// ListActivityType returns a slice of all activityTypes meeting the filter criteria.
func (m *ActivityTypeManager) ListActivityType(ctx context.Context, filters []*v1.ActivityTypeFilterRule, orderings []*v1.ActivityTypeOrdering, limit int64) ([]*ActivityType, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ActivityType Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildActivityTypeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ActivityType-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ActivityType{}
	for rows.Next() {
		activityType := new(ActivityType)
		if err := rows.Scan(&activityType.ID, &activityType.CreatedAt, &activityType.UpdatedAt, &activityType.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ActivityType row-> "+err.Error())
		}
		list = append(list, activityType)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ActivityType-> "+err.Error())
	}
	return list, nil
}

// UpdateActivityType runs an update query on the provided db and returns the rows affected as an int64.
func (m *ActivityTypeManager) UpdateActivityType(ctx context.Context, item *v1.ActivityType) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE activity_type SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ActivityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ActivityType with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteActivityType creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ActivityTypeManager) DeleteActivityType(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM activityType WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ActivityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ActivityType with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToActivityTypeProto accepts a activityType struct and returns a protobuf activityType struct.
func convertToActivityTypeProto(c *ActivityType) *v1.ActivityType {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ActivityType{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a activityType.
func (*ActivityTypeManager) GetProtoList(l []*ActivityType) []*v1.ActivityType {
	list := []*v1.ActivityType{}
	for _, v := range l {
		list = append(list, convertToActivityTypeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a activityType.
func (*ActivityTypeManager) GetProto(c *ActivityType) *v1.ActivityType {
	return convertToActivityTypeProto(c)
}

// BuildActivityTypeListQuery takes a filter and ordering object for a activityType.
// and returns an SQL string
func BuildActivityTypeListQuery(filters []*v1.ActivityTypeFilterRule, orderings []*v1.ActivityTypeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM activity_type"
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
func (m *ActivityTypeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
