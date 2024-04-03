package proxy

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/interceptor"
	jsonrpc "github.com/starton-io/tyrscale/gateway/pkg/jsonrpc"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
	"github.com/valyala/fasthttp"
)

type UpstreamClient struct {
	Client              *fasthttp.HostClient
	RequestInterceptor  interceptor.InterceptorRequestChain
	ResponseInterceptor interceptor.InterceptorResponseChain
	Healthy             bool
}

//ype DefaultProxyController struct {
//   mu             sync.Mutex
//   MapClient      map[string]*UpstreamClient
//   Balancer       balancer.IBalancer
//   CircuitBreaker circuitbreaker.ProxyCircuitBreaker
//}//

// Ensure DefaultProxyController implements the ProxyManager interface
//var _ ProxyManager = &DefaultProxyController{}

//type ProxyControllerOption func(*DefaultProxyController)

//func (m *DefaultProxyController) AddUpstream(upstream *upstream.UpstreamPublishUpsertModel) error {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//	fmt.Println("AddUpstream", upstream)
//	chain := interceptor.InterceptorRequestChain{}
//	chain.Add(&interceptor.DefaultRequestInterceptor{
//		Host:   upstream.Host,
//		Path:   upstream.Path,
//		Scheme: upstream.Scheme,
//		Port:   upstream.Port,
//	})
//	chainRes := interceptor.InterceptorResponseChain{}
//	chainRes.Add(&interceptor.DefaultResponseInterceptor{})
//
//	proxy := &UpstreamClient{
//		Client: &fasthttp.HostClient{
//			ReadTimeout:  5 * time.Second,
//			WriteTimeout: 5 * time.Second,
//			Addr:         upstream.Host,
//			Name:         upstream.Uuid,
//			IsTLS:        upstream.Scheme == "https",
//			TLSConfig: &tls.Config{
//				InsecureSkipVerify: true,
//			},
//			DisablePathNormalizing: true,
//		},
//		RequestInterceptor:  chain,
//		ResponseInterceptor: chainRes,
//	}
//	m.MapClient[upstream.Uuid] = proxy
//	m.Balancer.AddServer(&balancer.Server{
//		Uuid:   upstream.Uuid,
//		Weight: int(upstream.Weight),
//	})
//
//	if m.CircuitBreaker != nil {
//		m.CircuitBreaker.Add(upstream.Uuid)
//	}
//
//	return nil
//}
//
//func (m *DefaultProxyController) ExistUpstream(upstreamUuid string) bool {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//
//	_, ok := m.MapClient[upstreamUuid]
//	return ok
//}
//
//func (m *DefaultProxyController) UpdateUpstream(upstream *upstream.UpstreamPublishUpsertModel) error {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//
//	if m.ExistUpstream(upstream.Uuid) {
//		return m.Balancer.UpdateWeight(upstream.Uuid, int(upstream.Weight))
//	}
//	return fmt.Errorf("upstream not found")
//}
//
//func (m *DefaultProxyController) RemoveUpstream(uuid string) error {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//
//	if clientProxy, ok := m.MapClient[uuid]; ok {
//		clientProxy.Client.CloseIdleConnections() // Gracefully close the connections
//	}
//
//	err := m.Balancer.RemoveServer(uuid)
//	if err != nil {
//		return err
//	}
//	delete(m.MapClient, uuid)
//
//	if m.CircuitBreaker != nil {
//		m.CircuitBreaker.Remove(uuid)
//	}
//	return nil
//}
//
//func (m *DefaultProxyController) GetClient(uuid string) (*UpstreamClient, bool) {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//
//	client, ok := m.MapClient[uuid]
//	return client, ok
//}
//
//func (m *DefaultProxyController) GetCircuitBreaker(uuid string) *gobreaker.CircuitBreaker {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//	return m.CircuitBreaker.Get(uuid)
//}
//
//func (m *DefaultProxyController) GetBalancer() balancer.IBalancer {
//	return m.Balancer
//}
//
//func (m *DefaultProxyController) CloseAll() {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//
//	for _, client := range m.MapClient {
//		client.Client.CloseIdleConnections()
//	}
//	if m.CircuitBreaker != nil {
//		m.CircuitBreaker.Clean()
//	}
//}
//
//func (m *DefaultProxyController) Close(uuid string) {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//
//	if client, ok := m.MapClient[uuid]; ok {
//		client.Client.CloseIdleConnections()
//	}
//
//}
//
//func (p *DefaultProxyController) Do(client *fasthttp.HostClient, req *fasthttp.Request, res *fasthttp.Response) error {
//	return client.Do(req, res)
//}

//func (p *ProxyController) ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
//	req := &ctx.Request
//	res := &ctx.Response
//
//	if ip, _, err := net.SplitHostPort(ctx.RemoteAddr().String()); err == nil {
//		req.Header.Add("X-Forwarded-For", ip)
//	}
//
//	// remove hop-by-hop headers
//	for _, hopHeader := range hopHeaders {
//		req.Header.DelBytes(hopHeader)
//	}
//
//	p.Handler.Handle(ctx)
//
//	// remove hop-by-hop headers
//	for _, hopHeader := range hopHeaders {
//		res.Header.DelBytes(hopHeader)
//	}
//}

