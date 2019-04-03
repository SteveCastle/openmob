package models

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid"
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
