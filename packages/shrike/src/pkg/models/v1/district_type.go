package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// DistrictType is a type for district_type db element.
type DistrictType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
