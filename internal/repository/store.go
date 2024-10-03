package repository

import (
	"errors"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"golang.org/x/sync/syncmap"
	"time"
)

var StoreErrorNotFound = errors.New("key not found")

type StoreInterface interface {
	Save(item domain.ItemStore) error
	Read(key string) (*domain.ItemStore, error)
	Delete(key string) error
	List() (chan *domain.ItemStore, error)
	DeleteDead(date time.Time) (int, error)
}

type StoreRepository struct {
	store syncmap.Map
}

func NewStoreRepository() *StoreRepository {
	return &StoreRepository{
		store: syncmap.Map{},
	}
}

func (r *StoreRepository) Save(item domain.ItemStore) error {
	r.store.Store(item.Key, item)
	return nil
}

func (r *StoreRepository) Read(key string) (*domain.ItemStore, error) {
	if val, exist := r.store.Load(key); exist {
		result := val.(domain.ItemStore)
		return &result, nil
	}

	return nil, StoreErrorNotFound
}

func (r *StoreRepository) Delete(key string) error {
	r.store.Delete(key)
	return nil
}

func (r *StoreRepository) DeleteDead(date time.Time) (int, error) {
	var deleted int
	r.store.Range(func(key, value interface{}) bool {
		item := value.(domain.ItemStore)
		if item.IsDead(date) {
			r.store.Delete(key)
			deleted++
		}
		return true
	})
	return deleted, nil
}

func (r *StoreRepository) List() (chan *domain.ItemStore, error) {
	var result = make(chan *domain.ItemStore, 1000)

	go func() {
		r.store.Range(func(key, value interface{}) bool {
			item := value.(domain.ItemStore)
			result <- &item
			return true
		})
		close(result)
	}()

	return result, nil
}
