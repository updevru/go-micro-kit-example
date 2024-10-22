package cluster

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

func NewGrpcReplicator(ctx context.Context, logger *slog.Logger, tracer trace.Tracer, addresses []string) (*Replicator, error) {
	cluster := &Replicator{
		servers: make([]Client, 0, len(addresses)),
		tracer:  tracer,
	}

	for _, address := range addresses {
		client, err := NewGrpcClient(ctx, logger, tracer, address)
		if err != nil {
			return nil, err
		}
		cluster.servers = append(cluster.servers, client)
		logger.InfoContext(ctx, "connected server", slog.String("server", address))
	}

	return cluster, nil
}
