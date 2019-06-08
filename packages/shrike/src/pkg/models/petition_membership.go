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

// PetitionMembership is a type for petition_membership db element.
type PetitionMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Petition  uuid.UUID
}

// PetitionMembershipManager manages queries returning a petitionMembership or list of petitionMemberships.
// It is configured with a db field to contain the db driver.
type PetitionMembershipManager struct {
	db *sql.DB
}

// NewPetitionMembershipManager creates a petitionMembership manager
func NewPetitionMembershipManager(db *sql.DB) *PetitionMembershipManager {
	return &PetitionMembershipManager{db: db}
}

// CRUD Methods for the PetitionMembershipManager.

// Create creates a petitionMembership.
func (m *PetitionMembershipManager) Create(ctx context.Context, item *v1.CreatePetitionMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO petition_membership (cause, petition) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.Petition).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PetitionMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PetitionMembership-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single petitionMembership from the database by ID.
func (m *PetitionMembershipManager) Get(ctx context.Context, id string) (*PetitionMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query PetitionMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, petition FROM petition_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%s' is not found", id))
	}

	// scan PetitionMembership data into protobuf model
	var petitionMembership PetitionMembership

	if err := rows.Scan(&petitionMembership.ID, &petitionMembership.CreatedAt, &petitionMembership.UpdatedAt, &petitionMembership.Cause, &petitionMembership.Petition); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PetitionMembership rows with ID='%s'",
			id))
	}
	return &petitionMembership, nil
}

// List returns a slice of all petitionMemberships meeting the filter criteria.
func (m *PetitionMembershipManager) List(ctx context.Context, filters []*v1.PetitionMembershipFilterRule, orderings []*v1.PetitionMembershipOrdering, limit int64) ([]*PetitionMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in PetitionMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPetitionMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*PetitionMembership{}
	for rows.Next() {
		petitionMembership := new(PetitionMembership)
		if err := rows.Scan(&petitionMembership.ID, &petitionMembership.CreatedAt, &petitionMembership.UpdatedAt, &petitionMembership.Cause, &petitionMembership.Petition); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionMembership row-> "+err.Error())
		}
		list = append(list, petitionMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionMembership-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *PetitionMembershipManager) Update(ctx context.Context, item *v1.PetitionMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE petition_membership SET cause=$2, petition=$3 WHERE id=$1",
		item.ID, item.Cause, item.Petition)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PetitionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PetitionMembershipManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM petition_membership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PetitionMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPetitionMembershipProto accepts a petitionMembership struct and returns a protobuf petitionMembership struct.
func convertToPetitionMembershipProto(c *PetitionMembership) *v1.PetitionMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.PetitionMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		Petition:  c.Petition.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a petitionMembership.
func (*PetitionMembershipManager) GetProtoList(l []*PetitionMembership) []*v1.PetitionMembership {
	list := []*v1.PetitionMembership{}
	for _, v := range l {
		list = append(list, convertToPetitionMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a petitionMembership.
func (*PetitionMembershipManager) GetProto(c *PetitionMembership) *v1.PetitionMembership {
	return convertToPetitionMembershipProto(c)
}

// BuildPetitionMembershipListQuery takes a filter and ordering object for a petitionMembership.
// and returns an SQL string
func BuildPetitionMembershipListQuery(filters []*v1.PetitionMembershipFilterRule, orderings []*v1.PetitionMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, petition FROM petition_membership"
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
func (m *PetitionMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
