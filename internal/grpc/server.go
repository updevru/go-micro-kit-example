package grpc

import (
	pbClock "github.com/updevru/go-micro-kit-example/gen/clock"
	pbStore "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/handler/clock"
	"github.com/updevru/go-micro-kit-example/internal/handler/store"
	"github.com/updevru/go-micro-kit/server"
	"google.golang.org/grpc"
)

func NewGRPCServer(clockServer *clock.Handler, storeServer *store.Handler) server.GrpcHandler {
	return func(server *grpc.Server) {
		pbClock.RegisterClockServer(server, clockServer)
		pbStore.RegisterStoreServer(server, storeServer)
	}
}
