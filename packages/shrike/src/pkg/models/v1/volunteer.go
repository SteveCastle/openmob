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

// Volunteer is a type for volunteer db element.
type Volunteer struct {
	ID                   uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	VolunteerOpportunity uuid.UUID
	Contact              uuid.UUID
	Cause                uuid.UUID
}

// VolunteerManager manages queries returning a volunteer or list of volunteers.
// It is configured with a db field to contain the db driver.
type VolunteerManager struct {
	db *sql.DB
}

// NewVolunteerManager creates a volunteer manager
func NewVolunteerManager(db *sql.DB) *VolunteerManager {
	return &VolunteerManager{db: db}
}

// CRUD Methods for the VolunteerManager.

// Create creates a volunteer.
func (m *VolunteerManager) Create(ctx context.Context, item *v1.CreateVolunteer) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO volunteer (volunteer_opportunity, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		item.VolunteerOpportunity, item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Volunteer-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Volunteer-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single volunteer from the database by ID.
func (m *VolunteerManager) Get(ctx context.Context, id string) (*Volunteer, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Volunteer by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, volunteer_opportunity, contact, cause FROM volunteer WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Volunteer-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Volunteer-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Volunteer with ID='%s' is not found", id))
	}

	// scan Volunteer data into protobuf model
	var volunteer Volunteer

	if err := rows.Scan(&volunteer.ID, &volunteer.CreatedAt, &volunteer.UpdatedAt, &volunteer.VolunteerOpportunity, &volunteer.Contact, &volunteer.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Volunteer row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Volunteer rows with ID='%s'",
			id))
	}
	return &volunteer, nil
}

// List returns a slice of all volunteers meeting the filter criteria.
func (m *VolunteerManager) List(ctx context.Context, filters []*v1.VolunteerFilterRule, orderings []*v1.VolunteerOrdering, limit int64) ([]*Volunteer, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Volunteer Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildVolunteerListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Volunteer-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Volunteer{}
	for rows.Next() {
		volunteer := new(Volunteer)
		if err := rows.Scan(&volunteer.ID, &volunteer.CreatedAt, &volunteer.UpdatedAt, &volunteer.VolunteerOpportunity, &volunteer.Contact, &volunteer.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Volunteer row-> "+err.Error())
		}
		list = append(list, volunteer)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Volunteer-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *VolunteerManager) Update(ctx context.Context, item *v1.Volunteer) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE volunteer SET volunteer_opportunity=$2, contact=$3, cause=$4 WHERE id=$1",
		item.ID, item.VolunteerOpportunity, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Volunteer-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Volunteer with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *VolunteerManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM volunteer WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Volunteer-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Volunteer with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToVolunteerProto accepts a volunteer struct and returns a protobuf volunteer struct.
func convertToVolunteerProto(c *Volunteer) *v1.Volunteer {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Volunteer{
		ID:                   c.ID.String(),
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
		VolunteerOpportunity: c.VolunteerOpportunity.String(),
		Contact:              c.Contact.String(),
		Cause:                c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a volunteer.
func (*VolunteerManager) GetProtoList(l []*Volunteer) []*v1.Volunteer {
	list := []*v1.Volunteer{}
	for _, v := range l {
		list = append(list, convertToVolunteerProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a volunteer.
func (*VolunteerManager) GetProto(c *Volunteer) *v1.Volunteer {
	return convertToVolunteerProto(c)
}

// BuildVolunteerListQuery takes a filter and ordering object for a volunteer.
// and returns an SQL string
func BuildVolunteerListQuery(filters []*v1.VolunteerFilterRule, orderings []*v1.VolunteerOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, volunteer_opportunity, contact, cause FROM volunteer"
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
func (m *VolunteerManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
