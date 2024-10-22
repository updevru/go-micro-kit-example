package cluster

import (
	"context"
	"github.com/updevru/go-micro-kit-example/internal/cluster/mocks"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestReplicator_DeleteItem(t *testing.T) {
	type args struct {
		ctx  context.Context
		item *domain.ItemStore
	}
	tests := []struct {
		name  string
		args  args
		sends int
	}{
		{
			name: "Replicate DeleteItem send",
			args: args{
				ctx:  context.Background(),
				item: &domain.ItemStore{},
			},
			sends: 1,
		},
		{
			name: "Replicate DeleteItem not send",
			args: args{
				ctx:  metadata.NewIncomingContext(context.Background(), metadata.Pairs(metaHeader, metaValue)),
				item: &domain.ItemStore{},
			},
			sends: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.NewClient(t)
			if tt.sends > 0 {
				client.On("DeleteLog", tt.args.item).Return()
			} else {
				client.On("DeleteLog", tt.args.item).Unset()
			}

			c := &Replicator{
				servers: []Client{client},
				tracer:  otel.Tracer("main"),
			}
			c.DeleteItem(tt.args.ctx, tt.args.item)

			client.AssertNumberOfCalls(t, "DeleteLog", tt.sends)
		})
	}
}

func TestReplicator_SaveItem(t *testing.T) {
	type args struct {
		ctx  context.Context
		item *domain.ItemStore
	}
	tests := []struct {
		name  string
		args  args
		sends int
	}{
		{
			name: "Replicate SaveItem send",
			args: args{
				ctx:  context.Background(),
				item: &domain.ItemStore{},
			},
			sends: 1,
		},
		{
			name: "Replicate SaveItem not send",
			args: args{
				ctx:  metadata.NewIncomingContext(context.Background(), metadata.Pairs(metaHeader, metaValue)),
				item: &domain.ItemStore{},
			},
			sends: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.NewClient(t)
			if tt.sends > 0 {
				client.On("SaveLog", tt.args.item).Return()
			} else {
				client.On("SaveLog", tt.args.item).Unset()
			}

			c := &Replicator{
				servers: []Client{client},
				tracer:  otel.Tracer("main"),
			}
			c.SaveItem(tt.args.ctx, tt.args.item)

			client.AssertNumberOfCalls(t, "SaveLog", tt.sends)
		})
	}
}
