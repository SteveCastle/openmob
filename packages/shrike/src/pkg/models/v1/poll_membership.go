package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// PollMembership is a type for poll_membership db element.
type PollMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Petition  uuid.UUID
}
