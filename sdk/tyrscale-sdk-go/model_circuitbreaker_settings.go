/*
Tyrscale Manager API

This is the manager service for Tyrscale

API version: 1.0
Contact: support@starton.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tyrscalesdkgo

import (
	"encoding/json"
)

// checks if the CircuitbreakerSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitbreakerSettings{}

// CircuitbreakerSettings struct for CircuitbreakerSettings
type CircuitbreakerSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
	Interval *int32 `json:"interval,omitempty"`
	MaxConsecutiveFailures *int32 `json:"max_consecutive_failures,omitempty"`
	MaxRequests *int32 `json:"max_requests,omitempty"`
	Name *string `json:"name,omitempty"`
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewCircuitbreakerSettings instantiates a new CircuitbreakerSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitbreakerSettings() *CircuitbreakerSettings {
	this := CircuitbreakerSettings{}
	return &this
}

// NewCircuitbreakerSettingsWithDefaults instantiates a new CircuitbreakerSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitbreakerSettingsWithDefaults() *CircuitbreakerSettings {
	this := CircuitbreakerSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CircuitbreakerSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitbreakerSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CircuitbreakerSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CircuitbreakerSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *CircuitbreakerSettings) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitbreakerSettings) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *CircuitbreakerSettings) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *CircuitbreakerSettings) SetInterval(v int32) {
	o.Interval = &v
}

// GetMaxConsecutiveFailures returns the MaxConsecutiveFailures field value if set, zero value otherwise.
func (o *CircuitbreakerSettings) GetMaxConsecutiveFailures() int32 {
	if o == nil || IsNil(o.MaxConsecutiveFailures) {
		var ret int32
		return ret
	}
	return *o.MaxConsecutiveFailures
}

// GetMaxConsecutiveFailuresOk returns a tuple with the MaxConsecutiveFailures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitbreakerSettings) GetMaxConsecutiveFailuresOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConsecutiveFailures) {
		return nil, false
	}
	return o.MaxConsecutiveFailures, true
}

// HasMaxConsecutiveFailures returns a boolean if a field has been set.
func (o *CircuitbreakerSettings) HasMaxConsecutiveFailures() bool {
	if o != nil && !IsNil(o.MaxConsecutiveFailures) {
		return true
	}

	return false
}

// SetMaxConsecutiveFailures gets a reference to the given int32 and assigns it to the MaxConsecutiveFailures field.
func (o *CircuitbreakerSettings) SetMaxConsecutiveFailures(v int32) {
	o.MaxConsecutiveFailures = &v
}

// GetMaxRequests returns the MaxRequests field value if set, zero value otherwise.
func (o *CircuitbreakerSettings) GetMaxRequests() int32 {
	if o == nil || IsNil(o.MaxRequests) {
		var ret int32
		return ret
	}
	return *o.MaxRequests
}

// GetMaxRequestsOk returns a tuple with the MaxRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitbreakerSettings) GetMaxRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRequests) {
		return nil, false
	}
	return o.MaxRequests, true
}

// HasMaxRequests returns a boolean if a field has been set.
func (o *CircuitbreakerSettings) HasMaxRequests() bool {
	if o != nil && !IsNil(o.MaxRequests) {
		return true
	}

	return false
}

// SetMaxRequests gets a reference to the given int32 and assigns it to the MaxRequests field.
func (o *CircuitbreakerSettings) SetMaxRequests(v int32) {
	o.MaxRequests = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CircuitbreakerSettings) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitbreakerSettings) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CircuitbreakerSettings) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CircuitbreakerSettings) SetName(v string) {
	o.Name = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *CircuitbreakerSettings) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitbreakerSettings) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *CircuitbreakerSettings) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *CircuitbreakerSettings) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o CircuitbreakerSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitbreakerSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.MaxConsecutiveFailures) {
		toSerialize["max_consecutive_failures"] = o.MaxConsecutiveFailures
	}
	if !IsNil(o.MaxRequests) {
		toSerialize["max_requests"] = o.MaxRequests
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

type NullableCircuitbreakerSettings struct {
	value *CircuitbreakerSettings
	isSet bool
}

func (v NullableCircuitbreakerSettings) Get() *CircuitbreakerSettings {
	return v.value
}

func (v *NullableCircuitbreakerSettings) Set(val *CircuitbreakerSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitbreakerSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitbreakerSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitbreakerSettings(val *CircuitbreakerSettings) *NullableCircuitbreakerSettings {
	return &NullableCircuitbreakerSettings{value: val, isSet: true}
}

func (v NullableCircuitbreakerSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitbreakerSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


