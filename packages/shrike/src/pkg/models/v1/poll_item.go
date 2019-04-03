package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// PollItem is a type for poll_item db element.
type PollItem struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Poll      uuid.UUID
}
