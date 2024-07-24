# ListPluginsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**map[string]Plugin**](Plugin.md) |  | [optional] 

## Methods

### NewListPluginsResponse

`func NewListPluginsResponse() *ListPluginsResponse`

NewListPluginsResponse instantiates a new ListPluginsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPluginsResponseWithDefaults

`func NewListPluginsResponseWithDefaults() *ListPluginsResponse`

NewListPluginsResponseWithDefaults instantiates a new ListPluginsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugins

`func (o *ListPluginsResponse) GetPlugins() map[string]Plugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ListPluginsResponse) GetPluginsOk() (*map[string]Plugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ListPluginsResponse) SetPlugins(v map[string]Plugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ListPluginsResponse) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


