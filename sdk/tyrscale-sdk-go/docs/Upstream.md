# Upstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RpcUuid** | Pointer to **string** |  | [optional] 
**Scheme** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Weight** | **float32** |  | 

## Methods

### NewUpstream

`func NewUpstream(weight float32, ) *Upstream`

NewUpstream instantiates a new Upstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamWithDefaults

`func NewUpstreamWithDefaults() *Upstream`

NewUpstreamWithDefaults instantiates a new Upstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *Upstream) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Upstream) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Upstream) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Upstream) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPath

`func (o *Upstream) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Upstream) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Upstream) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Upstream) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *Upstream) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Upstream) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Upstream) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Upstream) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRpcUuid

`func (o *Upstream) GetRpcUuid() string`

GetRpcUuid returns the RpcUuid field if non-nil, zero value otherwise.

### GetRpcUuidOk

`func (o *Upstream) GetRpcUuidOk() (*string, bool)`

GetRpcUuidOk returns a tuple with the RpcUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcUuid

`func (o *Upstream) SetRpcUuid(v string)`

SetRpcUuid sets RpcUuid field to given value.

### HasRpcUuid

`func (o *Upstream) HasRpcUuid() bool`

HasRpcUuid returns a boolean if a field has been set.

### GetScheme

`func (o *Upstream) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *Upstream) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *Upstream) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *Upstream) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetUuid

`func (o *Upstream) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Upstream) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Upstream) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Upstream) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *Upstream) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Upstream) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Upstream) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


