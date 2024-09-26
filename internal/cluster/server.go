package cluster

import (
	"context"
	proto_demo_store_v1 "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type ServiceClient struct {
	conn   proto_demo_store_v1.StoreClient
	queue  chan *domain.ItemStore
	logger *slog.Logger
}

func NewServiceClient(logger *slog.Logger, address string) (*ServiceClient, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}

	client := proto_demo_store_v1.NewStoreClient(conn)
	return &ServiceClient{
		conn:   client,
		queue:  make(chan *domain.ItemStore, 1000),
		logger: logger.With(slog.String("server", address)),
	}, nil
}

func (s *ServiceClient) PublishItem(item *domain.ItemStore) {
	s.queue <- item
}

func (s *ServiceClient) Worker(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-s.queue:
				_, err := s.conn.Save(ctxWithMetadata(ctx), &proto_demo_store_v1.SaveRequest{
					Key:   item.Key,
					Value: item.Value,
				})
				if err != nil {
					s.logger.ErrorContext(ctx, "failed send item", slog.String("error", err.Error()), slog.String("key", item.Key))
				} else {
					s.logger.InfoContext(ctx, "success send item", slog.String("key", item.Key))
				}
			}
		}
	}()
}
