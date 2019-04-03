package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// LiveEventType is a type for live_event_type db element.
type LiveEventType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
