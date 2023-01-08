package users

import (
	"context"
	"fmt"

	"github.com/magartifex/kapran/constants"
	"github.com/magartifex/kapran/entities/user"
)

func (s *Store) Get(ctx context.Context) ([]*user.Entity, error) {
	var entities []*user.Entity

	var err error

	getUserFromValue := func(value any) (*user.Entity, error) {
		m, ok := value.(*model)
		if !ok {
			return nil, fmt.Errorf("value() value = %#v\terror=%w", value, constants.ErrNoDataType)
		}

		entity, er := m.marshal()
		if er != nil {
			return nil, fmt.Errorf("m.marshal() model = %#v\terror=%w", m, er)
		}

		return entity, nil
	}

	s.database.Range(func(_, value any) bool {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return false

		default:
			entity, er := getUserFromValue(value)
			if err != nil {
				err = er
				return false
			}

			entities = append(entities, entity)
			return true
		}
	})

	return entities, err
}
