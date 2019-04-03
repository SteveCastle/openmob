package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// PollRespondant is a type for poll_respondant db element.
type PollRespondant struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Poll      uuid.UUID
	Contact   uuid.UUID
	Cause     uuid.UUID
}
