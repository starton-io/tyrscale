# \NetworksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetwork**](NetworksAPI.md#CreateNetwork) | **Post** /networks | Create a network
[**DeleteNetwork**](NetworksAPI.md#DeleteNetwork) | **Delete** /networks/{name} | Delete a network
[**ListNetworks**](NetworksAPI.md#ListNetworks) | **Get** /networks | Get list networks



## CreateNetwork

> ResponsesCreatedSuccessResponseCreateNetworkRes CreateNetwork(ctx).Network(network).Execute()

Create a network



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
	network := *openapiclient.NewNetwork("ethereum", int32(1), "eth-mainnet") // Network | Network request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetwork(context.Background()).Network(network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetwork`: ResponsesCreatedSuccessResponseCreateNetworkRes
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network** | [**Network**](Network.md) | Network request | 

### Return type

[**ResponsesCreatedSuccessResponseCreateNetworkRes**](ResponsesCreatedSuccessResponseCreateNetworkRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> ResponsesDefaultSuccessResponseDeleteNetworkRes DeleteNetwork(ctx, name).Execute()

Delete a network



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
	name := "name_example" // string | Network Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetwork(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetwork`: ResponsesDefaultSuccessResponseDeleteNetworkRes
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Network Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesDefaultSuccessResponseDeleteNetworkRes**](ResponsesDefaultSuccessResponseDeleteNetworkRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworks

> ResponsesDefaultSuccessResponseListNetworkRes ListNetworks(ctx).Blockchain(blockchain).ChainId(chainId).Name(name).Execute()

Get list networks



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
	blockchain := "blockchain_example" // string |  (optional)
	chainId := "chainId_example" // string |  (optional)
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworks(context.Background()).Blockchain(blockchain).ChainId(chainId).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworks`: ResponsesDefaultSuccessResponseListNetworkRes
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockchain** | **string** |  | 
 **chainId** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**ResponsesDefaultSuccessResponseListNetworkRes**](ResponsesDefaultSuccessResponseListNetworkRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

