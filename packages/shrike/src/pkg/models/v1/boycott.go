package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Boycott is a type for boycott db element.
type Boycott struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
