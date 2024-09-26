package cluster

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const metaHeader = "i-client"
const metaValue = "internal"

func ctxWithMetadata(ctx context.Context) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(metaHeader, metaValue))
}

func isInternalRequest(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	if value, ok := md[metaHeader]; ok {
		return value[0] == metaValue
	}

	return false
}
