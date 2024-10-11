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

// checks if the UpstreamFastHTTPSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpstreamFastHTTPSettings{}

// UpstreamFastHTTPSettings struct for UpstreamFastHTTPSettings
type UpstreamFastHTTPSettings struct {
	ProxyHost *string `json:"proxy_host,omitempty"`
}

// NewUpstreamFastHTTPSettings instantiates a new UpstreamFastHTTPSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpstreamFastHTTPSettings() *UpstreamFastHTTPSettings {
	this := UpstreamFastHTTPSettings{}
	return &this
}

// NewUpstreamFastHTTPSettingsWithDefaults instantiates a new UpstreamFastHTTPSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpstreamFastHTTPSettingsWithDefaults() *UpstreamFastHTTPSettings {
	this := UpstreamFastHTTPSettings{}
	return &this
}

// GetProxyHost returns the ProxyHost field value if set, zero value otherwise.
func (o *UpstreamFastHTTPSettings) GetProxyHost() string {
	if o == nil || IsNil(o.ProxyHost) {
		var ret string
		return ret
	}
	return *o.ProxyHost
}

// GetProxyHostOk returns a tuple with the ProxyHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamFastHTTPSettings) GetProxyHostOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyHost) {
		return nil, false
	}
	return o.ProxyHost, true
}

// HasProxyHost returns a boolean if a field has been set.
func (o *UpstreamFastHTTPSettings) HasProxyHost() bool {
	if o != nil && !IsNil(o.ProxyHost) {
		return true
	}

	return false
}

// SetProxyHost gets a reference to the given string and assigns it to the ProxyHost field.
func (o *UpstreamFastHTTPSettings) SetProxyHost(v string) {
	o.ProxyHost = &v
}

func (o UpstreamFastHTTPSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpstreamFastHTTPSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProxyHost) {
		toSerialize["proxy_host"] = o.ProxyHost
	}
	return toSerialize, nil
}

type NullableUpstreamFastHTTPSettings struct {
	value *UpstreamFastHTTPSettings
	isSet bool
}

func (v NullableUpstreamFastHTTPSettings) Get() *UpstreamFastHTTPSettings {
	return v.value
}

func (v *NullableUpstreamFastHTTPSettings) Set(val *UpstreamFastHTTPSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpstreamFastHTTPSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpstreamFastHTTPSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpstreamFastHTTPSettings(val *UpstreamFastHTTPSettings) *NullableUpstreamFastHTTPSettings {
	return &NullableUpstreamFastHTTPSettings{value: val, isSet: true}
}

func (v NullableUpstreamFastHTTPSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpstreamFastHTTPSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


