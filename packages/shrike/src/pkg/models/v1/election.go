package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Election is a type for election db element.
type Election struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
