package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Poll is a type for poll db element.
type Poll struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
