package models

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/lib/pq"
)

// FieldType is a type for field_type db element.
type FieldType struct {
	ID                   uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Title                string
	DataType             string
	PropName             string
	StringValueDefault   sql.NullString
	IntValueDefault      sql.NullInt64
	FloatValueDefault    sql.NullFloat64
	BooleanValueDefault  sql.NullBool
	DateTimeValueDefault pq.NullTime
	ComponentType        uuid.NullUUID
}
