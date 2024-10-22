package cluster

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const metaHeader = "i-client"
const metaValue = "internal"

func ctxWithMetadata(ctx context.Context) context.Context {
	c := metadata.NewOutgoingContext(ctx, metadata.Pairs(metaHeader, metaValue))
	return c
}

func (c *Replicator) isInternalRequest(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	if value, ok := md[metaHeader]; ok {
		return value[0] == metaValue
	}

	return false
}
