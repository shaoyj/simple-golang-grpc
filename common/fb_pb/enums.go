package fb_pb

import "strings"

type ParamsEnum string

const (
	TargetKey ParamsEnum = "target"
	MethodKey ParamsEnum = "method"

	AuthorizationKey ParamsEnum = "authorization"
	UidKey           ParamsEnum = "uid"
	TraceIdKey       ParamsEnum = "trace_id"
	TimezoneKey      ParamsEnum = "x-timezone"
	AppTypeKey       ParamsEnum = "x-app-type"
	AppVersionKey    ParamsEnum = "appversion"

	DeviceNumKey ParamsEnum = "x-device-number"
	ProductIdKey ParamsEnum = "x-product-id"

	RequestIdKey ParamsEnum = "x-request-id"

	PlatformTypeKey ParamsEnum = "x-platform-type"
	TerminalKey     ParamsEnum = "app-os"
	OsVersionKey    ParamsEnum = "osversion"

	FbTimestampKey ParamsEnum = "fbTimestamp"
)

func NewParamsEnum(val string) ParamsEnum {
	enumVal := strings.ToLower(val)
	switch enumVal {
	case "authorization":
		return AuthorizationKey
	case "trace_id":
		return TraceIdKey
	case "x-timezone":
		return TimezoneKey

	case "x-app-type":
		return AppTypeKey
	case "x-product-id":
		return ProductIdKey
	case "x-device-number":
		return DeviceNumKey

	case "x-platform-type":
		return PlatformTypeKey
	case "x-request-id":
		return RequestIdKey
	default:
		return ParamsEnum(strings.ToLower(val))
	}
}

func (v ParamsEnum) Val() string {
	return string(v)
}
