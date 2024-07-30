# PluginListPluginsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**map[string]PluginPluginList**](PluginPluginList.md) | Map of plugin lists, keyed by an integer ID | [optional] 

## Methods

### NewPluginListPluginsResponse

`func NewPluginListPluginsResponse() *PluginListPluginsResponse`

NewPluginListPluginsResponse instantiates a new PluginListPluginsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginListPluginsResponseWithDefaults

`func NewPluginListPluginsResponseWithDefaults() *PluginListPluginsResponse`

NewPluginListPluginsResponseWithDefaults instantiates a new PluginListPluginsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugins

`func (o *PluginListPluginsResponse) GetPlugins() map[string]PluginPluginList`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *PluginListPluginsResponse) GetPluginsOk() (*map[string]PluginPluginList, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *PluginListPluginsResponse) SetPlugins(v map[string]PluginPluginList)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *PluginListPluginsResponse) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