//type Proxy struct {
//	Closer  ConnectionManager
//	Handler ReverseProxyHandler
//}

//func (p *Proxy) GetBalancer() balancer.IBalancer {
//	return p.Controller.GetBalancer()
//}
//
//func (p *Proxy) GetClient(uuid string) (*UpstreamClient, bool) {
//	return p.Controller.GetClient(uuid)
//}
//
//func (p *Proxy) GetCircuitBreaker(uuid string) *gobreaker.CircuitBreaker {
//	return p.Controller.CircuitBreaker.Get(uuid)
//}
//
//func (p *Proxy) DoWithTimeout(client *fasthttp.HostClient, req *fasthttp.Request, res *fasthttp.Response) error {
//	if p.Timeout <= 0 {
//		return client.Do(req, res)
//	}
//	return client.DoTimeout(req, res, p.Timeout)
//}

//func NewProxy(handler, closer ConnectionManager) *Proxy {
//	return &Proxy{
//		Handler: handler,
//		Closer:  closer,
//	}
//}

//func (p *Proxy) ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
//	req := &ctx.Request
//	res := &ctx.Response
//
//	if ip, _, err := net.SplitHostPort(ctx.RemoteAddr().String()); err == nil {
//		req.Header.Add("X-Forwarded-For", ip)
//	}
//
//	// remove hop-by-hop headers
//	for _, hopHeader := range hopHeaders {
//		req.Header.DelBytes(hopHeader)
//	}
//
//	p.Handler.Handle(ctx)
//
//	// remove hop-by-hop headers
//	for _, hopHeader := range hopHeaders {
//		res.Header.DelBytes(hopHeader)
//	}
//}
//
//func (p *Proxy) Close(uuid string) {
//	p.Closer.Close(uuid)
//}

//// -- proxyV2
//
//type ProxyV2 struct {
//	MapClient map[string]*UpstreamClient
//	Balancer  balancer.IBalancer
//	Handler   ReverseProxyHandlerV2
//}
//
//func (p *ProxyV2) ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
//	req := &ctx.Request
//	res := &ctx.Response
//
//	if ip, _, err := net.SplitHostPort(ctx.RemoteAddr().String()); err == nil {
//		req.Header.Add("X-Forwarded-For", ip)
//	}
//
//	// remove hop-by-hop headers
//	for _, hopHeader := range hopHeaders {
//		req.Header.DelBytes(hopHeader)
//	}
//
//	p.Handler.Handle(ctx, p)
//
//	// remove hop-by-hop headers
//	for _, hopHeader := range hopHeaders {
//		res.Header.DelBytes(hopHeader)
//	}
//}

// ClientManager manages a map of UpstreamClients.
type ClientManager interface {
	AddClient(uuid string, client *UpstreamClient)
	GetClient(uuid string) (*UpstreamClient, bool)
	GetAllClients() map[string]*UpstreamClient
	RemoveClient(uuid string)
	SetHealthy(uuid string, healthy bool)
	IsHealthy(uuid string) bool
	Close()
}

// DefaultClientManager is a default implementation of ClientManager.
type DefaultClientManager struct {
	mu           sync.Mutex
	MapClient    map[string]*UpstreamClient
	HighestBlock uint64
}

func NewDefaultClientManager() *DefaultClientManager {
	return &DefaultClientManager{
		MapClient: make(map[string]*UpstreamClient),
	}
}

func (m *DefaultClientManager) AddClient(uuid string, client *UpstreamClient) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.MapClient[uuid] = client
}

func (m *DefaultClientManager) GetClient(uuid string) (*UpstreamClient, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	client, ok := m.MapClient[uuid]
	return client, ok
}

func (m *DefaultClientManager) GetAllClients() map[string]*UpstreamClient {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.MapClient
}

func (m *DefaultClientManager) RemoveClient(uuid string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.MapClient, uuid)
}

func (m *DefaultClientManager) IsHealthy(uuid string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.MapClient[uuid].Healthy
}

func (m *DefaultClientManager) SetHealthy(uuid string, healthy bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, ok := m.MapClient[uuid]; ok {
		client.Healthy = healthy
	}
}

func GetBlockNumber(c *UpstreamClient) (uint64, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	c.RequestInterceptor.Intercept(req)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody([]byte(`{"method":"eth_blockNumber","params":[],"id":1,"jsonrpc":"2.0"}`))

	if err := c.Client.Do(req, resp); err != nil {
		return 0, err
	}

	if resp.StatusCode() != http.StatusOK {
		return 0, fmt.Errorf("unhealthy service status code: %d", resp.StatusCode())
	}

	var rpcResp jsonrpc.JsonrpcMessage
	if err := json.Unmarshal(resp.Body(), &rpcResp); err != nil {
		return 0, err
	}
	log.Println("rpcResp", rpcResp)

	if rpcResp.Error != nil {
		return 0, fmt.Errorf("JSON-RPC error: %s", rpcResp.Error.Message)
	}

	if rpcResp.Result == nil {
		return 0, fmt.Errorf("invalid JSON response from the service: result field is missing")
	}
	return hexutil.DecodeUint64(string(rpcResp.Result[1 : len(rpcResp.Result)-1]))
}

