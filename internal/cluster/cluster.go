package cluster

import (
	"context"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"log/slog"
)

type Cluster struct {
	servers []*ServiceClient
}

func New(ctx context.Context, logger *slog.Logger, addresses []string) (*Cluster, error) {
	cluster := &Cluster{
		servers: make([]*ServiceClient, 0, len(addresses)),
	}

	for _, address := range addresses {
		client, err := NewServiceClient(logger, address)
		if err != nil {
			return nil, err
		}

		client.Worker(ctx)
		cluster.servers = append(cluster.servers, client)
		logger.InfoContext(ctx, "connected server", slog.String("server", address))
	}

	return cluster, nil
}

func (c Cluster) PublishItem(ctx context.Context, item *domain.ItemStore) {
	if isInternalRequest(ctx) {
		return
	}

	for _, server := range c.servers {
		server.PublishItem(item)
	}
}
