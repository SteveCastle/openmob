package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// DonationCampaignMembership is a type for donation_campaign_membership db element.
type DonationCampaignMembership struct {
	ID               uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Cause            uuid.UUID
	DonationCampaign uuid.UUID
}
