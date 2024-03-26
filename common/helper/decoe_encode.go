package helper

import (
	"context"
	"encoding/json"
	"simple-go-grpc/common/pb"
	"simple-go-grpc/common/tool"
)

func DecodeDefault(ctx context.Context, req interface{}) (request interface{}, err error) {
	request = req
	return
}
func EncodeDefault(ctx context.Context, rep interface{}) (response interface{}, err error) {
	var res pb.ByteResult
	firstRes, ok := rep.(*pb.ByteResult)
	if ok {
		return firstRes, nil
	}

	data := rep.(tool.Result)
	res.Code = data.Code
	res.Msg = data.Msg

	if data.Code == 0 && data.Data != nil {
		bytes, err := json.Marshal(data.Data)
		if err != nil {
			res.Code = 500
			res.Msg = err.Error()
			response = res
		} else {
			res.Data = bytes
		}
	}

	response = &res
	return
}
