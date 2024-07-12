# HealthcheckHealthCheckConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CombinedWithCircuitBreaker** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**Request** | Pointer to [**HealthcheckRequest**](HealthcheckRequest.md) |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**HealthcheckHealthCheckType**](HealthcheckHealthCheckType.md) |  | [optional] 

## Methods

### NewHealthcheckHealthCheckConfig

`func NewHealthcheckHealthCheckConfig() *HealthcheckHealthCheckConfig`

NewHealthcheckHealthCheckConfig instantiates a new HealthcheckHealthCheckConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthcheckHealthCheckConfigWithDefaults

`func NewHealthcheckHealthCheckConfigWithDefaults() *HealthcheckHealthCheckConfig`

NewHealthcheckHealthCheckConfigWithDefaults instantiates a new HealthcheckHealthCheckConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinedWithCircuitBreaker

`func (o *HealthcheckHealthCheckConfig) GetCombinedWithCircuitBreaker() bool`

GetCombinedWithCircuitBreaker returns the CombinedWithCircuitBreaker field if non-nil, zero value otherwise.

### GetCombinedWithCircuitBreakerOk

`func (o *HealthcheckHealthCheckConfig) GetCombinedWithCircuitBreakerOk() (*bool, bool)`

GetCombinedWithCircuitBreakerOk returns a tuple with the CombinedWithCircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedWithCircuitBreaker

`func (o *HealthcheckHealthCheckConfig) SetCombinedWithCircuitBreaker(v bool)`

SetCombinedWithCircuitBreaker sets CombinedWithCircuitBreaker field to given value.

### HasCombinedWithCircuitBreaker

`func (o *HealthcheckHealthCheckConfig) HasCombinedWithCircuitBreaker() bool`

HasCombinedWithCircuitBreaker returns a boolean if a field has been set.

### GetEnabled

`func (o *HealthcheckHealthCheckConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HealthcheckHealthCheckConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HealthcheckHealthCheckConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HealthcheckHealthCheckConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInterval

`func (o *HealthcheckHealthCheckConfig) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HealthcheckHealthCheckConfig) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *HealthcheckHealthCheckConfig) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *HealthcheckHealthCheckConfig) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRequest

`func (o *HealthcheckHealthCheckConfig) GetRequest() HealthcheckRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *HealthcheckHealthCheckConfig) GetRequestOk() (*HealthcheckRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *HealthcheckHealthCheckConfig) SetRequest(v HealthcheckRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *HealthcheckHealthCheckConfig) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetTimeout

`func (o *HealthcheckHealthCheckConfig) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HealthcheckHealthCheckConfig) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HealthcheckHealthCheckConfig) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HealthcheckHealthCheckConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *HealthcheckHealthCheckConfig) GetType() HealthcheckHealthCheckType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthcheckHealthCheckConfig) GetTypeOk() (*HealthcheckHealthCheckType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthcheckHealthCheckConfig) SetType(v HealthcheckHealthCheckType)`

SetType sets Type field to given value.

### HasType

`func (o *HealthcheckHealthCheckConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


