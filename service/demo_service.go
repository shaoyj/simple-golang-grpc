package service

import (
	"context"
	"encoding/json"
	"fmt"
	"simple-go-grpc/api"
	"simple-go-grpc/common/dto"
	"simple-go-grpc/common/pb"
	"strings"
)

type fbRpcService struct{}

var (
	Fb          api.DemoRpc = new(fbRpcService)
	FbRpcRouter pb.RpcFunc  = new(fbRpcService)
)

func (fbRpcService) Target() string { return strings.ToLower("DemoRpc") }

func (fbRpcService) RoutingTable() map[string]pb.FbRpcFunc {
	dataMap := make(map[string]pb.FbRpcFunc)
	dataMap["FindInfo"] = Fb.FindInfo

	return dataMap
}

func (fbRpcService) FindInfo(ctx context.Context, req []byte) pb.FbResponse {
	fmt.Println("authorization ", ctx.Value(pb.AuthorizationKey))

	var request dto.FbRpcReq
	if err := json.Unmarshal(req, &request); err != nil {
		return pb.Fail(err.Error())
	}

	return pb.Suc(request)
}
