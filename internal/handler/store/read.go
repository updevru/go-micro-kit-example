package store

import (
	"context"
	"errors"
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Handler) Read(ctx context.Context, in *pb.ReadRequest) (*pb.StorageResponse, error) {
	_, span := s.tracer.Start(ctx, "handler.Store.Read")
	defer span.End()

	item, err := s.store.Read(in.Key)
	if err != nil {
		span.RecordError(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if item.IsDead(time.Now()) {
		span.RecordError(errors.New("key expired"))
		return nil, status.Error(codes.NotFound, "key not found")
	}

	return mapResponse(item), err
}
