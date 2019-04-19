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

// Note is a type for note db element.
type Note struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Contact   uuid.UUID
	Cause     uuid.UUID
	Body      sql.NullString
}

// NoteManager manages queries returning a note or list of notes.
// It is configured with a db field to contain the db driver.
type NoteManager struct {
	db *sql.DB
}

// NewNoteManager creates a note manager
func NewNoteManager(db *sql.DB) *NoteManager {
	return &NoteManager{db: db}
}

// CRUD Methods for the NoteManager.

// Create creates a note.
func (m *NoteManager) Create(ctx context.Context, item *v1.CreateNote) (*string, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id string
	// Execute INSERT query and then scan the resulting id into id string.
	err = c.QueryRowContext(ctx, "INSERT INTO note (contact, cause, body) VALUES($1, $2, $3)  RETURNING id;",
		item.Contact, item.Cause, item.Body).Scan(&id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Note-> "+err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Note-> "+err.Error())
	}
	return &id, nil
}

// Get gets a single note from the database by ID.
func (m *NoteManager) Get(ctx context.Context, id string) (*Note, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Note by ID
	rows, err := c.QueryContext(ctx, "SELECT id, created_at, updated_at, contact, cause, body FROM note WHERE id=$1",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Note-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Note-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%s' is not found", id))
	}

	// scan Note data into protobuf model
	var note Note

	if err := rows.Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt, &note.Contact, &note.Cause, &note.Body); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Note row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Note rows with ID='%s'",
			id))
	}
	return &note, nil
}

// List returns a slice of all notes meeting the filter criteria.
func (m *NoteManager) List(ctx context.Context, filters []*v1.NoteFilterRule, orderings []*v1.NoteOrdering, limit int64) ([]*Note, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Generate SQL to select all columns in Note Table
	// TODO: Allow column selection.
	// Then generate filtering and ordering sql and finally run query.
	querySQL := BuildNoteListQuery(filters, orderings, limit)

	// Execute query and scan into return type.
	rows, err := c.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Note-> "+err.Error())
	}
	defer rows.Close()

	// Scan the results into a slice.
	list := []*Note{}
	for rows.Next() {
		note := new(Note)
		if err := rows.Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt, &note.Contact, &note.Cause, &note.Body); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Note row-> "+err.Error())
		}
		list = append(list, note)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Note-> "+err.Error())
	}
	return list, nil
}

// Update runs an update query on the provided db and returns the rows affected as an int64.
func (m *NoteManager) Update(ctx context.Context, item *v1.Note) (*int64, error) {

	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "UPDATE note SET contact=$2, cause=$3, body=$4 WHERE id=$1",
		item.ID, item.Contact, item.Cause, item.Body)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update Note-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%s' is not found",
			item.ID))
	}
	return &rows, nil
}

//Delete creates and executes DELETE sql on a provided id and returns the number of rows affected.
func (m *NoteManager) Delete(ctx context.Context, id string) (*int64, error) {
	c, err := m.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM note WHERE id=$1", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete Note-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Note with ID='%s' is not found",
			id))
	}

	return &rows, nil
}

// convertToNoteProto accepts a note struct and returns a protobuf note struct.
func convertToNoteProto(c *Note) *v1.Note {
	createdAt, _ := convertTimeToProto(c.CreatedAt)
	updatedAt, _ := convertTimeToProto(c.UpdatedAt)

	return &v1.Note{
		ID:        c.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Contact:   c.Contact.String(),
		Cause:     c.Cause.String(),
		Body:      *safeNullString(c.Body),
	}
}

//GetProtoList returns a slice of protobuf typed struct of a note.
func (*NoteManager) GetProtoList(l []*Note) []*v1.Note {
	list := []*v1.Note{}
	for _, v := range l {
		list = append(list, convertToNoteProto(v))
	}
	return list
}

//GetProto returns a single protobuf typed struct of a note.
func (*NoteManager) GetProto(c *Note) *v1.Note {
	return convertToNoteProto(c)
}

// BuildNoteListQuery takes a filter and ordering object for a note.
// and returns an SQL string
func BuildNoteListQuery(filters []*v1.NoteFilterRule, orderings []*v1.NoteOrdering, limit int64) string {
	baseSQL := "SELECT id, created_at, updated_at, contact, cause, body FROM note"
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
func (m *NoteManager) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := m.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
