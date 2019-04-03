package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// DonationCampaign is a type for donation_campaign db element.
type DonationCampaign struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
