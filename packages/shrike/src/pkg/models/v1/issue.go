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

// Issue is a type for issue db element.
type Issue struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Election  uuid.UUID
}

// IssueManager manages queries returning a issue or list of issues.
// It is configured with a db field to contain the db driver.
type IssueManager struct {
	db *sql.DB
}

// NewIssueManager creates a issue manager
func NewIssueManager(db *sql.DB) *IssueManager {
	return &IssueManager{db: db}
}

// CRUD Methods for the IssueManager.

// CreateIssue creates a issue.
func (m *IssueManager) CreateIssue(ctx context.Context, item *v1.CreateIssue) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO issue (title, election) VALUES($1, $2)  RETURNING id;",
		item.Title, item.Election).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Issue-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Issue-> "+err.Error())
	}
	return &id, nil
}

// GetIssue gets a single issue from the database by ID.
func (m *IssueManager) GetIssue(ctx context.Context, id string) (*Issue, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Issue by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, election FROM issue WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Issue-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Issue-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%s' is not found", id))
	}

	// scan Issue data into protobuf model
	var issue Issue

	if err := rows.Scan(&issue.ID, &issue.CreatedAt, &issue.UpdatedAt, &issue.Title, &issue.Election); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Issue row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Issue rows with ID='%s'",
			id))
	}
	return &issue, nil
}

// ListIssue returns a slice of all issues meeting the filter criteria.
func (m *IssueManager) ListIssue(ctx context.Context, filters []*v1.IssueFilterRule, orderings []*v1.IssueOrdering, limit int64) ([]*Issue, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Issue Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildIssueListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Issue-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Issue{}
	for rows.Next() {
		issue := new(Issue)
		if err := rows.Scan(&issue.ID, &issue.CreatedAt, &issue.UpdatedAt, &issue.Title, &issue.Election); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Issue row-> "+err.Error())
		}
		list = append(list, issue)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Issue-> "+err.Error())
	}
	return list, nil
}

// UpdateIssue runs an update query on the provided db and returns the rows affected as an int64.
func (m *IssueManager) UpdateIssue(ctx context.Context, item *v1.Issue) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE issue SET title=$2, election=$3 WHERE id=$1",
		item.ID, item.Title, item.Election)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Issue-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteIssue creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *IssueManager) DeleteIssue(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM issue WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Issue-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Issue with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToIssueProto accepts a issue struct and returns a protobuf issue struct.
func convertToIssueProto(c *Issue) *v1.Issue {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Issue{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		Election:  c.Election.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a issue.
func (*IssueManager) GetProtoList(l []*Issue) []*v1.Issue {
	list := []*v1.Issue{}
	for _, v := range l {
		list = append(list, convertToIssueProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a issue.
func (*IssueManager) GetProto(c *Issue) *v1.Issue {
	return convertToIssueProto(c)
}

// BuildIssueListQuery takes a filter and ordering object for a issue.
// and returns an SQL string
func BuildIssueListQuery(filters []*v1.IssueFilterRule, orderings []*v1.IssueOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, election FROM issue"
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
func (m *IssueManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
