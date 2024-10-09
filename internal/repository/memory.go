package repository

import (
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"golang.org/x/sync/syncmap"
	"time"
)

type MemoryRepository struct {
	store syncmap.Map
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		store: syncmap.Map{},
	}
}

func (r *MemoryRepository) Save(item domain.ItemStore) error {
	r.store.Store(item.Key, item)
	return nil
}

func (r *MemoryRepository) Read(key string) (*domain.ItemStore, error) {
	if val, exist := r.store.Load(key); exist {
		result := val.(domain.ItemStore)
		return &result, nil
	}

	return nil, StoreErrorNotFound
}

func (r *MemoryRepository) Delete(key string) error {
	r.store.Delete(key)
	return nil
}

func (r *MemoryRepository) DeleteDead(date time.Time) (int, error) {
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
