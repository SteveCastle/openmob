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

// Office is a type for office db element.
type Office struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Election  uuid.NullUUID
}

// OfficeManager manages queries returning a office or list of offices.
// It is configured with a db field to contain the db driver.
type OfficeManager struct {
	db *sql.DB
}

// NewOfficeManager creates a office manager
func NewOfficeManager(db *sql.DB) *OfficeManager {
	return &OfficeManager{db: db}
}

// CRUD Methods for the OfficeManager.

// CreateOffice creates a office.
func (m *OfficeManager) CreateOffice(ctx context.Context, item *v1.CreateOffice) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO office (title, election) VALUES($1, $2)  RETURNING id;",
		item.Title, item.Election).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Office-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Office-> "+err.Error())
	}
	return &id, nil
}

// GetOffice gets a single office from the database by ID.
func (m *OfficeManager) GetOffice(ctx context.Context, id string) (*Office, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Office by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, election FROM office WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Office-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Office-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Office with ID='%s' is not found", id))
	}

	// scan Office data into protobuf model
	var office Office

	if err := rows.Scan(&office.ID, &office.CreatedAt, &office.UpdatedAt, &office.Title, &office.Election); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Office row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Office rows with ID='%s'",
			id))
	}
	return &office, nil
}

// ListOffice returns a slice of all offices meeting the filter criteria.
func (m *OfficeManager) ListOffice(ctx context.Context, filters []*v1.OfficeFilterRule, orderings []*v1.OfficeOrdering, limit int64) ([]*Office, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Office Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildOfficeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Office-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Office{}
	for rows.Next() {
		office := new(Office)
		if err := rows.Scan(&office.ID, &office.CreatedAt, &office.UpdatedAt, &office.Title, &office.Election); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Office row-> "+err.Error())
		}
		list = append(list, office)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Office-> "+err.Error())
	}
	return list, nil
}

// UpdateOffice runs an update query on the provided db and returns the rows affected as an int64.
func (m *OfficeManager) UpdateOffice(ctx context.Context, item *v1.Office) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE office SET title=$2, election=$3 WHERE id=$1",
		item.ID, item.Title, item.Election)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Office-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Office with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteOffice creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *OfficeManager) DeleteOffice(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM office WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Office-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Office with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToOfficeProto accepts a office struct and returns a protobuf office struct.
func convertToOfficeProto(c *Office) *v1.Office {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Office{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
		Election:  *safeNullUUID(c.Election),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a office.
func (*OfficeManager) GetProtoList(l []*Office) []*v1.Office {
	list := []*v1.Office{}
	for _, v := range l {
		list = append(list, convertToOfficeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a office.
func (*OfficeManager) GetProto(c *Office) *v1.Office {
	return convertToOfficeProto(c)
}

// BuildOfficeListQuery takes a filter and ordering object for a office.
// and returns an SQL string
func BuildOfficeListQuery(filters []*v1.OfficeFilterRule, orderings []*v1.OfficeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, election FROM office"
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
func (m *OfficeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
