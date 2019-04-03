package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// ProductMembership is a type for product_membership db element.
type ProductMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Product   uuid.UUID
}
