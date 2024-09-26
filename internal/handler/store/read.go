package store

import (
	"context"
	pb "github.com/updevru/go-micro-kit-example/gen/store"
)

func (s *Handler) Read(ctx context.Context, in *pb.ReadRequest) (*pb.StorageResponse, error) {
	_, span := s.tracer.Start(ctx, "repository.Read")
	defer span.End()

	value, err := s.store.Read(in.Key)
	if err != nil {
		span.RecordError(err)
	}

	return mapResponse(value), err
}
