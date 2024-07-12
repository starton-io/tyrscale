# AttachPluginReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]interface{}** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Priority** | **int32** |  | 
**Type** | [**PluginPluginType**](PluginPluginType.md) |  | 

## Methods

### NewAttachPluginReq

`func NewAttachPluginReq(config map[string]interface{}, name string, priority int32, type_ PluginPluginType, ) *AttachPluginReq`

NewAttachPluginReq instantiates a new AttachPluginReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachPluginReqWithDefaults

`func NewAttachPluginReqWithDefaults() *AttachPluginReq`

NewAttachPluginReqWithDefaults instantiates a new AttachPluginReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AttachPluginReq) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AttachPluginReq) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AttachPluginReq) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetDescription

`func (o *AttachPluginReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttachPluginReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttachPluginReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttachPluginReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AttachPluginReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachPluginReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachPluginReq) SetName(v string)`

SetName sets Name field to given value.


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


### GetType

`func (o *AttachPluginReq) GetType() PluginPluginType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachPluginReq) GetTypeOk() (*PluginPluginType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachPluginReq) SetType(v PluginPluginType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


