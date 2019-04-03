package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// HomePage is a type for home_page db element.
type HomePage struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Layout    uuid.NullUUID
}
