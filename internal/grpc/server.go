package grpc

import (
	pbStore "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/handler/store"
	"github.com/updevru/go-micro-kit/server"
	"google.golang.org/grpc"
)

func NewGRPCServer(storeServer *store.Handler) server.GrpcHandler {
	return func(server *grpc.Server) {
		pbStore.RegisterStoreServer(server, storeServer)
	}
}
