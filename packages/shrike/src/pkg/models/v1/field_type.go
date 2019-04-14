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

// FieldType is a type for field_type db element.
type FieldType struct {
	ID                   uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Title                string
	DataType             string
	PropName             string
	StringValueDefault   sql.NullString
	IntValueDefault      sql.NullInt64
	FloatValueDefault    sql.NullFloat64
	BooleanValueDefault  sql.NullBool
	DateTimeValueDefault pq.NullTime
	DataPath             sql.NullString
}

// FieldTypeManager manages queries returning a fieldType or list of fieldTypes.
// It is configured with a db field to contain the db driver.
type FieldTypeManager struct {
	db *sql.DB
}

// NewFieldTypeManager creates a fieldType manager
func NewFieldTypeManager(db *sql.DB) *FieldTypeManager {
	return &FieldTypeManager{db: db}
}

// CRUD Methods for the FieldTypeManager.

// CreateFieldType creates a fieldType.
func (m *FieldTypeManager) CreateFieldType(ctx context.Context, item *v1.CreateFieldType) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO field_type (title, data_type, prop_name, string_value_default, int_value_default, float_value_default, boolean_value_default, date_time_value_default, data_path) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)  RETURNING id;",
		item.Title, item.DataType, item.PropName, item.StringValueDefault, item.IntValueDefault, item.FloatValueDefault, item.BooleanValueDefault, item.DateTimeValueDefault, item.DataPath).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into FieldType-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created FieldType-> "+err.Error())
	}
	return &id, nil
}

// GetFieldType gets a single fieldType from the database by ID.
func (m *FieldTypeManager) GetFieldType(ctx context.Context, id string) (*FieldType, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query FieldType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, data_type, prop_name, string_value_default, int_value_default, float_value_default, boolean_value_default, date_time_value_default, data_path FROM field_type WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from FieldType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from FieldType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%s' is not found", id))
	}

	// scan FieldType data into protobuf model
	var fieldType FieldType

	if err := rows.Scan(&fieldType.ID, &fieldType.CreatedAt, &fieldType.UpdatedAt, &fieldType.Title, &fieldType.DataType, &fieldType.PropName, &fieldType.StringValueDefault, &fieldType.IntValueDefault, &fieldType.FloatValueDefault, &fieldType.BooleanValueDefault, &fieldType.DateTimeValueDefault, &fieldType.DataPath); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from FieldType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple FieldType rows with ID='%s'",
			id))
	}
	return &fieldType, nil
}

// ListFieldType returns a slice of all fieldTypes meeting the filter criteria.
func (m *FieldTypeManager) ListFieldType(ctx context.Context, filters []*v1.FieldTypeFilterRule, orderings []*v1.FieldTypeOrdering, limit int64) ([]*FieldType, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in FieldType Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildFieldTypeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from FieldType-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*FieldType{}
	for rows.Next() {
		fieldType := new(FieldType)
		if err := rows.Scan(&fieldType.ID, &fieldType.CreatedAt, &fieldType.UpdatedAt, &fieldType.Title, &fieldType.DataType, &fieldType.PropName, &fieldType.StringValueDefault, &fieldType.IntValueDefault, &fieldType.FloatValueDefault, &fieldType.BooleanValueDefault, &fieldType.DateTimeValueDefault, &fieldType.DataPath); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from FieldType row-> "+err.Error())
		}
		list = append(list, fieldType)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from FieldType-> "+err.Error())
	}
	return list, nil
}

// UpdateFieldType runs an update query on the provided db and returns the rows affected as an int64.
func (m *FieldTypeManager) UpdateFieldType(ctx context.Context, item *v1.FieldType) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE field_type SET title=$2, data_type=$3, prop_name=$4, string_value_default=$5, int_value_default=$6, float_value_default=$7, boolean_value_default=$8, date_time_value_default=$9, data_path=$10 WHERE id=$1",
		item.ID, item.Title, item.DataType, item.PropName, item.StringValueDefault, item.IntValueDefault, item.FloatValueDefault, item.BooleanValueDefault, item.DateTimeValueDefault, item.DataPath)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update FieldType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteFieldType creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *FieldTypeManager) DeleteFieldType(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM fieldType WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete FieldType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("FieldType with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToFieldTypeProto accepts a fieldType struct and returns a protobuf fieldType struct.
func convertToFieldTypeProto(c *FieldType) *v1.FieldType {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.FieldType{
		ID:                   c.ID.String(),
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
		Title:                c.Title,
		DataType:             c.DataType,
		PropName:             c.PropName,
		StringValueDefault:   *safeNullString(c.StringValueDefault),
		IntValueDefault:      *safeNullInt64(c.IntValueDefault),
		FloatValueDefault:    *safeNullFloat64(c.FloatValueDefault),
		BooleanValueDefault:  *safeNullBool(c.BooleanValueDefault),
		DateTimeValueDefault: safeNullTime(c.DateTimeValueDefault),
		DataPath:             *safeNullString(c.DataPath),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a fieldType.
func (*FieldTypeManager) GetProtoList(l []*FieldType) []*v1.FieldType {
	list := []*v1.FieldType{}
	for _, v := range l {
		list = append(list, convertToFieldTypeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a fieldType.
func (*FieldTypeManager) GetProto(c *FieldType) *v1.FieldType {
	return convertToFieldTypeProto(c)
}

// BuildFieldTypeListQuery takes a filter and ordering object for a fieldType.
// and returns an SQL string
func BuildFieldTypeListQuery(filters []*v1.FieldTypeFilterRule, orderings []*v1.FieldTypeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, data_type, prop_name, string_value_default, int_value_default, float_value_default, boolean_value_default, date_time_value_default, data_path FROM field_type"
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
func (m *FieldTypeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
