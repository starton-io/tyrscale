package middleware

import (
	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"
	"github.com/valyala/fasthttp"
)

func NewCors() types.MiddlewareFunc {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
			ctx.Response.Header.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")
			next(ctx)
		}
	}
}
