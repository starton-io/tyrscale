# \UpstreamsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUpstream**](UpstreamsAPI.md#DeleteUpstream) | **Delete** /routes/{route_uuid}/upstreams/{uuid} | Delete a upstream
[**ListUpstreams**](UpstreamsAPI.md#ListUpstreams) | **Get** /routes/{route_uuid}/upstreams | Get list upstreams
[**UpsertUpstream**](UpstreamsAPI.md#UpsertUpstream) | **Put** /routes/{route_uuid}/upstreams | Create or update a upstream



## DeleteUpstream

> ResponsesDefaultSuccessResponseUpstreamUpsertRes DeleteUpstream(ctx, routeUuid, uuid).Execute()

Delete a upstream



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
	routeUuid := "routeUuid_example" // string | Route UUID
	uuid := "uuid_example" // string | Upstream UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamsAPI.DeleteUpstream(context.Background(), routeUuid, uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamsAPI.DeleteUpstream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUpstream`: ResponsesDefaultSuccessResponseUpstreamUpsertRes
	fmt.Fprintf(os.Stdout, "Response from `UpstreamsAPI.DeleteUpstream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeUuid** | **string** | Route UUID | 
**uuid** | **string** | Upstream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUpstreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponsesDefaultSuccessResponseUpstreamUpsertRes**](ResponsesDefaultSuccessResponseUpstreamUpsertRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUpstreams

> ResponsesDefaultSuccessResponseListUpstreamRes ListUpstreams(ctx, routeUuid).Execute()

Get list upstreams



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
	routeUuid := "routeUuid_example" // string | Route UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamsAPI.ListUpstreams(context.Background(), routeUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamsAPI.ListUpstreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUpstreams`: ResponsesDefaultSuccessResponseListUpstreamRes
	fmt.Fprintf(os.Stdout, "Response from `UpstreamsAPI.ListUpstreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeUuid** | **string** | Route UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUpstreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesDefaultSuccessResponseListUpstreamRes**](ResponsesDefaultSuccessResponseListUpstreamRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertUpstream

> ResponsesCreatedSuccessResponseUpstreamUpsertRes UpsertUpstream(ctx, routeUuid).Upstream(upstream).Execute()

Create or update a upstream



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
	routeUuid := "routeUuid_example" // string | Route UUID
	upstream := *openapiclient.NewUpstream(float32(123)) // Upstream | Upstream request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamsAPI.UpsertUpstream(context.Background(), routeUuid).Upstream(upstream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamsAPI.UpsertUpstream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertUpstream`: ResponsesCreatedSuccessResponseUpstreamUpsertRes
	fmt.Fprintf(os.Stdout, "Response from `UpstreamsAPI.UpsertUpstream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeUuid** | **string** | Route UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertUpstreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstream** | [**Upstream**](Upstream.md) | Upstream request | 

### Return type

[**ResponsesCreatedSuccessResponseUpstreamUpsertRes**](ResponsesCreatedSuccessResponseUpstreamUpsertRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

