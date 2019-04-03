package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Candidate is a type for candidate db element.
type Candidate struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Election  uuid.UUID
}
