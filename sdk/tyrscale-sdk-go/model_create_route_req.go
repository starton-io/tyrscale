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
	"bytes"
	"fmt"
)

// checks if the CreateRouteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRouteReq{}

// CreateRouteReq struct for CreateRouteReq
type CreateRouteReq struct {
	CircuitBreaker *CircuitbreakerSettings `json:"circuit_breaker,omitempty"`
	HealthCheck *HealthcheckHealthCheckConfig `json:"health_check,omitempty"`
	Host string `json:"host"`
	LoadBalancerStrategy BalancerLoadBalancerStrategy `json:"load_balancer_strategy"`
	Path *string `json:"path,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

type _CreateRouteReq CreateRouteReq

// NewCreateRouteReq instantiates a new CreateRouteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRouteReq(host string, loadBalancerStrategy BalancerLoadBalancerStrategy) *CreateRouteReq {
	this := CreateRouteReq{}
	this.Host = host
	this.LoadBalancerStrategy = loadBalancerStrategy
	return &this
}

// NewCreateRouteReqWithDefaults instantiates a new CreateRouteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRouteReqWithDefaults() *CreateRouteReq {
	this := CreateRouteReq{}
	return &this
}

// GetCircuitBreaker returns the CircuitBreaker field value if set, zero value otherwise.
func (o *CreateRouteReq) GetCircuitBreaker() CircuitbreakerSettings {
	if o == nil || IsNil(o.CircuitBreaker) {
		var ret CircuitbreakerSettings
		return ret
	}
	return *o.CircuitBreaker
}

// GetCircuitBreakerOk returns a tuple with the CircuitBreaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteReq) GetCircuitBreakerOk() (*CircuitbreakerSettings, bool) {
	if o == nil || IsNil(o.CircuitBreaker) {
		return nil, false
	}
	return o.CircuitBreaker, true
}

// HasCircuitBreaker returns a boolean if a field has been set.
func (o *CreateRouteReq) HasCircuitBreaker() bool {
	if o != nil && !IsNil(o.CircuitBreaker) {
		return true
	}

	return false
}

// SetCircuitBreaker gets a reference to the given CircuitbreakerSettings and assigns it to the CircuitBreaker field.
func (o *CreateRouteReq) SetCircuitBreaker(v CircuitbreakerSettings) {
	o.CircuitBreaker = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *CreateRouteReq) GetHealthCheck() HealthcheckHealthCheckConfig {
	if o == nil || IsNil(o.HealthCheck) {
		var ret HealthcheckHealthCheckConfig
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteReq) GetHealthCheckOk() (*HealthcheckHealthCheckConfig, bool) {
	if o == nil || IsNil(o.HealthCheck) {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *CreateRouteReq) HasHealthCheck() bool {
	if o != nil && !IsNil(o.HealthCheck) {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given HealthcheckHealthCheckConfig and assigns it to the HealthCheck field.
func (o *CreateRouteReq) SetHealthCheck(v HealthcheckHealthCheckConfig) {
	o.HealthCheck = &v
}

// GetHost returns the Host field value
func (o *CreateRouteReq) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateRouteReq) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateRouteReq) SetHost(v string) {
	o.Host = v
}

// GetLoadBalancerStrategy returns the LoadBalancerStrategy field value
func (o *CreateRouteReq) GetLoadBalancerStrategy() BalancerLoadBalancerStrategy {
	if o == nil {
		var ret BalancerLoadBalancerStrategy
		return ret
	}

	return o.LoadBalancerStrategy
}

// GetLoadBalancerStrategyOk returns a tuple with the LoadBalancerStrategy field value
// and a boolean to check if the value has been set.
func (o *CreateRouteReq) GetLoadBalancerStrategyOk() (*BalancerLoadBalancerStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerStrategy, true
}

// SetLoadBalancerStrategy sets field value
func (o *CreateRouteReq) SetLoadBalancerStrategy(v BalancerLoadBalancerStrategy) {
	o.LoadBalancerStrategy = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CreateRouteReq) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteReq) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CreateRouteReq) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CreateRouteReq) SetPath(v string) {
	o.Path = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CreateRouteReq) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteReq) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CreateRouteReq) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CreateRouteReq) SetUuid(v string) {
	o.Uuid = &v
}

func (o CreateRouteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRouteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CircuitBreaker) {
		toSerialize["circuit_breaker"] = o.CircuitBreaker
	}
	if !IsNil(o.HealthCheck) {
		toSerialize["health_check"] = o.HealthCheck
	}
	toSerialize["host"] = o.Host
	toSerialize["load_balancer_strategy"] = o.LoadBalancerStrategy
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

func (o *CreateRouteReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"load_balancer_strategy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateRouteReq := _CreateRouteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRouteReq)

	if err != nil {
		return err
	}

	*o = CreateRouteReq(varCreateRouteReq)

	return err
}

type NullableCreateRouteReq struct {
	value *CreateRouteReq
	isSet bool
}

func (v NullableCreateRouteReq) Get() *CreateRouteReq {
	return v.value
}

func (v *NullableCreateRouteReq) Set(val *CreateRouteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRouteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRouteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRouteReq(val *CreateRouteReq) *NullableCreateRouteReq {
	return &NullableCreateRouteReq{value: val, isSet: true}
}

func (v NullableCreateRouteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRouteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


