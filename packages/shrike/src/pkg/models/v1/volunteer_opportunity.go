package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// VolunteerOpportunity is a type for volunteer_opportunity db element.
type VolunteerOpportunity struct {
	ID                       uuid.UUID
	CreatedAt                time.Time
	UpdatedAt                time.Time
	Title                    string
	VolunteerOpportunityType uuid.NullUUID
}
