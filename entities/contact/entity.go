package contact

import "github.com/google/uuid"

type Entity struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Type   string
	Value  string
}
