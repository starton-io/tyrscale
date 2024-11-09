package handler

import (
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/normalizer"
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

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
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
		setErrorResponse(res, fasthttp.StatusNotFound, "Balancer error: "+err.Error())
		return
	}

	if len(listUpstream) == 0 {
		logger.Info("No upstream servers found")
		setErrorResponse(res, fasthttp.StatusNotFound, "No upstream servers available")
		return
	}

	normalizedRequest := normalizer.NewNormalizedRequest(req.Body())
	method, err := normalizedRequest.Method()
	if err != nil {
		logger.Errorf("Failed to get method: %v", err)
		setErrorResponse(res, fasthttp.StatusBadRequest, "Invalid JSON-RPC request")
		return
	}

	for _, upstreamUuid := range listUpstream {
		upstreamClient, ok := h.proxyController.ClientManager.GetClient(upstreamUuid)
		if !ok || !upstreamClient.Healthy || upstreamClient.IgnoreMethod(method, res) {
			logger.Infof("Skipping upstream %s", upstreamUuid)
			continue
		}

		routeUrl := string(ctx.URI().Host()) + string(ctx.URI().Path())
		listLabelsValues := []string{upstreamClient.Client.Addr, upstreamUuid, routeUrl, h.proxyController.GetLabelValue("route_uuid")}
		requestContext := &RequestContext{
			req:              req,
			res:              res,
			upstreamClient:   upstreamClient,
			upstreamUuid:     upstreamUuid,
			method:           method,
			listLabelsValues: listLabelsValues,
			startTime:        time.Now(),
		}
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
	res.SetStatusCode(res.StatusCode())
	res.SetBody(res.Body())
}
