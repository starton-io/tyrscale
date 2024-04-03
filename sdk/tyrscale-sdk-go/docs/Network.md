# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** |  | 
**ChainId** | **int32** |  | 
**Name** | **string** |  | 

## Methods

### NewNetwork

`func NewNetwork(blockchain string, chainId int32, name string, ) *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *Network) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *Network) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *Network) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetChainId

`func (o *Network) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Network) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Network) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetName

`func (o *Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Network) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


