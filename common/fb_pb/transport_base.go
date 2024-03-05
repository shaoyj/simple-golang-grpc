package fb_pb

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
)

type baseServer struct {
	execute      grpctransport.Handler
}

func NewBaseServer(baseExecute endpoint.Endpoint, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) BaseServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	if zipkinTracer != nil {
		options = append(options, zipkin.GRPCServerTrace(zipkinTracer))
	}

	return &baseServer{
		execute: grpctransport.NewServer(
			baseExecute,
			DecodeDefault,
			EncodeDefault,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "execute", logger)))...,
		),
	}

}

func (s *baseServer) Execute(ctx context.Context, request *ComReq) (*ByteResult, error) {
	_, rep, err := s.execute.ServeGRPC(ctx, request)
	return rep.(*ByteResult), err
}
