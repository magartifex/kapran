package users

import (
	"context"
	"fmt"

	"github.com/magartifex/kapran/entities/user"
)

func (s *Store) Set(_ context.Context, entity *user.Entity) error {
	err := entity.GenerateID()
	if err != nil {
		return fmt.Errorf("entity.GenerateID() error = %w", err)
	}

	m := &model{}

	err = m.unmarshal(entity)
	if err != nil {
		return fmt.Errorf("m.unmarshal() entity = %#v\terror = %w", entity, err)
	}

	s.database.Store(m.ID, m)

	return nil
}
