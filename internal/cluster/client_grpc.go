package cluster

import (
	"context"
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type GrpcClient struct {
	conn   pb.LogClient
	queue  chan *pb.LogRequest
	logger *slog.Logger
	tracer trace.Tracer
}

func NewGrpcClient(ctx context.Context, logger *slog.Logger, tracer trace.Tracer, address string) (*GrpcClient, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewLogClient(conn)
	service := &GrpcClient{
		conn:   client,
		queue:  make(chan *pb.LogRequest, 1000),
		logger: logger.With(slog.String("server", address)),
		tracer: tracer,
	}
	service.worker(ctx)
	return service, nil
}

func (s *GrpcClient) SaveLog(item *domain.ItemStore) {
	s.queue <- &pb.LogRequest{
		Action:   pb.LogRequest_SAVE,
		Key:      item.Key,
		Value:    item.Value,
		Deadline: item.Deadline,
	}
}

func (s *GrpcClient) DeleteLog(item *domain.ItemStore) {
	s.queue <- &pb.LogRequest{
		Action: pb.LogRequest_DELETE,
		Key:    item.Key,
	}
}

func (s *GrpcClient) worker(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-s.queue:
				ctxSpan, span := s.tracer.Start(ctx, "worker.SaveItem", trace.WithAttributes(attribute.String("key", item.Key)))
				_, err := s.conn.Save(ctxWithMetadata(ctxSpan), item)
				if err != nil {
					span.RecordError(err)
					s.logger.ErrorContext(ctxSpan, "failed send item", slog.String("error", err.Error()), slog.String("key", item.Key))
				} else {
					s.logger.InfoContext(ctxSpan, "success send item", slog.String("key", item.Key))
				}
				span.End()
			}
		}
	}()
}
