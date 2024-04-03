# CircuitBreakerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Settings** | Pointer to [**CircuitbreakerSettings**](CircuitbreakerSettings.md) |  | [optional] 

## Methods

### NewCircuitBreakerConfig

`func NewCircuitBreakerConfig() *CircuitBreakerConfig`

NewCircuitBreakerConfig instantiates a new CircuitBreakerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitBreakerConfigWithDefaults

`func NewCircuitBreakerConfigWithDefaults() *CircuitBreakerConfig`

NewCircuitBreakerConfigWithDefaults instantiates a new CircuitBreakerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CircuitBreakerConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CircuitBreakerConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CircuitBreakerConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CircuitBreakerConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *CircuitBreakerConfig) GetSettings() CircuitbreakerSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CircuitBreakerConfig) GetSettingsOk() (*CircuitbreakerSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CircuitBreakerConfig) SetSettings(v CircuitbreakerSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CircuitBreakerConfig) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


