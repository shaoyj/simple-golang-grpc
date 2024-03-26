package service

import (
	"context"
	"encoding/json"
	"fmt"
	"simple-go-grpc/api"
	"simple-go-grpc/common/dto"
	"simple-go-grpc/common/helper"
	"simple-go-grpc/common/tool"
	"strings"
)

type fbRpcService struct{}

var (
	Fb          api.DemoRpc    = new(fbRpcService)
	FbRpcRouter helper.RpcFunc = new(fbRpcService)
)

func (fbRpcService) Target() string { return strings.ToLower("DemoRpc") }

func (fbRpcService) RoutingTable() map[string]helper.FbRpcFunc {
	dataMap := make(map[string]helper.FbRpcFunc)
	dataMap["FindInfo"] = Fb.FindInfo

	return dataMap
}

func (fbRpcService) FindInfo(ctx context.Context, req []byte) tool.FbResponse {
	fmt.Println("authorization ", ctx.Value(tool.AuthorizationKey))

	var request dto.FbRpcReq
	if err := json.Unmarshal(req, &request); err != nil {
		return tool.Fail(err.Error())
	}

	return tool.Suc(request)
}
