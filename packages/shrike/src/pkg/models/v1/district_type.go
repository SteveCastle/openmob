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

// DistrictType is a type for district_type db element.
type DistrictType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// DistrictTypeManager manages queries returning a districtType or list of districtTypes.
// It is configured with a db field to contain the db driver.
type DistrictTypeManager struct {
	db *sql.DB
}

// NewDistrictTypeManager creates a districtType manager
func NewDistrictTypeManager(db *sql.DB) *DistrictTypeManager {
	return &DistrictTypeManager{db: db}
}

// CRUD Methods for the DistrictTypeManager.

// CreateDistrictType creates a districtType.
func (m *DistrictTypeManager) CreateDistrictType(ctx context.Context, item *v1.CreateDistrictType) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO district_type (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into DistrictType-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created DistrictType-> "+err.Error())
	}
	return &id, nil
}

// GetDistrictType gets a single districtType from the database by ID.
func (m *DistrictTypeManager) GetDistrictType(ctx context.Context, id string) (*DistrictType, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query DistrictType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM district_type WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DistrictType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from DistrictType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DistrictType with ID='%s' is not found", id))
	}

	// scan DistrictType data into protobuf model
	var districtType DistrictType

	if err := rows.Scan(&districtType.ID, &districtType.CreatedAt, &districtType.UpdatedAt, &districtType.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from DistrictType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple DistrictType rows with ID='%s'",
			id))
	}
	return &districtType, nil
}

// ListDistrictType returns a slice of all districtTypes meeting the filter criteria.
func (m *DistrictTypeManager) ListDistrictType(ctx context.Context, filters []*v1.DistrictTypeFilterRule, orderings []*v1.DistrictTypeOrdering, limit int64) ([]*DistrictType, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in DistrictType Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildDistrictTypeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from DistrictType-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*DistrictType{}
	for rows.Next() {
		districtType := new(DistrictType)
		if err := rows.Scan(&districtType.ID, &districtType.CreatedAt, &districtType.UpdatedAt, &districtType.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from DistrictType row-> "+err.Error())
		}
		list = append(list, districtType)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from DistrictType-> "+err.Error())
	}
	return list, nil
}

// UpdateDistrictType runs an update query on the provided db and returns the rows affected as an int64.
func (m *DistrictTypeManager) UpdateDistrictType(ctx context.Context, item *v1.DistrictType) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE district_type SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update DistrictType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DistrictType with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteDistrictType creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *DistrictTypeManager) DeleteDistrictType(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM districtType WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete DistrictType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("DistrictType with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToDistrictTypeProto accepts a districtType struct and returns a protobuf districtType struct.
func convertToDistrictTypeProto(c *DistrictType) *v1.DistrictType {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.DistrictType{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a districtType.
func (*DistrictTypeManager) GetProtoList(l []*DistrictType) []*v1.DistrictType {
	list := []*v1.DistrictType{}
	for _, v := range l {
		list = append(list, convertToDistrictTypeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a districtType.
func (*DistrictTypeManager) GetProto(c *DistrictType) *v1.DistrictType {
	return convertToDistrictTypeProto(c)
}

// BuildDistrictTypeListQuery takes a filter and ordering object for a districtType.
// and returns an SQL string
func BuildDistrictTypeListQuery(filters []*v1.DistrictTypeFilterRule, orderings []*v1.DistrictTypeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM district_type"
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
func (m *DistrictTypeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
