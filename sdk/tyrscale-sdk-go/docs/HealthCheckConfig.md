# HealthCheckConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CombinedWithCircuitBreaker** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**HealthcheckHealthCheckType**](HealthcheckHealthCheckType.md) |  | [optional] 

## Methods

### NewHealthCheckConfig

`func NewHealthCheckConfig() *HealthCheckConfig`

NewHealthCheckConfig instantiates a new HealthCheckConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckConfigWithDefaults

`func NewHealthCheckConfigWithDefaults() *HealthCheckConfig`

NewHealthCheckConfigWithDefaults instantiates a new HealthCheckConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinedWithCircuitBreaker

`func (o *HealthCheckConfig) GetCombinedWithCircuitBreaker() bool`

GetCombinedWithCircuitBreaker returns the CombinedWithCircuitBreaker field if non-nil, zero value otherwise.

### GetCombinedWithCircuitBreakerOk

`func (o *HealthCheckConfig) GetCombinedWithCircuitBreakerOk() (*bool, bool)`

GetCombinedWithCircuitBreakerOk returns a tuple with the CombinedWithCircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedWithCircuitBreaker

`func (o *HealthCheckConfig) SetCombinedWithCircuitBreaker(v bool)`

SetCombinedWithCircuitBreaker sets CombinedWithCircuitBreaker field to given value.

### HasCombinedWithCircuitBreaker

`func (o *HealthCheckConfig) HasCombinedWithCircuitBreaker() bool`

HasCombinedWithCircuitBreaker returns a boolean if a field has been set.

### GetEnabled

`func (o *HealthCheckConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HealthCheckConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HealthCheckConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HealthCheckConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInterval

`func (o *HealthCheckConfig) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HealthCheckConfig) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *HealthCheckConfig) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *HealthCheckConfig) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTimeout

`func (o *HealthCheckConfig) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HealthCheckConfig) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HealthCheckConfig) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HealthCheckConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *HealthCheckConfig) GetType() HealthcheckHealthCheckType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthCheckConfig) GetTypeOk() (*HealthcheckHealthCheckType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthCheckConfig) SetType(v HealthcheckHealthCheckType)`

SetType sets Type field to given value.

### HasType

`func (o *HealthCheckConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


