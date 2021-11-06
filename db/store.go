package db

import (
	"sync"
)

type Store struct {
	sync.RWMutex
	data map[string]string
}

func New() *Store {
	return &Store{data: make(map[string]string)}
}

func (store *Store) Get(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (store *Store) Set(key string, value string) error {
	store.Lock()
	defer store.Unlock()
	store.data[key] = value

	return nil
}
