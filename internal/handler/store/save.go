package store

import (
	"context"
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Handler) Save(ctx context.Context, in *pb.SaveRequest) (*pb.StorageResponse, error) {
	ctxSpan, span := s.tracer.Start(ctx, "handler.Store.Save")
	defer span.End()

	if in.GetTtl() < 0 {
		return nil, status.Error(codes.InvalidArgument, "ttl must be greater than zero")
	}

	item := domain.NewItemStore(in.Key, in.Value, int64(in.Ttl))
	err := s.store.Save(item)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	s.log.InfoContext(ctxSpan, "success save item", slog.String("key", in.Key))
	s.replication.SaveItem(ctxSpan, &item)

	return mapResponse(&item), err
}
