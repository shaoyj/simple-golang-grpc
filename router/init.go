package router

import (
	"context"
	"simple-go-grpc/common/pb"
	"simple-go-grpc/service"
)

// target->method,value
var rpcRouterHelper = initRpc()

func initRpc() *pb.RpcRouterHelper {
	routerHelper := pb.NewRpcRouterHelper(16)
	//FbRpcRouter
	routerHelper.AddRouter(service.FbRpcRouter)
	//todo

	return routerHelper
}

func BaseEndPointFunc(ctx context.Context, request interface{}) (response interface{}, err error) {
	return rpcRouterHelper.RpcHandler(ctx, request)
}
