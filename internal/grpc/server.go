package grpc

import (
	pbStore "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit-example/internal/handler/log"
	"github.com/updevru/go-micro-kit-example/internal/handler/store"
	"github.com/updevru/go-micro-kit/server"
	"google.golang.org/grpc"
)

func NewGRPCServer(storeHandler *store.Handler, logHandler *log.StoreLog) server.GrpcHandler {
	return func(server *grpc.Server) {
		pbStore.RegisterStoreServer(server, storeHandler)
		pbStore.RegisterLogServer(server, logHandler)
	}
}
