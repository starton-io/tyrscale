package interceptor

import (
	"sort"

	"github.com/starton-io/tyrscale/gateway/pkg/interceptor/types"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

// InterceptorResponseChain is a chain of response interceptors
type InterceptorResponseChain struct {
	interceptors []*orderedInterceptor
}

type orderedInterceptor struct {
	interceptor types.IResponseInterceptor
	order       int
}

func NewInterceptorResponseChain() InterceptorResponseChain {
	return InterceptorResponseChain{
		interceptors: make([]*orderedInterceptor, 0),
	}
}

func (chain *InterceptorResponseChain) Intercept(res *fasthttp.Response) error {
	logger.Debugf("intercepting response with %d interceptors", len(chain.interceptors))
	for _, oi := range chain.interceptors {
		if err := oi.interceptor.Intercept(res); err != nil {
			return err
		}
	}
	return nil
}

func (chain *InterceptorResponseChain) AddFirst(interceptor types.IResponseInterceptor) {
	chain.interceptors = append([]*orderedInterceptor{{interceptor, 0}}, chain.interceptors...)
}

func (chain *InterceptorResponseChain) AddLast(interceptor types.IResponseInterceptor) {
	chain.interceptors = append(chain.interceptors, &orderedInterceptor{interceptor, len(chain.interceptors)})
}

func (chain *InterceptorResponseChain) AddOrdered(interceptor types.IResponseInterceptor, order int) {
	chain.interceptors = append(chain.interceptors, &orderedInterceptor{interceptor, order})
	sort.SliceStable(chain.interceptors, func(i, j int) bool {
		return chain.interceptors[i].order < chain.interceptors[j].order
	})
}

func (chain *InterceptorResponseChain) Remove(interceptor types.IResponseInterceptor) {
	for i, oi := range chain.interceptors {
		if oi.interceptor == interceptor {
			chain.interceptors = append(chain.interceptors[:i], chain.interceptors[i+1:]...)
			return
		}
	}
}
