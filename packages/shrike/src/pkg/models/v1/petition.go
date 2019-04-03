package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Petition is a type for petition db element.
type Petition struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
