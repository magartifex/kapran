package users_test

import (
	"context"
	"testing"

	"github.com/magartifex/kapran/entities/user"
	"github.com/magartifex/kapran/repository/users"
)

func TestStore_GetByID(t *testing.T) {
	store := users.New()

	type args struct {
		store   repoUser
		user    *user.Entity
		gotUser *user.Entity
		ifEqual bool
		err     error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"happy path",
			args{
				store: store,
				user: &user.Entity{
					Name: "Alex",
					Age:  35,
				},
				gotUser: &user.Entity{
					Name: "Alex",
					Age:  35,
				},
				ifEqual: true,
			},
		},
		{
			"name not equal",
			args{
				store: store,
				user: &user.Entity{
					Name: "Alex",
					Age:  35,
				},
				gotUser: &user.Entity{
					Name: "Magomed",
					Age:  31,
				},
				ifEqual: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.store.Set(context.Background(), tt.args.user)
			if err != nil {
				t.Errorf("user = %#v\terror = %s", tt.args.user, err)
			}

			got, err := tt.args.store.GetByID(context.Background(), tt.args.user.ID)
			if err != nil {
				t.Errorf("userID = %s\terror = %s", tt.args.user.ID, err)
			}

			if (tt.args.gotUser.Name == got.Name) != tt.args.ifEqual {
				t.Errorf("name user name = %s\tgot name = %s\tifEqual = %v", tt.args.gotUser.Name, got.Name, tt.args.ifEqual)
			}

			if (tt.args.gotUser.Age == got.Age) != tt.args.ifEqual {
				t.Errorf("name user age = %d\tgot age = %d\tifEqual = %v", tt.args.gotUser.Age, got.Age, tt.args.ifEqual)
			}
		})
	}
}
