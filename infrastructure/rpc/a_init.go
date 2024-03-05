package rpc

import (
	grpc "google.golang.org/grpc"
	"simple-go-grpc/common/fb_pb"
)

func initC(address string, port string) *fb_pb.BaseServiceClient {
	conn, e := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if e != nil {
		panic(e)
	}
	// defer conn.Close()

	// 新建一个客户端，方法为：NewXXXClinent(conn),XXX为你在proto定义的服务的名字
	c := fb_pb.NewBaseServiceClient(conn)

	return &c
}
