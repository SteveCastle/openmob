package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// MailingAddress is a type for mailing_address db element.
type MailingAddress struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	StreetAddress string
	City          string
	State         string
	ZipCode       string
}
