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
	name        string
	order       int
}

func NewInterceptorRequestChain() InterceptorRequestChain {
	return InterceptorRequestChain{}
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

func (chain *InterceptorRequestChain) AddFirst(interceptor types.IRequestInterceptor, name string) {
	chain.interceptors = append([]*orderedRequestInterceptor{{interceptor, name, 0}}, chain.interceptors...)
}

func (chain *InterceptorRequestChain) AddLast(interceptor types.IRequestInterceptor, name string) {
	chain.interceptors = append(chain.interceptors, &orderedRequestInterceptor{interceptor, name, len(chain.interceptors)})
}

func (chain *InterceptorRequestChain) Has(name string) bool {
	for _, oi := range chain.interceptors {
		if oi.name == name {
			return true
		}
	}
	return false
}

func (chain *InterceptorRequestChain) AddOrdered(interceptor types.IRequestInterceptor, name string, order int) {
	chain.interceptors = append(chain.interceptors, &orderedRequestInterceptor{interceptor, name, order})
	sort.SliceStable(chain.interceptors, func(i, j int) bool {
		return chain.interceptors[i].order < chain.interceptors[j].order
	})
}

func (chain *InterceptorRequestChain) KeepFirstAndReplaceOthers(newChain InterceptorRequestChain) {
	if len(chain.interceptors) > 0 {
		firstInterceptor := chain.interceptors[0]
		chain.interceptors = append([]*orderedRequestInterceptor{firstInterceptor}, newChain.interceptors...)
	} else {
		chain.interceptors = newChain.interceptors
	}
}

func (chain *InterceptorRequestChain) Remove(name string) {
	for i, oi := range chain.interceptors {
		if oi.name == name {
			chain.interceptors = append(chain.interceptors[:i], chain.interceptors[i+1:]...)
			return
		}
	}
}
