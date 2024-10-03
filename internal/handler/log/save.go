package log

import (
	"context"
	pb "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/domain"
	"log/slog"
)

func (l *StoreLog) Save(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	_, span := l.tracer.Start(ctx, "handler.Log.Save")
	defer span.End()

	var err error
	switch req.GetAction() {
	case pb.LogRequest_SAVE:
		err = l.storeRepository.Save(domain.NewItemStore(req.GetKey(), req.GetValue(), req.GetDeadline()))
	case pb.LogRequest_DELETE:
		err = l.storeRepository.Delete(req.GetKey())
	}

	if err != nil {
		span.RecordError(err)
		l.log.ErrorContext(ctx, "failed save log", slog.String("key", req.GetKey()), slog.String("error", err.Error()))
	} else {
		l.log.InfoContext(ctx, "success save log", slog.String("key", req.GetKey()))
	}

	return &pb.LogResponse{}, err
}
