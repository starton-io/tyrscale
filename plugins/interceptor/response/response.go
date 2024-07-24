package main

import (
	"log"

	"github.com/starton-io/tyrscale/gateway/pkg/interceptor/types"
	"github.com/valyala/fasthttp"
)

type MyResponseInterceptor struct {
}

func (m *MyResponseInterceptor) Intercept(response *fasthttp.Response) error {
	response.Header.Set("X-Intercepted", "true")
	log.Println("MyResponseInterceptor")
	return nil
}

type MyResponseInterceptorRegister struct{}

func (p *MyResponseInterceptorRegister) RegisterResponseInterceptors(registerFunc func(name string, interceptor types.IResponseInterceptor)) {

	registerFunc("MyResponseInterceptor", &MyResponseInterceptor{})
}

// Ensure MyPlugin implements the Registerer interface
var ResponseInterceptor MyResponseInterceptorRegister

func main() {}
