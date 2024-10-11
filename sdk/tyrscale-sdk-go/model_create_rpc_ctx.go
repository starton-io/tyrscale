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

// checks if the CreateRpcCtx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRpcCtx{}

// CreateRpcCtx struct for CreateRpcCtx
type CreateRpcCtx struct {
	Uuid *string `json:"uuid,omitempty"`
}

// NewCreateRpcCtx instantiates a new CreateRpcCtx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRpcCtx() *CreateRpcCtx {
	this := CreateRpcCtx{}
	return &this
}

// NewCreateRpcCtxWithDefaults instantiates a new CreateRpcCtx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRpcCtxWithDefaults() *CreateRpcCtx {
	this := CreateRpcCtx{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CreateRpcCtx) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRpcCtx) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CreateRpcCtx) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CreateRpcCtx) SetUuid(v string) {
	o.Uuid = &v
}

func (o CreateRpcCtx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRpcCtx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableCreateRpcCtx struct {
	value *CreateRpcCtx
	isSet bool
}

func (v NullableCreateRpcCtx) Get() *CreateRpcCtx {
	return v.value
}

func (v *NullableCreateRpcCtx) Set(val *CreateRpcCtx) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRpcCtx) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRpcCtx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRpcCtx(val *CreateRpcCtx) *NullableCreateRpcCtx {
	return &NullableCreateRpcCtx{value: val, isSet: true}
}

func (v NullableCreateRpcCtx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRpcCtx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

