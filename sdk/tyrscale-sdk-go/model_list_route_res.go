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

// checks if the ListRouteRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRouteRes{}

// ListRouteRes struct for ListRouteRes
type ListRouteRes struct {
	Items []Route `json:"items,omitempty"`
}

// NewListRouteRes instantiates a new ListRouteRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRouteRes() *ListRouteRes {
	this := ListRouteRes{}
	return &this
}

// NewListRouteResWithDefaults instantiates a new ListRouteRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRouteResWithDefaults() *ListRouteRes {
	this := ListRouteRes{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListRouteRes) GetItems() []Route {
	if o == nil || IsNil(o.Items) {
		var ret []Route
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRouteRes) GetItemsOk() ([]Route, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListRouteRes) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Route and assigns it to the Items field.
func (o *ListRouteRes) SetItems(v []Route) {
	o.Items = v
}

func (o ListRouteRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRouteRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableListRouteRes struct {
	value *ListRouteRes
	isSet bool
}

func (v NullableListRouteRes) Get() *ListRouteRes {
	return v.value
}

func (v *NullableListRouteRes) Set(val *ListRouteRes) {
	v.value = val
	v.isSet = true
}

func (v NullableListRouteRes) IsSet() bool {
	return v.isSet
}

func (v *NullableListRouteRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRouteRes(val *ListRouteRes) *NullableListRouteRes {
	return &NullableListRouteRes{value: val, isSet: true}
}

func (v NullableListRouteRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRouteRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


