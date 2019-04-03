package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// LiveEvent is a type for live_event db element.
type LiveEvent struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Title         string
	LiveEventType uuid.UUID
}
