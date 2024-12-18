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

// checks if the PluginListPluginsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginListPluginsResponse{}

// PluginListPluginsResponse struct for PluginListPluginsResponse
type PluginListPluginsResponse struct {
	// Map of plugin lists, keyed by an integer ID
	Plugins *map[string]PluginPluginList `json:"plugins,omitempty"`
}

// NewPluginListPluginsResponse instantiates a new PluginListPluginsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginListPluginsResponse() *PluginListPluginsResponse {
	this := PluginListPluginsResponse{}
	return &this
}

// NewPluginListPluginsResponseWithDefaults instantiates a new PluginListPluginsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginListPluginsResponseWithDefaults() *PluginListPluginsResponse {
	this := PluginListPluginsResponse{}
	return &this
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *PluginListPluginsResponse) GetPlugins() map[string]PluginPluginList {
	if o == nil || IsNil(o.Plugins) {
		var ret map[string]PluginPluginList
		return ret
	}
	return *o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginListPluginsResponse) GetPluginsOk() (*map[string]PluginPluginList, bool) {
	if o == nil || IsNil(o.Plugins) {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *PluginListPluginsResponse) HasPlugins() bool {
	if o != nil && !IsNil(o.Plugins) {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given map[string]PluginPluginList and assigns it to the Plugins field.
func (o *PluginListPluginsResponse) SetPlugins(v map[string]PluginPluginList) {
	o.Plugins = &v
}

func (o PluginListPluginsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginListPluginsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plugins) {
		toSerialize["plugins"] = o.Plugins
	}
	return toSerialize, nil
}

type NullablePluginListPluginsResponse struct {
	value *PluginListPluginsResponse
	isSet bool
}

func (v NullablePluginListPluginsResponse) Get() *PluginListPluginsResponse {
	return v.value
}

func (v *NullablePluginListPluginsResponse) Set(val *PluginListPluginsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginListPluginsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginListPluginsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginListPluginsResponse(val *PluginListPluginsResponse) *NullablePluginListPluginsResponse {
	return &NullablePluginListPluginsResponse{value: val, isSet: true}
}

func (v NullablePluginListPluginsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginListPluginsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


