package middleware

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starton-io/tyrscale/gateway/pkg/metrics"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type Prometheus struct {
	RouteUuid string
}

func NewPrometheus(prometheus *Prometheus) types.MiddlewareFunc {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			start := time.Now()
			next(ctx)
			duration := time.Since(start)

			routeUuid := string(prometheus.RouteUuid)
			method := string(ctx.Method())
			status := ctx.Response.StatusCode()
			host := string(ctx.Host())
			scheme := string(ctx.URI().Scheme())
			statusStr := strconv.Itoa(status) // Convert status to string
			//requestCount.WithLabelValues(routeUuid, method, statusStr, host, scheme).Inc()
			metrics.RouteRequestCount.WithLabelValues(routeUuid, method, statusStr, host, scheme).Inc()
			//requestDuration.WithLabelValues(routeUuid, method, statusStr, host, scheme).Observe(duration.Seconds())
			metrics.RouteRequestDuration.WithLabelValues(routeUuid, method, statusStr, host, scheme).Observe(duration.Seconds())
		}
	}
}

// PrometheusHandler returns an HTTP handler for Prometheus metrics
func PrometheusHandler() fasthttp.RequestHandler {
	return fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
}
