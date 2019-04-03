package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// PetitionMembership is a type for petition_membership db element.
type PetitionMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Petition  uuid.UUID
}
