package main

import (
	"context"
	"errors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/updevru/go-micro-kit-example/internal/config"
	"github.com/updevru/go-micro-kit-example/internal/cron"
	"github.com/updevru/go-micro-kit-example/internal/grpc"
	"github.com/updevru/go-micro-kit-example/internal/handler/clock"
	"github.com/updevru/go-micro-kit-example/internal/handler/store"
	"github.com/updevru/go-micro-kit-example/internal/repository"
	"github.com/updevru/go-micro-kit-example/internal/rest"
	configPkg "github.com/updevru/go-micro-kit/config"
	"github.com/updevru/go-micro-kit/discovery"
	"github.com/updevru/go-micro-kit/server"
	"github.com/updevru/go-micro-kit/telemetry"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var cfg config.Config
	if err := configPkg.CreateConfig(ctx, &cfg); err != nil {
		panic(err)
	}

	// Set up OpenTelemetry.
	otelShutdown, err := telemetry.SetupOTel(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	logger := telemetry.CreateLogger()
	tracer := telemetry.CreateTracer()
	meter := telemetry.CreateMeter()

	repositoryStore := repository.NewStoreRepository()
	storeServer := store.NewHandler(logger, tracer, repositoryStore)
	clockServer := clock.NewHandler(logger, tracer)

	app := server.NewServer(ctx, logger, tracer, meter)
	app.Grpc(&cfg.Grpc, grpc.NewGRPCServer(clockServer, storeServer))
	app.Http(&cfg.Http, &cfg.Grpc, rest.NewRestServer()...)
	app.Cron(cron.NewCron(cron.NewCleaner(logger, tracer)))

	consul, err := discovery.NewConsul(&cfg.App, &cfg.Http, &cfg.Grpc)
	app.AddDiscovery(consul)

	if err := app.Run(); err != nil {
		logger.ErrorContext(ctx, "Failed to run server: %v", err)
		panic(err)
	}
}
