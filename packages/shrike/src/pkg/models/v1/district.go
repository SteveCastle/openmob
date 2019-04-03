package models

import (
	"time"

	"github.com/SteveCastle/openmob/packages/shrike/src/geography"
	uuid "github.com/gofrs/uuid"
)

// District is a type for district db element.
type District struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Geom         geography.NullRegion
	Title        string
	DistrictType uuid.UUID
}
