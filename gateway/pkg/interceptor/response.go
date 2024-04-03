package interceptor

import (
	"errors"

	"github.com/valyala/fasthttp"
)

// var (
//
//	requestCount = prometheus.NewCounterVec(
//		prometheus.CounterOpts{
//			Name: "http_requests_total",
//			Help: "Number of HTTP requests",
//		},
//		[]string{"route_uuid", "method", "status", "host", "scheme"},
//	)
//	requestDuration = prometheus.NewHistogramVec(
//		prometheus.HistogramOpts{
//			Name:    "http_request_duration_seconds",
//			Help:    "Duration of HTTP requests",
//			Buckets: prometheus.DefBuckets,
//		},
//		[]string{"route_uuid", "method", "status", "host", "scheme"},
//	)
//
// )
//
// //func init() {
// //	prometheus.MustRegister(requestCount)
// //	prometheus.MustRegister(requestDuration)
// //}

type DefaultResponseInterceptor struct {
}

func (r *DefaultResponseInterceptor) Intercept(res *fasthttp.Response) error {
	if res.StatusCode() == fasthttp.StatusTooManyRequests {
		return errors.New("too many requests")
	}
	if res.StatusCode() >= fasthttp.StatusInternalServerError {
		return errors.New("invalid response")
	}

	return nil
}
