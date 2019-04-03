package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Issue is a type for issue db element.
type Issue struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Election  uuid.UUID
}
