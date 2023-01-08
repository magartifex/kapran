package users

import (
	"context"
	"fmt"

	"github.com/magartifex/kapran/constants"
	"github.com/magartifex/kapran/entities/user"
)

func (s *Store) GetByID(_ context.Context, userID user.ID) (*user.Entity, error) {
	value, ok := s.database.Load(userID.String())
	if !ok {
		return nil, fmt.Errorf("s.database.Load() userID = %s\terror=%w", userID, constants.ErrNoDataRow)
	}

	m, ok := value.(*model)
	if !ok {
		return nil, fmt.Errorf("value() value = %#v\terror=%w", value, constants.ErrNoDataType)
	}

	entity, err := m.marshal()
	if err != nil {
		return nil, fmt.Errorf("m.marshal() model = %#v\terror=%w", m, err)
	}

	return entity, nil
}
