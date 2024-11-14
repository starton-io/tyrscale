package handler

import (
	"errors"

	"github.com/starton-io/tyrscale/gateway/pkg/metrics"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

func setErrorResponse(res *fasthttp.Response, statusCode int, message string) {
	res.SetStatusCode(statusCode)
	res.SetBody([]byte(message))
}

func handleClientError(res *fasthttp.Response, err error) {
	if errors.Is(err, fasthttp.ErrTimeout) {
		setErrorResponse(res, fasthttp.StatusRequestTimeout, err.Error())
	} else {
		setErrorResponse(res, fasthttp.StatusInternalServerError, err.Error())
	}
}

func handleCircuitBreakerError(ctx *RequestContext) bool {
	statusCodeActions := map[int]func() bool{
		fasthttp.StatusTooManyRequests: func() bool {
			metrics.Status429Responses.WithLabelValues(ctx.listLabelsValues...).Inc()
			return true
		},
		fasthttp.StatusMethodNotAllowed: func() bool {
			logger.Debugf("Status 405, ignoring method %s", ctx.ethMethod)
			ctx.upstreamClient.AddIgnoreMethod(ctx.ethMethod)
			return true
		},
		fasthttp.StatusInternalServerError: func() bool {
			return false
		},
		fasthttp.StatusBadGateway: func() bool {
			return false
		},
		fasthttp.StatusServiceUnavailable: func() bool {
			return false
		},
		fasthttp.StatusGatewayTimeout: func() bool {
			return false
		},
	}

	statusCode := ctx.res.StatusCode()
	logger.Debugf("Handling status code: %d", statusCode)

	if action, exists := statusCodeActions[statusCode]; exists {
		return action()
	}

	logger.Debugf("No action found for status code: %d", statusCode)
	return false
}
