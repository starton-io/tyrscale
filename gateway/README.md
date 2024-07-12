flowchart LR 
    RPCJson(RPC Json 2.0) --> RPCGateway
    WSC(WS) --> RPCGateway
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