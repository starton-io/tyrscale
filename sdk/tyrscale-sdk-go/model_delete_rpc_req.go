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

// checks if the DeleteRpcReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRpcReq{}

// DeleteRpcReq struct for DeleteRpcReq
type DeleteRpcReq struct {
	CascadeDeleteUpstream *bool `json:"cascadeDeleteUpstream,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewDeleteRpcReq instantiates a new DeleteRpcReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRpcReq() *DeleteRpcReq {
	this := DeleteRpcReq{}
	return &this
}

// NewDeleteRpcReqWithDefaults instantiates a new DeleteRpcReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRpcReqWithDefaults() *DeleteRpcReq {
	this := DeleteRpcReq{}
	return &this
}

// GetCascadeDeleteUpstream returns the CascadeDeleteUpstream field value if set, zero value otherwise.
func (o *DeleteRpcReq) GetCascadeDeleteUpstream() bool {
	if o == nil || IsNil(o.CascadeDeleteUpstream) {
		var ret bool
		return ret
	}
	return *o.CascadeDeleteUpstream
}

// GetCascadeDeleteUpstreamOk returns a tuple with the CascadeDeleteUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRpcReq) GetCascadeDeleteUpstreamOk() (*bool, bool) {
	if o == nil || IsNil(o.CascadeDeleteUpstream) {
		return nil, false
	}
	return o.CascadeDeleteUpstream, true
}

// HasCascadeDeleteUpstream returns a boolean if a field has been set.
func (o *DeleteRpcReq) HasCascadeDeleteUpstream() bool {
	if o != nil && !IsNil(o.CascadeDeleteUpstream) {
		return true
	}

	return false
}

// SetCascadeDeleteUpstream gets a reference to the given bool and assigns it to the CascadeDeleteUpstream field.
func (o *DeleteRpcReq) SetCascadeDeleteUpstream(v bool) {
	o.CascadeDeleteUpstream = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DeleteRpcReq) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRpcReq) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DeleteRpcReq) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DeleteRpcReq) SetUuid(v string) {
	o.Uuid = &v
}

func (o DeleteRpcReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRpcReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CascadeDeleteUpstream) {
		toSerialize["cascadeDeleteUpstream"] = o.CascadeDeleteUpstream
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableDeleteRpcReq struct {
	value *DeleteRpcReq
	isSet bool
}

func (v NullableDeleteRpcReq) Get() *DeleteRpcReq {
	return v.value
}

func (v *NullableDeleteRpcReq) Set(val *DeleteRpcReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRpcReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRpcReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRpcReq(val *DeleteRpcReq) *NullableDeleteRpcReq {
	return &NullableDeleteRpcReq{value: val, isSet: true}
}

func (v NullableDeleteRpcReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRpcReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


