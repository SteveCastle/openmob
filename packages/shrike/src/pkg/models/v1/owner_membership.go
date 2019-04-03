package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// OwnerMembership is a type for owner_membership db element.
type OwnerMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Account   uuid.UUID
}
