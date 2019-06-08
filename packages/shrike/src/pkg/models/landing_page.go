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

// LandingPage is a type for landing_page db element.
type LandingPage struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Cause     uuid.UUID
	Layout    uuid.NullUUID
}

// LandingPageManager manages queries returning a landingPage or list of landingPages.
// It is configured with a db field to contain the db driver.
type LandingPageManager struct {
	db *sql.DB
}

// NewLandingPageManager creates a landingPage manager
func NewLandingPageManager(db *sql.DB) *LandingPageManager {
	return &LandingPageManager{db: db}
}

// CRUD Methods for the LandingPageManager.

// Create creates a landingPage.
func (m *LandingPageManager) Create(ctx context.Context, item *v1.CreateLandingPage) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO landing_page (title, cause, layout) VALUES($1, $2, $3)  RETURNING id;",
		item.Title, item.Cause, item.Layout).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into LandingPage-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created LandingPage-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single landingPage from the database by ID.
func (m *LandingPageManager) Get(ctx context.Context, id string) (*LandingPage, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query LandingPage by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, cause, layout FROM landing_page WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LandingPage-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from LandingPage-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LandingPage with ID='%s' is not found", id))
	}

	// scan LandingPage data into protobuf model
	var landingPage LandingPage

	if err := rows.Scan(&landingPage.ID, &landingPage.CreatedAt, &landingPage.UpdatedAt, &landingPage.Title, &landingPage.Cause, &landingPage.Layout); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from LandingPage row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple LandingPage rows with ID='%s'",
			id))
	}
	return &landingPage, nil
}

// List returns a slice of all landingPages meeting the filter criteria.
func (m *LandingPageManager) List(ctx context.Context, filters []*v1.LandingPageFilterRule, orderings []*v1.LandingPageOrdering, limit int64) ([]*LandingPage, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in LandingPage Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildLandingPageListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from LandingPage-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*LandingPage{}
	for rows.Next() {
		landingPage := new(LandingPage)
		if err := rows.Scan(&landingPage.ID, &landingPage.CreatedAt, &landingPage.UpdatedAt, &landingPage.Title, &landingPage.Cause, &landingPage.Layout); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from LandingPage row-> "+err.Error())
		}
		list = append(list, landingPage)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from LandingPage-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *LandingPageManager) Update(ctx context.Context, item *v1.LandingPage) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE landing_page SET title=$2, cause=$3, layout=$4 WHERE id=$1",
		item.ID, item.Title, item.Cause, item.Layout)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update LandingPage-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LandingPage with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *LandingPageManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM landing_page WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete LandingPage-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("LandingPage with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToLandingPageProto accepts a landingPage struct and returns a protobuf landingPage struct.
func convertToLandingPageProto(c *LandingPage) *v1.LandingPage {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.LandingPage{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		Cause:     c.Cause.String(),
		Layout:    *safeNullUUID(c.Layout),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a landingPage.
func (*LandingPageManager) GetProtoList(l []*LandingPage) []*v1.LandingPage {
	list := []*v1.LandingPage{}
	for _, v := range l {
		list = append(list, convertToLandingPageProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a landingPage.
func (*LandingPageManager) GetProto(c *LandingPage) *v1.LandingPage {
	return convertToLandingPageProto(c)
}

// BuildLandingPageListQuery takes a filter and ordering object for a landingPage.
// and returns an SQL string
func BuildLandingPageListQuery(filters []*v1.LandingPageFilterRule, orderings []*v1.LandingPageOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, cause, layout FROM landing_page"
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
func (m *LandingPageManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
