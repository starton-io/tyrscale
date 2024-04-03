# CircuitbreakerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**MaxConsecutiveFailures** | Pointer to **int32** |  | [optional] 
**MaxRequests** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewCircuitbreakerSettings

`func NewCircuitbreakerSettings() *CircuitbreakerSettings`

NewCircuitbreakerSettings instantiates a new CircuitbreakerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitbreakerSettingsWithDefaults

`func NewCircuitbreakerSettingsWithDefaults() *CircuitbreakerSettings`

NewCircuitbreakerSettingsWithDefaults instantiates a new CircuitbreakerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CircuitbreakerSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CircuitbreakerSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CircuitbreakerSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CircuitbreakerSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInterval

`func (o *CircuitbreakerSettings) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CircuitbreakerSettings) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CircuitbreakerSettings) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CircuitbreakerSettings) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMaxConsecutiveFailures

`func (o *CircuitbreakerSettings) GetMaxConsecutiveFailures() int32`

GetMaxConsecutiveFailures returns the MaxConsecutiveFailures field if non-nil, zero value otherwise.

### GetMaxConsecutiveFailuresOk

`func (o *CircuitbreakerSettings) GetMaxConsecutiveFailuresOk() (*int32, bool)`

GetMaxConsecutiveFailuresOk returns a tuple with the MaxConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConsecutiveFailures

`func (o *CircuitbreakerSettings) SetMaxConsecutiveFailures(v int32)`

SetMaxConsecutiveFailures sets MaxConsecutiveFailures field to given value.

### HasMaxConsecutiveFailures

`func (o *CircuitbreakerSettings) HasMaxConsecutiveFailures() bool`

HasMaxConsecutiveFailures returns a boolean if a field has been set.

### GetMaxRequests

`func (o *CircuitbreakerSettings) GetMaxRequests() int32`

GetMaxRequests returns the MaxRequests field if non-nil, zero value otherwise.

### GetMaxRequestsOk

`func (o *CircuitbreakerSettings) GetMaxRequestsOk() (*int32, bool)`

GetMaxRequestsOk returns a tuple with the MaxRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequests

`func (o *CircuitbreakerSettings) SetMaxRequests(v int32)`

SetMaxRequests sets MaxRequests field to given value.

### HasMaxRequests

`func (o *CircuitbreakerSettings) HasMaxRequests() bool`

HasMaxRequests returns a boolean if a field has been set.

### GetName

`func (o *CircuitbreakerSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CircuitbreakerSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CircuitbreakerSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CircuitbreakerSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeout

`func (o *CircuitbreakerSettings) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CircuitbreakerSettings) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CircuitbreakerSettings) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CircuitbreakerSettings) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


