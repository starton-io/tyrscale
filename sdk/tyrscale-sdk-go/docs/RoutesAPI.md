# \RoutesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoute**](RoutesAPI.md#CreateRoute) | **Post** /routes | Create a route
[**DeleteRoute**](RoutesAPI.md#DeleteRoute) | **Delete** /routes/{uuid} | Delete a route
[**ListRoutes**](RoutesAPI.md#ListRoutes) | **Get** /routes | Get list routes
[**UpdateRoute**](RoutesAPI.md#UpdateRoute) | **Patch** /routes/{uuid} | Update a route



## CreateRoute

> ResponsesCreatedSuccessResponseCreateRouteRes CreateRoute(ctx).Route(route).Execute()

Create a route



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
	route := *openapiclient.NewCreateRouteReq("Host_example", openapiclient.balancer.LoadBalancerStrategy("weight-round-robin")) // CreateRouteReq | Route request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.CreateRoute(context.Background()).Route(route).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.CreateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoute`: ResponsesCreatedSuccessResponseCreateRouteRes
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.CreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | [**CreateRouteReq**](CreateRouteReq.md) | Route request | 

### Return type

[**ResponsesCreatedSuccessResponseCreateRouteRes**](ResponsesCreatedSuccessResponseCreateRouteRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> ResponsesDefaultSuccessResponseWithoutData DeleteRoute(ctx, uuid).Execute()

Delete a route



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
	resp, r, err := apiClient.RoutesAPI.DeleteRoute(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.DeleteRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoute`: ResponsesDefaultSuccessResponseWithoutData
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.DeleteRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Route UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesDefaultSuccessResponseWithoutData**](ResponsesDefaultSuccessResponseWithoutData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoutes

> ResponsesDefaultSuccessResponseListRouteRes ListRoutes(ctx).Host(host).LoadBalancerStrategy(loadBalancerStrategy).Path(path).Uuid(uuid).Execute()

Get list routes



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
	host := "host_example" // string |  (optional)
	loadBalancerStrategy := "loadBalancerStrategy_example" // string |  (optional)
	path := "path_example" // string |  (optional)
	uuid := "uuid_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.ListRoutes(context.Background()).Host(host).LoadBalancerStrategy(loadBalancerStrategy).Path(path).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.ListRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoutes`: ResponsesDefaultSuccessResponseListRouteRes
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.ListRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** |  | 
 **loadBalancerStrategy** | **string** |  | 
 **path** | **string** |  | 
 **uuid** | **string** |  | 

### Return type

[**ResponsesDefaultSuccessResponseListRouteRes**](ResponsesDefaultSuccessResponseListRouteRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoute

> ResponsesDefaultSuccessResponseWithoutData UpdateRoute(ctx, uuid).Route(route).Execute()

Update a route



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
	uuid := "uuid_example" // string | UUID
	route := *openapiclient.NewUpdateRouteReq() // UpdateRouteReq | Route request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.UpdateRoute(context.Background(), uuid).Route(route).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.UpdateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoute`: ResponsesDefaultSuccessResponseWithoutData
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **route** | [**UpdateRouteReq**](UpdateRouteReq.md) | Route request | 

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

