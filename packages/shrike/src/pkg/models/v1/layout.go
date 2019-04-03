package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Layout is a type for layout db element.
type Layout struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	LayoutType uuid.NullUUID
}
