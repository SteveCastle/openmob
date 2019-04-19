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

// Cause is a type for cause db element.
type Cause struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Slug      string
	Summary   sql.NullString
	HomePage  uuid.NullUUID
	Photo     uuid.NullUUID
}

// CauseManager manages queries returning a cause or list of causes.
// It is configured with a db field to contain the db driver.
type CauseManager struct {
	db *sql.DB
}

// NewCauseManager creates a cause manager
func NewCauseManager(db *sql.DB) *CauseManager {
	return &CauseManager{db: db}
}

// CRUD Methods for the CauseManager.

// Create creates a cause.
func (m *CauseManager) Create(ctx context.Context, item *v1.CreateCause) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO cause (title, slug, summary, home_page, photo) VALUES($1, $2, $3, $4, $5)  RETURNING id;",
		item.Title, item.Slug, item.Summary, item.HomePage, item.Photo).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Cause-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Cause-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single cause from the database by ID.
func (m *CauseManager) Get(ctx context.Context, id string) (*Cause, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Cause by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, slug, summary, home_page, photo FROM cause WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cause-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Cause-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%s' is not found", id))
	}

	// scan Cause data into protobuf model
	var cause Cause

	if err := rows.Scan(&cause.ID, &cause.CreatedAt, &cause.UpdatedAt, &cause.Title, &cause.Slug, &cause.Summary, &cause.HomePage, &cause.Photo); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cause row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Cause rows with ID='%s'",
			id))
	}
	return &cause, nil
}

// List returns a slice of all causes meeting the filter criteria.
func (m *CauseManager) List(ctx context.Context, filters []*v1.CauseFilterRule, orderings []*v1.CauseOrdering, limit int64) ([]*Cause, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Cause Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildCauseListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cause-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Cause{}
	for rows.Next() {
		cause := new(Cause)
		if err := rows.Scan(&cause.ID, &cause.CreatedAt, &cause.UpdatedAt, &cause.Title, &cause.Slug, &cause.Summary, &cause.HomePage, &cause.Photo); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cause row-> "+err.Error())
		}
		list = append(list, cause)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Cause-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *CauseManager) Update(ctx context.Context, item *v1.Cause) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE cause SET title=$2, slug=$3, summary=$4, home_page=$5, photo=$6 WHERE id=$1",
		item.ID, item.Title, item.Slug, item.Summary, item.HomePage, item.Photo)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Cause-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *CauseManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM cause WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Cause-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToCauseProto accepts a cause struct and returns a protobuf cause struct.
func convertToCauseProto(c *Cause) *v1.Cause {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Cause{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		Slug:      c.Slug,
		Summary:   *safeNullString(c.Summary),
		HomePage:  *safeNullUUID(c.HomePage),
		Photo:     *safeNullUUID(c.Photo),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a cause.
func (*CauseManager) GetProtoList(l []*Cause) []*v1.Cause {
	list := []*v1.Cause{}
	for _, v := range l {
		list = append(list, convertToCauseProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a cause.
func (*CauseManager) GetProto(c *Cause) *v1.Cause {
	return convertToCauseProto(c)
}

// BuildCauseListQuery takes a filter and ordering object for a cause.
// and returns an SQL string
func BuildCauseListQuery(filters []*v1.CauseFilterRule, orderings []*v1.CauseOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, slug, summary, home_page, photo FROM cause"
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
func (m *CauseManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
