package pb

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/metadata"
)

func HandlerParams(ctx context.Context, req *ComReq) (context.Context, *ComReq) {
	//req.Headers 如果不为空 则 优先从这里解析 请求头信息
	if len(req.Headers) > 0 {
		var headerParams map[string]interface{}
		if hpEr := json.Unmarshal(req.Headers, &headerParams); hpEr == nil {
			for k, v := range headerParams {
				ctx = context.WithValue(ctx, NewParamsEnum(k), v)
			}
		}
	} else {
		//md
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			//authorization
			authorizationList := md.Get(AuthorizationKey.Val())
			if len(authorizationList) > 0 {
				ctx = context.WithValue(ctx, AuthorizationKey, authorizationList[0])
			}
			//TraceId
			TraceIdList := md.Get(TraceIdKey.Val())
			if len(TraceIdList) > 0 {
				ctx = context.WithValue(ctx, TraceIdKey, TraceIdList[0])
			}
			//uid
			uidList := md.Get(UidKey.Val())
			if len(uidList) > 0 {
				ctx = context.WithValue(ctx, UidKey, uidList[0])
			}

			//trace_id
			traceIdList := md.Get(TraceIdKey.Val())
			if len(traceIdList) > 0 {
				ctx = context.WithValue(ctx, TraceIdKey, traceIdList[0])
			}

			//x-timezone
			timeZoneList := md.Get(TimezoneKey.Val())
			if len(timeZoneList) > 0 {
				ctx = context.WithValue(ctx, TimezoneKey, timeZoneList[0])
			}

			//appVersion
			appVersionList := md.Get(AppVersionKey.Val())
			if len(appVersionList) > 0 {
				ctx = context.WithValue(ctx, AppVersionKey, appVersionList[0])
			}

			//appId
			appIdList := md.Get(AppTypeKey.Val())
			if len(appIdList) > 0 {
				ctx = context.WithValue(ctx, AppTypeKey, appIdList[0])
			}

			//platform
			platformList := md.Get(PlatformTypeKey.Val())
			if len(platformList) > 0 {
				ctx = context.WithValue(ctx, PlatformTypeKey, platformList[0])
			}

			//Terminal
			terminalList := md.Get(TerminalKey.Val())
			if len(terminalList) > 0 {
				ctx = context.WithValue(ctx, TerminalKey, terminalList[0])
			}

			//DeviceNum
			deviceNumList := md.Get(DeviceNumKey.Val())
			if len(deviceNumList) > 0 {
				ctx = context.WithValue(ctx, DeviceNumKey, deviceNumList[0])
			}

			//ProductId
			ProductIdList := md.Get(ProductIdKey.Val())
			if len(ProductIdList) > 0 {
				ctx = context.WithValue(ctx, ProductIdKey, ProductIdList[0])
			}
		}
	}

	//parse target&method
	if req.Target == "" || req.Method == "" {
		var params map[string]interface{}
		if paramsErr := json.Unmarshal(req.Body, &params); paramsErr == nil {
			//target 参数
			if req.Target == "" {
				target, targetOk := params[TargetKey.Val()]
				if targetOk {
					req.Target = ToString(target)
				}
			}
			//method 参数
			if req.Method == "" {
				method, methodOk := params[MethodKey.Val()]
				if methodOk {
					req.Method = ToString(method)
				}
			}
		}
	}

	return ctx, req
}
