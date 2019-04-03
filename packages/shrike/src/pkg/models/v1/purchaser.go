package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Purchaser is a type for purchaser db element.
type Purchaser struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CustomerOrder uuid.UUID
	Contact       uuid.UUID
	Cause         uuid.UUID
}
