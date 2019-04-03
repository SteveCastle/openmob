package models

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid"
)

// LayoutRow is a type for layout_row db element.
type LayoutRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Layout    uuid.UUID
	Container sql.NullBool
}
