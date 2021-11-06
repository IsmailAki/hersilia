package db

import (
	"errors"
	"sync"
)

type StoreData struct {
	data map[string]string
	sync.RWMutex
}

func New() *StoreData {
	return &StoreData{data: make(map[string]string)}
}

func (store *StoreData) Get(key string) (string, error) {
	store.RLock()
	defer store.RUnlock()
	value, ok := store.data[key]
	if !ok {
		return "", errors.New("Key not found")
	}
	return value, nil
}

func (store *StoreData) Set(key string, value string) error {
	store.Lock()
	defer store.Unlock()
	store.data[key] = value

	return nil
}
