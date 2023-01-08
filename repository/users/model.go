package users

import (
	"fmt"

	"github.com/magartifex/kapran/entities/user"
)

type model struct {
	ID   string
	Name string
	Age  int
}

func (m *model) marshal() (*user.Entity, error) {
	userID, err := user.ParseID(m.ID)
	if err != nil {
		return nil, fmt.Errorf("user.ParseID() id = %s\terror = %w", m.ID, err)
	}

	return &user.Entity{
		ID:   userID,
		Name: m.Name,
		Age:  m.Age,
	}, nil
}

func (m *model) unmarshal(entity *user.Entity) error {
	m.ID = entity.ID.String()
	m.Name = entity.Name
	m.Age = entity.Age

	return nil
}
