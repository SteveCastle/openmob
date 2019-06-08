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

// ProductMembership is a type for product_membership db element.
type ProductMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Product   uuid.UUID
}

// ProductMembershipManager manages queries returning a productMembership or list of productMemberships.
// It is configured with a db field to contain the db driver.
type ProductMembershipManager struct {
	db *sql.DB
}

// NewProductMembershipManager creates a productMembership manager
func NewProductMembershipManager(db *sql.DB) *ProductMembershipManager {
	return &ProductMembershipManager{db: db}
}

// CRUD Methods for the ProductMembershipManager.

// Create creates a productMembership.
func (m *ProductMembershipManager) Create(ctx context.Context, item *v1.CreateProductMembership) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO product_membership (cause, product) VALUES($1, $2)  RETURNING id;",
		item.Cause, item.Product).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ProductMembership-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ProductMembership-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single productMembership from the database by ID.
func (m *ProductMembershipManager) Get(ctx context.Context, id string) (*ProductMembership, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ProductMembership by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, cause, product FROM product_membership WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductMembership-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductMembership-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%s' is not found", id))
	}

	// scan ProductMembership data into protobuf model
	var productMembership ProductMembership

	if err := rows.Scan(&productMembership.ID, &productMembership.CreatedAt, &productMembership.UpdatedAt, &productMembership.Cause, &productMembership.Product); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductMembership row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ProductMembership rows with ID='%s'",
			id))
	}
	return &productMembership, nil
}

// List returns a slice of all productMemberships meeting the filter criteria.
func (m *ProductMembershipManager) List(ctx context.Context, filters []*v1.ProductMembershipFilterRule, orderings []*v1.ProductMembershipOrdering, limit int64) ([]*ProductMembership, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ProductMembership Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildProductMembershipListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ProductMembership-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ProductMembership{}
	for rows.Next() {
		productMembership := new(ProductMembership)
		if err := rows.Scan(&productMembership.ID, &productMembership.CreatedAt, &productMembership.UpdatedAt, &productMembership.Cause, &productMembership.Product); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ProductMembership row-> "+err.Error())
		}
		list = append(list, productMembership)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ProductMembership-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *ProductMembershipManager) Update(ctx context.Context, item *v1.ProductMembership) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE product_membership SET cause=$2, product=$3 WHERE id=$1",
		item.ID, item.Cause, item.Product)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ProductMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ProductMembershipManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM product_membership WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ProductMembership-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ProductMembership with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToProductMembershipProto accepts a productMembership struct and returns a protobuf productMembership struct.
func convertToProductMembershipProto(c *ProductMembership) *v1.ProductMembership {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ProductMembership{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Cause:     c.Cause.String(),
		Product:   c.Product.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a productMembership.
func (*ProductMembershipManager) GetProtoList(l []*ProductMembership) []*v1.ProductMembership {
	list := []*v1.ProductMembership{}
	for _, v := range l {
		list = append(list, convertToProductMembershipProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a productMembership.
func (*ProductMembershipManager) GetProto(c *ProductMembership) *v1.ProductMembership {
	return convertToProductMembershipProto(c)
}

// BuildProductMembershipListQuery takes a filter and ordering object for a productMembership.
// and returns an SQL string
func BuildProductMembershipListQuery(filters []*v1.ProductMembershipFilterRule, orderings []*v1.ProductMembershipOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, cause, product FROM product_membership"
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
func (m *ProductMembershipManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
