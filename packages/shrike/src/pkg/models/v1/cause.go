package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/queries"
	"github.com/SteveCastle/structs"
	uuid "github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

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
type CauseManager struct {
	db *sql.DB
}

// GetCause gets a cause
func (m *CauseManager) GetCause(ctx context.Context, id string) (*Cause, error) {
	// get SQL connection from pool
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
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cause with ID='%s' is not found",
			id))
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

// ListCause lists causes.
func (m *CauseManager) ListCause(ctx context.Context, filters []*v1.CauseFilterRule, orderings []*v1.CauseOrdering, limit int64) ([]*Cause, error) {

	// get SQL connection from pool
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Cause Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := queries.BuildCauseListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cause-> "+err.Error())
	}
	defer rows.Close()

	// Variables to store results returned by database.
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

// CreateCause creates a cause.
func (m *CauseManager) CreateCause() {

}

// UpdateCause updates a cause.
func (m *CauseManager) UpdateCause() {

}

//DeleteCause deletes a cause.
func (m *CauseManager) DeleteCause() {

}

func convertTimeToProto(t time.Time) (*timestamp.Timestamp, error) {
	time, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return time, nil
}
func safeNullString(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

func safeNullUUID(u uuid.NullUUID) string {
	if !u.Valid {
		return ""
	}
	return u.UUID.String()
}
func convertToProto(c *Cause) *v1.Cause {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Cause{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		Slug:      c.Slug,
		Summary:   safeNullString(c.Summary),
		HomePage:  safeNullUUID(c.HomePage),
		Photo:     safeNullUUID(c.Photo),
	}
}

//GetProtoList returns a slice of proto struct versions of a cause.
func (*CauseManager) GetProtoList(l []*Cause) []*v1.Cause {
	list := []*v1.Cause{}
	for _, v := range l {
		list = append(list, convertToProto(v))
	}
	return list
}

//GetProto returns a proto struct version of a cause.
func (*CauseManager) GetProto(c *Cause) *v1.Cause {
	return convertToProto(c)
}

// BuildCauseListQuery takes a filter and ordering object for a cause.
// and returns an SQL string
func BuildCauseListQuery(filters []*v1.CauseFilterRule, orderings []*v1.CauseOrdering, limit int64) string {
	// SQL to get all Causes and all columns.
	baseSQL := "SELECT id, created_at, updated_at, title, slug, summary, home_page, photo FROM cause"
	// Generate WHERE clause from filters passed in request.
	for i, r := range filters {
		// Insert where clause before the first filter.
		// And the Logical operator of each successive filter.
		if i == 0 {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "WHERE")
		} else {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "AND")
		}
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			if f.IsExported() {
				baseSQL = fmt.Sprintf("%s %s %s '%s'", baseSQL, ToSnakeCase(f.Name()), Comparison["EQ"], f.Value())
			}
		}
	}
	// Generate ORDER BY clause from ordering passed in request.
	for _, r := range orderings {
		s := structs.New(r.GetField())
		for _, f := range s.Fields() {
			baseSQL = fmt.Sprintf("%s %s", baseSQL, "ORDER BY")
			if f.IsExported() {
				baseSQL = fmt.Sprintf("%s %s ASC", baseSQL, ToSnakeCase(f.Name()))
			}
		}

	}
	baseSQL = fmt.Sprintf("%s LIMIT %d;", baseSQL, limit)
	fmt.Printf("List SQL Executed: %v\n", baseSQL)
	return baseSQL
}

// connect returns SQL database connection from the pool
func (m *CauseManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// NewCauseManager creates a cause manager
func NewCauseManager(db *sql.DB) *CauseManager {
	return &CauseManager{db: db}
}
