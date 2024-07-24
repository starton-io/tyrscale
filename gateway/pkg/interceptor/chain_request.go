package interceptor

import (
	"sort"

	"github.com/starton-io/tyrscale/gateway/pkg/interceptor/types"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

// InterceptorResponseChain is a chain of response interceptors
type InterceptorRequestChain struct {
	interceptors []*orderedRequestInterceptor
}

type orderedRequestInterceptor struct {
	interceptor types.IRequestInterceptor
	order       int
}

func NewInterceptorRequestChain() InterceptorRequestChain {
	return InterceptorRequestChain{
		interceptors: make([]*orderedRequestInterceptor, 0),
	}
}

func (chain *InterceptorRequestChain) Intercept(req *fasthttp.Request) error {
	logger.Debugf("intercepting request with %d interceptors", len(chain.interceptors))
	for _, oi := range chain.interceptors {
		if err := oi.interceptor.Intercept(req); err != nil {
			return err
		}
	}
	return nil
}

func (chain *InterceptorRequestChain) AddFirst(interceptor types.IRequestInterceptor) {
	chain.interceptors = append([]*orderedRequestInterceptor{{interceptor, 0}}, chain.interceptors...)
}

func (chain *InterceptorRequestChain) AddLast(interceptor types.IRequestInterceptor) {
	chain.interceptors = append(chain.interceptors, &orderedRequestInterceptor{interceptor, len(chain.interceptors)})
}

func (chain *InterceptorRequestChain) AddOrdered(interceptor types.IRequestInterceptor, order int) {
	chain.interceptors = append(chain.interceptors, &orderedRequestInterceptor{interceptor, order})
	sort.SliceStable(chain.interceptors, func(i, j int) bool {
		return chain.interceptors[i].order < chain.interceptors[j].order
	})
}

func (chain *InterceptorRequestChain) Remove(interceptor types.IRequestInterceptor) {
	for i, oi := range chain.interceptors {
		if oi.interceptor == interceptor {
			chain.interceptors = append(chain.interceptors[:i], chain.interceptors[i+1:]...)
			return
		}
	}
}
