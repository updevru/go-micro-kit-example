package repository

import "errors"

var StoreErrorNotFound = errors.New("key not found")

type StoreInterface interface {
	Save(key string, value string) error
	Read(key string) (string, error)
	List() ([]ItemStore, error)
}

type ItemStore struct {
	Key   string
	Value string
}

type StoreRepository struct {
	store map[string]string
}

func NewStoreRepository() *StoreRepository {
	return &StoreRepository{
		store: make(map[string]string),
	}
}

func (r *StoreRepository) Save(key string, value string) error {
	r.store[key] = value
	return nil
}

func (r *StoreRepository) Read(key string) (string, error) {
	if val, exist := r.store[key]; exist {
		return val, nil
	}

	return "", StoreErrorNotFound
}

func (r *StoreRepository) List() ([]ItemStore, error) {
	var result []ItemStore
	for key, value := range r.store {
		result = append(result, ItemStore{Key: key, Value: value})
	}
	return result, nil
}

func (r *StoreRepository) Watch(con chan ItemStore) {

}
