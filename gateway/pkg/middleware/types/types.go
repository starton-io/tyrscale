package types

import "github.com/valyala/fasthttp"

type MiddlewareFunc func(next fasthttp.RequestHandler) fasthttp.RequestHandler
