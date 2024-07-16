package store

import (
	"context"
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/repository"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"time"
)

type Handler struct {
	pb.UnimplementedStoreServer
	log    *slog.Logger
	tracer trace.Tracer
	store  repository.StoreInterface
}

func NewHandler(log *slog.Logger, tracer trace.Tracer, store repository.StoreInterface) *Handler {
	return &Handler{log: log, tracer: tracer, store: store}
}

func (s *Handler) Save(ctx context.Context, in *pb.SaveRequest) (*pb.SaveResponse, error) {
	_, span := s.tracer.Start(ctx, "repository.Save")
	defer span.End()

	err := s.store.Save(in.Key, in.Value)
	if err != nil {
		span.RecordError(err)
	}

	return &pb.SaveResponse{Status: err == nil}, err
}

func (s *Handler) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	_, span := s.tracer.Start(ctx, "repository.Read")
	defer span.End()

	value, err := s.store.Read(in.Key)
	if err != nil {
		span.RecordError(err)
	}

	return &pb.ReadResponse{Value: value}, err
}

func (s *Handler) List(in *pb.ListRequest, stream pb.Store_ListServer) error {

	ticker := time.NewTicker(time.Second)
	go func() {
		s.log.Info("Ticker started")
		<-stream.Context().Done()
		ticker.Stop()
		s.log.Info("Ticker stop")
	}()

	for range ticker.C {
		items, err := s.store.List()
		if err != nil {
			return err
		}
		for _, item := range items {
			err := stream.Send(&pb.ListResponse{Key: item.Key, Value: item.Value})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
