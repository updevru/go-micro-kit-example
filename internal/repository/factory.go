package repository

import (
	"errors"
	"github.com/updevru/go-micro-kit-example/internal/config"
)

const (
	StoreNameMemory = "memory"
	StoreNameBolt   = "bolt"
)

func Factory(config config.Storage) (StoreInterface, error) {
	switch config.Name {
	case StoreNameMemory:
		return NewMemoryRepository(), nil
	case StoreNameBolt:
		if config.Bolt.File == "" {
			return nil, errors.New("bolt file is empty")
		}
		return NewBoltRepository(&config.Bolt)
	}

	return nil, errors.New("unknown store")
}
