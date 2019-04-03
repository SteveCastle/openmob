package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// LiveEventMembership is a type for live_event_membership db element.
type LiveEventMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	LiveEvent uuid.UUID
}
