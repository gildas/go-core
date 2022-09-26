package core

import "github.com/google/uuid"

// Identifiable describes that can get their Identifier as a UUID
type Identifiable interface {
	GetID() uuid.UUID
}
