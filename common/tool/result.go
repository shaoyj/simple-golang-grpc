package tool

import (
	"encoding/json"
)

// 空集合
var emptyList = make([]interface{}, 0)

type PageResult[T any] struct {
	List   []T    `json:"list"`
	Anchor string `json:"anchor,omitempty"`
	End    *bool  `json:"end,omitempty"`
}

func IsJSON(in []byte) bool {
	var js map[string]interface{}
	return json.Unmarshal(in, &js) == nil

}

func OkBool(data bool) Result {
	return Result{
		Code: ok.ToInt(),
		Data: data,
	}
}

func Ok[T any](data T) Result {
	return Result{
		Code: ok.ToInt(),
		Data: data,
	}
}

func Failed(msg string) Result {
	return Result{
		Code: failed.ToInt(),
		Msg:  msg,
	}
}

func Error(err error) Result {
	return Result{
		Code: failed.ToInt(),
		Msg:  err.Error(),
	}
}

type Code int32

const ok Code = 0
const failed Code = -1

func (code Code) ToInt() int32 {
	return int32(code)
}

type Result struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
