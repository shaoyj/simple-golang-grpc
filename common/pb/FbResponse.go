package pb

type FbResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data"`
}

func Suc(data any) FbResponse {
	return FbResponse{
		Code: ok.ToInt(),
		Data: data,
	}
}

func Fail(msg string) FbResponse {
	return FbResponse{
		Code: failed.ToInt(),
		Msg:  msg,
	}
}
