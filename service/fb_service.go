package service

import (
	"context"
	"encoding/json"
	"fmt"
	"simple-go-grpc/api"
	"simple-go-grpc/common/dto"
	"simple-go-grpc/common/fb_pb"
	"strings"
)

type fbRpcService struct{}

var (
	Fb          api.DemoRpc   = new(fbRpcService)
	FbRpcRouter fb_pb.RpcFunc = new(fbRpcService)
)

func (fbRpcService) Target() string { return strings.ToLower("DemoRpc") }

func (fbRpcService) RoutingTable() map[string]fb_pb.FbRpcFunc {
	dataMap := make(map[string]fb_pb.FbRpcFunc)
	dataMap["FindInfo"] = Fb.FindInfo
	// dataMap["FindInfoV2"] = Fb.FindInfoV2
	return dataMap
}

func (fbRpcService) FindInfo(ctx context.Context, req []byte) fb_pb.FbResponse {
	fmt.Println("authorization ", ctx.Value(fb_pb.AuthorizationKey))

	var request dto.FbRpcReq
	if err := json.Unmarshal(req, &request); err != nil {
		return fb_pb.Fail(err.Error())
	}

	return fb_pb.Suc(request)
}
