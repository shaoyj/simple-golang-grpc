package helper

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/metadata"
	"simple-go-grpc/common/logs"
	"simple-go-grpc/common/pb"
	"simple-go-grpc/common/tool"
	"time"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net/http"
	"strings"
)

// InitServer  创建grpc服务
func InitServer(address string, tlsConfig *tls.Config, baseExecute endpoint.Endpoint) *http.Server {
	//server := grpc.NewServer()
	server := grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		// 在这里拦截下行程序抛出的异常
		defer func() {

			diff := time.Now().UnixMilli() - start.UnixMilli()
			logs.NewLog("").Infof("cost %d", diff)
			//md
			traceId := ""
			md, ok := metadata.FromIncomingContext(ctx)
			if ok {
				traceIdList := md.Get(tool.RequestIdKey.Val())
				if len(traceIdList) > 0 {
					traceId = traceIdList[0]
				}
			}
			//header
			comReq, ok := req.(pb.ComReq)
			if ok && len(comReq.Headers) > 2 && traceId == "" {
				var headerMap map[string]string
				jsonErr := json.Unmarshal(comReq.Headers, &headerMap)
				if jsonErr != nil {
					reqId, ok := headerMap[tool.RequestIdKey.Val()]
					if ok {
						traceId = reqId
					}
				}
			}

			if err := recover(); err != nil {
				logs.NewLog("").Error(err)
			}

		}()

		// 执行对应的业务方法
		resp, err = handler(ctx, req)
		return resp, err
	}))

	// register rpc service
	initRpc(server, baseExecute)

	// reflection register (这行代码不是必须的，加入反射是为了后续使用grpcurl和grpcui。)
	reflection.Register(server)

	// 注册网关服务
	mux := http.NewServeMux()
	mux.Handle("/", initGw(address, []grpc.DialOption{grpc.WithInsecure()}))

	// register checkHealth
	mux.HandleFunc("/health", checkHealth)

	return &http.Server{
		Addr:           address,
		Handler:        GrpcHandlerFunc(server, mux),
		TLSConfig:      tlsConfig,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

/*
*
* 健康检查
 */
func checkHealth(w http.ResponseWriter, r *http.Request) {
	if strings.EqualFold(r.URL.Path, "/health") {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
		return
	}

	http.NotFound(w, r)
}
