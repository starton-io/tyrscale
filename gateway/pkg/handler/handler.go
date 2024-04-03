package handler

import (
	"fmt"

	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/valyala/fasthttp"
)

type ProxyHandler interface {
	Handle(ctx *fasthttp.RequestCtx)
}

func NewFactory(proxyController *proxy.ProxyController) (ProxyHandler, error) {
	switch proxyController.Balancer.GetStrategy() {
	case balancer.BalancerPriority:
		return NewFailoverHandler(proxyController), nil
	case balancer.BalancerLeastLoad:
		return NewDefaultHandler(proxyController), nil
	case balancer.BalancerWeightRoundRobin:
		return NewDefaultHandler(proxyController), nil
	default:
		return nil, fmt.Errorf("invalid balancer strategy: %v", proxyController.Balancer.GetStrategy())
	}
}
