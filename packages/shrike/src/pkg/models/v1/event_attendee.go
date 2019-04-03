package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// EventAttendee is a type for event_attendee db element.
type EventAttendee struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	LiveEvent uuid.UUID
	Contact   uuid.UUID
	Cause     uuid.UUID
}
