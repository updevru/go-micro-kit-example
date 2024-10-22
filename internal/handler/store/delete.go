package store

import (
	"context"
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Handler) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	ctxSpan, span := s.tracer.Start(ctx, "handler.Store.Delete")
	defer span.End()

	item, err := s.store.Read(in.Key)
	if err != nil {
		span.RecordError(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	s.log.InfoContext(ctxSpan, "success delete item", slog.String("key", in.Key))
	s.replication.DeleteItem(ctxSpan, item)

	return &pb.DeleteResponse{}, err
}
