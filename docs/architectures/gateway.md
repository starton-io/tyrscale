# Gateway Architecture

```mermaid
flowchart LR 
    RPCJson(RPC Json 2.0) --> RPCGateway
    subgraph RPC Gateway
      RPCGateway(Handler) --> Routing
        subgraph fasthttp
         Routing --> Middleware(Middleware)
         Middleware --> ProxyHandler
        end
         ProxyHandler --> Balancer
    end
    subgraph Circuit Breaker
        Balancer --> CB[Circuit Breaker]
    end
    subgraph Pool Blockchain RPC
        CB --> PoolBlockchainRPC
        PoolBlockchainRPC --> NodeRPC1
        PoolBlockchainRPC --> NodeRPC2
        PoolBlockchainRPC --> NodeRPC3
    end
    subgraph Health Checks
        HC[Health Check Manager] -->|Update Circuit Breaker| CB
        HC[Healt Check Block Number] --> |Get the highest Block| PoolBlockchainRPC
    end
```

For now, the gateway only support RPC Json 2.0 and not support websocket (may be in the future).
The Request will be sent to the RPC Gateway. The request will be routed to the right proxy handler. The proxy handler will be sent to the right node RPC depend on the load balancer algorithm. The node RPC will send the response to the gateway. The gateway will send the response to the client.


Each load balancer have its own proxy handler.

| Load Balancer       | Proxy Handler  | Description |  
|---------------------|----------------|-------------|
| Weight-Round-Robin  | Default        | Distributes requests to upstream servers based on their weights. Each server is assigned a weight, and the load balancer cycles through the servers, sending more requests to servers with higher weights. |
| Least-Loaded        | Default        | Routes requests to the server with the least current load. This strategy helps in balancing the load more evenly across servers. |
| None-Concurrent     | Concurrent     | Ensures that requests are handled concurrently, optimizing for scenarios where multiple requests need to be processed simultaneously. |
| Failover-Priority   | Failover       | Routes requests to the server with the highest priority. This strategy helps in ensuring high availability and failover support. |


This is the sequence diagram of the proxy handler.

```mermaid
sequenceDiagram
    participant Client as fasthttp.RequestCtx
    participant Balancer as Balancer
    participant ReqInterceptor as RequestInterceptor
    participant CircuitBreaker as CircuitBreaker
    participant NodeRPC as Node RPC
    participant ResInterceptor as ResponseInterceptor

    Client->>Balancer: Request to balance
    Balancer-->>Client: List of upstream nodes
    loop For each upstream node
        Client->>ReqInterceptor: Intercept request
        ReqInterceptor-->>Client: Interception result
        alt CircuitBreaker exists
            Client->>CircuitBreaker: Execute with CircuitBreaker
            CircuitBreaker->>NodeRPC: Send request
            NodeRPC-->>CircuitBreaker: Response
            CircuitBreaker-->>Client: Execution result
        else No CircuitBreaker
            Client->>NodeRPC: Send request
            NodeRPC-->>Client: Response
        end
        Client->>ResInterceptor: Intercept response
        ResInterceptor-->>Client: Interception result
        alt Successful execution
            Client-->>Client: Return response
        else Continue to next node
        end
    end
    Client-->>Client: All nodes unhealthy/dead
```

In the proxy handler, the request will be sent to the balancer for get the list of node RPC depend on the load balancer algorithm. Then the request will be intercept. will sent to the circuit breaker for check the circuit breaker status. If the circuit breaker is open, it will return the error to the client. If the circuit breaker is closed, it will sent to right node RPC for get the response.