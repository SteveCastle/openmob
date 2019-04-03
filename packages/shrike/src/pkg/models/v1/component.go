package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Component is a type for component db element.
type Component struct {
	ID                      uuid.UUID
	CreatedAt               time.Time
	UpdatedAt               time.Time
	ComponentType           uuid.UUID
	ComponentImplementation uuid.UUID
	LayoutColumn            uuid.NullUUID
}
