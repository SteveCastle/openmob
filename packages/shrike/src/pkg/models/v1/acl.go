package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// ACL is a type for acl db element.
type ACL struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
