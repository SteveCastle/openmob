package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Volunteer is a type for volunteer db element.
type Volunteer struct {
	ID                   uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	VolunteerOpportunity uuid.UUID
	Contact              uuid.UUID
	Cause                uuid.UUID
}
