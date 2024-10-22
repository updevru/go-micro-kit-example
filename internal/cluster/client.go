package cluster

import (
	"github.com/updevru/go-micro-kit-example/internal/domain"
)

//go:generate mockery --name=Client --filename=client_store.go
type Client interface {
	SaveLog(item *domain.ItemStore)
	DeleteLog(item *domain.ItemStore)
}
