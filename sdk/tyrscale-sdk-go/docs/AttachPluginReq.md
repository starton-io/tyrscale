# AttachPluginReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** |  | 
**PluginType** | [**PluginPluginType**](PluginPluginType.md) |  | 
**Priority** | **int32** |  | 

## Methods

### NewAttachPluginReq

`func NewAttachPluginReq(pluginName string, pluginType PluginPluginType, priority int32, ) *AttachPluginReq`

NewAttachPluginReq instantiates a new AttachPluginReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachPluginReqWithDefaults

`func NewAttachPluginReqWithDefaults() *AttachPluginReq`

NewAttachPluginReqWithDefaults instantiates a new AttachPluginReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AttachPluginReq) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AttachPluginReq) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AttachPluginReq) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetPluginType

`func (o *AttachPluginReq) GetPluginType() PluginPluginType`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AttachPluginReq) GetPluginTypeOk() (*PluginPluginType, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AttachPluginReq) SetPluginType(v PluginPluginType)`

SetPluginType sets PluginType field to given value.


### GetPriority

`func (o *AttachPluginReq) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AttachPluginReq) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AttachPluginReq) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


