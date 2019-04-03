package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Office is a type for office db element.
type Office struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Election  uuid.NullUUID
}
