package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// ComponentType is a type for component_type db element.
type ComponentType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
