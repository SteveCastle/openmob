package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Company is a type for company db element.
type Company struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
