package repository

import (
	"errors"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"time"
)

var StoreErrorNotFound = errors.New("key not found")

type StoreInterface interface {
	Save(item domain.ItemStore) error
	Read(key string) (*domain.ItemStore, error)
	Delete(key string) error
	DeleteDead(date time.Time) (int, error)
}
