package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Payment is a type for payment db element.
type Payment struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CustomerOrder uuid.UUID
}
