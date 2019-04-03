package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Delivery is a type for delivery db element.
type Delivery struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
