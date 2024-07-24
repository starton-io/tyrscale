package types

import "github.com/valyala/fasthttp"

type IRequestInterceptor interface {
	Intercept(*fasthttp.Request) error
}

type IResponseInterceptor interface {
	Intercept(*fasthttp.Response) error
}
