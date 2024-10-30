package handler

import (
	"errors"
	"time"

	"github.com/sony/gobreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/metrics"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

type RequestContext struct {
	req              *fasthttp.Request
	res              *fasthttp.Response
	upstreamClient   *proxy.UpstreamClient
	upstreamUuid     string
	method           string
	listLabelsValues []string
	startTime        time.Time
}

func processRequest(ctx *RequestContext, circuitBreaker circuitbreaker.ProxyCircuitBreaker) error {
	err := ctx.upstreamClient.RequestInterceptor.Intercept(ctx.req)
	if err != nil {
		logger.Errorf("Request interception failed for %s: %v", ctx.upstreamUuid, err)
		setErrorResponse(ctx.res, fasthttp.StatusInternalServerError, "Interception error: "+err.Error())
		return err
	}

	if circuitBreaker != nil {
		cb := circuitBreaker.GetTwoStep(ctx.upstreamUuid)
		if cb != nil {
			return executeWithCircuitBreaker(ctx, cb)
		}
	}

	return executeRequest(ctx)
}

func executeWithCircuitBreaker(ctx *RequestContext, cb *gobreaker.TwoStepCircuitBreaker) error {
	done, err := cb.Allow()
	if err != nil {
		logger.Errorf("Circuit breaker is still open for %s", ctx.upstreamUuid)
		return err
	}

	if err := ctx.upstreamClient.Client.Do(ctx.req, ctx.res); err != nil {
		done(false)
		logger.Errorf("Request failed for %s: %v", ctx.upstreamUuid, err)
		return err
	}

	err = ctx.upstreamClient.ResponseInterceptor.Intercept(ctx.res)
	if err != nil {
		done(false)
		logger.Errorf("Response interception failed for %s: %v", ctx.upstreamUuid, err)
		metrics.UpstreamFailures.WithLabelValues(ctx.listLabelsValues...).Inc()
		setErrorResponse(ctx.res, fasthttp.StatusInternalServerError, "Interception error: "+err.Error())
		return err
	}

	logger.Debugf("context Res: %v", ctx.res)

	success := handleCircuitBreakerError(ctx)
	if success {
		done(true)
		metrics.UpstreamDuration.WithLabelValues(ctx.listLabelsValues...).Observe(time.Since(ctx.startTime).Seconds())
		metrics.UpstreamSuccesses.WithLabelValues(ctx.listLabelsValues...).Inc()
		return nil
	}

	done(false)
	metrics.UpstreamFailures.WithLabelValues(ctx.listLabelsValues...).Inc()
	return errors.New("circuit breaker error")
}

func executeRequest(ctx *RequestContext) error {
	if err := ctx.upstreamClient.Client.Do(ctx.req, ctx.res); err != nil {
		handleClientError(ctx.res, err)
		metrics.UpstreamFailures.WithLabelValues(ctx.listLabelsValues...).Inc()
		return err
	}

	err := ctx.upstreamClient.ResponseInterceptor.Intercept(ctx.res)
	if err != nil {
		logger.Errorf("Response interception failed for %s: %v", ctx.upstreamUuid, err)
		if ctx.res.StatusCode() == fasthttp.StatusTooManyRequests {
			metrics.Status429Responses.WithLabelValues(ctx.listLabelsValues...).Inc()
		}
		metrics.UpstreamFailures.WithLabelValues(ctx.listLabelsValues...).Inc()
		setErrorResponse(ctx.res, fasthttp.StatusInternalServerError, "Interception error: "+err.Error())
		return err
	}

	metrics.UpstreamDuration.WithLabelValues(ctx.listLabelsValues...).Observe(time.Since(ctx.startTime).Seconds())
	metrics.UpstreamSuccesses.WithLabelValues(ctx.listLabelsValues...).Inc()
	return nil
}
