package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// LayoutColumn is a type for layout_column db element.
type LayoutColumn struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	LayoutRow uuid.UUID
	Width     int
}
