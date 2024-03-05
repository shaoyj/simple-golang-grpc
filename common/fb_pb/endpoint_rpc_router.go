package fb_pb

import (
	"context"
	"encoding/json"
	"errors"

	"strings"
)

// RpcRouterHelper
type RpcRouterHelper struct {
	allRpcRouterMap map[string]map[string]FbRpcFunc
}

func NewRpcRouterHelper(size int) *RpcRouterHelper {
	return &RpcRouterHelper{
		allRpcRouterMap: make(map[string]map[string]FbRpcFunc, size),
	}
}

func (rpcRh *RpcRouterHelper) AddRouter(rpcFunc RpcFunc) {
	rpcRh.allRpcRouterMap[rpcFunc.Target()] = rpcFunc.RoutingTable()
}

func (rpcRh *RpcRouterHelper) FindFbRpcFunc(target, method string) (FbRpcFunc, error) {
	reqTarget := strings.ToLower(target)
	targetMap, ok := rpcRh.allRpcRouterMap[reqTarget]
	if ok {
		funcRpc, ok2 := targetMap[method]
		if ok2 {
			return funcRpc, nil
		}
	}

	return nil, errors.New("routing path that does not exist")
}

func (rpcRh *RpcRouterHelper) ParseByteResult(res *FbResponse) *ByteResult {
	var finalRes ByteResult
	//code !=0
	if res.Code != 0 {
		finalRes.Code = res.Code
		finalRes.Msg = res.Msg
		return &finalRes
	}

	//code==0
	btys, err := json.Marshal(res.Data)
	if err != nil {
		finalRes.Code = -1
		finalRes.Msg = err.Error()
		return &finalRes
	}

	//other
	finalRes.Code = res.Code
	finalRes.Msg = res.Msg
	finalRes.Data = btys
	return &finalRes
}

// RpcFuncHandler 解析 rpc函数返回
func (rpcRh *RpcRouterHelper) RpcFuncHandler(ctx context.Context, req *ComReq) *ByteResult {
	var finalRes ByteResult
	//parse funcRpc
	funcRpc, rpcErr := rpcRh.FindFbRpcFunc(strings.ToLower(req.Target), req.Method)
	if rpcErr != nil {
		finalRes.Code = -1
		finalRes.Msg = rpcErr.Error()
		return &finalRes
	}

	//execute
	res := funcRpc(ctx, req.Body)

	//result
	return rpcRh.ParseByteResult(&res)
}

func (rpcRh *RpcRouterHelper) RpcHandler(ctx context.Context, request interface{}) (response interface{}, err error) {
	var finalRes ByteResult

	//参数类型校验
	req, ok := request.(*ComReq)
	if !ok {
		finalRes.Code = -1
		finalRes.Msg = "Unsupported parameter type"
		response = &finalRes
		return response, nil
	}

	return rpcRh.RpcFuncHandler(HandlerParams(ctx, req)), nil
}
