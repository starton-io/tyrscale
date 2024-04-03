# \RpcsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRpc**](RpcsAPI.md#CreateRpc) | **Post** /rpcs | Create a new Rpc
[**DeleteRPC**](RpcsAPI.md#DeleteRPC) | **Delete** /rpcs/{uuid} | Delete a RPC
[**ListRPCs**](RpcsAPI.md#ListRPCs) | **Get** /rpcs | List RPCs
[**UpdateRPC**](RpcsAPI.md#UpdateRPC) | **Put** /rpcs | Update a RPC



## CreateRpc

> ResponsesCreatedSuccessResponseCreateRpcRes CreateRpc(ctx).Rpc(rpc).Execute()

Create a new Rpc



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
	rpc := *openapiclient.NewCreateRpcReq("NetworkName_example", "Provider_example", openapiclient.RPCType("private"), "Url_example") // CreateRpcReq | Rpc request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RpcsAPI.CreateRpc(context.Background()).Rpc(rpc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpcsAPI.CreateRpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRpc`: ResponsesCreatedSuccessResponseCreateRpcRes
	fmt.Fprintf(os.Stdout, "Response from `RpcsAPI.CreateRpc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rpc** | [**CreateRpcReq**](CreateRpcReq.md) | Rpc request | 

### Return type

[**ResponsesCreatedSuccessResponseCreateRpcRes**](ResponsesCreatedSuccessResponseCreateRpcRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRPC

> ResponsesDefaultSuccessResponseWithoutData DeleteRPC(ctx, uuid).Rpc(rpc).Execute()

Delete a RPC



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
	rpc := *openapiclient.NewDeleteRpcOptReq() // DeleteRpcOptReq | Delete Rpc request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RpcsAPI.DeleteRPC(context.Background(), uuid).Rpc(rpc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpcsAPI.DeleteRPC``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRPC`: ResponsesDefaultSuccessResponseWithoutData
	fmt.Fprintf(os.Stdout, "Response from `RpcsAPI.DeleteRPC`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRPCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpc** | [**DeleteRpcOptReq**](DeleteRpcOptReq.md) | Delete Rpc request | 

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


## ListRPCs

> ResponsesDefaultSuccessResponseListRpcRes ListRPCs(ctx).Uuid(uuid).ChainId(chainId).Provider(provider).Type_(type_).NetworkName(networkName).SortBy(sortBy).SortAscending(sortAscending).Execute()

List RPCs



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
	uuid := "uuid_example" // string | UUID (optional)
	chainId := "chainId_example" // string | Chain ID (optional)
	provider := "provider_example" // string | provider (optional)
	type_ := "type__example" // string | type (optional)
	networkName := "networkName_example" // string | network_name (optional)
	sortBy := "sortBy_example" // string | sort_by (optional)
	sortAscending := true // bool | sort_ascending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RpcsAPI.ListRPCs(context.Background()).Uuid(uuid).ChainId(chainId).Provider(provider).Type_(type_).NetworkName(networkName).SortBy(sortBy).SortAscending(sortAscending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpcsAPI.ListRPCs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRPCs`: ResponsesDefaultSuccessResponseListRpcRes
	fmt.Fprintf(os.Stdout, "Response from `RpcsAPI.ListRPCs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** | UUID | 
 **chainId** | **string** | Chain ID | 
 **provider** | **string** | provider | 
 **type_** | **string** | type | 
 **networkName** | **string** | network_name | 
 **sortBy** | **string** | sort_by | 
 **sortAscending** | **bool** | sort_ascending | 

### Return type

[**ResponsesDefaultSuccessResponseListRpcRes**](ResponsesDefaultSuccessResponseListRpcRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRPC

> ResponsesDefaultSuccessResponseWithoutData UpdateRPC(ctx).Rpc(rpc).Execute()

Update a RPC



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
	rpc := *openapiclient.NewRpc("NetworkName_example", "Provider_example", openapiclient.RPCType("private"), "Url_example") // Rpc | RPC Object request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RpcsAPI.UpdateRPC(context.Background()).Rpc(rpc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpcsAPI.UpdateRPC``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRPC`: ResponsesDefaultSuccessResponseWithoutData
	fmt.Fprintf(os.Stdout, "Response from `RpcsAPI.UpdateRPC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRPCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rpc** | [**Rpc**](Rpc.md) | RPC Object request | 

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

