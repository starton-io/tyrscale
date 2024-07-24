# \PluginsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPlugins**](PluginsAPI.md#ListPlugins) | **Get** /plugins | Get list plugins



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

