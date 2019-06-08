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

// VolunteerOpportunityMembership is a type for volunteer_opportunity_membership db element.
type VolunteerOpportunityMembership struct {
	ID                   uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Cause                uuid.UUID
	VolunteerOpportunity uuid.UUID
}

// VolunteerOpportunityMembershipManager manages queries returning a volunteerOpportunityMembership or list of volunteerOpportunityMemberships.
// It is configured with a db field to contain the db driver.
type VolunteerOpportunityMembershipManager struct {
	db *sql.DB
}

// NewVolunteerOpportunityMembershipManager creates a volunteerOpportunityMembership manager
func NewVolunteerOpportunityMembershipManager(db *sql.DB) *VolunteerOpportunityMembershipManager {
	return &VolunteerOpportunityMembershipManager{db: db}
}

// CRUD Methods for the VolunteerOpportunityMembershipManager.

// Create creates a volunteerOpportunityMembership.
func (m *VolunteerOpportunityMembershipManager) Create(ctx context.Context, item *v1.CreateVolunteerOpportunityMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer_opportunity_membership (cause, volunteer_opportunity) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.VolunteerOpportunity).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into VolunteerOpportunityMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created VolunteerOpportunityMembership-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single volunteerOpportunityMembership from the database by ID.
func (m *VolunteerOpportunityMembershipManager) Get(ctx context.Context, id string) (*VolunteerOpportunityMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query VolunteerOpportunityMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, volunteer_opportunity FROM volunteer_opportunity_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%s' is not found", id))
	}

	// scan VolunteerOpportunityMembership data into protobuf model
	var volunteerOpportunityMembership VolunteerOpportunityMembership

	if err := rows.Scan(&volunteerOpportunityMembership.ID, &volunteerOpportunityMembership.CreatedAt, &volunteerOpportunityMembership.UpdatedAt, &volunteerOpportunityMembership.Cause, &volunteerOpportunityMembership.VolunteerOpportunity); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple VolunteerOpportunityMembership rows with ID='%s'",
			id))
	}
	return &volunteerOpportunityMembership, nil
}

// List returns a slice of all volunteerOpportunityMemberships meeting the filter criteria.
func (m *VolunteerOpportunityMembershipManager) List(ctx context.Context, filters []*v1.VolunteerOpportunityMembershipFilterRule, orderings []*v1.VolunteerOpportunityMembershipOrdering, limit int64) ([]*VolunteerOpportunityMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in VolunteerOpportunityMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildVolunteerOpportunityMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from VolunteerOpportunityMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*VolunteerOpportunityMembership{}
	for rows.Next() {
		volunteerOpportunityMembership := new(VolunteerOpportunityMembership)
		if err := rows.Scan(&volunteerOpportunityMembership.ID, &volunteerOpportunityMembership.CreatedAt, &volunteerOpportunityMembership.UpdatedAt, &volunteerOpportunityMembership.Cause, &volunteerOpportunityMembership.VolunteerOpportunity); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from VolunteerOpportunityMembership row-> "+err.Error())
		}
		list = append(list, volunteerOpportunityMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from VolunteerOpportunityMembership-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *VolunteerOpportunityMembershipManager) Update(ctx context.Context, item *v1.VolunteerOpportunityMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE volunteer_opportunity_membership SET cause=$2, volunteer_opportunity=$3 WHERE id=$1",
		item.ID, item.Cause, item.VolunteerOpportunity)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update VolunteerOpportunityMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *VolunteerOpportunityMembershipManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM volunteer_opportunity_membership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete VolunteerOpportunityMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("VolunteerOpportunityMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToVolunteerOpportunityMembershipProto accepts a volunteerOpportunityMembership struct and returns a protobuf volunteerOpportunityMembership struct.
func convertToVolunteerOpportunityMembershipProto(c *VolunteerOpportunityMembership) *v1.VolunteerOpportunityMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.VolunteerOpportunityMembership{
		ID:                   c.ID.String(),
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
		Cause:                c.Cause.String(),
		VolunteerOpportunity: c.VolunteerOpportunity.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a volunteerOpportunityMembership.
func (*VolunteerOpportunityMembershipManager) GetProtoList(l []*VolunteerOpportunityMembership) []*v1.VolunteerOpportunityMembership {
	list := []*v1.VolunteerOpportunityMembership{}
	for _, v := range l {
		list = append(list, convertToVolunteerOpportunityMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a volunteerOpportunityMembership.
func (*VolunteerOpportunityMembershipManager) GetProto(c *VolunteerOpportunityMembership) *v1.VolunteerOpportunityMembership {
	return convertToVolunteerOpportunityMembershipProto(c)
}

// BuildVolunteerOpportunityMembershipListQuery takes a filter and ordering object for a volunteerOpportunityMembership.
// and returns an SQL string
func BuildVolunteerOpportunityMembershipListQuery(filters []*v1.VolunteerOpportunityMembershipFilterRule, orderings []*v1.VolunteerOpportunityMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, volunteer_opportunity FROM volunteer_opportunity_membership"
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
func (m *VolunteerOpportunityMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
