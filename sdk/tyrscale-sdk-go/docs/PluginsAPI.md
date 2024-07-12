# \PluginsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachPlugin**](PluginsAPI.md#AttachPlugin) | **Post** /routes/{uuid}/attach-plugin | Attach plugin to route
[**DetachPlugin**](PluginsAPI.md#DetachPlugin) | **Post** /routes/{uuid}/detach-plugin | Detach plugin from route
[**ListPlugins**](PluginsAPI.md#ListPlugins) | **Get** /plugins | Get list plugins
[**ListPluginsFromRoute**](PluginsAPI.md#ListPluginsFromRoute) | **Get** /routes/{uuid}/plugins | Get list plugins from route



## AttachPlugin

> ResponsesDefaultSuccessResponseWithoutData AttachPlugin(ctx, uuid).Body(body).Execute()

Attach plugin to route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/starton-io/tyrscale/sdk/tyrscale-sdk-go"
)

func main() {
	uuid := "uuid_example" // string | Route UUID
	body := *openapiclient.NewAttachPluginReq(map[string]interface{}(123), "Name_example", int32(123), openapiclient.plugin.PluginType("ResponseInterceptor")) // AttachPluginReq | Attach plugin request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.AttachPlugin(context.Background(), uuid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.AttachPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachPlugin`: ResponsesDefaultSuccessResponseWithoutData
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.AttachPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Route UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AttachPluginReq**](AttachPluginReq.md) | Attach plugin request | 

### Return type

[**ResponsesDefaultSuccessResponseWithoutData**](ResponsesDefaultSuccessResponseWithoutData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachPlugin

> ResponsesDefaultSuccessResponseWithoutData DetachPlugin(ctx, uuid).Body(body).Execute()

Detach plugin from route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/starton-io/tyrscale/sdk/tyrscale-sdk-go"
)

func main() {
	uuid := "uuid_example" // string | Route UUID
	body := *openapiclient.NewDetachPluginReq("Name_example", openapiclient.plugin.PluginType("ResponseInterceptor")) // DetachPluginReq | Detach plugin request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.DetachPlugin(context.Background(), uuid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.DetachPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachPlugin`: ResponsesDefaultSuccessResponseWithoutData
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.DetachPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Route UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DetachPluginReq**](DetachPluginReq.md) | Detach plugin request | 

### Return type

[**ResponsesDefaultSuccessResponseWithoutData**](ResponsesDefaultSuccessResponseWithoutData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlugins

> ResponsesDefaultSuccessResponsePluginListPluginsResponse ListPlugins(ctx).Execute()

Get list plugins



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/starton-io/tyrscale/sdk/tyrscale-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.ListPlugins(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.ListPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlugins`: ResponsesDefaultSuccessResponsePluginListPluginsResponse
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.ListPlugins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPluginsRequest struct via the builder pattern


### Return type

[**ResponsesDefaultSuccessResponsePluginListPluginsResponse**](ResponsesDefaultSuccessResponsePluginListPluginsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPluginsFromRoute

> ResponsesDefaultSuccessResponsePlugins ListPluginsFromRoute(ctx, uuid).Execute()

Get list plugins from route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/starton-io/tyrscale/sdk/tyrscale-sdk-go"
)

func main() {
	uuid := "uuid_example" // string | Route UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.ListPluginsFromRoute(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.ListPluginsFromRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPluginsFromRoute`: ResponsesDefaultSuccessResponsePlugins
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.ListPluginsFromRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Route UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPluginsFromRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesDefaultSuccessResponsePlugins**](ResponsesDefaultSuccessResponsePlugins.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

