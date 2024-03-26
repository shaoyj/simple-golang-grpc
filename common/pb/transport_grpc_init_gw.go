package pb

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// init_gw 注册网关服务
func initGw(address string, dialOpt []grpc.DialOption) *runtime.ServeMux {
	ctx := context.Background()
	//为了与前端交互，使用了grpc-gateway，在json 序列化时保证 code=0 正常显示
	gateWayMux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard,
		&FbJSONPb{OrigName: true, EmitDefaults: true}))

	// base
	if err := RegisterBaseServiceHandlerFromEndpoint(ctx, gateWayMux, address, dialOpt); err != nil {
		panic(err)
	}

	return gateWayMux
}
