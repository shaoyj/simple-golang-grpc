package dto

// PageReq 请求参数
type PageReq struct {
	Page int64 `json:"page"` // 当前页
	Size int64 `json:"size"` // 尺寸
}

// KvDTO
type KvDTO struct {
	Key   string `json:"key,omitempty"`
	Value int64  `json:"value,omitempty"`
}

// req
type FbRpcReq struct {
	PageReq
	Name string `json:"name"`
	Age  int64  `json:"age"`

	K1 int64 `json:"k1"`
	K2 int64 `json:"k2"`

	KeyList []KvDTO `json:"keyList"`

	Info map[string]int64 `json:"info"`
}
