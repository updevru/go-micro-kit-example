package rest

import (
	pbClock "github.com/updevru/go-micro-kit-example/gen/clock"
	pbStore "github.com/updevru/go-micro-kit-example/gen/store"
	"github.com/updevru/go-micro-kit/server"
)

func NewRestServer() []server.HttpHandler {
	return []server.HttpHandler{
		pbClock.RegisterClockHandler,
		pbStore.RegisterStoreHandler,
	}
}
