package users

import (
	"sync"
)

type Store struct {
	database sync.Map
}

func New() *Store {
	return &Store{
		database: sync.Map{},
	}
}
