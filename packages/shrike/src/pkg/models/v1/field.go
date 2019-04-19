package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/structs"
	uuid "github.com/gofrs/uuid"
	"github.com/lib/pq"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Field is a type for field db element.
type Field struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	FieldType     uuid.UUID
	StringValue   sql.NullString
	IntValue      sql.NullInt64
	FloatValue    sql.NullFloat64
	BooleanValue  sql.NullBool
	DateTimeValue pq.NullTime
	DataPath      sql.NullString
	Component     uuid.NullUUID
}

// FieldManager manages queries returning a field or list of fields.
// It is configured with a db field to contain the db driver.
type FieldManager struct {
	db *sql.DB
}

// NewFieldManager creates a field manager
func NewFieldManager(db *sql.DB) *FieldManager {
	return &FieldManager{db: db}
}

// CRUD Methods for the FieldManager.

// Create creates a field.
func (m *FieldManager) Create(ctx context.Context, item *v1.CreateField) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO field (field_type, string_value, int_value, float_value, boolean_value, date_time_value, data_path, component) VALUES($1, $2, $3, $4, $5, $6, $7, $8)  RETURNING id;",
		item.FieldType, item.StringValue, item.IntValue, item.FloatValue, item.BooleanValue, item.DateTimeValue, item.DataPath, item.Component).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Field-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Field-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single field from the database by ID.
func (m *FieldManager) Get(ctx context.Context, id string) (*Field, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Field by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, field_type, string_value, int_value, float_value, boolean_value, date_time_value, data_path, component FROM field WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Field-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Field-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%s' is not found", id))
	}

	// scan Field data into protobuf model
	var field Field

	if err := rows.Scan(&field.ID, &field.CreatedAt, &field.UpdatedAt, &field.FieldType, &field.StringValue, &field.IntValue, &field.FloatValue, &field.BooleanValue, &field.DateTimeValue, &field.DataPath, &field.Component); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Field row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Field rows with ID='%s'",
			id))
	}
	return &field, nil
}

// List returns a slice of all fields meeting the filter criteria.
func (m *FieldManager) List(ctx context.Context, filters []*v1.FieldFilterRule, orderings []*v1.FieldOrdering, limit int64) ([]*Field, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Field Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildFieldListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Field-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Field{}
	for rows.Next() {
		field := new(Field)
		if err := rows.Scan(&field.ID, &field.CreatedAt, &field.UpdatedAt, &field.FieldType, &field.StringValue, &field.IntValue, &field.FloatValue, &field.BooleanValue, &field.DateTimeValue, &field.DataPath, &field.Component); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Field row-> "+err.Error())
		}
		list = append(list, field)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Field-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *FieldManager) Update(ctx context.Context, item *v1.Field) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE field SET field_type=$2, string_value=$3, int_value=$4, float_value=$5, boolean_value=$6, date_time_value=$7, data_path=$8, component=$9 WHERE id=$1",
		item.ID, item.FieldType, item.StringValue, item.IntValue, item.FloatValue, item.BooleanValue, item.DateTimeValue, item.DataPath, item.Component)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Field-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *FieldManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM field WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Field-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Field with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToFieldProto accepts a field struct and returns a protobuf field struct.
func convertToFieldProto(c *Field) *v1.Field {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Field{
		ID:            c.ID.String(),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		FieldType:     c.FieldType.String(),
		StringValue:   *safeNullString(c.StringValue),
		IntValue:      *safeNullInt64(c.IntValue),
		FloatValue:    *safeNullFloat64(c.FloatValue),
		BooleanValue:  *safeNullBool(c.BooleanValue),
		DateTimeValue: safeNullTime(c.DateTimeValue),
		DataPath:      *safeNullString(c.DataPath),
		Component:     *safeNullUUID(c.Component),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a field.
func (*FieldManager) GetProtoList(l []*Field) []*v1.Field {
	list := []*v1.Field{}
	for _, v := range l {
		list = append(list, convertToFieldProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a field.
func (*FieldManager) GetProto(c *Field) *v1.Field {
	return convertToFieldProto(c)
}

// BuildFieldListQuery takes a filter and ordering object for a field.
// and returns an SQL string
func BuildFieldListQuery(filters []*v1.FieldFilterRule, orderings []*v1.FieldOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, field_type, string_value, int_value, float_value, boolean_value, date_time_value, data_path, component FROM field"
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
func (m *FieldManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
