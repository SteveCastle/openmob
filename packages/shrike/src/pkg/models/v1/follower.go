package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Follower is a type for follower db element.
type Follower struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Contact   uuid.UUID
	Cause     uuid.UUID
}
