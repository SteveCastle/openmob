package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// Product is a type for product db element.
type Product struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	ProductType uuid.UUID
}
