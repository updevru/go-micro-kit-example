package log

import (
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/repository"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

type StoreLog struct {
	pb.UnimplementedLogServer
	tracer          trace.Tracer
	log             *slog.Logger
	storeRepository repository.StoreInterface
}

func NewHandler(log *slog.Logger, tracer trace.Tracer, storeRepository repository.StoreInterface) *StoreLog {
	return &StoreLog{tracer: tracer, log: log, storeRepository: storeRepository}
}
