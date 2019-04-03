package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// VolunteerOpportunityType is a type for volunteer_opportunity_type db element.
type VolunteerOpportunityType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
