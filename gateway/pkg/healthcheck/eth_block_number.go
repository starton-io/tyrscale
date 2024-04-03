package healthcheck

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"

	jsonrpc "github.com/starton-io/tyrscale/gateway/pkg/jsonrpc"
	"github.com/valyala/fasthttp"
)

type EthBlockNumber struct {
	clientManager  proxy.ClientManager
	CircuitBreaker circuitbreaker.ProxyCircuitBreaker
	highestBlock   uint64
	interval       time.Duration
	timeout        time.Duration
}

type EthBlockNumberOption func(*EthBlockNumber)

func NewHealthEthBlockNumber(clientManager proxy.ClientManager, opts ...EthBlockNumberOption) HealthCheckInterface {
	h := &EthBlockNumber{
		clientManager: clientManager,
		interval:      time.Duration(10 * time.Second),
		timeout:       time.Duration(10 * time.Second),
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func WithInterval(interval time.Duration) EthBlockNumberOption {
	return func(h *EthBlockNumber) {
		h.interval = interval
	}
}

func WithTimeout(timeout time.Duration) EthBlockNumberOption {
	return func(h *EthBlockNumber) {
		h.timeout = timeout
	}
}

func (h *EthBlockNumber) GetInterval() time.Duration {
	return h.interval
}

func (h *EthBlockNumber) SetCircuitBreaker(circuitBreaker circuitbreaker.ProxyCircuitBreaker) {
	h.CircuitBreaker = circuitBreaker
}

func GetBlockNumber(c *proxy.UpstreamClient, timeout time.Duration) (uint64, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	c.RequestInterceptor.Intercept(req)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody([]byte(`{"method":"eth_blockNumber","params":[],"id":1,"jsonrpc":"2.0"}`))

	if err := c.Client.DoTimeout(req, resp, timeout); err != nil {
		return 0, err
	}

	if resp.StatusCode() != http.StatusOK {
		return 0, fmt.Errorf("unhealthy service status code: %d", resp.StatusCode())
	}

	var rpcResp jsonrpc.JsonrpcMessage
	if err := json.Unmarshal(resp.Body(), &rpcResp); err != nil {
		return 0, err
	}

	if rpcResp.Error != nil {
		return 0, fmt.Errorf("JSON-RPC error: %s", rpcResp.Error.Message)
	}

	if rpcResp.Result == nil {
		return 0, fmt.Errorf("invalid JSON response from the service: result field is missing")
	}
	return hexutil.DecodeUint64(string(rpcResp.Result[1 : len(rpcResp.Result)-1]))
}

func (h *EthBlockNumber) CheckHealth() error {
	var wg sync.WaitGroup
	clients := h.clientManager.GetAllClients()
	mapBlockNumber := make(map[string]uint64)
	var highestBlock uint64
	var muBlock sync.Mutex // Mutex for safe access to mapBlockNumber and highestBlock

	for _, client := range clients {
		wg.Add(1)

		// no limit on goroutines because we want to check the current block number for all clients at the same time
		go func(client *proxy.UpstreamClient) {
			defer wg.Done()
			var (
				blockNumber interface{}
				err         error
			)
			if h.CircuitBreaker != nil {
				cb := h.CircuitBreaker.Get(client.Client.Name)
				blockNumber, err = cb.Execute(func() (interface{}, error) {
					return GetBlockNumber(client, h.timeout)
				})
			} else {
				blockNumber, err = GetBlockNumber(client, h.timeout)
			}
			if err != nil {
				client.Healthy = false
				return
			}
			blockNumberUint, ok := blockNumber.(uint64)
			if !ok {
				client.Healthy = false
				return
			}
			muBlock.Lock()
			defer muBlock.Unlock()
			mapBlockNumber[client.Client.Name] = blockNumberUint
			if blockNumberUint > highestBlock {
				highestBlock = blockNumberUint
			}
		}(client)
	}

	wg.Wait()
	// Update the highest block safely
	h.highestBlock = highestBlock

	// Update client health status
	for _, client := range clients {
		h.clientManager.SetHealthy(client.Client.Name, client.Healthy && mapBlockNumber[client.Client.Name] >= h.highestBlock && mapBlockNumber[client.Client.Name] != 0)
	}
	return nil
}
