package cluster

import (
	"context"
	"github.com/updevru/go-micro-kit-example/internal/domain"
)

type Replicate interface {
	SaveItem(context.Context, *domain.ItemStore)
	DeleteItem(context.Context, *domain.ItemStore)
}
