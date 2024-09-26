package rest

import (
	pbStore "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit/server"
)

func NewRestServer() []server.HttpHandler {
	return []server.HttpHandler{
		pbStore.RegisterStoreHandler,
	}
}