func (m *DefaultClientManager) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, client := range m.MapClient {
		client.Client.CloseIdleConnections()
	}
}

type ProxyController struct {
	Labels         map[string]string
	ClientManager  *DefaultClientManager
	Balancer       balancer.IBalancer
	CircuitBreaker circuitbreaker.ProxyCircuitBreaker
}

func NewProxyController(typeBalancer balancer.LoadBalancerStrategy, labels map[string]string) *ProxyController {
	controller := &ProxyController{
		ClientManager: NewDefaultClientManager(),
		Balancer:      balancer.NewBalancer(typeBalancer),
		Labels:        labels,
	}
	return controller
}

func (m *ProxyController) GetLabelValue(key string) string {
	return m.Labels[key]
}

func (m *ProxyController) AddUpstream(upstream *upstream.UpstreamPublishUpsertModel) {
	chain := interceptor.InterceptorRequestChain{}
	chain.Add(&interceptor.DefaultRequestInterceptor{
		Host:   upstream.Host,
		Path:   upstream.Path,
		Scheme: upstream.Scheme,
		Port:   upstream.Port,
	})
	chainRes := interceptor.InterceptorResponseChain{}
	chainRes.Add(&interceptor.DefaultResponseInterceptor{})

	proxy := &UpstreamClient{
		Healthy: true,
		Client: &fasthttp.HostClient{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			Addr:         upstream.Host,
			Name:         upstream.Uuid,
			IsTLS:        upstream.Scheme == "https",
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DisablePathNormalizing: true,
		},
		RequestInterceptor:  chain,
		ResponseInterceptor: chainRes,
	}

	m.ClientManager.AddClient(upstream.Uuid, proxy)
	m.Balancer.AddServer(&balancer.Server{
		Uuid:   upstream.Uuid,
		Weight: int(upstream.Weight),
	})
	if m.CircuitBreaker != nil {
		m.CircuitBreaker.Add(upstream.Uuid)
	}
}

//func (m *ProxyController) addScheduler(interval time.Duration, taskFunc func() error) {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//
//	m.Scheduler.Add(&tasks.Task{
//		RunSingleInstance: true,
//		Interval:          interval,
//		TaskFunc:          taskFunc,
//	})
//}

//func (p *ProxyController) PerformGlobalCheck() {
//	var wg sync.WaitGroup
//	mapBlockNumber := make(map[string]uint64)
//	var highestBlock uint64
//	var muBlock sync.Mutex // Mutex for safe access to mapBlockNumber and highestBlock
//
//	clients := p.ClientManager.GetAllClients()
//	for _, client := range clients {
//		wg.Add(1)
//		go func(c *UpstreamClient) {
//			defer wg.Done()
//			// get circuit breaker
//			var (
//				blockNumber interface{}
//				err         error
//			)
//			if p.CircuitBreaker != nil {
//				cb := p.CircuitBreaker.Get(c.Client.Name)
//				blockNumber, err = cb.Execute(func() (interface{}, error) {
//					return GetBlockNumber(c)
//				})
//			} else {
//				blockNumber, err = GetBlockNumber(c)
//			}
//			if err != nil {
//				c.Healthy = false
//				return
//			}
//			blockNumberUint, ok := blockNumber.(uint64)
//			if !ok {
//				c.Healthy = false
//				return
//			}
//			muBlock.Lock()
//			defer muBlock.Unlock()
//			mapBlockNumber[c.Client.Name] = blockNumberUint
//			if blockNumberUint > highestBlock {
//				highestBlock = blockNumberUint
//			}
//		}(client)
//	}
//	wg.Wait()
//
//	// Update the highest block safely
//	p.ClientManager.HighestBlock = highestBlock
//
//	// Update client health status
//	for _, client := range clients {
//		client.Healthy = mapBlockNumber[client.Client.Name] >= highestBlock
//	}
//}

func (m *ProxyController) CloseAll() {
}

func (m *ProxyController) RemoveUpstream(uuid string) {
	client, ok := m.ClientManager.GetClient(uuid)
	if ok {
		client.Client.CloseIdleConnections() // Gracefully close the connections
	}
	m.ClientManager.RemoveClient(uuid)
	_ = m.Balancer.RemoveServer(uuid)
	if m.CircuitBreaker != nil {
		m.CircuitBreaker.Remove(uuid)
	}
}

func (m *ProxyController) UpdateUpstream(upstream *upstream.UpstreamPublishUpsertModel) error {
	if _, ok := m.ClientManager.GetClient(upstream.Uuid); ok {
		logger.Infof("update upstream %s, weight: %f", upstream.Uuid, upstream.Weight)
		return m.Balancer.UpdateWeight(upstream.Uuid, int(upstream.Weight))
	}
	logger.Error("upstream not found")
	return fmt.Errorf("upstream not found")
}
