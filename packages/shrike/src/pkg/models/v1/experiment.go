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

// Experiment is a type for experiment db element.
type Experiment struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

// ExperimentManager manages queries returning a experiment or list of experiments.
// It is configured with a db field to contain the db driver.
type ExperimentManager struct {
	db *sql.DB
}

// NewExperimentManager creates a experiment manager
func NewExperimentManager(db *sql.DB) *ExperimentManager {
	return &ExperimentManager{db: db}
}

// CRUD Methods for the ExperimentManager.

// CreateExperiment creates a experiment.
func (m *ExperimentManager) CreateExperiment(ctx context.Context, item *v1.CreateExperiment) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO experiment (title) VALUES($1)  RETURNING id;",
		item.Title).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Experiment-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Experiment-> "+err.Error())
	}
	return &id, nil
}

// GetExperiment gets a single experiment from the database by ID.
func (m *ExperimentManager) GetExperiment(ctx context.Context, id string) (*Experiment, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Experiment by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title FROM experiment WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Experiment-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Experiment-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Experiment with ID='%s' is not found", id))
	}

	// scan Experiment data into protobuf model
	var experiment Experiment

	if err := rows.Scan(&experiment.ID, &experiment.CreatedAt, &experiment.UpdatedAt, &experiment.Title); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Experiment row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Experiment rows with ID='%s'",
			id))
	}
	return &experiment, nil
}

// ListExperiment returns a slice of all experiments meeting the filter criteria.
func (m *ExperimentManager) ListExperiment(ctx context.Context, filters []*v1.ExperimentFilterRule, orderings []*v1.ExperimentOrdering, limit int64) ([]*Experiment, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Experiment Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildExperimentListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Experiment-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Experiment{}
	for rows.Next() {
		experiment := new(Experiment)
		if err := rows.Scan(&experiment.ID, &experiment.CreatedAt, &experiment.UpdatedAt, &experiment.Title); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Experiment row-> "+err.Error())
		}
		list = append(list, experiment)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Experiment-> "+err.Error())
	}
	return list, nil
}

// UpdateExperiment runs an update query on the provided db and returns the rows affected as an int64.
func (m *ExperimentManager) UpdateExperiment(ctx context.Context, item *v1.Experiment) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE experiment SET title=$2 WHERE id=$1",
		item.ID, item.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Experiment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Experiment with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteExperiment creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ExperimentManager) DeleteExperiment(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM experiment WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Experiment-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Experiment with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToExperimentProto accepts a experiment struct and returns a protobuf experiment struct.
func convertToExperimentProto(c *Experiment) *v1.Experiment {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Experiment{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     c.Title,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a experiment.
func (*ExperimentManager) GetProtoList(l []*Experiment) []*v1.Experiment {
	list := []*v1.Experiment{}
	for _, v := range l {
		list = append(list, convertToExperimentProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a experiment.
func (*ExperimentManager) GetProto(c *Experiment) *v1.Experiment {
	return convertToExperimentProto(c)
}

// BuildExperimentListQuery takes a filter and ordering object for a experiment.
// and returns an SQL string
func BuildExperimentListQuery(filters []*v1.ExperimentFilterRule, orderings []*v1.ExperimentOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title FROM experiment"
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
func (m *ExperimentManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
