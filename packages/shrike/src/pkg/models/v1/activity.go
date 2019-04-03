package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Activity is a type for activity db element.
type Activity struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Title        string
	ActivityType uuid.UUID
	Contact      uuid.UUID
	Cause        uuid.UUID
}
