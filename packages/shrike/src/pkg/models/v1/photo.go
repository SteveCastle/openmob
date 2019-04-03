package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Photo is a type for photo db element.
type Photo struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	URI       string
	Width     int
	Height    int
}
