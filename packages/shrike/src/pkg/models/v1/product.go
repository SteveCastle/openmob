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

// Product is a type for product db element.
type Product struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	ProductType uuid.UUID
}

// ProductManager manages queries returning a product or list of products.
// It is configured with a db field to contain the db driver.
type ProductManager struct {
	db *sql.DB
}

// NewProductManager creates a product manager
func NewProductManager(db *sql.DB) *ProductManager {
	return &ProductManager{db: db}
}

// CRUD Methods for the ProductManager.

// CreateProduct creates a product.
func (m *ProductManager) CreateProduct(ctx context.Context, item *v1.CreateProduct) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO product (title, product_type) VALUES($1, $2)  RETURNING id;",
		item.Title, item.ProductType).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Product-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Product-> "+err.Error())
	}
	return &id, nil
}

// GetProduct gets a single product from the database by ID.
func (m *ProductManager) GetProduct(ctx context.Context, id string) (*Product, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Product by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, product_type FROM product WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Product-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Product-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%s' is not found", id))
	}

	// scan Product data into protobuf model
	var product Product

	if err := rows.Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt, &product.Title, &product.ProductType); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Product row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Product rows with ID='%s'",
			id))
	}
	return &product, nil
}

// ListProduct returns a slice of all products meeting the filter criteria.
func (m *ProductManager) ListProduct(ctx context.Context, filters []*v1.ProductFilterRule, orderings []*v1.ProductOrdering, limit int64) ([]*Product, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Product Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildProductListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Product-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Product{}
	for rows.Next() {
		product := new(Product)
		if err := rows.Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt, &product.Title, &product.ProductType); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Product row-> "+err.Error())
		}
		list = append(list, product)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Product-> "+err.Error())
	}
	return list, nil
}

// UpdateProduct runs an update query on the provided db and returns the rows affected as an int64.
func (m *ProductManager) UpdateProduct(ctx context.Context, item *v1.Product) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE product SET title=$2, product_type=$3 WHERE id=$1",
		item.ID, item.Title, item.ProductType)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Product-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteProduct creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ProductManager) DeleteProduct(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM product WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Product-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Product with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToProductProto accepts a product struct and returns a protobuf product struct.
func convertToProductProto(c *Product) *v1.Product {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Product{
		ID:          c.ID.String(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Title:       c.Title,
		ProductType: c.ProductType.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a product.
func (*ProductManager) GetProtoList(l []*Product) []*v1.Product {
	list := []*v1.Product{}
	for _, v := range l {
		list = append(list, convertToProductProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a product.
func (*ProductManager) GetProto(c *Product) *v1.Product {
	return convertToProductProto(c)
}

// BuildProductListQuery takes a filter and ordering object for a product.
// and returns an SQL string
func BuildProductListQuery(filters []*v1.ProductFilterRule, orderings []*v1.ProductOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, product_type FROM product"
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
func (m *ProductManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
