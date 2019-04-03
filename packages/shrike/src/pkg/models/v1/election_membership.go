package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// ElectionMembership is a type for election_membership db element.
type ElectionMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Election  uuid.UUID
}
