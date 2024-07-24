package main

import (
	"log"

	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"

	"github.com/valyala/fasthttp"
)

// LoggingMiddleware is an example middleware that logs requests
func LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Printf("Request: %s %s", ctx.Method(), ctx.RequestURI())
		next(ctx)
		log.Printf("Response: %s %d", ctx.Method(), ctx.Response.StatusCode())
	}
}

type LoggingMiddlewareRegister struct{}

func (p *LoggingMiddlewareRegister) RegisterMiddleware(registerFunc func(name string, interceptor types.MiddlewareFunc)) {
	registerFunc("LoggingMiddleware", LoggingMiddleware)
}

// Exported symbol
var Middleware LoggingMiddlewareRegister

func main() {}
