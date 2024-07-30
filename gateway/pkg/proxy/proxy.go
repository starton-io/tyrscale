package proxy

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/interceptor"
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

func (m *DefaultClientManager) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, client := range m.MapClient {
		client.Client.CloseIdleConnections()
	}
}

type ProxyController struct {
	mu                    sync.Mutex
	Labels                map[string]string
	ClientManager         *DefaultClientManager
	Balancer              balancer.IBalancer
	CircuitBreaker        circuitbreaker.ProxyCircuitBreaker
	ResponsesInterceptors interceptor.InterceptorResponseChain
	RequestsInterceptors  interceptor.InterceptorRequestChain
}

func NewProxyController(typeBalancer balancer.LoadBalancerStrategy, labels map[string]string) *ProxyController {
	controller := &ProxyController{
		ClientManager:         NewDefaultClientManager(),
		Balancer:              balancer.NewBalancer(typeBalancer),
		Labels:                labels,
		ResponsesInterceptors: interceptor.NewInterceptorResponseChain(),
		RequestsInterceptors:  interceptor.NewInterceptorRequestChain(),
	}
	return controller
}

func (m *ProxyController) GetRequestsInterceptors() interceptor.InterceptorRequestChain {
	return m.RequestsInterceptors
}

func (m *ProxyController) GetResponsesInterceptors() interceptor.InterceptorResponseChain {
	return m.ResponsesInterceptors
}

func (m *ProxyController) SetResponsesInterceptors(interceptor interceptor.InterceptorResponseChain) {
	m.ResponsesInterceptors = interceptor
}

func (m *ProxyController) SetRequestsInterceptors(interceptor interceptor.InterceptorRequestChain) {
	m.RequestsInterceptors = interceptor
}

func (m *ProxyController) GetLabelValue(key string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.Labels[key]
}

func (m *ProxyController) AddUpstream(upstream *upstream.UpstreamPublishUpsertModel) {
	chainReq := m.RequestsInterceptors
	chainReq.AddFirst(&interceptor.DefaultRequestInterceptor{
		Host:   upstream.Host,
		Path:   upstream.Path,
		Scheme: upstream.Scheme,
		Port:   upstream.Port,
	}, "default")
	chainRes := m.ResponsesInterceptors
	chainRes.AddFirst(&interceptor.DefaultResponseInterceptor{}, "default")
	//chainRes.AddLast(&interceptor.DefaultLastResponseInterceptor{})

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
		RequestInterceptor:  chainReq,
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

// New methods to update interceptors for all upstream clients
func (m *ProxyController) UpdateRequestInterceptors(newInterceptor interceptor.InterceptorRequestChain) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.RequestsInterceptors = newInterceptor
	for _, client := range m.ClientManager.GetAllClients() {
		client.RequestInterceptor.KeepFirstAndReplaceOthers(newInterceptor)
	}
}

func (m *ProxyController) UpdateResponseInterceptors(newInterceptor interceptor.InterceptorResponseChain) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ResponsesInterceptors = newInterceptor
	for _, client := range m.ClientManager.GetAllClients() {
		client.ResponseInterceptor.KeepFirstAndReplaceOthers(newInterceptor)
	}
}
