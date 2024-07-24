package healthcheck

import (
	"sync"
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

type CustomHealthCheck struct {
	clientManager  proxy.ClientManager
	CircuitBreaker circuitbreaker.ProxyCircuitBreaker
	interval       time.Duration
	timeout        time.Duration
	request        *Request
}

type CustomHealthCheckOption func(*CustomHealthCheck)

func NewHealthCustom(clientManager proxy.ClientManager, opts ...CustomHealthCheckOption) HealthCheckInterface {
	h := &CustomHealthCheck{
		clientManager: clientManager,
		interval:      time.Duration(10 * time.Second),
		timeout:       time.Duration(10 * time.Second),
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func WithCustomHealthCheckInterval(interval time.Duration) CustomHealthCheckOption {
	return func(h *CustomHealthCheck) {
		h.interval = interval
	}
}

func WithCustomHealthCheckTimeout(timeout time.Duration) CustomHealthCheckOption {
	return func(h *CustomHealthCheck) {
		h.timeout = timeout
	}
}

func WithCustomHealthCheckRequest(request *Request) CustomHealthCheckOption {
	return func(h *CustomHealthCheck) {
		h.request = request
	}
}

func (h *CustomHealthCheck) GetInterval() time.Duration {
	return h.interval
}

func (h *CustomHealthCheck) SetCircuitBreaker(circuitBreaker circuitbreaker.ProxyCircuitBreaker) {
	h.CircuitBreaker = circuitBreaker
}

func CustomRequest(c *proxy.UpstreamClient, requestConfig *Request, timeout time.Duration) (int, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	err := c.RequestInterceptor.Intercept(req)
	if err != nil {
		return fasthttp.StatusInternalServerError, err
	}
	req.Header.SetMethod(requestConfig.Method)

	// set Content-Type by default application/json
	req.Header.SetContentType("application/json")

	// check if request headers have content type
	if requestConfig.Headers != nil {
		for key, value := range requestConfig.Headers {
			req.Header.Set(key, value)
		}
	}

	req.SetBody([]byte(requestConfig.Body))

	if err := c.Client.DoTimeout(req, resp, timeout); err != nil {
		logger.Errorf("error doing request: %v", err)
		return resp.StatusCode(), err
	}

	return resp.StatusCode(), nil
}

func (h *CustomHealthCheck) CheckHealth() error {
	var wg sync.WaitGroup
	clients := h.clientManager.GetAllClients()

	for _, client := range clients {
		wg.Add(1)

		// no limit on goroutines because we want to check the syncing status for all clients at the same time
		go func(client *proxy.UpstreamClient) {
			defer wg.Done()
			var (
				statusCode interface{}
				err        error
			)
			if h.CircuitBreaker != nil {
				cb := h.CircuitBreaker.Get(client.Client.Name)
				statusCode, err = cb.Execute(func() (interface{}, error) {
					return CustomRequest(client, h.request, h.timeout)
				})
			} else {
				statusCode, err = CustomRequest(client, h.request, h.timeout)
			}
			if err != nil {
				client.Healthy = false
				return
			}
			statusCodeInt, ok := statusCode.(int)
			if !ok {
				client.Healthy = false
				return
			}
			client.Healthy = statusCodeInt == int(h.request.StatusCode)
		}(client)
	}

	wg.Wait()

	// Update client health status
	for _, client := range clients {
		logger.Debugf("upstream %s health status: %v", client.Client.Name, client.Healthy)
		h.clientManager.SetHealthy(client.Client.Name, client.Healthy)
	}
	return nil
}
