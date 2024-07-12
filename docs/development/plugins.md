# Plugins (Experimental feature)

Tyrscale is a gateway that can be extended with plugins. In fact there are two types of plugins that can be used to extend the gateway:

- **Interceptors**: Interceptors allow you to manipulate requests and responses. They can be used to add headers, modify the body, or handle errors. For example, you can create a response interceptor to add custom headers to the response or log specific information. Interceptors are implemented by defining structs that implement the `IRequestInterceptor` or `IResponseInterceptor` interfaces and registering them with the gateway.

- **Middlewares**: Middlewares are used to process requests and responses in a chainable manner. They can be used for tasks such as logging, authentication, CORS handling, rate limiting, etc. Middlewares are implemented as functions that take a `fasthttp.RequestHandler` and return a new `fasthttp.RequestHandler`. They can be composed and applied to the request handling pipeline to add additional functionality.

## Interceptors

### Response Interceptors

How to create a response plugin interceptor:

1. Create a new file `my_response_interceptor.go`
2. Implement the response interceptor plugin that implements the `RegistererResponse` interface

```go
type RegistererResponse interface {
	RegisterResponseInterceptor(func(
		name string,
		interceptor types.IResponseInterceptor,
	), []byte) error
	Validate([]byte) error
}
```

```go
package main

import (
	"encoding/json"

	"github.com/starton-io/tyrscale/gateway/pkg/interceptor/types"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/valyala/fasthttp"
)

var validate validation.Validation

func init() {
	validate = validation.New()
}

type Config struct {
	Headers map[string]string `json:"headers" validate:"required"`
}

type MyResponseInterceptor struct {
	Config Config
}

func (m *MyResponseInterceptor) Intercept(response *fasthttp.Response) error {
	response.Header.Set("X-Intercepted", "true")
	for k, v := range m.Config.Headers {
		response.Header.Set(k, v)
	}
	return nil
}

type MyResponseInterceptorRegister struct{}

func (p *MyResponseInterceptorRegister) RegisterResponseInterceptor(registerFunc func(name string, interceptor types.IResponseInterceptor), payload []byte) error {
	var config Config
	if err := json.Unmarshal(payload, &config); err != nil {
		return err
	}
	if err := validate.ValidateStruct(config); err != nil {
		return err
	}
	registerFunc("MyResponseInterceptor", &MyResponseInterceptor{Config: config})
	return nil
}

func (p *MyResponseInterceptorRegister) Validate(configPayload []byte) error {
	var config Config
	if err := json.Unmarshal(configPayload, &config); err != nil {
		return err
	}
	return validate.ValidateStruct(config)
}

var ResponseInterceptor MyResponseInterceptorRegister

func main() {}

```
### Explanation:
- **Config**: A struct that holds the configuration for the interceptor, including headers to be added to the response.
- **MyResponseInterceptor**: A struct that will handle the response interception using the provided configuration.
- **Intercept**: A method that sets a custom header and additional headers from the configuration when a response is intercepted.
- **MyResponseInterceptorRegister**: A struct that registers `MyResponseInterceptor` using a provided function. This struct also handles the validation of the configuration.
- **RegisterResponseInterceptor**: A method that registers `MyResponseInterceptor` by unmarshalling the configuration payload and validating it.
- **Validate**: A method that validates the configuration payload.
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
  - name: MyResponseInterceptor
    path: ./plugins/interceptor/response/my_response_interceptor.so
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
curl -X POST http://localhost:8888/v1/routes/{uuid}/attach-plugin --data '{
    "name": "MyResponseInterceptor",
    "type": "ResponseInterceptor",
    "description": "My response interceptor",
    "config": {
        "headers": {
            "X-Custom-Header": "Custom value"
        }
    },
    "priority": 1
  }'
```

8. Disable the plugin with the manager API

```bash
curl -X POST http://localhost:8888/v1/routes/{uuid}/detach-plugin --data '{"name": "MyResponseInterceptor","type": "ResponseInterceptor"}'
```



## Middlewares

In this example, we will create a middleware that logs the request method and URI before calling the next handler, and logs the response status code after the handler has processed the request.

How to create a middleware plugin:

1. Create a new file `logging_middleware.go`
2. Implement the middleware plugin that implements the `RegistererMiddleware` interface
```go
type RegistererMiddleware interface {
	RegisterMiddleware(func(
		name string,
		middleware typesMiddleware.MiddlewareFunc,
	), []byte) error
	Validate([]byte) error
}
```

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

func (p *LoggingMiddlewareRegister) RegisterMiddleware(registerFunc func(name string, middleware types.MiddlewareFunc), payload []byte) error {
	registerFunc("LoggingMiddleware", LoggingMiddleware)
	return nil
}

func (p *LoggingMiddlewareRegister) Validate(configPayload []byte) error {
	return nil
}

// Exported symbol
var Middleware LoggingMiddlewareRegister

func main() {}
```
### Explanation:
- **LoggingMiddleware**: This function acts as a middleware that logs the HTTP request method and URI before passing the request to the next handler in the chain. After the next handler processes the request, it logs the response status code.
- **LoggingMiddlewareRegister**: A struct responsible for registering the `LoggingMiddleware` with the gateway.
- **RegisterMiddleware**: A method of `LoggingMiddlewareRegister` that registers the `LoggingMiddleware` function using a provided registration function. This method is called by the gateway to add the middleware to the processing pipeline.
- **Validate**: A method of `LoggingMiddlewareRegister` that validates the configuration payload. In this example, it does nothing and always returns `nil`, indicating no validation errors.
- **Middleware**: An exported variable of type `LoggingMiddlewareRegister` that the gateway uses to access and register the middleware.

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
    path: ./plugins/middleware/example/logging_middleware.so
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
curl -X POST http://localhost:8888/v1/routes/{uuid}/attach-plugin --data '{"name": "LoggingMiddleware","type": "Middleware", "priority": 1}'
```


8. Disable the plugin with the manager API

```bash
curl -X POST http://localhost:8888/v1/routes/{uuid}/detach-plugin --data '{"name": "LoggingMiddleware","type": "Middleware"}'
```