# HealthcheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewHealthcheckRequest

`func NewHealthcheckRequest() *HealthcheckRequest`

NewHealthcheckRequest instantiates a new HealthcheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthcheckRequestWithDefaults

`func NewHealthcheckRequestWithDefaults() *HealthcheckRequest`

NewHealthcheckRequestWithDefaults instantiates a new HealthcheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *HealthcheckRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HealthcheckRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HealthcheckRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HealthcheckRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *HealthcheckRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HealthcheckRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HealthcheckRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HealthcheckRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *HealthcheckRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HealthcheckRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HealthcheckRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HealthcheckRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetStatusCode

`func (o *HealthcheckRequest) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HealthcheckRequest) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HealthcheckRequest) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *HealthcheckRequest) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


