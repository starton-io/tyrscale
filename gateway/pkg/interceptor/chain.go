package interceptor

import (
	"github.com/valyala/fasthttp"
)

type InterceptorRequestChain struct {
	interceptors []IRequestInterceptor
}

func NewInterceptorRequestChain() *InterceptorRequestChain {
	return &InterceptorRequestChain{}
}

func (chain *InterceptorRequestChain) Intercept(req *fasthttp.Request) error {
	for _, interceptor := range chain.interceptors {
		if err := interceptor.Intercept(req); err != nil {
			return err
		}
	}
	return nil
}

func (chain *InterceptorRequestChain) Add(interceptor IRequestInterceptor) {
	chain.interceptors = append(chain.interceptors, interceptor)
}

// InterceptorResponseChain is a chain of response interceptors
type InterceptorResponseChain struct {
	interceptors []IResponseInterceptor
}

func NewInterceptorResponseChain() *InterceptorResponseChain {
	return &InterceptorResponseChain{}
}

func (chain *InterceptorResponseChain) Intercept(res *fasthttp.Response) error {
	for _, interceptor := range chain.interceptors {
		if err := interceptor.Intercept(res); err != nil {
			return err
		}
	}
	return nil
}

func (chain *InterceptorResponseChain) Add(interceptor IResponseInterceptor) {
	chain.interceptors = append(chain.interceptors, interceptor)
}
