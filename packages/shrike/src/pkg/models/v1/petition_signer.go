package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// PetitionSigner is a type for petition_signer db element.
type PetitionSigner struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Petition  uuid.UUID
	Contact   uuid.UUID
	Cause     uuid.UUID
}
