package models

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/lib/pq"
)

// Field is a type for field db element.
type Field struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	FieldType     uuid.UUID
	StringValue   sql.NullString
	IntValue      sql.NullInt64
	FloatValue    sql.NullFloat64
	BooleanValue  sql.NullBool
	DateTimeValue pq.NullTime
	Component     uuid.NullUUID
}
