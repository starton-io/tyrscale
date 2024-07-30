# Plugins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Middleware** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 
**RequestInterceptor** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 
**ResponseInterceptor** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 

## Methods

### NewPlugins

`func NewPlugins() *Plugins`

NewPlugins instantiates a new Plugins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsWithDefaults

`func NewPluginsWithDefaults() *Plugins`

NewPluginsWithDefaults instantiates a new Plugins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMiddleware

`func (o *Plugins) GetMiddleware() []Plugin`

GetMiddleware returns the Middleware field if non-nil, zero value otherwise.

### GetMiddlewareOk

`func (o *Plugins) GetMiddlewareOk() (*[]Plugin, bool)`

GetMiddlewareOk returns a tuple with the Middleware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleware

`func (o *Plugins) SetMiddleware(v []Plugin)`

SetMiddleware sets Middleware field to given value.

### HasMiddleware

`func (o *Plugins) HasMiddleware() bool`

HasMiddleware returns a boolean if a field has been set.

### GetRequestInterceptor

`func (o *Plugins) GetRequestInterceptor() []Plugin`

GetRequestInterceptor returns the RequestInterceptor field if non-nil, zero value otherwise.

### GetRequestInterceptorOk

`func (o *Plugins) GetRequestInterceptorOk() (*[]Plugin, bool)`

GetRequestInterceptorOk returns a tuple with the RequestInterceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestInterceptor

`func (o *Plugins) SetRequestInterceptor(v []Plugin)`

SetRequestInterceptor sets RequestInterceptor field to given value.

### HasRequestInterceptor

`func (o *Plugins) HasRequestInterceptor() bool`

HasRequestInterceptor returns a boolean if a field has been set.

### GetResponseInterceptor

`func (o *Plugins) GetResponseInterceptor() []Plugin`

GetResponseInterceptor returns the ResponseInterceptor field if non-nil, zero value otherwise.

### GetResponseInterceptorOk

`func (o *Plugins) GetResponseInterceptorOk() (*[]Plugin, bool)`

GetResponseInterceptorOk returns a tuple with the ResponseInterceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseInterceptor

`func (o *Plugins) SetResponseInterceptor(v []Plugin)`

SetResponseInterceptor sets ResponseInterceptor field to given value.

### HasResponseInterceptor

`func (o *Plugins) HasResponseInterceptor() bool`

HasResponseInterceptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


