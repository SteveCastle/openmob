package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// LandingPage is a type for landing_page db element.
type LandingPage struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Cause     uuid.UUID
	Layout    uuid.NullUUID
}
