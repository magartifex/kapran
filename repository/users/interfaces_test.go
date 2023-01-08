package users_test

import (
	"context"

	"github.com/magartifex/kapran/entities/user"
)

type repoUser interface {
	Delete(context.Context, user.ID) error
	Get(context.Context) ([]*user.Entity, error)
	GetByID(context.Context, user.ID) (*user.Entity, error)
	Set(context.Context, *user.Entity) error
}
