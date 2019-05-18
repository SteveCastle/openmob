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

// Donor is a type for donor db element.
type Donor struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CustomerOrder uuid.UUID
	Contact       uuid.UUID
	Cause         uuid.UUID
}

// DonorManager manages queries returning a donor or list of donors.
// It is configured with a db field to contain the db driver.
type DonorManager struct {
	db *sql.DB
}

// NewDonorManager creates a donor manager
func NewDonorManager(db *sql.DB) *DonorManager {
	return &DonorManager{db: db}
}

// CRUD Methods for the DonorManager.

// Create creates a donor.
func (m *DonorManager) Create(ctx context.Context, item *v1.CreateDonor) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO donor (customer_order, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		item.CustomerOrder, item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Donor-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Donor-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single donor from the database by ID.
func (m *DonorManager) Get(ctx context.Context, id string) (*Donor, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Donor by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, customer_order, contact, cause FROM donor WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Donor-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Donor-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Donor with ID='%s' is not found", id))
	}

	// scan Donor data into protobuf model
	var donor Donor

	if err := rows.Scan(&donor.ID, &donor.CreatedAt, &donor.UpdatedAt, &donor.CustomerOrder, &donor.Contact, &donor.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Donor row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Donor rows with ID='%s'",
			id))
	}
	return &donor, nil
}

// List returns a slice of all donors meeting the filter criteria.
func (m *DonorManager) List(ctx context.Context, filters []*v1.DonorFilterRule, orderings []*v1.DonorOrdering, limit int64) ([]*Donor, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Donor Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildDonorListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Donor-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Donor{}
	for rows.Next() {
		donor := new(Donor)
		if err := rows.Scan(&donor.ID, &donor.CreatedAt, &donor.UpdatedAt, &donor.CustomerOrder, &donor.Contact, &donor.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Donor row-> "+err.Error())
		}
		list = append(list, donor)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Donor-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *DonorManager) Update(ctx context.Context, item *v1.Donor) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE donor SET customer_order=$2, contact=$3, cause=$4 WHERE id=$1",
		item.ID, item.CustomerOrder, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Donor-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Donor with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *DonorManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM donor WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Donor-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Donor with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToDonorProto accepts a donor struct and returns a protobuf donor struct.
func convertToDonorProto(c *Donor) *v1.Donor {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Donor{
		ID:            c.ID.String(),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		CustomerOrder: c.CustomerOrder.String(),
		Contact:       c.Contact.String(),
		Cause:         c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a donor.
func (*DonorManager) GetProtoList(l []*Donor) []*v1.Donor {
	list := []*v1.Donor{}
	for _, v := range l {
		list = append(list, convertToDonorProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a donor.
func (*DonorManager) GetProto(c *Donor) *v1.Donor {
	return convertToDonorProto(c)
}

// BuildDonorListQuery takes a filter and ordering object for a donor.
// and returns an SQL string
func BuildDonorListQuery(filters []*v1.DonorFilterRule, orderings []*v1.DonorOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, customer_order, contact, cause FROM donor"
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
func (m *DonorManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
