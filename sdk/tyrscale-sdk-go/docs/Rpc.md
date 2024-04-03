# Rpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int32** |  | [optional] 
**Collectors** | Pointer to **[]string** |  | [optional] 
**NetworkName** | **string** |  | 
**Provider** | **string** |  | 
**Type** | [**RPCType**](RPCType.md) |  | 
**Url** | **string** |  | 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewRpc

`func NewRpc(networkName string, provider string, type_ RPCType, url string, ) *Rpc`

NewRpc instantiates a new Rpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcWithDefaults

`func NewRpcWithDefaults() *Rpc`

NewRpcWithDefaults instantiates a new Rpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *Rpc) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Rpc) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Rpc) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Rpc) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetCollectors

`func (o *Rpc) GetCollectors() []string`

GetCollectors returns the Collectors field if non-nil, zero value otherwise.

### GetCollectorsOk

`func (o *Rpc) GetCollectorsOk() (*[]string, bool)`

GetCollectorsOk returns a tuple with the Collectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectors

`func (o *Rpc) SetCollectors(v []string)`

SetCollectors sets Collectors field to given value.

### HasCollectors

`func (o *Rpc) HasCollectors() bool`

HasCollectors returns a boolean if a field has been set.

### GetNetworkName

`func (o *Rpc) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *Rpc) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *Rpc) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetProvider

`func (o *Rpc) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Rpc) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Rpc) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetType

`func (o *Rpc) GetType() RPCType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rpc) GetTypeOk() (*RPCType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rpc) SetType(v RPCType)`

SetType sets Type field to given value.


### GetUrl

`func (o *Rpc) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Rpc) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Rpc) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUuid

`func (o *Rpc) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Rpc) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Rpc) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Rpc) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


