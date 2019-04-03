package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Territory is a type for territory db element.
type Territory struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
