package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// PhoneNumber is a type for phone_number db element.
type PhoneNumber struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PhoneNumber string
}
