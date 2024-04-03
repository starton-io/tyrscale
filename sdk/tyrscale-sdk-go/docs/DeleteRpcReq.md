# DeleteRpcReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CascadeDeleteUpstream** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteRpcReq

`func NewDeleteRpcReq() *DeleteRpcReq`

NewDeleteRpcReq instantiates a new DeleteRpcReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRpcReqWithDefaults

`func NewDeleteRpcReqWithDefaults() *DeleteRpcReq`

NewDeleteRpcReqWithDefaults instantiates a new DeleteRpcReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCascadeDeleteUpstream

`func (o *DeleteRpcReq) GetCascadeDeleteUpstream() bool`

GetCascadeDeleteUpstream returns the CascadeDeleteUpstream field if non-nil, zero value otherwise.

### GetCascadeDeleteUpstreamOk

`func (o *DeleteRpcReq) GetCascadeDeleteUpstreamOk() (*bool, bool)`

GetCascadeDeleteUpstreamOk returns a tuple with the CascadeDeleteUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeDeleteUpstream

`func (o *DeleteRpcReq) SetCascadeDeleteUpstream(v bool)`

SetCascadeDeleteUpstream sets CascadeDeleteUpstream field to given value.

### HasCascadeDeleteUpstream

`func (o *DeleteRpcReq) HasCascadeDeleteUpstream() bool`

HasCascadeDeleteUpstream returns a boolean if a field has been set.

### GetUuid

`func (o *DeleteRpcReq) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeleteRpcReq) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeleteRpcReq) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeleteRpcReq) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


