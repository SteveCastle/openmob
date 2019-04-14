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

// Candidate is a type for candidate db element.
type Candidate struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Election  uuid.UUID
}

// CandidateManager manages queries returning a candidate or list of candidates.
// It is configured with a db field to contain the db driver.
type CandidateManager struct {
	db *sql.DB
}

// NewCandidateManager creates a candidate manager
func NewCandidateManager(db *sql.DB) *CandidateManager {
	return &CandidateManager{db: db}
}

// CRUD Methods for the CandidateManager.

// CreateCandidate creates a candidate.
func (m *CandidateManager) CreateCandidate(ctx context.Context, item *v1.CreateCandidate) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO candidate (election) VALUES($1)  RETURNING id;",
		item.Election).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Candidate-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Candidate-> "+err.Error())
	}
	return &id, nil
}

// GetCandidate gets a single candidate from the database by ID.
func (m *CandidateManager) GetCandidate(ctx context.Context, id string) (*Candidate, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Candidate by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, election FROM candidate WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Candidate-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Candidate-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%s' is not found", id))
	}

	// scan Candidate data into protobuf model
	var candidate Candidate

	if err := rows.Scan(&candidate.ID, &candidate.CreatedAt, &candidate.UpdatedAt, &candidate.Election); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Candidate row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Candidate rows with ID='%s'",
			id))
	}
	return &candidate, nil
}

// ListCandidate returns a slice of all candidates meeting the filter criteria.
func (m *CandidateManager) ListCandidate(ctx context.Context, filters []*v1.CandidateFilterRule, orderings []*v1.CandidateOrdering, limit int64) ([]*Candidate, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Candidate Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildCandidateListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Candidate-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Candidate{}
	for rows.Next() {
		candidate := new(Candidate)
		if err := rows.Scan(&candidate.ID, &candidate.CreatedAt, &candidate.UpdatedAt, &candidate.Election); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Candidate row-> "+err.Error())
		}
		list = append(list, candidate)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Candidate-> "+err.Error())
	}
	return list, nil
}

// UpdateCandidate runs an update query on the provided db and returns the rows affected as an int64.
func (m *CandidateManager) UpdateCandidate(ctx context.Context, item *v1.Candidate) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE candidate SET election=$2 WHERE id=$1",
		item.ID, item.Election)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Candidate-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteCandidate creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *CandidateManager) DeleteCandidate(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM candidate WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Candidate-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Candidate with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToCandidateProto accepts a candidate struct and returns a protobuf candidate struct.
func convertToCandidateProto(c *Candidate) *v1.Candidate {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Candidate{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Election:  c.Election.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a candidate.
func (*CandidateManager) GetProtoList(l []*Candidate) []*v1.Candidate {
	list := []*v1.Candidate{}
	for _, v := range l {
		list = append(list, convertToCandidateProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a candidate.
func (*CandidateManager) GetProto(c *Candidate) *v1.Candidate {
	return convertToCandidateProto(c)
}

// BuildCandidateListQuery takes a filter and ordering object for a candidate.
// and returns an SQL string
func BuildCandidateListQuery(filters []*v1.CandidateFilterRule, orderings []*v1.CandidateOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, election FROM candidate"
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
func (m *CandidateManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
