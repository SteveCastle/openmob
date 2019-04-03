package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// BoycottMembership is a type for boycott_membership db element.
type BoycottMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Boycott   uuid.UUID
}
