# Manager Service (Alpha)

## Overview

The Manager Service is a core component of the TyrScale project, a blockchain load balancer designed to manage the load between multiple blockchain RPC nodes, currently supporting EVM-based chains. It handles the addition, removal, and updating of RPC nodes and networks, configuring the service monitor for optimal performance.

## Features

- **Network Management**: Create, list, and delete blockchain networks.

- **RPC Node Management**: Add, update, and remove RPC nodes, ensuring optimal load distribution and network performance.

- **Health Checks**: Implements health checks for both the service itself and the underlying infrastructure, ensuring high availability and reliability.

- **Swagger Documentation**: Comprehensive API documentation using Swagger, facilitating easy integration and usage.


## Getting Started

### Prerequisites

- Go 1.20+ (this service use generics)
- make
- docker
- docker-compose
- Redis instance for caching and message queuing


### Installation
1. Clone the repository:

```bash
git clone https://github.com/starton-io/tyrscale.git
```

2. Navigate to the manager service directory:

```bash
cd tyrscale/manager
```

3. Run the service:

```bash
make setup
```

### Configuration

The Manager Service configuration is managed through environment variables, with the option to override defaults using a `.env` file located at the path specified by the `--config` flag when starting the service. Below are the configurable environment variables and their default values:

- `environment`: Sets the running environment of the service. Default is `"production"`.
- `http_port`: The HTTP port on which the service listens. Default is `8888`.
- `redis_uri`: The URI for connecting to the Redis instance. Default is `"localhost:6379"`.
- `redis_global_prefix`: The global prefix used for Redis keys. Default is `"evm:stream"`.
- `redis_password`: The password for connecting to the Redis instance. No default provided.
- `redis_db`: The Redis database number to use. Default is `0`.
- `read_timeout`: The read timeout (in seconds) for the server. No default provided.
- `write_timeout`: The write timeout (in seconds) for the server. No default provided.
- `server_name`: The name of the server. No default provided.
- `log_level`: The logging level of the application. Default is `"debug"`.
- `otlp_enabled`: Flag to enable or disable OTLP. Default is `true`.
- `otlp_endpoint`: The endpoint for OTLP traces. Default is `"http://localhost:4318/v1/traces"`.
- `app_version`: The version of the application. No default provided.

To override any of these defaults, you can create or modify the `.env` file at the specified configuration path (`--config` flag) or set them directly in your environment before starting the service.### Running the service


- To start the service using Docker Compose:
```bash
make run
```

- To build the service manually:
```bash
make build
```



## API Documentation

API documentation is available in Swagger format. After starting the service, visit `/docs` endpoint for the Swagger UI.

## Contributing

Contributions are welcome! Please refer to [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines on how to contribute to the project.


## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](../LICENSE.md) file for details.


## Contact


- Starton: [support@starton.com](mailto:support@starton.com)
- Ghislain Cheng


For more information, visit the [project repository](https://github.com/starton-io/tyrscale).