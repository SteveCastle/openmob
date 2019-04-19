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

// Petition is a type for petition db element.
type Petition struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// PetitionManager manages queries returning a petition or list of petitions.
// It is configured with a db field to contain the db driver.
type PetitionManager struct {
	db *sql.DB
}

// NewPetitionManager creates a petition manager
func NewPetitionManager(db *sql.DB) *PetitionManager {
	return &PetitionManager{db: db}
}

// CRUD Methods for the PetitionManager.

// Create creates a petition.
func (m *PetitionManager) Create(ctx context.Context, item *v1.CreatePetition) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO petition (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Petition-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Petition-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single petition from the database by ID.
func (m *PetitionManager) Get(ctx context.Context, id string) (*Petition, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Petition by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM petition WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Petition-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Petition-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%s' is not found", id))
	}

	// scan Petition data into protobuf model
	var petition Petition

	if err := rows.Scan(&petition.ID, &petition.CreatedAt, &petition.UpdatedAt, &petition.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Petition row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Petition rows with ID='%s'",
			id))
	}
	return &petition, nil
}

// List returns a slice of all petitions meeting the filter criteria.
func (m *PetitionManager) List(ctx context.Context, filters []*v1.PetitionFilterRule, orderings []*v1.PetitionOrdering, limit int64) ([]*Petition, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Petition Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPetitionListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Petition-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Petition{}
	for rows.Next() {
		petition := new(Petition)
		if err := rows.Scan(&petition.ID, &petition.CreatedAt, &petition.UpdatedAt, &petition.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Petition row-> "+err.Error())
		}
		list = append(list, petition)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Petition-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *PetitionManager) Update(ctx context.Context, item *v1.Petition) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE petition SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Petition-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PetitionManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM petition WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Petition-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Petition with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPetitionProto accepts a petition struct and returns a protobuf petition struct.
func convertToPetitionProto(c *Petition) *v1.Petition {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Petition{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a petition.
func (*PetitionManager) GetProtoList(l []*Petition) []*v1.Petition {
	list := []*v1.Petition{}
	for _, v := range l {
		list = append(list, convertToPetitionProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a petition.
func (*PetitionManager) GetProto(c *Petition) *v1.Petition {
	return convertToPetitionProto(c)
}

// BuildPetitionListQuery takes a filter and ordering object for a petition.
// and returns an SQL string
func BuildPetitionListQuery(filters []*v1.PetitionFilterRule, orderings []*v1.PetitionOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM petition"
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
func (m *PetitionManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
