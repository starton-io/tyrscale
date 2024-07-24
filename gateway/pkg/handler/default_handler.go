package handler

import (
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/metrics"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

type DefaultHandler struct {
	proxyController *proxy.ProxyController
	maxRetries      int
}

func NewDefaultHandler(proxyController *proxy.ProxyController) *DefaultHandler {
	return &DefaultHandler{
		proxyController: proxyController,
		maxRetries:      3,
	}
}

func (h *DefaultHandler) Handle(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response

	for retries := 0; retries < h.maxRetries; retries++ {
		upstreamUuid, err := h.proxyController.Balancer.Balance()
		if err != nil {
			setErrorResponse(res, fasthttp.StatusNotFound, err.Error())
			return
		}

		if len(upstreamUuid) == 0 {
			setErrorResponse(res, fasthttp.StatusNotFound, "upstream not found")
			logger.Error("upstream not found...")
			continue
		}

		upstreamClient, ok := h.proxyController.ClientManager.GetClient(upstreamUuid[0])
		if !ok {
			setErrorResponse(res, fasthttp.StatusNotFound, "upstream not found")
			logger.Error("upstream client not found...")
			continue
		}
		if !upstreamClient.Healthy {
			setErrorResponse(res, fasthttp.StatusServiceUnavailable, "upstream not healthy")
			logger.Error("upstream not healthy")
			continue
		} else {
			logger.Debugf("upstream UUID: %s, healthy: %t", upstreamUuid[0], upstreamClient.Healthy)
		}
		err = upstreamClient.RequestInterceptor.Intercept(req)
		if err != nil {
			setErrorResponse(res, fasthttp.StatusInternalServerError, err.Error())
			return
		}
		start := time.Now()
		metrics.UpstreamTotalRequests.WithLabelValues(upstreamUuid[0], h.proxyController.GetLabelValue("route_uuid")).Inc()

		if h.proxyController.CircuitBreaker != nil {
			cb := h.proxyController.CircuitBreaker.Get(upstreamUuid[0])
			if cb != nil {
				_, err = cb.Execute(func() (interface{}, error) {
					err := upstreamClient.Client.Do(req, res)
					if err != nil {
						return nil, err
					}
					return nil, upstreamClient.ResponseInterceptor.Intercept(res)
				})
				if err == nil {
					metrics.UpstreamSuccesses.WithLabelValues(upstreamUuid[0], h.proxyController.GetLabelValue("route_uuid")).Inc()
					metrics.UpstreamDuration.WithLabelValues(upstreamUuid[0], h.proxyController.GetLabelValue("route_uuid")).Observe(time.Since(start).Seconds())
					return // Successful execution
				}
				continue
			}
		}
		if err := upstreamClient.Client.Do(req, res); err != nil {
			handleClientError(res, err)
			metrics.UpstreamFailures.WithLabelValues(upstreamUuid[0], h.proxyController.GetLabelValue("route_uuid")).Inc()
			continue
		}
		err = upstreamClient.ResponseInterceptor.Intercept(res)
		if err != nil {
			setErrorResponse(res, fasthttp.StatusInternalServerError, err.Error())
			if res.StatusCode() == fasthttp.StatusTooManyRequests {
				metrics.Status429Responses.WithLabelValues(upstreamUuid[0], h.proxyController.GetLabelValue("route_uuid")).Inc()
			}
			metrics.UpstreamFailures.WithLabelValues(upstreamUuid[0], h.proxyController.GetLabelValue("route_uuid")).Inc()
			continue
		}
		metrics.UpstreamSuccesses.WithLabelValues(upstreamUuid[0], h.proxyController.GetLabelValue("route_uuid")).Inc()
		metrics.UpstreamDuration.WithLabelValues(upstreamUuid[0], h.proxyController.GetLabelValue("route_uuid")).Observe(time.Since(start).Seconds())
		return
	}
	logger.Error("all upstream nodes are unhealthy/dead after %d retries", h.maxRetries)
	setErrorResponse(res, fasthttp.StatusServiceUnavailable, "upstream not healthy")
}
