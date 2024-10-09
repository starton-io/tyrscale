package handler

import (
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/metrics"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

type FailoverHandler struct {
	proxyController *proxy.ProxyController
}

func NewFailoverHandler(proxyController *proxy.ProxyController) ProxyHandler {
	return &FailoverHandler{
		proxyController: proxyController,
	}
}

func (h *FailoverHandler) CloseConnections() {
}

func (h *FailoverHandler) Handle(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response
	var lastErr error

	listUpstream, err := h.proxyController.Balancer.Balance()
	if err != nil {
		logger.Errorf("Failed to balance: %v", err)
		res.SetStatusCode(fasthttp.StatusNotFound)
		res.SetBody([]byte("Balancer error: " + err.Error()))
		return
	}

	if len(listUpstream) == 0 {
		logger.Info("No upstream servers found")
		res.SetStatusCode(fasthttp.StatusNotFound)
		res.SetBody([]byte("No upstream servers available"))
		return
	}

	for _, upstreamUuid := range listUpstream {
		upstreamClient, ok := h.proxyController.ClientManager.GetClient(upstreamUuid)
		if !ok {
			logger.Infof("Upstream %s is not found", upstreamUuid)
			continue
		}
		if !upstreamClient.Healthy {
			logger.Infof("Upstream %s is not healthy", upstreamUuid)
			continue
		}

		err := upstreamClient.RequestInterceptor.Intercept(req)
		if err != nil {
			logger.Errorf("Request interception failed for %s: %v", upstreamUuid, err)
			setErrorResponse(res, fasthttp.StatusInternalServerError, "Interception error: "+err.Error())
			return
		}
		routeUrl := string(ctx.URI().Host()) + string(ctx.URI().Path())
		listLabelsValues := []string{upstreamClient.Client.Addr, upstreamUuid, routeUrl, h.proxyController.GetLabelValue("route_uuid")}
		start := time.Now()
		metrics.UpstreamTotalRequests.WithLabelValues(listLabelsValues...).Inc()
		if h.proxyController.CircuitBreaker != nil {
			cb := h.proxyController.CircuitBreaker.Get(upstreamUuid)
			if cb != nil {
				_, err = cb.Execute(func() (interface{}, error) {
					if err := upstreamClient.Client.Do(req, res); err != nil {
						lastErr = err
						return nil, err
					}
					return nil, upstreamClient.ResponseInterceptor.Intercept(res)
				})
				if err == nil {
					metrics.UpstreamDuration.WithLabelValues(listLabelsValues...).Observe(time.Since(start).Seconds())
					metrics.UpstreamSuccesses.WithLabelValues(listLabelsValues...).Inc()
					return // Successful execution
				}
				if res.StatusCode() == fasthttp.StatusTooManyRequests {
					metrics.Status429Responses.WithLabelValues(listLabelsValues...).Inc()
				}
				logger.Errorf("Circuit breaker execution failed for %s: %v", upstreamUuid, err)
				metrics.UpstreamFailures.WithLabelValues(listLabelsValues...).Inc()
				continue
			}
		} else {
			if err := upstreamClient.Client.Do(req, res); err != nil {
				handleClientError(res, err)
				metrics.UpstreamFailures.WithLabelValues(listLabelsValues...).Inc()
				continue
			}
			err = upstreamClient.ResponseInterceptor.Intercept(res)
			if err != nil {
				logger.Errorf("Response interception failed for %s: %v", upstreamUuid, err)
				if res.StatusCode() == fasthttp.StatusTooManyRequests {
					metrics.Status429Responses.WithLabelValues(listLabelsValues...).Inc()
				}
				metrics.UpstreamFailures.WithLabelValues(listLabelsValues...).Inc()
				setErrorResponse(res, fasthttp.StatusInternalServerError, "Interception error: "+err.Error())
				continue
			}
			metrics.UpstreamDuration.WithLabelValues(listLabelsValues...).Observe(time.Since(start).Seconds())
			metrics.UpstreamSuccesses.WithLabelValues(listLabelsValues...).Inc()
			return
		}
	}
	logger.Error("All upstream nodes are unhealthy/dead")
	if lastErr != nil {
		logger.Errorf("last listUpstream error: %v, body: %s", lastErr, res.Body())
	}
	res.SetStatusCode(res.StatusCode())
	res.SetBody(res.Body())
}
