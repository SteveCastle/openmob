package models

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid"
)

// Contact is a type for contact db element.
type Contact struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FirstName   sql.NullString
	MiddleName  sql.NullString
	LastName    sql.NullString
	Email       sql.NullString
	PhoneNumber sql.NullString
}
