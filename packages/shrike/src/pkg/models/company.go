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

// Company is a type for company db element.
type Company struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// CompanyManager manages queries returning a company or list of companys.
// It is configured with a db field to contain the db driver.
type CompanyManager struct {
	db *sql.DB
}

// NewCompanyManager creates a company manager
func NewCompanyManager(db *sql.DB) *CompanyManager {
	return &CompanyManager{db: db}
}

// CRUD Methods for the CompanyManager.

// Create creates a company.
func (m *CompanyManager) Create(ctx context.Context, item *v1.CreateCompany) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO company (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Company-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Company-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single company from the database by ID.
func (m *CompanyManager) Get(ctx context.Context, id string) (*Company, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Company by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM company WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Company-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Company-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Company with ID='%s' is not found", id))
	}

	// scan Company data into protobuf model
	var company Company

	if err := rows.Scan(&company.ID, &company.CreatedAt, &company.UpdatedAt, &company.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Company row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Company rows with ID='%s'",
			id))
	}
	return &company, nil
}

// List returns a slice of all companys meeting the filter criteria.
func (m *CompanyManager) List(ctx context.Context, filters []*v1.CompanyFilterRule, orderings []*v1.CompanyOrdering, limit int64) ([]*Company, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Company Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildCompanyListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Company-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Company{}
	for rows.Next() {
		company := new(Company)
		if err := rows.Scan(&company.ID, &company.CreatedAt, &company.UpdatedAt, &company.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Company row-> "+err.Error())
		}
		list = append(list, company)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Company-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *CompanyManager) Update(ctx context.Context, item *v1.Company) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE company SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Company-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Company with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *CompanyManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM company WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Company-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Company with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToCompanyProto accepts a company struct and returns a protobuf company struct.
func convertToCompanyProto(c *Company) *v1.Company {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Company{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a company.
func (*CompanyManager) GetProtoList(l []*Company) []*v1.Company {
	list := []*v1.Company{}
	for _, v := range l {
		list = append(list, convertToCompanyProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a company.
func (*CompanyManager) GetProto(c *Company) *v1.Company {
	return convertToCompanyProto(c)
}

// BuildCompanyListQuery takes a filter and ordering object for a company.
// and returns an SQL string
func BuildCompanyListQuery(filters []*v1.CompanyFilterRule, orderings []*v1.CompanyOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM company"
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
func (m *CompanyManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
