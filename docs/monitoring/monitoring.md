# Monitoring Metrics

The Gateway service exposes Prometheus metrics for upstream and route performance.

## Upstream Metrics

### UpstreamSuccesses
- **Name**: `upstream_successes_total`
- **Help**: Total number of requests to upstreams
- **Labels**: `upstream_uuid`, `route_uuid`

### UpstreamFailures
- **Name**: `upstream_failures_total`
- **Help**: Total number of failed requests to upstreams
- **Labels**: `upstream_uuid`, `route_uuid`

### UpstreamDuration
- **Name**: `upstream_duration_seconds`
- **Help**: Duration of requests to upstreams
- **Labels**: `upstream_uuid`, `route_uuid`
- **Buckets**: Default Prometheus buckets

### Status429Responses
- **Name**: `status_429_responses_total`
- **Help**: Total number of status code 429 responses
- **Labels**: `upstream_uuid`, `route_uuid`

### UpstreamTotalRequests
- **Name**: `upstream_total_requests`
- **Help**: Total number of responses
- **Labels**: `upstream_uuid`, `route_uuid`

## Route Metrics

### RouteRequestCount
- **Name**: `route_requests_total`
- **Help**: Number of HTTP requests
- **Labels**: `route_uuid`, `method`, `status`, `host`, `scheme`

### RouteRequestDuration
- **Name**: `route_http_request_duration_seconds`
- **Help**: Duration of HTTP requests
- **Labels**: `route_uuid`, `method`, `status`, `host`, `scheme`
- **Buckets**: Default Prometheus buckets
