package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Voter is a type for voter db element.
type Voter struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Contact   uuid.UUID
	Cause     uuid.UUID
}
