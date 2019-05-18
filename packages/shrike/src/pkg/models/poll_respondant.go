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

// PollRespondant is a type for poll_respondant db element.
type PollRespondant struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Poll      uuid.UUID
	Contact   uuid.UUID
	Cause     uuid.UUID
}

// PollRespondantManager manages queries returning a pollRespondant or list of pollRespondants.
// It is configured with a db field to contain the db driver.
type PollRespondantManager struct {
	db *sql.DB
}

// NewPollRespondantManager creates a pollRespondant manager
func NewPollRespondantManager(db *sql.DB) *PollRespondantManager {
	return &PollRespondantManager{db: db}
}

// CRUD Methods for the PollRespondantManager.

// Create creates a pollRespondant.
func (m *PollRespondantManager) Create(ctx context.Context, item *v1.CreatePollRespondant) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO poll_respondant (poll, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		item.Poll, item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PollRespondant-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PollRespondant-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single pollRespondant from the database by ID.
func (m *PollRespondantManager) Get(ctx context.Context, id string) (*PollRespondant, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query PollRespondant by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, poll, contact, cause FROM poll_respondant WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollRespondant-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PollRespondant-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%s' is not found", id))
	}

	// scan PollRespondant data into protobuf model
	var pollRespondant PollRespondant

	if err := rows.Scan(&pollRespondant.ID, &pollRespondant.CreatedAt, &pollRespondant.UpdatedAt, &pollRespondant.Poll, &pollRespondant.Contact, &pollRespondant.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollRespondant row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PollRespondant rows with ID='%s'",
			id))
	}
	return &pollRespondant, nil
}

// List returns a slice of all pollRespondants meeting the filter criteria.
func (m *PollRespondantManager) List(ctx context.Context, filters []*v1.PollRespondantFilterRule, orderings []*v1.PollRespondantOrdering, limit int64) ([]*PollRespondant, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in PollRespondant Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPollRespondantListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PollRespondant-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*PollRespondant{}
	for rows.Next() {
		pollRespondant := new(PollRespondant)
		if err := rows.Scan(&pollRespondant.ID, &pollRespondant.CreatedAt, &pollRespondant.UpdatedAt, &pollRespondant.Poll, &pollRespondant.Contact, &pollRespondant.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PollRespondant row-> "+err.Error())
		}
		list = append(list, pollRespondant)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PollRespondant-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *PollRespondantManager) Update(ctx context.Context, item *v1.PollRespondant) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE poll_respondant SET poll=$2, contact=$3, cause=$4 WHERE id=$1",
		item.ID, item.Poll, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PollRespondant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PollRespondantManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM pollRespondant WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PollRespondant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PollRespondant with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPollRespondantProto accepts a pollRespondant struct and returns a protobuf pollRespondant struct.
func convertToPollRespondantProto(c *PollRespondant) *v1.PollRespondant {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.PollRespondant{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Poll:      c.Poll.String(),
		Contact:   c.Contact.String(),
		Cause:     c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a pollRespondant.
func (*PollRespondantManager) GetProtoList(l []*PollRespondant) []*v1.PollRespondant {
	list := []*v1.PollRespondant{}
	for _, v := range l {
		list = append(list, convertToPollRespondantProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a pollRespondant.
func (*PollRespondantManager) GetProto(c *PollRespondant) *v1.PollRespondant {
	return convertToPollRespondantProto(c)
}

// BuildPollRespondantListQuery takes a filter and ordering object for a pollRespondant.
// and returns an SQL string
func BuildPollRespondantListQuery(filters []*v1.PollRespondantFilterRule, orderings []*v1.PollRespondantOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, poll, contact, cause FROM poll_respondant"
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
func (m *PollRespondantManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
