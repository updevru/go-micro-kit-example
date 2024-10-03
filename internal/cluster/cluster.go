package cluster

import (
	"context"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

type Cluster struct {
	servers []*ServiceClient
	tracer  trace.Tracer
}

func New(ctx context.Context, logger *slog.Logger, tracer trace.Tracer, addresses []string) (*Cluster, error) {
	cluster := &Cluster{
		servers: make([]*ServiceClient, 0, len(addresses)),
		tracer:  tracer,
	}

	for _, address := range addresses {
		client, err := NewServiceClient(logger, tracer, address)
		if err != nil {
			return nil, err
		}

		client.Worker(ctx)
		cluster.servers = append(cluster.servers, client)
		logger.InfoContext(ctx, "connected server", slog.String("server", address))
	}

	return cluster, nil
}

func (c *Cluster) SaveItem(ctx context.Context, item *domain.ItemStore) {
	if isInternalRequest(ctx) {
		return
	}

	for _, server := range c.servers {
		server.SaveLog(item)
	}
}

func (c *Cluster) DeleteItem(ctx context.Context, item *domain.ItemStore) {
	if isInternalRequest(ctx) {
		return
	}

	for _, server := range c.servers {
		server.SaveLog(item)
	}
}
