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

// ProductType is a type for product_type db element.
type ProductType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// ProductTypeManager manages queries returning a productType or list of productTypes.
// It is configured with a db field to contain the db driver.
type ProductTypeManager struct {
	db *sql.DB
}

// NewProductTypeManager creates a productType manager
func NewProductTypeManager(db *sql.DB) *ProductTypeManager {
	return &ProductTypeManager{db: db}
}

// CRUD Methods for the ProductTypeManager.

// Create creates a productType.
func (m *ProductTypeManager) Create(ctx context.Context, item *v1.CreateProductType) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO product_type (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ProductType-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ProductType-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single productType from the database by ID.
func (m *ProductTypeManager) Get(ctx context.Context, id string) (*ProductType, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ProductType by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM product_type WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductType-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductType-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%s' is not found", id))
	}

	// scan ProductType data into protobuf model
	var productType ProductType

	if err := rows.Scan(&productType.ID, &productType.CreatedAt, &productType.UpdatedAt, &productType.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductType row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ProductType rows with ID='%s'",
			id))
	}
	return &productType, nil
}

// List returns a slice of all productTypes meeting the filter criteria.
func (m *ProductTypeManager) List(ctx context.Context, filters []*v1.ProductTypeFilterRule, orderings []*v1.ProductTypeOrdering, limit int64) ([]*ProductType, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ProductType Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildProductTypeListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductType-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ProductType{}
	for rows.Next() {
		productType := new(ProductType)
		if err := rows.Scan(&productType.ID, &productType.CreatedAt, &productType.UpdatedAt, &productType.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductType row-> "+err.Error())
		}
		list = append(list, productType)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductType-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *ProductTypeManager) Update(ctx context.Context, item *v1.ProductType) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE product_type SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ProductType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ProductTypeManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM product_type WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ProductType-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductType with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToProductTypeProto accepts a productType struct and returns a protobuf productType struct.
func convertToProductTypeProto(c *ProductType) *v1.ProductType {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ProductType{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a productType.
func (*ProductTypeManager) GetProtoList(l []*ProductType) []*v1.ProductType {
	list := []*v1.ProductType{}
	for _, v := range l {
		list = append(list, convertToProductTypeProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a productType.
func (*ProductTypeManager) GetProto(c *ProductType) *v1.ProductType {
	return convertToProductTypeProto(c)
}

// BuildProductTypeListQuery takes a filter and ordering object for a productType.
// and returns an SQL string
func BuildProductTypeListQuery(filters []*v1.ProductTypeFilterRule, orderings []*v1.ProductTypeOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM product_type"
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
func (m *ProductTypeManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
