package clock

import (
	"context"
	pb "github.com/updevru/go-micro-kit-example/gen/clock"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

type Handler struct {
	pb.UnimplementedClockServer
	log    *slog.Logger
	tracer trace.Tracer
}

func NewHandler(log *slog.Logger, tracer trace.Tracer) *Handler {
	return &Handler{log: log, tracer: tracer}
}

func (s *Handler) Now(ctx context.Context, in *pb.ClockRequest) (*pb.ClockResponse, error) {
	_, span := s.tracer.Start(ctx, "handler.Now")
	defer span.End()

	s.log.InfoContext(ctx, "Call Now()", slog.String("timezone", in.GetTimezone()))

	timezone, err := time.LoadLocation(in.GetTimezone())
	if err != nil {
		span.RecordError(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	now := time.Now().In(timezone)

	return &pb.ClockResponse{Time: now.String()}, nil
}
