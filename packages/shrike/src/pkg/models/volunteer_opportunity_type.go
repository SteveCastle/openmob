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

// VolunteerOpportunityType is a type for volunteer_opportunity_type db element.
type VolunteerOpportunityType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// VolunteerOpportunityTypeManager manages queries returning a volunteerOpportunityType or list of volunteerOpportunityTypes.
// It is configured with a db field to contain the db driver.
type VolunteerOpportunityTypeManager struct {
	db *sql.DB
}

// NewVolunteerOpportunityTypeManager creates a volunteerOpportunityType manager
func NewVolunteerOpportunityTypeManager(db *sql.DB) *VolunteerOpportunityTypeManager {
	return &VolunteerOpportunityTypeManager{db: db}
}

// CRUD Methods for the VolunteerOpportunityTypeManager.

// Create creates a volunteerOpportunityType.
func (m *VolunteerOpportunityTypeManager) Create(ctx context.Context, item *v1.CreateVolunteerOpportunityType) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer_opportunity_type (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunityType-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunityType-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single volunteerOpportunityType from the database by ID.
func (m *VolunteerOpportunityTypeManager) Get(ctx context.Context, id string) (*VolunteerOpportunityType, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query VolunteerOpportunityType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM volunteer_opportunity_type WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%s' is not found", id))
	}

	// scan VolunteerOpportunityType data into protobuf model
	var volunteerOpportunityType VolunteerOpportunityType

	if err := rows.Scan(&volunteerOpportunityType.ID, &volunteerOpportunityType.CreatedAt, &volunteerOpportunityType.UpdatedAt, &volunteerOpportunityType.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunityType rows with ID='%s'",
			id))
	}
	return &volunteerOpportunityType, nil
}

// List returns a slice of all volunteerOpportunityTypes meeting the filter criteria.
func (m *VolunteerOpportunityTypeManager) List(ctx context.Context, filters []*v1.VolunteerOpportunityTypeFilterRule, orderings []*v1.VolunteerOpportunityTypeOrdering, limit int64) ([]*VolunteerOpportunityType, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in VolunteerOpportunityType Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildVolunteerOpportunityTypeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityType-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*VolunteerOpportunityType{}
	for rows.Next() {
		volunteerOpportunityType := new(VolunteerOpportunityType)
		if err := rows.Scan(&volunteerOpportunityType.ID, &volunteerOpportunityType.CreatedAt, &volunteerOpportunityType.UpdatedAt, &volunteerOpportunityType.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityType row-> "+err.Error())
		}
		list = append(list, volunteerOpportunityType)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityType-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *VolunteerOpportunityTypeManager) Update(ctx context.Context, item *v1.VolunteerOpportunityType) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE volunteer_opportunity_type SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update VolunteerOpportunityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *VolunteerOpportunityTypeManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM volunteer_opportunity_type WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete VolunteerOpportunityType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityType with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToVolunteerOpportunityTypeProto accepts a volunteerOpportunityType struct and returns a protobuf volunteerOpportunityType struct.
func convertToVolunteerOpportunityTypeProto(c *VolunteerOpportunityType) *v1.VolunteerOpportunityType {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.VolunteerOpportunityType{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a volunteerOpportunityType.
func (*VolunteerOpportunityTypeManager) GetProtoList(l []*VolunteerOpportunityType) []*v1.VolunteerOpportunityType {
	list := []*v1.VolunteerOpportunityType{}
	for _, v := range l {
		list = append(list, convertToVolunteerOpportunityTypeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a volunteerOpportunityType.
func (*VolunteerOpportunityTypeManager) GetProto(c *VolunteerOpportunityType) *v1.VolunteerOpportunityType {
	return convertToVolunteerOpportunityTypeProto(c)
}

// BuildVolunteerOpportunityTypeListQuery takes a filter and ordering object for a volunteerOpportunityType.
// and returns an SQL string
func BuildVolunteerOpportunityTypeListQuery(filters []*v1.VolunteerOpportunityTypeFilterRule, orderings []*v1.VolunteerOpportunityTypeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM volunteer_opportunity_type"
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
func (m *VolunteerOpportunityTypeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
