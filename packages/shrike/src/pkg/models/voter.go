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

// Voter is a type for voter db element.
type Voter struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Contact   uuid.UUID
	Cause     uuid.UUID
}

// VoterManager manages queries returning a voter or list of voters.
// It is configured with a db field to contain the db driver.
type VoterManager struct {
	db *sql.DB
}

// NewVoterManager creates a voter manager
func NewVoterManager(db *sql.DB) *VoterManager {
	return &VoterManager{db: db}
}

// CRUD Methods for the VoterManager.

// Create creates a voter.
func (m *VoterManager) Create(ctx context.Context, item *v1.CreateVoter) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO voter (contact, cause) VALUES($1, $2)  RETURNING id;",
		item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Voter-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Voter-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single voter from the database by ID.
func (m *VoterManager) Get(ctx context.Context, id string) (*Voter, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Voter by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, contact, cause FROM voter WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Voter-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Voter-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%s' is not found", id))
	}

	// scan Voter data into protobuf model
	var voter Voter

	if err := rows.Scan(&voter.ID, &voter.CreatedAt, &voter.UpdatedAt, &voter.Contact, &voter.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Voter row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Voter rows with ID='%s'",
			id))
	}
	return &voter, nil
}

// List returns a slice of all voters meeting the filter criteria.
func (m *VoterManager) List(ctx context.Context, filters []*v1.VoterFilterRule, orderings []*v1.VoterOrdering, limit int64) ([]*Voter, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Voter Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildVoterListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Voter-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Voter{}
	for rows.Next() {
		voter := new(Voter)
		if err := rows.Scan(&voter.ID, &voter.CreatedAt, &voter.UpdatedAt, &voter.Contact, &voter.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Voter row-> "+err.Error())
		}
		list = append(list, voter)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Voter-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *VoterManager) Update(ctx context.Context, item *v1.Voter) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE voter SET contact=$2, cause=$3 WHERE id=$1",
		item.ID, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Voter-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *VoterManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM voter WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Voter-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Voter with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToVoterProto accepts a voter struct and returns a protobuf voter struct.
func convertToVoterProto(c *Voter) *v1.Voter {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Voter{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Contact:   c.Contact.String(),
		Cause:     c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a voter.
func (*VoterManager) GetProtoList(l []*Voter) []*v1.Voter {
	list := []*v1.Voter{}
	for _, v := range l {
		list = append(list, convertToVoterProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a voter.
func (*VoterManager) GetProto(c *Voter) *v1.Voter {
	return convertToVoterProto(c)
}

// BuildVoterListQuery takes a filter and ordering object for a voter.
// and returns an SQL string
func BuildVoterListQuery(filters []*v1.VoterFilterRule, orderings []*v1.VoterOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, contact, cause FROM voter"
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
func (m *VoterManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
