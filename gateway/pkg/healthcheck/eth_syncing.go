package healthcheck

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"

	jsonrpc "github.com/starton-io/tyrscale/gateway/pkg/jsonrpc"
	"github.com/valyala/fasthttp"
)

type EthSyncing struct {
	clientManager  proxy.ClientManager
	CircuitBreaker circuitbreaker.ProxyCircuitBreaker
	interval       time.Duration
	timeout        time.Duration
}

type EthSyncingOption func(*EthSyncing)

func NewHealthEthSyncing(clientManager proxy.ClientManager, opts ...EthSyncingOption) HealthCheckInterface {
	h := &EthSyncing{
		clientManager: clientManager,
		interval:      time.Duration(10 * time.Second),
		timeout:       time.Duration(10 * time.Second),
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func WithEthSyncingInterval(interval time.Duration) EthSyncingOption {
	return func(h *EthSyncing) {
		h.interval = interval
	}
}

func WithEthSyncingTimeout(timeout time.Duration) EthSyncingOption {
	return func(h *EthSyncing) {
		h.timeout = timeout
	}
}

func (h *EthSyncing) GetInterval() time.Duration {
	return h.interval
}

func (h *EthSyncing) SetCircuitBreaker(circuitBreaker circuitbreaker.ProxyCircuitBreaker) {
	h.CircuitBreaker = circuitBreaker
}

func GetSyncingStatus(c *proxy.UpstreamClient, timeout time.Duration) (bool, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	c.RequestInterceptor.Intercept(req)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody([]byte(`{"method":"eth_syncing","params":[],"id":1,"jsonrpc":"2.0"}`))

	if err := c.Client.DoTimeout(req, resp, timeout); err != nil {
		logger.Errorf("error doing request: %v", err)
		return false, err
	}

	if resp.StatusCode() != http.StatusOK {
		logger.Errorf("unhealthy service status code: %d", resp.StatusCode())
		return false, fmt.Errorf("unhealthy service status code: %d", resp.StatusCode())
	}

	var rpcResp jsonrpc.JsonrpcMessage
	if err := json.Unmarshal(resp.Body(), &rpcResp); err != nil {
		logger.Errorf("error unmarshalling response: %v", err)
		return false, err
	}

	if rpcResp.Error != nil {
		logger.Errorf("JSON-RPC error: %s", rpcResp.Error.Message)
		return false, fmt.Errorf("JSON-RPC error: %s", rpcResp.Error.Message)
	}

	if rpcResp.Result == nil {
		logger.Errorf("invalid JSON response from the service: result field is missing")
		return false, fmt.Errorf("invalid JSON response from the service: result field is missing")
	}

	isSyncing := string(rpcResp.Result) != "false"
	logger.Debugf("upstream UUID: %s, syncing: %t", c.Client.Name, isSyncing)
	return isSyncing, nil
}

func (h *EthSyncing) CheckHealth() error {
	var wg sync.WaitGroup
	clients := h.clientManager.GetAllClients()

	for _, client := range clients {
		wg.Add(1)

		// no limit on goroutines because we want to check the syncing status for all clients at the same time
		go func(client *proxy.UpstreamClient) {
			defer wg.Done()
			var (
				isSyncing interface{}
				err       error
			)
			if h.CircuitBreaker != nil {
				cb := h.CircuitBreaker.Get(client.Client.Name)
				isSyncing, err = cb.Execute(func() (interface{}, error) {
					return GetSyncingStatus(client, h.timeout)
				})
			} else {
				isSyncing, err = GetSyncingStatus(client, h.timeout)
			}
			if err != nil {
				client.Healthy = false
				return
			}
			isSyncingBool, ok := isSyncing.(bool)
			if !ok {
				client.Healthy = false
				return
			}
			client.Healthy = !isSyncingBool
		}(client)
	}

	wg.Wait()

	// Update client health status
	for _, client := range clients {
		log.Println(client.Client.Name, client.Healthy)
		h.clientManager.SetHealthy(client.Client.Name, client.Healthy)
	}
	return nil
}
