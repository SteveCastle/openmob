package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// CustomerOrder is a type for customer_order db element.
type CustomerOrder struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CustomerCart uuid.UUID
}
