package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// VolunteerOpportunityMembership is a type for volunteer_opportunity_membership db element.
type VolunteerOpportunityMembership struct {
	ID                   uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Cause                uuid.UUID
	VolunteerOpportunity uuid.UUID
}
