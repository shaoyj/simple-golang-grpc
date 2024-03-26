package rpc

import (
	"context"
	"fmt"
	"simple-go-grpc/common/dto"

	// "encoding/json"
	// "fmt"
	"testing"
)

func TestClient(t *testing.T) {
	var req dto.FbRpcReq
	req.K1 = 18101029476241408
	req.K2 = 18101029476241410
	result := FindInfo(context.Background(), req)
	fmt.Println(result)
}
