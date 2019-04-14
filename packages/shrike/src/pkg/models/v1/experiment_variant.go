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

// ExperimentVariant is a type for experiment_variant db element.
type ExperimentVariant struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	VariantType sql.NullString
	Experiment  uuid.UUID
	LandingPage uuid.NullUUID
	Field       uuid.NullUUID
	Component   uuid.NullUUID
}

// ExperimentVariantManager manages queries returning a experimentVariant or list of experimentVariants.
// It is configured with a db field to contain the db driver.
type ExperimentVariantManager struct {
	db *sql.DB
}

// NewExperimentVariantManager creates a experimentVariant manager
func NewExperimentVariantManager(db *sql.DB) *ExperimentVariantManager {
	return &ExperimentVariantManager{db: db}
}

// CRUD Methods for the ExperimentVariantManager.

// CreateExperimentVariant creates a experimentVariant.
func (m *ExperimentVariantManager) CreateExperimentVariant(ctx context.Context, item *v1.CreateExperimentVariant) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO experiment_variant (title, variant_type, experiment, landing_page, field, component) VALUES($1, $2, $3, $4, $5, $6)  RETURNING id;",
		item.Title, item.VariantType, item.Experiment, item.LandingPage, item.Field, item.Component).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ExperimentVariant-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ExperimentVariant-> "+err.Error())
	}
	return &id, nil
}

// GetExperimentVariant gets a single experimentVariant from the database by ID.
func (m *ExperimentVariantManager) GetExperimentVariant(ctx context.Context, id string) (*ExperimentVariant, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query ExperimentVariant by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, title, variant_type, experiment, landing_page, field, component FROM experiment_variant WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ExperimentVariant-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ExperimentVariant-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ExperimentVariant with ID='%s' is not found", id))
	}

	// scan ExperimentVariant data into protobuf model
	var experimentVariant ExperimentVariant

	if err := rows.Scan(&experimentVariant.ID, &experimentVariant.CreatedAt, &experimentVariant.UpdatedAt, &experimentVariant.Title, &experimentVariant.VariantType, &experimentVariant.Experiment, &experimentVariant.LandingPage, &experimentVariant.Field, &experimentVariant.Component); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ExperimentVariant row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ExperimentVariant rows with ID='%s'",
			id))
	}
	return &experimentVariant, nil
}

// ListExperimentVariant returns a slice of all experimentVariants meeting the filter criteria.
func (m *ExperimentVariantManager) ListExperimentVariant(ctx context.Context, filters []*v1.ExperimentVariantFilterRule, orderings []*v1.ExperimentVariantOrdering, limit int64) ([]*ExperimentVariant, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in ExperimentVariant Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildExperimentVariantListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ExperimentVariant-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*ExperimentVariant{}
	for rows.Next() {
		experimentVariant := new(ExperimentVariant)
		if err := rows.Scan(&experimentVariant.ID, &experimentVariant.CreatedAt, &experimentVariant.UpdatedAt, &experimentVariant.Title, &experimentVariant.VariantType, &experimentVariant.Experiment, &experimentVariant.LandingPage, &experimentVariant.Field, &experimentVariant.Component); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ExperimentVariant row-> "+err.Error())
		}
		list = append(list, experimentVariant)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ExperimentVariant-> "+err.Error())
	}
	return list, nil
}

// UpdateExperimentVariant runs an update query on the provided db and returns the rows affected as an int64.
func (m *ExperimentVariantManager) UpdateExperimentVariant(ctx context.Context, item *v1.ExperimentVariant) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE experiment_variant SET title=$2, variant_type=$3, experiment=$4, landing_page=$5, field=$6, component=$7 WHERE id=$1",
		item.ID, item.Title, item.VariantType, item.Experiment, item.LandingPage, item.Field, item.Component)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ExperimentVariant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ExperimentVariant with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//DeleteExperimentVariant creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *ExperimentVariantManager) DeleteExperimentVariant(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM experimentVariant WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ExperimentVariant-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ExperimentVariant with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToExperimentVariantProto accepts a experimentVariant struct and returns a protobuf experimentVariant struct.
func convertToExperimentVariantProto(c *ExperimentVariant) *v1.ExperimentVariant {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.ExperimentVariant{
		ID:          c.ID.String(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Title:       c.Title,
		VariantType: *safeNullString(c.VariantType),
		Experiment:  c.Experiment.String(),
		LandingPage: *safeNullUUID(c.LandingPage),
		Field:       *safeNullUUID(c.Field),
		Component:   *safeNullUUID(c.Component),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a experimentVariant.
func (*ExperimentVariantManager) GetProtoList(l []*ExperimentVariant) []*v1.ExperimentVariant {
	list := []*v1.ExperimentVariant{}
	for _, v := range l {
		list = append(list, convertToExperimentVariantProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a experimentVariant.
func (*ExperimentVariantManager) GetProto(c *ExperimentVariant) *v1.ExperimentVariant {
	return convertToExperimentVariantProto(c)
}

// BuildExperimentVariantListQuery takes a filter and ordering object for a experimentVariant.
// and returns an SQL string
func BuildExperimentVariantListQuery(filters []*v1.ExperimentVariantFilterRule, orderings []*v1.ExperimentVariantOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, title, variant_type, experiment, landing_page, field, component FROM experiment_variant"
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
func (m *ExperimentVariantManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
