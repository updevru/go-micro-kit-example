package repository

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/updevru/go-micro-kit-example/internal/config"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"go.etcd.io/bbolt"
	"time"
)

type BoltRepository struct {
	store      *bbolt.DB
	bucketName []byte
}

func (b *BoltRepository) Save(item domain.ItemStore) error {
	return b.store.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", b.bucketName)
		}

		var err error
		if err = bucket.Put(b.keyValue(item.Key), []byte(item.Value)); err != nil {
			return err
		}

		if item.Deadline > 0 {
			buf := make([]byte, 8)
			binary.BigEndian.PutUint64(buf, uint64(item.Deadline))
			err = bucket.Put(b.keyDeadline(item.Key), buf)
		}

		return err
	})
}

func (b *BoltRepository) Read(key string) (*domain.ItemStore, error) {
	var result domain.ItemStore
	err := b.store.View(func(tx *bbolt.Tx) (err error) {
		bucket := tx.Bucket(b.bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", b.bucketName)
		}
		value := bucket.Get(b.keyValue(key))
		if len(value) == 0 {
			return StoreErrorNotFound
		}

		deadline := bucket.Get(b.keyDeadline(key))
		if len(deadline) > 0 {
			result.Deadline = int64(binary.BigEndian.Uint64(deadline))
		}

		result.Key = key
		result.Value = string(value)

		return nil
	})

	return &result, err
}

func (b *BoltRepository) Delete(key string) error {
	return b.store.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", b.bucketName)
		}
		err := bucket.Delete(b.keyValue(key))
		_ = bucket.Delete(b.keyDeadline(key))

		return err
	})
}

func (b *BoltRepository) keyValue(key string) []byte {
	return []byte("_v." + key)
}

func (b *BoltRepository) keyDeadline(key string) []byte {
	return []byte("_d." + key)
}

func (b *BoltRepository) DeleteDead(date time.Time) (int, error) {
	var deleted int
	err := b.store.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", b.bucketName)
		}

		c := bucket.Cursor()
		prefix := b.keyDeadline("")
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			deadline := int64(binary.BigEndian.Uint64(v))

			if deadline < date.Unix() {
				_ = bucket.Delete(b.keyValue(string(k)))
				_ = bucket.Delete(b.keyDeadline(string(k)))
				deleted++
			}
		}

		return nil
	})

	return deleted, err
}

func NewBoltRepository(config *config.StorageBolt) (*BoltRepository, error) {
	db, err := bbolt.Open(config.File, 0600, nil)
	if err != nil {
		return nil, err
	}

	bucketName := []byte("store")
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &BoltRepository{store: db, bucketName: bucketName}, nil
}
