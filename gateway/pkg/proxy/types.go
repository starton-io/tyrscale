package proxy

import (
	"github.com/sony/gobreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
	"github.com/valyala/fasthttp"
)

// ReverseProxyHandler is the interface that wraps the Handle method.
type ReverseProxyHandler interface {
	// Handle is the method that handles the request.
	Handle(ctx *fasthttp.RequestCtx)
}

type ProxyManager interface {
	AddUpstream(upstream *upstream.UpstreamPublishUpsertModel) error
	RemoveUpstream(uuid string) error
	ExistUpstream(upstreamUuid string) bool
	CloseAll()
	Close(uuid string)
}

type ProxyContext interface {
	GetBalancer() balancer.IBalancer
	GetClient(uuid string) (*UpstreamClient, bool)
	Do(client *fasthttp.HostClient, req *fasthttp.Request, res *fasthttp.Response) error
	GetCircuitBreaker(uuid string) *gobreaker.CircuitBreaker
	CloseAll()
	Close(uuid string)
}

type ConnectionManager interface {
	CloseAll()
	Close(uuid string)
}
