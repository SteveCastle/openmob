package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// CustomerCart is a type for customer_cart db element.
type CustomerCart struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
