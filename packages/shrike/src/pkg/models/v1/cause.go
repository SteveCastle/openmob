package models

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid"
)

// Cause is a type for cause db element.
type Cause struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Slug      string
	Summary   sql.NullString
	HomePage  uuid.NullUUID
	Photo     uuid.NullUUID
}
