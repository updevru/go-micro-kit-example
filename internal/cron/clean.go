package cron

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

type Cleaner struct {
	logger *slog.Logger
	tracer trace.Tracer
}

func NewCleaner(logger *slog.Logger, tracer trace.Tracer) *Cleaner {
	return &Cleaner{
		logger: logger,
		tracer: tracer,
	}
}

func (c *Cleaner) Clean(ctx context.Context) error {

	_, span := c.tracer.Start(ctx, "start long operation")
	c.logger.InfoContext(ctx, "Cleaning data")
	span.End()

	return nil
}
