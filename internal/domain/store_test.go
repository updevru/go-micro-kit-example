package domain

import (
	"testing"
	"time"
)

func TestItemStore_IsDead(t *testing.T) {
	type fields struct {
		Key      string
		Value    string
		Deadline int64
	}
	type args struct {
		date time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "not dead",
			fields: fields{
				Key:   "key",
				Value: "value",
			},
			args: args{
				date: time.Now(),
			},
			want: false,
		},
		{
			name: "dead",
			fields: fields{
				Key:      "key",
				Value:    "value",
				Deadline: time.Now().Add(-1 * time.Hour).Unix(),
			},
			args: args{
				date: time.Now(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ItemStore{
				Key:      tt.fields.Key,
				Value:    tt.fields.Value,
				Deadline: tt.fields.Deadline,
			}
			if got := i.IsDead(tt.args.date); got != tt.want {
				t.Errorf("IsDead() = %v, want %v", got, tt.want)
			}
		})
	}
}
