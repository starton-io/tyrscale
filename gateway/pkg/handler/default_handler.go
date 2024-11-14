package handler

import (
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/normalizer"
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
	var lastErr error
	var requestContext *RequestContext

	routeUrl := string(req.URI().Host()) + string(req.URI().Path())
	logger.Debugf("routeUrl: %s", routeUrl)
	normalizedRequest := normalizer.NewNormalizedRequest(req.Body())
	method, err := normalizedRequest.Method()
	if err != nil {
		logger.Errorf("Failed to get method: %v", err)
		setErrorResponse(res, fasthttp.StatusBadRequest, "Invalid JSON-RPC request")
		return
	}

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
		if !ok || !upstreamClient.Healthy || upstreamClient.IgnoreMethod(method, res) {
			logger.Infof("Skipping upstream %s", upstreamUuid)
			continue
		}

		listLabelsValues := []string{upstreamClient.Client.Addr, upstreamUuid[0], routeUrl, h.proxyController.GetLabelValue("route_uuid")}
		requestContext = &RequestContext{
			req:              req,
			res:              res,
			upstreamClient:   upstreamClient,
			upstreamUuid:     upstreamUuid[0],
			ethMethod:        method,
			listLabelsValues: listLabelsValues,
			startTime:        time.Now(),
		}
		logger.Debugf("circuitBreaker: %v", h.proxyController.CircuitBreaker)
		if err := processRequest(requestContext, h.proxyController.CircuitBreaker); err != nil {
			lastErr = err
			continue
		}
		return
	}
	logger.Error("All upstream nodes are unhealthy/dead")
	if lastErr != nil {
		logger.Errorf("last listUpstream error: %v, body: %s", lastErr, res.Body())
	}
}
