package store

import (
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/cluster"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"github.com/updevru/go-micro-kit-example/internal/repository"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

type Handler struct {
	pb.UnimplementedStoreServer
	log         *slog.Logger
	tracer      trace.Tracer
	store       repository.StoreInterface
	replication *cluster.Replicator
}

func NewHandler(log *slog.Logger, tracer trace.Tracer, store repository.StoreInterface, replication *cluster.Replicator) *Handler {
	return &Handler{log: log, tracer: tracer, store: store, replication: replication}
}

func mapResponse(item *domain.ItemStore) *pb.StorageResponse {
	if item == nil {
		return nil
	}

	return &pb.StorageResponse{Key: item.Key, Value: item.Value}
}
