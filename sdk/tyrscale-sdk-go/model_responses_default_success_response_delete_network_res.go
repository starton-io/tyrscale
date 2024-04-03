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

// checks if the ResponsesDefaultSuccessResponseDeleteNetworkRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesDefaultSuccessResponseDeleteNetworkRes{}

// ResponsesDefaultSuccessResponseDeleteNetworkRes struct for ResponsesDefaultSuccessResponseDeleteNetworkRes
type ResponsesDefaultSuccessResponseDeleteNetworkRes struct {
	Code *int32 `json:"code,omitempty"`
	Data *DeleteNetworkRes `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

// NewResponsesDefaultSuccessResponseDeleteNetworkRes instantiates a new ResponsesDefaultSuccessResponseDeleteNetworkRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesDefaultSuccessResponseDeleteNetworkRes() *ResponsesDefaultSuccessResponseDeleteNetworkRes {
	this := ResponsesDefaultSuccessResponseDeleteNetworkRes{}
	return &this
}

// NewResponsesDefaultSuccessResponseDeleteNetworkResWithDefaults instantiates a new ResponsesDefaultSuccessResponseDeleteNetworkRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesDefaultSuccessResponseDeleteNetworkResWithDefaults() *ResponsesDefaultSuccessResponseDeleteNetworkRes {
	this := ResponsesDefaultSuccessResponseDeleteNetworkRes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) GetData() DeleteNetworkRes {
	if o == nil || IsNil(o.Data) {
		var ret DeleteNetworkRes
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) GetDataOk() (*DeleteNetworkRes, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteNetworkRes and assigns it to the Data field.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) SetData(v DeleteNetworkRes) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ResponsesDefaultSuccessResponseDeleteNetworkRes) SetStatus(v int32) {
	o.Status = &v
}

func (o ResponsesDefaultSuccessResponseDeleteNetworkRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesDefaultSuccessResponseDeleteNetworkRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResponsesDefaultSuccessResponseDeleteNetworkRes struct {
	value *ResponsesDefaultSuccessResponseDeleteNetworkRes
	isSet bool
}

func (v NullableResponsesDefaultSuccessResponseDeleteNetworkRes) Get() *ResponsesDefaultSuccessResponseDeleteNetworkRes {
	return v.value
}

func (v *NullableResponsesDefaultSuccessResponseDeleteNetworkRes) Set(val *ResponsesDefaultSuccessResponseDeleteNetworkRes) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesDefaultSuccessResponseDeleteNetworkRes) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesDefaultSuccessResponseDeleteNetworkRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesDefaultSuccessResponseDeleteNetworkRes(val *ResponsesDefaultSuccessResponseDeleteNetworkRes) *NullableResponsesDefaultSuccessResponseDeleteNetworkRes {
	return &NullableResponsesDefaultSuccessResponseDeleteNetworkRes{value: val, isSet: true}
}

func (v NullableResponsesDefaultSuccessResponseDeleteNetworkRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesDefaultSuccessResponseDeleteNetworkRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


