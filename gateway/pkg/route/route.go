package route

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/starton-io/tyrscale/gateway/pkg/healthcheck"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/starton-io/tyrscale/gateway/pkg/reverseproxy"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

type IRouter interface {
	Upsert(route *Route) error
	Remove(host string, path string) error
	AddHealthCheck(id string, healthCheck healthcheck.HealthCheckInterface)
	RemoveHealthCheck(id string)
	ProxyRouter(ctx *fasthttp.RequestCtx)
	GetRoute(host string, path string) (*Route, error)
	GetProxyController(host string, path string) (*proxy.ProxyController, error)
}

type Route struct {
	NormalizeHostURI  string
	Uuid              string
	HealthCheckConfig *healthcheck.HealthCheckConfig

	ReverseProxy    reverseproxy.ProxyHandler
	ProxyController *proxy.ProxyController

	Middleware     types.MiddlewareFunc
	ListMiddleware []*middleware.MiddlewareWithPriority
	Host           string
	Path           string
}

type Router struct {
	routes             map[string]*Route
	HealthCheckManager healthcheck.HealthCheckManagerInterface
	port               int32
	mutex              sync.RWMutex
}

type RouterOption func(*Router)

func NewRouter(opts ...RouterOption) IRouter {
	r := &Router{
		routes: make(map[string]*Route),
		port:   80,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func WithHealthCheckManager(healthCheckManager healthcheck.HealthCheckManagerInterface) RouterOption {
	return func(r *Router) {
		r.HealthCheckManager = healthCheckManager
	}
}

type RouteOption func(*Route)

func WithMiddleware(middleware types.MiddlewareFunc) RouteOption {
	return func(r *Route) {
		r.Middleware = middleware
	}
}

func WithListMiddleware(listMiddleware []*middleware.MiddlewareWithPriority) RouteOption {
	return func(r *Route) {
		r.ListMiddleware = listMiddleware
		r.Middleware = middleware.MiddlewareWithPriorityComposer(listMiddleware...)
	}
}

func (r *Route) SetListMiddleware(listMiddleware []*middleware.MiddlewareWithPriority) {
	r.ListMiddleware = listMiddleware
	r.Middleware = middleware.MiddlewareWithPriorityComposer(listMiddleware...)
}

func WithHealthCheckConfig(healthCheckConfig *healthcheck.HealthCheckConfig) RouteOption {
	return func(r *Route) {
		r.HealthCheckConfig = healthCheckConfig
	}
}

func NewRoute(uuid string, host string, path string, reverseProxy reverseproxy.ProxyHandler, proxyController *proxy.ProxyController, opts ...RouteOption) *Route {
	r := &Route{
		Uuid:            uuid,
		Host:            host,
		Path:            path,
		ReverseProxy:    reverseProxy,
		ProxyController: proxyController,
		ListMiddleware:  make([]*middleware.MiddlewareWithPriority, 0),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func WithPort(port int32) RouterOption {
	return func(r *Router) {
		r.port = port
	}
}

func (r *Router) normalizeHostURI(host string, path string) string {
	hostURI := strings.TrimSuffix(host+":"+fmt.Sprintf("%d", r.port)+"/"+path, "/")
	return strings.ReplaceAll(hostURI, "//", "/")
}

func (r *Router) Upsert(route *Route) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	route.NormalizeHostURI = r.normalizeHostURI(route.Host, route.Path)
	r.routes[route.NormalizeHostURI] = route
	return nil
}

func (r *Router) AddHealthCheck(id string, healthCheck healthcheck.HealthCheckInterface) {
	r.HealthCheckManager.AddHealthCheck(id, healthCheck)
}

func (r *Router) RemoveHealthCheck(id string) {
	r.HealthCheckManager.RemoveHealthCheck(id)
}

func (r *Router) GetRoute(host string, path string) (*Route, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	hostUri := r.normalizeHostURI(host, path)
	route, ok := r.routes[hostUri]
	if !ok {
		return nil, errors.New("route not found")
	}
	return route, nil
}

func (r *Router) Remove(host string, path string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	hostUri := r.normalizeHostURI(host, path)
	route, ok := r.routes[hostUri]
	if !ok {
		return nil
	}
	route.ProxyController.CloseAll()
	delete(r.routes, hostUri)
	return nil
}

func (r *Router) GetProxyController(host string, path string) (*proxy.ProxyController, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	hostUri := r.normalizeHostURI(host, path)
	route, ok := r.routes[hostUri]
	if !ok {
		return nil, errors.New("route not found")
	}
	return route.ProxyController, nil
}

func (r *Router) ProxyRouter(ctx *fasthttp.RequestCtx) {
	logger.Infof("Request Host: %s, Path: %s", ctx.Host(), ctx.Path())
	hostURI := strings.TrimSuffix(string(ctx.Host())+string(ctx.Path()), "/")
	r.mutex.RLock()
	route, ok := r.routes[hostURI]
	if !ok {
		r.mutex.RUnlock()
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBody([]byte("404 Not Found"))
		return
	}
	r.mutex.RUnlock()
	handler := route.ReverseProxy.ReverseProxyHandler
	if route.Middleware != nil {
		handler = route.Middleware(handler)
	}
	handler(ctx)
}

//type MapRoute struct {
//	proxyClient     map[string]proxy.IProxyClient
//
//	// thread	safe
//	mutex sync.Mutex
//}
//
//func NewRoute() *MapRoute {
//	return &MapRoute{
//		proxyClient: make(map[string]proxy.IProxyClient),
//	}
//}
//
//func (r *MapRoute) ProxyRouter(ctx *fasthttp.RequestCtx) {
//	dns := ctx.Host()
//	proxy, ok := r.proxyClient[string(dns)]
//	if !ok {
//		ctx.SetStatusCode(fasthttp.StatusNotFound)
//		ctx.SetBody([]byte("404 Not Found"))
//		return
//	}
//	proxy.ReverseProxyHandler(ctx)
//}
//
//func (r *MapRoute) Add(host string, proxy proxy.IProxyClient) error {
//	r.mutex.Lock()
//	defer r.mutex.Unlock()
//
//	r.proxyClient[host] = proxy
//	return nil
//}
//
//func (r *MapRoute) Remove(host string, proxy proxy.IProxyClient) error {
//	r.mutex.Lock()
//	defer r.mutex.Unlock()
//	r.proxyClient[host].CloseAll()
//	delete(r.proxyClient, host)
//	return nil
//}
//
