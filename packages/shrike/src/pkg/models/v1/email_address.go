package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// EmailAddress is a type for email_address db element.
type EmailAddress struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Address   string
}
