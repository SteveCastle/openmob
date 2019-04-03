package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Agent is a type for agent db element.
type Agent struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Account   uuid.UUID
}
