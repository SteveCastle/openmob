package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Experiment is a type for experiment db element.
type Experiment struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	LandingPage uuid.NullUUID
}
