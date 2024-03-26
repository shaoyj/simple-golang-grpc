package helper

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/metadata"
	"simple-go-grpc/common/pb"
	"simple-go-grpc/common/tool"
)

func HandlerParams(ctx context.Context, req *pb.ComReq) (context.Context, *pb.ComReq) {
	//req.Headers 如果不为空 则 优先从这里解析 请求头信息
	if len(req.Headers) > 0 {
		var headerParams map[string]interface{}
		if hpEr := json.Unmarshal(req.Headers, &headerParams); hpEr == nil {
			for k, v := range headerParams {
				ctx = context.WithValue(ctx, tool.NewParamsEnum(k), v)
			}
		}
	} else {
		//md
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			//authorization
			authorizationList := md.Get(tool.AuthorizationKey.Val())
			if len(authorizationList) > 0 {
				ctx = context.WithValue(ctx, tool.AuthorizationKey, authorizationList[0])
			}
			//TraceId
			TraceIdList := md.Get(tool.TraceIdKey.Val())
			if len(TraceIdList) > 0 {
				ctx = context.WithValue(ctx, tool.TraceIdKey, TraceIdList[0])
			}
			//uid
			uidList := md.Get(tool.UidKey.Val())
			if len(uidList) > 0 {
				ctx = context.WithValue(ctx, tool.UidKey, uidList[0])
			}

			//trace_id
			traceIdList := md.Get(tool.TraceIdKey.Val())
			if len(traceIdList) > 0 {
				ctx = context.WithValue(ctx, tool.TraceIdKey, traceIdList[0])
			}

			//x-timezone
			timeZoneList := md.Get(tool.TimezoneKey.Val())
			if len(timeZoneList) > 0 {
				ctx = context.WithValue(ctx, tool.TimezoneKey, timeZoneList[0])
			}

			//appVersion
			appVersionList := md.Get(tool.AppVersionKey.Val())
			if len(appVersionList) > 0 {
				ctx = context.WithValue(ctx, tool.AppVersionKey, appVersionList[0])
			}

			//appId
			appIdList := md.Get(tool.AppTypeKey.Val())
			if len(appIdList) > 0 {
				ctx = context.WithValue(ctx, tool.AppTypeKey, appIdList[0])
			}

			//platform
			platformList := md.Get(tool.PlatformTypeKey.Val())
			if len(platformList) > 0 {
				ctx = context.WithValue(ctx, tool.PlatformTypeKey, platformList[0])
			}

			//Terminal
			terminalList := md.Get(tool.TerminalKey.Val())
			if len(terminalList) > 0 {
				ctx = context.WithValue(ctx, tool.TerminalKey, terminalList[0])
			}

			//DeviceNum
			deviceNumList := md.Get(tool.DeviceNumKey.Val())
			if len(deviceNumList) > 0 {
				ctx = context.WithValue(ctx, tool.DeviceNumKey, deviceNumList[0])
			}

			//ProductId
			ProductIdList := md.Get(tool.ProductIdKey.Val())
			if len(ProductIdList) > 0 {
				ctx = context.WithValue(ctx, tool.ProductIdKey, ProductIdList[0])
			}
		}
	}

	//parse target&method
	if req.Target == "" || req.Method == "" {
		var params map[string]interface{}
		if paramsErr := json.Unmarshal(req.Body, &params); paramsErr == nil {
			//target 参数
			if req.Target == "" {
				target, targetOk := params[tool.TargetKey.Val()]
				if targetOk {
					req.Target = tool.ToString(target)
				}
			}
			//method 参数
			if req.Method == "" {
				method, methodOk := params[tool.MethodKey.Val()]
				if methodOk {
					req.Method = tool.ToString(method)
				}
			}
		}
	}

	return ctx, req
}
