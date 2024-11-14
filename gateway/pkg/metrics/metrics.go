package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// upstream metrics
var (
	UpstreamSuccesses = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "upstream_successes_total",
			Help: "Total number of requests to upstreams",
		},
		[]string{"route_uuid", "route_url", "upstream_uuid", "upstream_url"},
	)
	UpstreamFailures = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "upstream_failures_total",
			Help: "Total number of failed requests to upstreams",
		},
		[]string{"route_uuid", "route_url", "upstream_uuid", "upstream_url"},
	)
	UpstreamDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "upstream_duration_seconds",
			Help:    "Duration of requests to upstreams",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"route_uuid", "route_url", "upstream_uuid", "upstream_url"},
	)
	Status429Responses = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "status_429_responses_total",
			Help: "Total number of status code 429 responses",
		},
		[]string{"route_uuid", "route_url", "upstream_uuid", "upstream_url"},
	)
	UpstreamTotalRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "upstream_total_requests",
			Help: "Total number of responses",
		},
		[]string{"route_uuid", "route_url", "upstream_uuid", "upstream_url", "code", "eth_method"},
	)
)

// route metrics
var (
	RouteRequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "route_requests_total",
			Help: "Number of HTTP requests",
		},
		[]string{"route_uuid", "method", "status", "host", "scheme"},
	)
	RouteRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "route_http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"route_uuid", "method", "status", "host", "scheme"},
	)
)

//func init() {
//	prometheus.MustRegister(UpstreamSuccesses, UpstreamFailures, UpstreamDuration, Status429Responses, UpstreamTotalRequests)
//	prometheus.MustRegister(RouteRequestCount, RouteRequestDuration)
//}
