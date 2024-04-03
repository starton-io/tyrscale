package middleware

import "github.com/valyala/fasthttp"

func HeadersMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Request.Header.Set("Accept-Encoding", "gzip")
		next(ctx)
	}
}
