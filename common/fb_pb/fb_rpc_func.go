package fb_pb

import (
	"context"
)

type FbRpcFunc func(headers context.Context, req []byte) FbResponse

type RpcFunc interface {
	Target() string
	RoutingTable() map[string]FbRpcFunc
}
