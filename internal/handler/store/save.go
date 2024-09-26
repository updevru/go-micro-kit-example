package store

import (
	"context"
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"log/slog"
)

func (s *Handler) Save(ctx context.Context, in *pb.SaveRequest) (*pb.StorageResponse, error) {
	ctxSpan, span := s.tracer.Start(ctx, "repository.Save")
	defer span.End()

	item := domain.ItemStore{Key: in.Key, Value: in.Value}
	err := s.store.Save(item)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	s.log.InfoContext(ctxSpan, "success save item", slog.String("key", in.Key))
	s.cluster.PublishItem(ctxSpan, &item)

	return mapResponse(&item), err
}
