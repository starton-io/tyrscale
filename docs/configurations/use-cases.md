# Use Cases

## 1: Use tyrscale behind public node RPC

Each public node RPC has its own rate limit per minute. It is possible to use tyrscale for load balance request with the least loaded algorithm. We don't recommend in this case to use any health check since the number of request will reach the rate limit of the public node RPC. Meanwhile, each public node RPC does not have the same rate limit.
Starton recommend to use circuit breaker for limiting the number of request to each public node RPC when it reach the rate limit. With enterprise edition, it is possible to use rate limit recommendation : the recommender analyze the load balancer metrics from prometheus then take the decision to update the weight of the public node RPC.

oss example configuration:
```
POST /api/v1/routes
{
    "host": "starton-local.com",
    "Path": "/eth/strat3",
    "load_balancer_strategy": "failover-priority",
    "circuit_breaker": {
        "enabled": true,
        "interval": 0,
        "max_consecutive_failures": 3,
        "max_requests": 1,
        "timeout": 120000
	}
}
```
Use failover-priority strategy to make sure to send the request to the node RPC with highest weight. Circuit breaker is used to limit the number of request to each public that are already failed. The timeout will be the highest rate limit of the public node RPC in millisecond that you have import with /route/:route_id/upstreams endpoint API. The interval is 0 to indicate that we never reset the circuit breaker metrics. 

How to configure recommendation enterprise feature:
```
POST /api/v1/routes
{
    "host": "starton-local.com",
    "Path": "/eth/strat3",
    "load_balancer_strategy": "failover-priority",
    "circuit_breaker": {
        "enabled": true,
        "interval": 0,
        "max_consecutive_failures": 3,
        "max_requests": 1,
        "timeout": 120000
	}
}
POST /api/v1/recommendations
{
  "network_name": "eth",
  "route_uuid": "route_id",
  "schedule": "* * * * *",
  "strategy": "STRATEGY_RATE_LIMIT"
}
```
This recommendations allow to determined the rate limit of the public node RPC through the metrics of the load balancer. The rate limit will be the highest rate limit of the public node RPC in millisecond that you have import with /route/:route_id/upstreams endpoint API.

## 2: Use tyrscale behind private Node RPC

The best algorithm will be weighted round robin with the same weight for each private node RPC. The reason is that the block height is probably the same for each node RPC. The chance that the node RPC with the highest block would be down is very low. Even the call fails, the request (max 3 retries) will be retried with the other node RPC if the weight is the same. With Enterprise edition, it is possible to use area of block number strategy to determine the weight of the private node RPC : it calculate the area under the curve of block number metrics.

```
POST /api/v1/routes
{
    "host": "starton-local.com",
    "Path": "/eth/strat3",
    "load_balancer_strategy": "weight-round-robin",
    "circuit_breaker": {
        "enabled": true,
        "interval": 0,
        "max_consecutive_failures": 3,
        "max_requests": 1,
        "timeout": 120000
	}
}
POST /api/v1/recommendations (enterprise feature)
{
  "network_name": "eth",
  "route_uuid": "route_id",
  "schedule": "* * * * *",
  "strategy": "STRATEGY_AREA_BLOCK_NUMBER"
}
```


## 3: Enable smart caching for optimize all Node RPC calls (In development)

Smart caching allows developers to cache the response of the Node RPC calls. This can be very benefit for optimize the external call. Be careful we recommend to adjust the expire cache to lower than the block time of the network.
