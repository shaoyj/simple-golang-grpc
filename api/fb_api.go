package api

import (
	"context"
	"simple-go-grpc/common/fb_pb"
)

type DemoRpc interface {
	FindInfo(ctx context.Context, req []byte) fb_pb.FbResponse
}
