package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// LayoutType is a type for layout_type db element.
type LayoutType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
