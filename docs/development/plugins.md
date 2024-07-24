# Plugins (Experimental feature)

Tyrscale is a gateway that can be extended with plugins. In fact there are two types of plugins that can be used to extend the gateway:

- **Interceptors**: Interceptors allow you to manipulate requests and responses. They can be used to add headers, modify the body, or handle errors. For example, you can create a response interceptor to add custom headers to the response or log specific information. Interceptors are implemented by defining structs that implement the `IRequestInterceptor` or `IResponseInterceptor` interfaces and registering them with the gateway.

- **Middlewares**: Middlewares are used to process requests and responses in a chainable manner. They can be used for tasks such as logging, authentication, CORS handling, rate limiting, etc. Middlewares are implemented as functions that take a `fasthttp.RequestHandler` and return a new `fasthttp.RequestHandler`. They can be composed and applied to the request handling pipeline to add additional functionality.

## Interceptors

### Response Interceptors

How to create a response plugin interceptor:

1. Create a new file `my_response_interceptor.go`
2. Implement the response interceptor plugin
```go
package main

import (
	"log"

	"github.com/starton-io/tyrscale/gateway/pkg/interceptor/types"
	"github.com/valyala/fasthttp"
)

type MyResponseInterceptor struct {
}

func (m *MyResponseInterceptor) Intercept(response *fasthttp.Response) error {
	response.Header.Set("X-Intercepted", "true")
	log.Println("MyResponseInterceptor")
	return nil
}

type MyPlugin struct{}

func (p *MyPlugin) RegisterResponseInterceptors(registerFunc func(name string, interceptor types.IResponseInterceptor)) {
	registerFunc("MyResponseInterceptor", &MyResponseInterceptor{})
}

// Ensure MyPlugin implements the Registerer interface
var ResponseInterceptor MyPlugin

func main() {}

```
### Explanation:
- **MyResponseInterceptor**: A struct that will handle the response interception.
- **Intercept**: A method that sets a custom header and logs a message when a response is intercepted.
- **RegisterResponseInterceptors**: A method that registers `MyResponseInterceptor` using a provided function. This method must be implemented by the plugin.
- **MyPlugin**: A struct that will register the interceptor to the gateway.
- **ResponseInterceptor**: An exported variable that will be used by the gateway to register the interceptor.

3. Make sure that the plugin go.mod have the same package version than the gateway go.mod. For example, if the gateway application uses `github.com/valyala/fasthttp v1.54.0`, the version of this package must be the same for the go.mod plugin.

4. Build the plugin

```bash
go build -buildmode=plugin -o ./plugins/interceptor/response/my_response_interceptor.so
```

5. Get the sha256 checksum of the plugin

With Linux:
```bash
sha256sum ./plugins/interceptor/response/my_response_interceptor.so
```
With MacOs:
```bash
shasum -a 256 ./plugins/interceptor/response/my_response_interceptor.so
```

6. Register the plugin in plugins.yaml and restart the gateway

```yaml
plugins:
  - name: my_response_interceptor
    sha256: 1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
```


7. It's possible to list the plugins with the manager API

```bash
curl -X GET http://localhost:8888/v1/plugins
```
```json
{
    "status": 200,
    "code": 0,
    "message": "Success",
    "data": {
        "plugins": {
            "Middleware": {
                "names": [
                    "LoggingMiddleware"
                ]
            },
            "ResponseInterceptor": {
                "names": [
                    "MyResponseInterceptor"
                ]
            }
        }
    }
}
```

7. Enable the plugin with the manager API

```bash
curl -X POST http://localhost:8888/v1/routes/{uuid}/attach-plugin --data '{"plugin_name": "my_response_interceptor","plugin_type": "ResponseInterceptor", "priority": 1}'
```

8. Disable the plugin with the manager API

```bash
curl -X POST http://localhost:8888/v1/routes/{uuid}/attach-plugin --data '{"plugin_name": "my_response_interceptor","plugin_type": "ResponseInterceptor"}'
```



## Middlewares

In this example, we will create a middleware that logs the request method and URI before calling the next handler, and logs the response status code after the handler has processed the request.

How to create a middleware plugin:

1. Create a new file `logging_middleware.go`
2. Implement the middleware plugin
```go
package main

import (
	"log"

	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"

	"github.com/valyala/fasthttp"
)

// LoggingMiddleware is an example middleware that logs requests
func LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Printf("Request: %s %s", ctx.Method(), ctx.RequestURI())
		next(ctx)
		log.Printf("Response: %s %d", ctx.Method(), ctx.Response.StatusCode())
	}
}

type LoggingMiddlewareRegister struct{}

func (p *LoggingMiddlewareRegister) RegisterMiddleware(registerFunc func(name string, interceptor types.MiddlewareFunc)) {
	registerFunc("LoggingMiddleware", LoggingMiddleware)
}

// Exported symbol
var Middleware LoggingMiddlewareRegister

func main() {}
```
### Explanation:
- **LoggingMiddleware**: A function that logs the request method and URI before calling the next handler, and logs the response status code after the handler has processed the request.
- **LoggingMiddlewareRegister**: A struct that registers the `LoggingMiddleware` function with the gateway.
- **RegisterMiddleware**: A method that registers `LoggingMiddleware` using a provided function. This method must be implemented by the plugin.
- **Middleware**: An exported variable that will be used by the gateway to register the middleware.

3. Make sure that the plugin go.mod have the same package version than the gateway go.mod. For example, if the gateway application uses `github.com/valyala/fasthttp v1.54.0`, the version of this package must be the same for the go.mod plugin.

4. Build the plugin

```bash
go build -buildmode=plugin -o ./plugins/middleware/example/logging_middleware.so
```

5. Get the sha256 checksum of the plugin

With Linux:
```bash
sha256sum ./plugins/middleware/example/logging_middleware.so
```
With MacOs:
```bash
shasum -a 256 ./plugins/middleware/example/logging_middleware.so
```

6. Register the plugin in plugins.yaml and restart the gateway

```yaml
plugins:
  - name: LoggingMiddleware
    sha256: 1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
```


7. It's possible to list the plugins with the manager API

```bash
curl -X GET http://localhost:8888/v1/plugins
```
```json
{
    "status": 200,
    "code": 0,
    "message": "Success",
    "data": {
        "plugins": {
            "Middleware": {
                "names": [
                    "LoggingMiddleware"
                ]
            },
            "ResponseInterceptor": {
                "names": [
                    "MyResponseInterceptor"
                ]
            }
        }
    }
}
```

7. Enable the plugin with the manager API

```bash
curl -X POST http://localhost:8888/v1/routes/{uuid}/attach-plugin --data '{"plugin_name": "LoggingMiddleware","plugin_type": "Middleware", "priority": 1}'
```


8. Disable the plugin with the manager API

```bash
curl -X POST http://localhost:8888/v1/routes/{uuid}/attach-plugin --data '{"plugin_name": "LoggingMiddleware","plugin_type": "Middleware"}'
```
