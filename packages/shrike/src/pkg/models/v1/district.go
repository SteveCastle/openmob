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

	"github.com/SteveCastle/openmob/packages/shrike/src/geography"
)

// District is a type for district db element.
type District struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Geom         geography.NullRegion
	Title        string
	DistrictType uuid.UUID
}

// DistrictManager manages queries returning a district or list of districts.
// It is configured with a db field to contain the db driver.
type DistrictManager struct {
	db *sql.DB
}

// NewDistrictManager creates a district manager
func NewDistrictManager(db *sql.DB) *DistrictManager {
	return &DistrictManager{db: db}
}

// CRUD Methods for the DistrictManager.

// CreateDistrict creates a district.
func (m *DistrictManager) CreateDistrict(ctx context.Context, item *v1.CreateDistrict) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO district (geom, title, district_type) VALUES($1, $2, $3)  RETURNING id;",
		item.Geom, item.Title, item.DistrictType).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into District-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created District-> "+err.Error())
	}
	return &id, nil
}

// GetDistrict gets a single district from the database by ID.
func (m *DistrictManager) GetDistrict(ctx context.Context, id string) (*District, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query District by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, geom, title, district_type FROM district WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from District-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from District-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("District with ID='%s' is not found", id))
	}

	// scan District data into protobuf model
	var district District

	if err := rows.Scan(&district.ID, &district.CreatedAt, &district.UpdatedAt, &district.Geom, &district.Title, &district.DistrictType); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from District row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple District rows with ID='%s'",
			id))
	}
	return &district, nil
}

// ListDistrict returns a slice of all districts meeting the filter criteria.
func (m *DistrictManager) ListDistrict(ctx context.Context, filters []*v1.DistrictFilterRule, orderings []*v1.DistrictOrdering, limit int64) ([]*District, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in District Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildDistrictListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from District-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*District{}
	for rows.Next() {
		district := new(District)
		if err := rows.Scan(&district.ID, &district.CreatedAt, &district.UpdatedAt, &district.Geom, &district.Title, &district.DistrictType); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from District row-> "+err.Error())
		}
		list = append(list, district)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from District-> "+err.Error())
	}
	return list, nil
}

// UpdateDistrict runs an update query on the provided db and returns the rows affected as an int64.
func (m *DistrictManager) UpdateDistrict(ctx context.Context, item *v1.District) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE district SET geom=$2, title=$3, district_type=$4 WHERE id=$1",
		item.ID, item.Geom, item.Title, item.DistrictType)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update District-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("District with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteDistrict creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *DistrictManager) DeleteDistrict(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM district WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete District-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("District with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToDistrictProto accepts a district struct and returns a protobuf district struct.
func convertToDistrictProto(c *District) *v1.District {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.District{
		ID:           c.ID.String(),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		Geom:         *safeNullRegion(c.Geom),
		Title:        c.Title,
		DistrictType: c.DistrictType.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a district.
func (*DistrictManager) GetProtoList(l []*District) []*v1.District {
	list := []*v1.District{}
	for _, v := range l {
		list = append(list, convertToDistrictProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a district.
func (*DistrictManager) GetProto(c *District) *v1.District {
	return convertToDistrictProto(c)
}

// BuildDistrictListQuery takes a filter and ordering object for a district.
// and returns an SQL string
func BuildDistrictListQuery(filters []*v1.DistrictFilterRule, orderings []*v1.DistrictOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, geom, title, district_type FROM district"
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
func (m *DistrictManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
