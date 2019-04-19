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

// ACL is a type for acl db element.
type ACL struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ACLManager manages queries returning a acl or list of acls.
// It is configured with a db field to contain the db driver.
type ACLManager struct {
	db *sql.DB
}

// NewACLManager creates a acl manager
func NewACLManager(db *sql.DB) *ACLManager {
	return &ACLManager{db: db}
}

// CRUD Methods for the ACLManager.

// Create creates a acl.
func (m *ACLManager) Create(ctx context.Context, item *v1.CreateACL) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO acl () VALUES()  RETURNING id;").Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ACL-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ACL-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single acl from the database by ID.
func (m *ACLManager) Get(ctx context.Context, id string) (*ACL, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ACL by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at FROM acl WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ACL-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ACL-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ACL with ID='%s' is not found", id))
	}

	// scan ACL data into protobuf model
	var acl ACL

	if err := rows.Scan(&acl.ID, &acl.CreatedAt, &acl.UpdatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ACL row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ACL rows with ID='%s'",
			id))
	}
	return &acl, nil
}

// List returns a slice of all acls meeting the filter criteria.
func (m *ACLManager) List(ctx context.Context, filters []*v1.ACLFilterRule, orderings []*v1.ACLOrdering, limit int64) ([]*ACL, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ACL Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildACLListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ACL-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ACL{}
	for rows.Next() {
		acl := new(ACL)
		if err := rows.Scan(&acl.ID, &acl.CreatedAt, &acl.UpdatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ACL row-> "+err.Error())
		}
		list = append(list, acl)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ACL-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *ACLManager) Update(ctx context.Context, item *v1.ACL) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE acl SET  WHERE id=$1",
		item.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ACL-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ACL with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ACLManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM acl WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ACL-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ACL with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToACLProto accepts a acl struct and returns a protobuf acl struct.
func convertToACLProto(c *ACL) *v1.ACL {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ACL{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a acl.
func (*ACLManager) GetProtoList(l []*ACL) []*v1.ACL {
	list := []*v1.ACL{}
	for _, v := range l {
		list = append(list, convertToACLProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a acl.
func (*ACLManager) GetProto(c *ACL) *v1.ACL {
	return convertToACLProto(c)
}

// BuildACLListQuery takes a filter and ordering object for a acl.
// and returns an SQL string
func BuildACLListQuery(filters []*v1.ACLFilterRule, orderings []*v1.ACLOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at FROM acl"
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
func (m *ACLManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
