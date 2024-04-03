package middleware

import (
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// NewLogger logs after a request
func NewLogger(logger *zap.Logger) MiddlewareFunc {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			var startTime = time.Now()
			next(ctx)
			logger.Info("access",
				zap.Int("code", ctx.Response.StatusCode()),
				zap.Duration("time", time.Since(startTime)),
				zap.ByteString("host", ctx.Host()),
				zap.ByteString("method", ctx.Method()),
			)
		}
	}
}
