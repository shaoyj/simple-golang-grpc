package api

import (
	"context"
	"simple-go-grpc/common/tool"
)

type DemoRpc interface {
	FindInfo(ctx context.Context, req []byte) tool.FbResponse
}
