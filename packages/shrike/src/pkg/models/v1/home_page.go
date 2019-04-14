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

// HomePage is a type for home_page db element.
type HomePage struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Layout    uuid.NullUUID
}

// HomePageManager manages queries returning a homePage or list of homePages.
// It is configured with a db field to contain the db driver.
type HomePageManager struct {
	db *sql.DB
}

// NewHomePageManager creates a homePage manager
func NewHomePageManager(db *sql.DB) *HomePageManager {
	return &HomePageManager{db: db}
}

// CRUD Methods for the HomePageManager.

// CreateHomePage creates a homePage.
func (m *HomePageManager) CreateHomePage(ctx context.Context, item *v1.CreateHomePage) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO home_page (title, layout) VALUES($1, $2)  RETURNING id;",
		item.Title, item.Layout).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into HomePage-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created HomePage-> "+err.Error())
	}
	return &id, nil
}

// GetHomePage gets a single homePage from the database by ID.
func (m *HomePageManager) GetHomePage(ctx context.Context, id string) (*HomePage, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query HomePage by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, layout FROM home_page WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from HomePage-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from HomePage-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("HomePage with ID='%s' is not found", id))
	}

	// scan HomePage data into protobuf model
	var homePage HomePage

	if err := rows.Scan(&homePage.ID, &homePage.CreatedAt, &homePage.UpdatedAt, &homePage.Title, &homePage.Layout); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from HomePage row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple HomePage rows with ID='%s'",
			id))
	}
	return &homePage, nil
}

// ListHomePage returns a slice of all homePages meeting the filter criteria.
func (m *HomePageManager) ListHomePage(ctx context.Context, filters []*v1.HomePageFilterRule, orderings []*v1.HomePageOrdering, limit int64) ([]*HomePage, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in HomePage Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildHomePageListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from HomePage-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*HomePage{}
	for rows.Next() {
		homePage := new(HomePage)
		if err := rows.Scan(&homePage.ID, &homePage.CreatedAt, &homePage.UpdatedAt, &homePage.Title, &homePage.Layout); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from HomePage row-> "+err.Error())
		}
		list = append(list, homePage)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from HomePage-> "+err.Error())
	}
	return list, nil
}

// UpdateHomePage runs an update query on the provided db and returns the rows affected as an int64.
func (m *HomePageManager) UpdateHomePage(ctx context.Context, item *v1.HomePage) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE home_page SET title=$2, layout=$3 WHERE id=$1",
		item.ID, item.Title, item.Layout)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update HomePage-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("HomePage with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteHomePage creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *HomePageManager) DeleteHomePage(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM homePage WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete HomePage-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("HomePage with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToHomePageProto accepts a homePage struct and returns a protobuf homePage struct.
func convertToHomePageProto(c *HomePage) *v1.HomePage {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.HomePage{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		Layout:    *safeNullUUID(c.Layout),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a homePage.
func (*HomePageManager) GetProtoList(l []*HomePage) []*v1.HomePage {
	list := []*v1.HomePage{}
	for _, v := range l {
		list = append(list, convertToHomePageProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a homePage.
func (*HomePageManager) GetProto(c *HomePage) *v1.HomePage {
	return convertToHomePageProto(c)
}

// BuildHomePageListQuery takes a filter and ordering object for a homePage.
// and returns an SQL string
func BuildHomePageListQuery(filters []*v1.HomePageFilterRule, orderings []*v1.HomePageOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, layout FROM home_page"
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
func (m *HomePageManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
