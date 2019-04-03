package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// ComponentImplementation is a type for component_implementation db element.
type ComponentImplementation struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Path      string
}
