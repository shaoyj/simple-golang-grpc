package helper

import (
	"context"
	"simple-go-grpc/common/tool"
)

type FbRpcFunc func(headers context.Context, req []byte) tool.FbResponse

type RpcFunc interface {
	Target() string
	RoutingTable() map[string]FbRpcFunc
}
