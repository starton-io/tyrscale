package main

import (
	"log"
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"

	"github.com/valyala/fasthttp"
)

// LoggingMiddleware is an example middleware that logs requests
func LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Printf("Request: %s, %s", ctx.Method(), ctx.RequestURI())
		start := time.Now()
		next(ctx)
		log.Printf("Response: %s ,%d , duration: %s", ctx.Method(), ctx.Response.StatusCode(), time.Since(start))
	}
}

type LoggingMiddlewareRegister struct{}

func (p *LoggingMiddlewareRegister) RegisterMiddleware(registerFunc func(name string, middleware types.MiddlewareFunc), payload []byte) error {
	registerFunc("LoggingMiddleware", LoggingMiddleware)
	return nil
}

func (p *LoggingMiddlewareRegister) Validate(configPayload []byte) error {
	return nil
}

// Exported symbol
var Middleware LoggingMiddlewareRegister

func main() {}
