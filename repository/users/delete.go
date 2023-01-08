package users

import (
	"context"
	"fmt"

	"github.com/magartifex/kapran/constants"
	"github.com/magartifex/kapran/entities/user"
)

func (s *Store) Delete(_ context.Context, userID user.ID) error {
	_, ok := s.database.LoadAndDelete(userID.String())
	if !ok {
		return fmt.Errorf("s.database.LoadAndDelete() userID = %s\terror = %w", userID, constants.ErrNoDataRow)
	}

	return nil
}
