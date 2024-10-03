package domain

import "time"

type ItemStore struct {
	Key      string
	Value    string
	Deadline int64
}

func NewItemStore(key, value string, ttl int64) ItemStore {
	if ttl > 0 {
		ttl = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}

	return ItemStore{
		Key:      key,
		Value:    value,
		Deadline: ttl,
	}
}

func (i *ItemStore) IsDead(date time.Time) bool {
	if i.Deadline == 0 {
		return false
	}

	return i.Deadline < date.Unix()
}
