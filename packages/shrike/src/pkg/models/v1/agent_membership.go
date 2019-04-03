package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

// AgentMembership is a type for agent_membership db element.
type AgentMembership struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Cause     uuid.UUID
	Agent     uuid.UUID
}
