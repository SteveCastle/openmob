package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// ActivityType is a type for activity_type db element.
type ActivityType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
