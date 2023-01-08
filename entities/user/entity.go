package user

import (
	"fmt"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func ParseID(id string) (ID, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return ID(uuid.Nil), fmt.Errorf("uuid.Parse() id = %s\terr = %w", id, err)
	}

	return ID(userID), nil
}

type Entity struct {
	ID   ID
	Name string
	Age  int
}

func (e *Entity) GenerateID() error {
	e.ID = ID(uuid.New())
	return nil
}
