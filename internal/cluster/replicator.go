package cluster

import (
	"context"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"go.opentelemetry.io/otel/trace"
)

type Replicator struct {
	servers []Client
	tracer  trace.Tracer
}

func (c *Replicator) SaveItem(ctx context.Context, item *domain.ItemStore) {
	if c.isInternalRequest(ctx) {
		return
	}

	for _, server := range c.servers {
		server.SaveLog(item)
	}
}

func (c *Replicator) DeleteItem(ctx context.Context, item *domain.ItemStore) {
	if c.isInternalRequest(ctx) {
		return
	}

	for _, server := range c.servers {
		server.DeleteLog(item)
	}
}
