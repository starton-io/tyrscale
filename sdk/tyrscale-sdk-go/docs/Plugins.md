# Plugins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterceptorRequest** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 
**InterceptorResponse** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 
**Middleware** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 

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

### GetInterceptorRequest

`func (o *Plugins) GetInterceptorRequest() []Plugin`

GetInterceptorRequest returns the InterceptorRequest field if non-nil, zero value otherwise.

### GetInterceptorRequestOk

`func (o *Plugins) GetInterceptorRequestOk() (*[]Plugin, bool)`

GetInterceptorRequestOk returns a tuple with the InterceptorRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterceptorRequest

`func (o *Plugins) SetInterceptorRequest(v []Plugin)`

SetInterceptorRequest sets InterceptorRequest field to given value.

### HasInterceptorRequest

`func (o *Plugins) HasInterceptorRequest() bool`

HasInterceptorRequest returns a boolean if a field has been set.

### GetInterceptorResponse

`func (o *Plugins) GetInterceptorResponse() []Plugin`

GetInterceptorResponse returns the InterceptorResponse field if non-nil, zero value otherwise.

### GetInterceptorResponseOk

`func (o *Plugins) GetInterceptorResponseOk() (*[]Plugin, bool)`

GetInterceptorResponseOk returns a tuple with the InterceptorResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterceptorResponse

`func (o *Plugins) SetInterceptorResponse(v []Plugin)`

SetInterceptorResponse sets InterceptorResponse field to given value.

### HasInterceptorResponse

`func (o *Plugins) HasInterceptorResponse() bool`

HasInterceptorResponse returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


