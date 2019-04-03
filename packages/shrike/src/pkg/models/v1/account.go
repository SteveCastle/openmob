package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Account is a type for account db element.
type Account struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
}
