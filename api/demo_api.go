package api

import (
	"context"
	"simple-go-grpc/common/pb"
)

type DemoRpc interface {
	FindInfo(ctx context.Context, req []byte) pb.FbResponse
}
