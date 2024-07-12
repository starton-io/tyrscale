![Starton Banner](https://github.com/starton-io/.github/blob/master/github-banner.jpg?raw=true)


# TyrScale

TyrScale is a blockchain load balancer that can be used to balance the load between multiple blockchain rpc nodes. For now, it only supports EVM based chains. 


## List of services

### 1. Manager API (OSS)

The Manager API manages TyrScale load balancer, handling the addition, removal, and updating of RPC nodes and networks. It also configures the service monitor by specifying which collectors should be used for particular chains. Finally, it also manages the configuration of the gateway by pushing to the message queue (redis stream: route and upstream topics).

[Manager Documentations](./docs/architectures/manager.md)


### 2. Prometheus evm exporter (enterprise)

The Prometheus evm exporter exports EVM metrics to Prometheus. This service consumes messages from the message queue (Redis stream) and dynamically updates the evm exporter configuration. For scalability, it utilizes go-dispatch (with consistent hashing) to distribute the monitor load between evm exporter nodes.

### 3. Recommender (enterprise)

The recommender is responsible for recommending the best RPC node to use based on the current load and the configuration of the load balancer. It is responsible for adding, removing, and updating the rpc nodes scoring. It is also responsible for managing the weight of each RPC node by pushing to the rest api of the manager service.

### 4. Gateway (OSS)

The Gateway is a service that is responsible for routing the request to the available RPC node. It is responsible for adding, removing, and updating the RPC nodes. It is also responsible for managing the load balancer's configuration of each route. Theses operations are pushing to the message queue (redis stream) and consumer by the manager service.

[Gateway Documentations](./docs/architectures/gateway.md)

#### 4.1 Gateway Performance 

The Gateway expose prometheus metrics in /metrics handler.

[Metrics Documentations](./docs/monitoring/monitoring.md)

## Getting Started

### Deploying TyrScale
```
make run
```

### Make sure starton-local.com is added to your /etc/hosts
```
127.0.0.1 starton-local.com
```

### Create a new route

```
POST http://localhost:8888/api/v1/routes
{
    "host": "starton-local.com",
    "Path": "/eth/strat3",
    "load_balancer_strategy": "failover-priority",
    "circuit_breaker": {
        "enabled": true,
        "interval": 120000,
        "max_consecutive_failures": 3,
        "max_requests": 1,
        "timeout": 60000
  },
  "health_check": {
    "combined_with_circuit_breaker": true,
    "enabled": true,
    "interval": 10000,
    "timeout": 5000,
    "type": "eth_block_number"
  }
}
```

circuit Breaker configurations:
- interval is the cyclic period of the closed state.  If Interval is less than or equal to 0, the CircuitBreaker doesn't clear internal Counts during the closed state.
- max_consecutive_failures is the number of consecutive failures before the circuit breaker opens. By default the max_consecutive_failures is 3
- max_requests is the maximum number of requests allowed to pass through when the CircuitBreaker is half-open. If max_requests is 0, the CircuitBreaker allows only 1 request.
- timeout is the period of the open state, after which the state of the CircuitBreaker becomes half-open. If timeout is less than or equal to 0, the timeout value of the CircuitBreaker is set to 60 seconds.

Health check configurations:
- combined_with_circuit_breaker is a boolean value to combine the circuit breaker with the health check system this is useful when you want to open the circuit breaker when the health check status is failed
- enabled is a boolean value to enable the health check system
- interval is the interval of the health check
- timeout is the max timeout request when checking the health status
- type is the type of the health check (for now, only eth_block_number is supported)

### Create a new blockchain network
```
POST http://localhost:8888/api/v1/networks
{
  "blockchain": "ethereum",
  "name": "ethereum-mainnet",
  "chain_id": 1
}
```

### Create a new RPC node
```
POST http://localhost:8888/api/v1/rpcs
{
  "collectors": [
    "eth_block_number"
  ],
  "network_name": "ethereum-mainnet",
  "provider": "alchemy",
  "type": "private",
  "url": "https://eth-mainnet.g.alchemy.com/v2/xxxxx"
}
```

### Create a upstream from RPC node
```
POST http://localhost:8888/api/v1/routes/:route_uuid/upstreams
{
  "rpc_uuid": "xxxxx-xxxx-xxxx-xxxx",
  "weight": 90.0
}
```

### Test the RPC node
```
POST http://starton-local.com/eth/strat3
{
	"jsonrpc":"2.0",
	"method":"eth_syncing",
	"params":[],
	"id":1
}
```

## Use cases

If you are using TyrScale for the first time, you can refer to the following use cases to get started after deploying the TyrScale Load Balancer:
[Configurations Use Cases](./docs/configurations/use-cases.md)


## Contributing

Feel free to explore, contribute, and shape the future of TyrScale with us! Your feedback and collaboration are invaluable as we continue to refine and enhance this tool.

To get started, see [CONTRIBUTING.md](./CONTRIBUTING.md).

Please adhere to Starton's [Code of Conduct](./CODE_OF_CONDUCT.md).


## License

Tyrscale is licensed under the [Apache License 2.0](./LICENSE.md).


## Authors

- Starton: [support@starton.com](mailto:support@starton.com)
- Ghislain Cheng