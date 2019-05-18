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

// Photo is a type for photo db element.
type Photo struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	URI       string
	Width     int64
	Height    int64
}

// PhotoManager manages queries returning a photo or list of photos.
// It is configured with a db field to contain the db driver.
type PhotoManager struct {
	db *sql.DB
}

// NewPhotoManager creates a photo manager
func NewPhotoManager(db *sql.DB) *PhotoManager {
	return &PhotoManager{db: db}
}

// CRUD Methods for the PhotoManager.

// Create creates a photo.
func (m *PhotoManager) Create(ctx context.Context, item *v1.CreatePhoto) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO photo (uri, width, height) VALUES($1, $2, $3)  RETURNING id;",
		item.URI, item.Width, item.Height).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Photo-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Photo-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single photo from the database by ID.
func (m *PhotoManager) Get(ctx context.Context, id string) (*Photo, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Photo by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, uri, width, height FROM photo WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Photo-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Photo-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Photo with ID='%s' is not found", id))
	}

	// scan Photo data into protobuf model
	var photo Photo

	if err := rows.Scan(&photo.ID, &photo.CreatedAt, &photo.UpdatedAt, &photo.URI, &photo.Width, &photo.Height); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Photo row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Photo rows with ID='%s'",
			id))
	}
	return &photo, nil
}

// List returns a slice of all photos meeting the filter criteria.
func (m *PhotoManager) List(ctx context.Context, filters []*v1.PhotoFilterRule, orderings []*v1.PhotoOrdering, limit int64) ([]*Photo, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Photo Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildPhotoListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Photo-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Photo{}
	for rows.Next() {
		photo := new(Photo)
		if err := rows.Scan(&photo.ID, &photo.CreatedAt, &photo.UpdatedAt, &photo.URI, &photo.Width, &photo.Height); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Photo row-> "+err.Error())
		}
		list = append(list, photo)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Photo-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *PhotoManager) Update(ctx context.Context, item *v1.Photo) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE photo SET uri=$2, width=$3, height=$4 WHERE id=$1",
		item.ID, item.URI, item.Width, item.Height)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Photo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Photo with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *PhotoManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM photo WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Photo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Photo with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToPhotoProto accepts a photo struct and returns a protobuf photo struct.
func convertToPhotoProto(c *Photo) *v1.Photo {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Photo{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		URI:       c.URI,
		Width:     c.Width,
		Height:    c.Height,
	}
}

//GetProtoList returns a slice of protobuf typed struct of a photo.
func (*PhotoManager) GetProtoList(l []*Photo) []*v1.Photo {
	list := []*v1.Photo{}
	for _, v := range l {
		list = append(list, convertToPhotoProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a photo.
func (*PhotoManager) GetProto(c *Photo) *v1.Photo {
	return convertToPhotoProto(c)
}

// BuildPhotoListQuery takes a filter and ordering object for a photo.
// and returns an SQL string
func BuildPhotoListQuery(filters []*v1.PhotoFilterRule, orderings []*v1.PhotoOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, uri, width, height FROM photo"
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
func (m *PhotoManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
