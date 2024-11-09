package normalizer

import (
	"sync"
)

type NormalizedResponse struct {
	sync.RWMutex
	RpcResponse *RPCResponse
}

type RPCResponse struct {
	JsonrpcVersion string      `json:"jsonrpc"`
	Result         interface{} `json:"result"`
	Error          *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID int `json:"id"`
}

func (nr *NormalizedResponse) ParseResponse(body []byte) error {
	nr.Lock()
	defer nr.Unlock()
	nr.RpcResponse = &RPCResponse{}
	if err := SonicCfg.Unmarshal(body, nr.RpcResponse); err != nil {
		return err
	}
	return nil
}
