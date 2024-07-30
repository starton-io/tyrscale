package interceptor

import (
	"sort"

	"github.com/starton-io/tyrscale/gateway/pkg/interceptor/types"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

//type InterceptorRequestChain struct {
//	interceptors []types.IRequestInterceptor
//}
//
//func NewInterceptorRequestChain() InterceptorRequestChain {
//	return InterceptorRequestChain{}
//}
//
//func (chain *InterceptorRequestChain) Intercept(req *fasthttp.Request) error {
//	for _, interceptor := range chain.interceptors {
//		if err := interceptor.Intercept(req); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (chain *InterceptorRequestChain) Add(interceptor types.IRequestInterceptor) {
//	chain.interceptors = append(chain.interceptors, interceptor)
//}

// InterceptorResponseChain is a chain of response interceptors
type InterceptorResponseChain struct {
	interceptors []*orderedInterceptor
}

type orderedInterceptor struct {
	interceptor types.IResponseInterceptor
	name        string
	order       int
}

func NewInterceptorResponseChain() InterceptorResponseChain {
	return InterceptorResponseChain{}
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

func (chain *InterceptorResponseChain) AddFirst(interceptor types.IResponseInterceptor, name string) {
	chain.interceptors = append([]*orderedInterceptor{{interceptor, name, 0}}, chain.interceptors...)
}

func (chain *InterceptorResponseChain) Has(name string) bool {
	for _, oi := range chain.interceptors {
		if oi.name == name {
			return true
		}
	}
	return false
}

func (chain *InterceptorResponseChain) AddLast(interceptor types.IResponseInterceptor, name string) {
	chain.interceptors = append(chain.interceptors, &orderedInterceptor{interceptor, name, len(chain.interceptors)})
}

func (chain *InterceptorResponseChain) AddOrdered(interceptor types.IResponseInterceptor, name string, order int) {
	chain.interceptors = append(chain.interceptors, &orderedInterceptor{interceptor, name, order})
	sort.SliceStable(chain.interceptors, func(i, j int) bool {
		return chain.interceptors[i].order < chain.interceptors[j].order
	})
}

func (chain *InterceptorResponseChain) KeepFirstAndReplaceOthers(interceptor InterceptorResponseChain) {
	if len(chain.interceptors) > 0 {
		firstInterceptor := chain.interceptors[0]
		chain.interceptors = append([]*orderedInterceptor{firstInterceptor}, interceptor.interceptors...)
	} else {
		chain.interceptors = interceptor.interceptors
	}

}

func (chain *InterceptorResponseChain) Remove(name string) {
	for i, oi := range chain.interceptors {
		if oi.name == name {
			logger.Debugf("removing interceptor %s", name)
			chain.interceptors = append(chain.interceptors[:i], chain.interceptors[i+1:]...)
			return
		}
	}
}
