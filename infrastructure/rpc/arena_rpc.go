package rpc

import (
	"context"
	"encoding/json"
	"simple-go-grpc/common/config"
	"simple-go-grpc/common/dto"
	"simple-go-grpc/common/fb_pb"
	"strings"

	"google.golang.org/grpc/metadata"
)

var (
	arenaClient *fb_pb.BaseServiceClient = initC(config.Instance.Rpc.ArenaAddress, "8081")
	target                               = "arenaActivityRpc"
)

func FindInfo(ctx context.Context, req dto.FbRpcReq) *dto.FbRpcReq {
	// 调用远程，并得到返回
	// ctx := context.Background()
	header := metadata.New(map[string]string{"authorization": "ae4c5fbe-8571-45d5-bb51-7fe081ef", "space": "", "org": "", "limit": "", "offset": ""})
	ctx = metadata.NewOutgoingContext(ctx, header)

	bt1, _ := json.Marshal(req)
	r, err := (*arenaClient).Execute(ctx, &fb_pb.ComReq{
		Target: strings.ToLower(target),
		Method: "FindInfo",
		Body:   bt1,
	})
	if err != nil {
		//fbl.Log().Sugar().Errorf("could not greet: %v", err)
		return nil
	}

	var result dto.FbRpcReq
	if resErr := json.Unmarshal(r.Data, &result); resErr != nil {
		//fbl.Log().Sugar().Errorf("could not greet2: %v", resErr)
		return nil
	}

	return &result
}
