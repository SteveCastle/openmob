package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// ContactMembership is a type for contact_membership db element.
type ContactMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Contact   uuid.UUID
}
