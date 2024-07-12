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

// checks if the PluginPluginList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginPluginList{}

// PluginPluginList struct for PluginPluginList
type PluginPluginList struct {
	// Repeated field of plugin names
	Names []string `json:"names,omitempty"`
}

// NewPluginPluginList instantiates a new PluginPluginList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginPluginList() *PluginPluginList {
	this := PluginPluginList{}
	return &this
}

// NewPluginPluginListWithDefaults instantiates a new PluginPluginList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginPluginListWithDefaults() *PluginPluginList {
	this := PluginPluginList{}
	return &this
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *PluginPluginList) GetNames() []string {
	if o == nil || IsNil(o.Names) {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginPluginList) GetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *PluginPluginList) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *PluginPluginList) SetNames(v []string) {
	o.Names = v
}

func (o PluginPluginList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginPluginList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	return toSerialize, nil
}

type NullablePluginPluginList struct {
	value *PluginPluginList
	isSet bool
}

func (v NullablePluginPluginList) Get() *PluginPluginList {
	return v.value
}

func (v *NullablePluginPluginList) Set(val *PluginPluginList) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginPluginList) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginPluginList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginPluginList(val *PluginPluginList) *NullablePluginPluginList {
	return &NullablePluginPluginList{value: val, isSet: true}
}

func (v NullablePluginPluginList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginPluginList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


