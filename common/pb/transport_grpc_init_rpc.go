package pb

import (
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	"google.golang.org/grpc"

	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func initRpc(server *grpc.Server, baseExecute endpoint.Endpoint) {
	// Create a single logger, which we'll use and give to other components.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	zipkinBridge := ToPtr(false)
	zipkinURL := ToPtr("")

	var zipkinTracer *zipkin.Tracer
	{
		if *zipkinURL != "" {
			var (
				err         error
				hostPort    = "localhost:80"
				serviceName = "addsvc"
				reporter    = zipkinhttp.NewReporter(*zipkinURL)
			)
			defer reporter.Close()
			zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
			zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(zEP))
			if err != nil {
				logger.Log("err", err)
				os.Exit(1)
			}
			if !(*zipkinBridge) {
				logger.Log("tracer", "Zipkin", "type", "Native", "URL", *zipkinURL)
			}
		}
	}

	var tracer stdopentracing.Tracer
	{
		if *zipkinBridge && zipkinTracer != nil {
			logger.Log("tracer", "Zipkin", "type", "OpenTracing", "URL", *zipkinURL)
			tracer = zipkinot.Wrap(zipkinTracer)
			zipkinTracer = nil
		} else {
			tracer = stdopentracing.GlobalTracer() // no-op
		}
	}

	//register baseServer
	RegisterBaseServiceServer(server, NewBaseServer(baseExecute, tracer, zipkinTracer, logger))
}
