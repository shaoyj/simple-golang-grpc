package router

import (
	"context"
	"simple-go-grpc/common/helper"
	"simple-go-grpc/service"
)

// target->method,value
var rpcRouterHelper = initRpc()

func initRpc() *helper.RpcRouterHelper {
	routerHelper := helper.NewRpcRouterHelper(16)
	//FbRpcRouter
	routerHelper.AddRouter(service.FbRpcRouter)
	//todo

	return routerHelper
}

func BaseEndPointFunc(ctx context.Context, request interface{}) (response interface{}, err error) {
	return rpcRouterHelper.RpcHandler(ctx, request)
}
