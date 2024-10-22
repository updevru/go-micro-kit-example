package repository

import (
	"errors"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"reflect"
	"testing"
	"time"
)

func TestMemoryRepository_Save(t *testing.T) {
	type args struct {
		item domain.ItemStore
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Save",
			args: args{
				item: domain.ItemStore{Key: "key", Value: "value"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMemoryRepository()
			if err := r.Save(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryRepository_Read(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.ItemStore
		wantErr bool
	}{
		{
			name: "Read key",
			args: args{
				key: "key",
			},
			want:    &domain.ItemStore{Key: "key", Value: "value"},
			wantErr: false,
		},
		{
			name: "Read key3",
			args: args{
				key: "key3",
			},
			want:    &domain.ItemStore{Key: "key3", Value: "value3"},
			wantErr: false,
		},
		{
			name: "Key not found",
			args: args{
				key: "key4",
			},
			want:    nil,
			wantErr: true,
		},
	}

	items := []domain.ItemStore{
		{Key: "key", Value: "value"},
		{Key: "key2", Value: "value2"},
		{Key: "key3", Value: "value3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMemoryRepository()
			for _, item := range items {
				_ = r.Save(item)
			}

			got, err := r.Read(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !errors.Is(err, StoreErrorNotFound) {
				t.Errorf("Read() error = %v, must be %v", err, StoreErrorNotFound)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryRepository_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Delete key",
			args: args{
				key: "key",
			},
			wantErr: false,
		},
		{
			name: "Delete not found key4",
			args: args{
				key: "key4",
			},
			wantErr: false,
		},
	}

	items := []domain.ItemStore{
		{Key: "key", Value: "value"},
		{Key: "key2", Value: "value2"},
		{Key: "key3", Value: "value3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMemoryRepository()
			for _, item := range items {
				_ = r.Save(item)
			}

			if err := r.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			if item, err := r.Read(tt.args.key); item != nil && !errors.Is(err, StoreErrorNotFound) {
				t.Errorf("Delete() item = %v, must be nil", item)
			}
		})
	}
}

func TestMemoryRepository_DeleteDead(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Delete 1 item",
			args: args{
				date: time.Now(),
			},
			want: 1,
		},
		{
			name: "Delete 3 item",
			args: args{
				date: time.Now().Add(2 * time.Hour),
			},
			want: 3,
		},
		{
			name: "Delete 0 item",
			args: args{
				date: time.Now().Add(-2 * time.Hour),
			},
			want: 0,
		},
	}

	items := []domain.ItemStore{
		{Key: "key", Value: "value", Deadline: time.Now().Add(1 * time.Hour).Unix()},
		{Key: "key2", Value: "value2", Deadline: time.Now().Add(1 * time.Minute).Unix()},
		{Key: "key3", Value: "value3", Deadline: time.Now().Add(-1 * time.Hour).Unix()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMemoryRepository()
			for _, item := range items {
				_ = r.Save(item)
			}

			got, err := r.DeleteDead(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteDead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteDead() got = %v, want %v", got, tt.want)
			}

			count := 0
			for _, item := range items {
				if _, err := r.Read(item.Key); err == nil {
					count++
				}
			}

			if len(items)-count != tt.want {
				t.Errorf("DeleteDead() count = %v, want %v", count, tt.want)
			}
		})
	}
}
