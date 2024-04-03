# CreateRpcReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collectors** | Pointer to **[]string** |  | [optional] 
**NetworkName** | **string** |  | 
**Provider** | **string** |  | 
**Type** | [**RPCType**](RPCType.md) |  | 
**Url** | **string** |  | 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRpcReq

`func NewCreateRpcReq(networkName string, provider string, type_ RPCType, url string, ) *CreateRpcReq`

NewCreateRpcReq instantiates a new CreateRpcReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRpcReqWithDefaults

`func NewCreateRpcReqWithDefaults() *CreateRpcReq`

NewCreateRpcReqWithDefaults instantiates a new CreateRpcReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectors

`func (o *CreateRpcReq) GetCollectors() []string`

GetCollectors returns the Collectors field if non-nil, zero value otherwise.

### GetCollectorsOk

`func (o *CreateRpcReq) GetCollectorsOk() (*[]string, bool)`

GetCollectorsOk returns a tuple with the Collectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectors

`func (o *CreateRpcReq) SetCollectors(v []string)`

SetCollectors sets Collectors field to given value.

### HasCollectors

`func (o *CreateRpcReq) HasCollectors() bool`

HasCollectors returns a boolean if a field has been set.

### GetNetworkName

`func (o *CreateRpcReq) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *CreateRpcReq) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *CreateRpcReq) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetProvider

`func (o *CreateRpcReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateRpcReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateRpcReq) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetType

`func (o *CreateRpcReq) GetType() RPCType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRpcReq) GetTypeOk() (*RPCType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRpcReq) SetType(v RPCType)`

SetType sets Type field to given value.


### GetUrl

`func (o *CreateRpcReq) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateRpcReq) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateRpcReq) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUuid

`func (o *CreateRpcReq) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateRpcReq) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateRpcReq) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CreateRpcReq) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


