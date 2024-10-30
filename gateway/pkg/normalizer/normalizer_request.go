package normalizer

import (
	"sync"

	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
)

// RPCRequest represents the structure of an RPC request
type RPCRequest struct {
	JsonrpcVersion string        `json:"jsonrpc"`
	Method         string        `json:"method"`
	Params         []interface{} `json:"params"`
	ID             int           `json:"id"`
}

// RPCErrorResponse represents the structure of an RPC error response
type RPCErrorResponse struct {
	JsonrpcVersion string `json:"jsonrpc"`
	Error          struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID int `json:"id"`
}

// NormalizedRequest handles the parsing and management of JSON-RPC requests
type NormalizedRequest struct {
	sync.RWMutex
	body       []byte
	rpcRequest *RPCRequest
}

// NewNormalizedRequest creates a new NormalizedRequest
func NewNormalizedRequest(body []byte) *NormalizedRequest {
	return &NormalizedRequest{
		body: body,
	}
}

// ParseRequest parses the request body into an RPCRequest struct
func (nr *NormalizedRequest) ParseRequest() error {
	nr.Lock()
	defer nr.Unlock()

	if nr.rpcRequest != nil {
		return nil // Already parsed
	}

	var rpcRequest RPCRequest
	if err := SonicCfg.Unmarshal(nr.body, &rpcRequest); err != nil {
		logger.Errorf("Failed to parse JSON body: %v", err)
		return err
	}
	nr.rpcRequest = &rpcRequest
	return nil
}

// Method returns the method of the RPC request
func (nr *NormalizedRequest) Method() (string, error) {
	if err := nr.ParseRequest(); err != nil {
		return "", err
	}
	return nr.rpcRequest.Method, nil
}

// NormalizeErrorResponse creates a standardized error response
func NormalizeErrorResponse(method string, id int) []byte {
	errorResponse := RPCErrorResponse{
		JsonrpcVersion: "2.0",
		Error: struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    -32601,
			Message: "Method " + method + " is not supported",
		},
		ID: id,
	}
	responseBody, _ := SonicCfg.Marshal(errorResponse)
	return responseBody
}
