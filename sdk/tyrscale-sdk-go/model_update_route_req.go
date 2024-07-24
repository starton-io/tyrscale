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

// checks if the UpdateRouteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRouteReq{}

// UpdateRouteReq struct for UpdateRouteReq
type UpdateRouteReq struct {
	CircuitBreaker *CircuitbreakerSettings `json:"circuit_breaker,omitempty"`
	HealthCheck *HealthcheckHealthCheckConfig `json:"health_check,omitempty"`
	// Uuid                 string                         `json:\"uuid\" validate:\"required,uuid\"`
	Host *string `json:"host,omitempty"`
	LoadBalancerStrategy *BalancerLoadBalancerStrategy `json:"load_balancer_strategy,omitempty"`
	Path *string `json:"path,omitempty"`
	Plugins *Plugins `json:"plugins,omitempty"`
}

// NewUpdateRouteReq instantiates a new UpdateRouteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRouteReq() *UpdateRouteReq {
	this := UpdateRouteReq{}
	return &this
}

// NewUpdateRouteReqWithDefaults instantiates a new UpdateRouteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRouteReqWithDefaults() *UpdateRouteReq {
	this := UpdateRouteReq{}
	return &this
}

// GetCircuitBreaker returns the CircuitBreaker field value if set, zero value otherwise.
func (o *UpdateRouteReq) GetCircuitBreaker() CircuitbreakerSettings {
	if o == nil || IsNil(o.CircuitBreaker) {
		var ret CircuitbreakerSettings
		return ret
	}
	return *o.CircuitBreaker
}

// GetCircuitBreakerOk returns a tuple with the CircuitBreaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteReq) GetCircuitBreakerOk() (*CircuitbreakerSettings, bool) {
	if o == nil || IsNil(o.CircuitBreaker) {
		return nil, false
	}
	return o.CircuitBreaker, true
}

// HasCircuitBreaker returns a boolean if a field has been set.
func (o *UpdateRouteReq) HasCircuitBreaker() bool {
	if o != nil && !IsNil(o.CircuitBreaker) {
		return true
	}

	return false
}

// SetCircuitBreaker gets a reference to the given CircuitbreakerSettings and assigns it to the CircuitBreaker field.
func (o *UpdateRouteReq) SetCircuitBreaker(v CircuitbreakerSettings) {
	o.CircuitBreaker = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *UpdateRouteReq) GetHealthCheck() HealthcheckHealthCheckConfig {
	if o == nil || IsNil(o.HealthCheck) {
		var ret HealthcheckHealthCheckConfig
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteReq) GetHealthCheckOk() (*HealthcheckHealthCheckConfig, bool) {
	if o == nil || IsNil(o.HealthCheck) {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *UpdateRouteReq) HasHealthCheck() bool {
	if o != nil && !IsNil(o.HealthCheck) {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given HealthcheckHealthCheckConfig and assigns it to the HealthCheck field.
func (o *UpdateRouteReq) SetHealthCheck(v HealthcheckHealthCheckConfig) {
	o.HealthCheck = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UpdateRouteReq) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteReq) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UpdateRouteReq) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UpdateRouteReq) SetHost(v string) {
	o.Host = &v
}

// GetLoadBalancerStrategy returns the LoadBalancerStrategy field value if set, zero value otherwise.
func (o *UpdateRouteReq) GetLoadBalancerStrategy() BalancerLoadBalancerStrategy {
	if o == nil || IsNil(o.LoadBalancerStrategy) {
		var ret BalancerLoadBalancerStrategy
		return ret
	}
	return *o.LoadBalancerStrategy
}

// GetLoadBalancerStrategyOk returns a tuple with the LoadBalancerStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteReq) GetLoadBalancerStrategyOk() (*BalancerLoadBalancerStrategy, bool) {
	if o == nil || IsNil(o.LoadBalancerStrategy) {
		return nil, false
	}
	return o.LoadBalancerStrategy, true
}

// HasLoadBalancerStrategy returns a boolean if a field has been set.
func (o *UpdateRouteReq) HasLoadBalancerStrategy() bool {
	if o != nil && !IsNil(o.LoadBalancerStrategy) {
		return true
	}

	return false
}

// SetLoadBalancerStrategy gets a reference to the given BalancerLoadBalancerStrategy and assigns it to the LoadBalancerStrategy field.
func (o *UpdateRouteReq) SetLoadBalancerStrategy(v BalancerLoadBalancerStrategy) {
	o.LoadBalancerStrategy = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *UpdateRouteReq) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteReq) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UpdateRouteReq) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *UpdateRouteReq) SetPath(v string) {
	o.Path = &v
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *UpdateRouteReq) GetPlugins() Plugins {
	if o == nil || IsNil(o.Plugins) {
		var ret Plugins
		return ret
	}
	return *o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteReq) GetPluginsOk() (*Plugins, bool) {
	if o == nil || IsNil(o.Plugins) {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *UpdateRouteReq) HasPlugins() bool {
	if o != nil && !IsNil(o.Plugins) {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given Plugins and assigns it to the Plugins field.
func (o *UpdateRouteReq) SetPlugins(v Plugins) {
	o.Plugins = &v
}

func (o UpdateRouteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRouteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CircuitBreaker) {
		toSerialize["circuit_breaker"] = o.CircuitBreaker
	}
	if !IsNil(o.HealthCheck) {
		toSerialize["health_check"] = o.HealthCheck
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.LoadBalancerStrategy) {
		toSerialize["load_balancer_strategy"] = o.LoadBalancerStrategy
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Plugins) {
		toSerialize["plugins"] = o.Plugins
	}
	return toSerialize, nil
}

type NullableUpdateRouteReq struct {
	value *UpdateRouteReq
	isSet bool
}

func (v NullableUpdateRouteReq) Get() *UpdateRouteReq {
	return v.value
}

func (v *NullableUpdateRouteReq) Set(val *UpdateRouteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRouteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRouteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRouteReq(val *UpdateRouteReq) *NullableUpdateRouteReq {
	return &NullableUpdateRouteReq{value: val, isSet: true}
}

func (v NullableUpdateRouteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRouteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


