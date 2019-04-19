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

// VolunteerOpportunity is a type for volunteer_opportunity db element.
type VolunteerOpportunity struct {
	ID                       uuid.UUID
	CreatedAt                time.Time
	UpdatedAt                time.Time
	Title                    string
	VolunteerOpportunityType uuid.NullUUID
}

// VolunteerOpportunityManager manages queries returning a volunteerOpportunity or list of volunteerOpportunitys.
// It is configured with a db field to contain the db driver.
type VolunteerOpportunityManager struct {
	db *sql.DB
}

// NewVolunteerOpportunityManager creates a volunteerOpportunity manager
func NewVolunteerOpportunityManager(db *sql.DB) *VolunteerOpportunityManager {
	return &VolunteerOpportunityManager{db: db}
}

// CRUD Methods for the VolunteerOpportunityManager.

// Create creates a volunteerOpportunity.
func (m *VolunteerOpportunityManager) Create(ctx context.Context, item *v1.CreateVolunteerOpportunity) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer_opportunity (title, volunteer_opportunity_type) VALUES($1, $2)  RETURNING id;",
		item.Title, item.VolunteerOpportunityType).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunity-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunity-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single volunteerOpportunity from the database by ID.
func (m *VolunteerOpportunityManager) Get(ctx context.Context, id string) (*VolunteerOpportunity, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query VolunteerOpportunity by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, volunteer_opportunity_type FROM volunteer_opportunity WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunity-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunity-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunity with ID='%s' is not found", id))
	}

	// scan VolunteerOpportunity data into protobuf model
	var volunteerOpportunity VolunteerOpportunity

	if err := rows.Scan(&volunteerOpportunity.ID, &volunteerOpportunity.CreatedAt, &volunteerOpportunity.UpdatedAt, &volunteerOpportunity.Title, &volunteerOpportunity.VolunteerOpportunityType); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunity row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunity rows with ID='%s'",
			id))
	}
	return &volunteerOpportunity, nil
}

// List returns a slice of all volunteerOpportunitys meeting the filter criteria.
func (m *VolunteerOpportunityManager) List(ctx context.Context, filters []*v1.VolunteerOpportunityFilterRule, orderings []*v1.VolunteerOpportunityOrdering, limit int64) ([]*VolunteerOpportunity, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in VolunteerOpportunity Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildVolunteerOpportunityListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunity-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*VolunteerOpportunity{}
	for rows.Next() {
		volunteerOpportunity := new(VolunteerOpportunity)
		if err := rows.Scan(&volunteerOpportunity.ID, &volunteerOpportunity.CreatedAt, &volunteerOpportunity.UpdatedAt, &volunteerOpportunity.Title, &volunteerOpportunity.VolunteerOpportunityType); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunity row-> "+err.Error())
		}
		list = append(list, volunteerOpportunity)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunity-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *VolunteerOpportunityManager) Update(ctx context.Context, item *v1.VolunteerOpportunity) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE volunteer_opportunity SET title=$2, volunteer_opportunity_type=$3 WHERE id=$1",
		item.ID, item.Title, item.VolunteerOpportunityType)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update VolunteerOpportunity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunity with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *VolunteerOpportunityManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM volunteerOpportunity WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete VolunteerOpportunity-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunity with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToVolunteerOpportunityProto accepts a volunteerOpportunity struct and returns a protobuf volunteerOpportunity struct.
func convertToVolunteerOpportunityProto(c *VolunteerOpportunity) *v1.VolunteerOpportunity {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.VolunteerOpportunity{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		VolunteerOpportunityType: *safeNullUUID(c.VolunteerOpportunityType),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a volunteerOpportunity.
func (*VolunteerOpportunityManager) GetProtoList(l []*VolunteerOpportunity) []*v1.VolunteerOpportunity {
	list := []*v1.VolunteerOpportunity{}
	for _, v := range l {
		list = append(list, convertToVolunteerOpportunityProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a volunteerOpportunity.
func (*VolunteerOpportunityManager) GetProto(c *VolunteerOpportunity) *v1.VolunteerOpportunity {
	return convertToVolunteerOpportunityProto(c)
}

// BuildVolunteerOpportunityListQuery takes a filter and ordering object for a volunteerOpportunity.
// and returns an SQL string
func BuildVolunteerOpportunityListQuery(filters []*v1.VolunteerOpportunityFilterRule, orderings []*v1.VolunteerOpportunityOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, volunteer_opportunity_type FROM volunteer_opportunity"
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
func (m *VolunteerOpportunityManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
