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

// PetitionSigner is a type for petition_signer db element.
type PetitionSigner struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Petition  uuid.UUID
	Contact   uuid.UUID
	Cause     uuid.UUID
}

// PetitionSignerManager manages queries returning a petitionSigner or list of petitionSigners.
// It is configured with a db field to contain the db driver.
type PetitionSignerManager struct {
	db *sql.DB
}

// NewPetitionSignerManager creates a petitionSigner manager
func NewPetitionSignerManager(db *sql.DB) *PetitionSignerManager {
	return &PetitionSignerManager{db: db}
}

// CRUD Methods for the PetitionSignerManager.

// Create creates a petitionSigner.
func (m *PetitionSignerManager) Create(ctx context.Context, item *v1.CreatePetitionSigner) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO petition_signer (petition, contact, cause) VALUES($1, $2, $3)  RETURNING id;",
		item.Petition, item.Contact, item.Cause).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into PetitionSigner-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created PetitionSigner-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single petitionSigner from the database by ID.
func (m *PetitionSignerManager) Get(ctx context.Context, id string) (*PetitionSigner, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query PetitionSigner by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, petition, contact, cause FROM petition_signer WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionSigner-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionSigner-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%s' is not found", id))
	}

	// scan PetitionSigner data into protobuf model
	var petitionSigner PetitionSigner

	if err := rows.Scan(&petitionSigner.ID, &petitionSigner.CreatedAt, &petitionSigner.UpdatedAt, &petitionSigner.Petition, &petitionSigner.Contact, &petitionSigner.Cause); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionSigner row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple PetitionSigner rows with ID='%s'",
			id))
	}
	return &petitionSigner, nil
}

// List returns a slice of all petitionSigners meeting the filter criteria.
func (m *PetitionSignerManager) List(ctx context.Context, filters []*v1.PetitionSignerFilterRule, orderings []*v1.PetitionSignerOrdering, limit int64) ([]*PetitionSigner, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in PetitionSigner Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPetitionSignerListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from PetitionSigner-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*PetitionSigner{}
	for rows.Next() {
		petitionSigner := new(PetitionSigner)
		if err := rows.Scan(&petitionSigner.ID, &petitionSigner.CreatedAt, &petitionSigner.UpdatedAt, &petitionSigner.Petition, &petitionSigner.Contact, &petitionSigner.Cause); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from PetitionSigner row-> "+err.Error())
		}
		list = append(list, petitionSigner)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from PetitionSigner-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *PetitionSignerManager) Update(ctx context.Context, item *v1.PetitionSigner) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE petition_signer SET petition=$2, contact=$3, cause=$4 WHERE id=$1",
		item.ID, item.Petition, item.Contact, item.Cause)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update PetitionSigner-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PetitionSignerManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM petitionSigner WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete PetitionSigner-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("PetitionSigner with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPetitionSignerProto accepts a petitionSigner struct and returns a protobuf petitionSigner struct.
func convertToPetitionSignerProto(c *PetitionSigner) *v1.PetitionSigner {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.PetitionSigner{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Petition:  c.Petition.String(),
		Contact:   c.Contact.String(),
		Cause:     c.Cause.String(),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a petitionSigner.
func (*PetitionSignerManager) GetProtoList(l []*PetitionSigner) []*v1.PetitionSigner {
	list := []*v1.PetitionSigner{}
	for _, v := range l {
		list = append(list, convertToPetitionSignerProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a petitionSigner.
func (*PetitionSignerManager) GetProto(c *PetitionSigner) *v1.PetitionSigner {
	return convertToPetitionSignerProto(c)
}

// BuildPetitionSignerListQuery takes a filter and ordering object for a petitionSigner.
// and returns an SQL string
func BuildPetitionSignerListQuery(filters []*v1.PetitionSignerFilterRule, orderings []*v1.PetitionSignerOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, petition, contact, cause FROM petition_signer"
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
func (m *PetitionSignerManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
