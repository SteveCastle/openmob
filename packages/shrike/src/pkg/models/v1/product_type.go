package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// ProductType is a type for product_type db element.
type ProductType struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}
