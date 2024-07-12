# Plugins (Experimental)

## Interceptors

Interceptors are like middleware for the gateway. They can be used to intercept requests and responses before/after the request is processed to upstream node rpc.

### Request Interceptors

How to create a request plugin interceptor:

1. Create a new file `my_request_interceptor.go`
2. Implement the `IRequestInterceptor` interface
```go
type IRequestInterceptor interface {
	Intercept(*fasthttp.Request) error
}
```
3. Build the plugin
```bash
go build -buildmode=plugin -o ./plugins/interceptor/requests/my_request_interceptor.so
```
4. Add the plugin.so to the plugins folder
```bash
mv my_request_interceptor.so ./plugins/interceptor/requests/
```

### Response Interceptors

How to create a response plugin interceptor:

1. Create a new file `my_response_interceptor.go`
2. Implement the `IResponseInterceptor` interface
```go
type IResponseInterceptor interface {
	Intercept(*fasthttp.Response) error
}
```
3. Build the plugin
```bash
go build -buildmode=plugin -o ./plugins/interceptor/responses/my_response_interceptor.so
```
4. Add the plugin.so to the plugins folder
```bash
mv my_responsoe_interceptor.so ./plugins/interceptor/rsponspses/
```

### Plugin Loader

The plugin loader is a helper to load plugins from the plugins folder. It is used to load plugins from the plugins folder and cache them.
Be aware that the plugin system will be load plugins with alphabetical order. So if you have a plugin named `001_request_interceptor.so` and `002_request_interceptor.so`, the request interceptor `001` will be loaded before the interceptor `002`.
