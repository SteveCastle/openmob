package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Donor is a type for donor db element.
type Donor struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CustomerOrder uuid.UUID
	Contact       uuid.UUID
	Cause         uuid.UUID
}
