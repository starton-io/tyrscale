# \RecommendationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecommendation**](RecommendationAPI.md#CreateRecommendation) | **Post** /recommendations | Create a recommendation
[**DeleteRecommendation**](RecommendationAPI.md#DeleteRecommendation) | **Delete** /recommendations/{route_uuid} | Delete a recommendation
[**ListRecommendations**](RecommendationAPI.md#ListRecommendations) | **Get** /recommendations | List recommendation
[**UpdateRecommendation**](RecommendationAPI.md#UpdateRecommendation) | **Put** /recommendations | Update a recommendation



## CreateRecommendation

> ResponsesCreatedSuccessResponseCreateRecommendationRes CreateRecommendation(ctx).Recommendation(recommendation).Execute()

Create a recommendation



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
	recommendation := *openapiclient.NewCreateRecommendationReq("NetworkName_example", "RouteUuid_example", "Schedule_example", openapiclient.StrategyName("STRATEGY_CUSTOM")) // CreateRecommendationReq | Recommendation Object request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecommendationAPI.CreateRecommendation(context.Background()).Recommendation(recommendation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecommendationAPI.CreateRecommendation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecommendation`: ResponsesCreatedSuccessResponseCreateRecommendationRes
	fmt.Fprintf(os.Stdout, "Response from `RecommendationAPI.CreateRecommendation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecommendationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommendation** | [**CreateRecommendationReq**](CreateRecommendationReq.md) | Recommendation Object request | 

### Return type

[**ResponsesCreatedSuccessResponseCreateRecommendationRes**](ResponsesCreatedSuccessResponseCreateRecommendationRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecommendation

> ResponsesDefaultSuccessResponseWithoutData DeleteRecommendation(ctx, routeUuid).Execute()

Delete a recommendation



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
	resp, r, err := apiClient.RecommendationAPI.DeleteRecommendation(context.Background(), routeUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecommendationAPI.DeleteRecommendation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRecommendation`: ResponsesDefaultSuccessResponseWithoutData
	fmt.Fprintf(os.Stdout, "Response from `RecommendationAPI.DeleteRecommendation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeUuid** | **string** | Route UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecommendationRequest struct via the builder pattern


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


## ListRecommendations

> ResponsesDefaultSuccessResponseListRecommendationRes ListRecommendations(ctx).Execute()

List recommendation



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
	resp, r, err := apiClient.RecommendationAPI.ListRecommendations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecommendationAPI.ListRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRecommendations`: ResponsesDefaultSuccessResponseListRecommendationRes
	fmt.Fprintf(os.Stdout, "Response from `RecommendationAPI.ListRecommendations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRecommendationsRequest struct via the builder pattern


### Return type

[**ResponsesDefaultSuccessResponseListRecommendationRes**](ResponsesDefaultSuccessResponseListRecommendationRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecommendation

> ResponsesDefaultSuccessResponseWithoutData UpdateRecommendation(ctx).Recommendation(recommendation).Execute()

Update a recommendation



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
	recommendation := *openapiclient.NewUpdateRecommendationReq("NetworkName_example", "RouteUuid_example", "Schedule_example", openapiclient.StrategyName("STRATEGY_CUSTOM")) // UpdateRecommendationReq | Recommendation Object request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecommendationAPI.UpdateRecommendation(context.Background()).Recommendation(recommendation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecommendationAPI.UpdateRecommendation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRecommendation`: ResponsesDefaultSuccessResponseWithoutData
	fmt.Fprintf(os.Stdout, "Response from `RecommendationAPI.UpdateRecommendation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecommendationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommendation** | [**UpdateRecommendationReq**](UpdateRecommendationReq.md) | Recommendation Object request | 

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

