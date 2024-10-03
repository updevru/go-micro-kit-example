package cron

import (
	"context"
	"github.com/updevru/go-micro-kit-example/internal/repository"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"time"
)

type Cleaner struct {
	logger          *slog.Logger
	tracer          trace.Tracer
	storeRepository repository.StoreInterface
}

func NewCleaner(logger *slog.Logger, tracer trace.Tracer, storeRepository repository.StoreInterface) *Cleaner {
	return &Cleaner{
		logger:          logger,
		tracer:          tracer,
		storeRepository: storeRepository,
	}
}

func (c *Cleaner) Clean(ctx context.Context) error {
	num, err := c.storeRepository.DeleteDead(time.Now())
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to clean data", slog.String("error", err.Error()))
		return err
	}

	if num > 0 {
		c.logger.InfoContext(ctx, "Cleaned", slog.Int("num", num))
	}

	return nil
}
