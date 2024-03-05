package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"simple-go-grpc/common/fb_pb"
	"syscall"
	"time"

	"golang.org/x/net/netutil"

	"simple-go-grpc/router"
)

func main() {
	Port := "8082"
	address := ":" + Port
	conn, err := net.Listen("tcp", address)
	if err != nil {
		//fbl.Log().Sugar().Infof("TCP Listen err: %v\n", err)
		panic(err)
	}

	//启动
	server := fb_pb.InitServer(address, nil, router.BaseEndPointFunc)
	//fbl.Log().Sugar().Infof("gRPC and http listen on:%s", Port)
	go func() {
		//最大连接数
		conn = netutil.LimitListener(conn, 10000)
		defer conn.Close()
		if err = server.Serve(conn); err != nil {
			//fbl.Log().Sugar().Infof("Listen and Server err: %v", err)
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the service with a timeout of 5 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		//fbl.Log().Sugar().Fatal(err.Error())
	}
}
