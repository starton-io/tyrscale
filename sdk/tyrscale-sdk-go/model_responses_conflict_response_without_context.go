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

// checks if the ResponsesConflictResponseWithoutContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesConflictResponseWithoutContext{}

// ResponsesConflictResponseWithoutContext struct for ResponsesConflictResponseWithoutContext
type ResponsesConflictResponseWithoutContext struct {
	Code *int32 `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

// NewResponsesConflictResponseWithoutContext instantiates a new ResponsesConflictResponseWithoutContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesConflictResponseWithoutContext() *ResponsesConflictResponseWithoutContext {
	this := ResponsesConflictResponseWithoutContext{}
	return &this
}

// NewResponsesConflictResponseWithoutContextWithDefaults instantiates a new ResponsesConflictResponseWithoutContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesConflictResponseWithoutContextWithDefaults() *ResponsesConflictResponseWithoutContext {
	this := ResponsesConflictResponseWithoutContext{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResponsesConflictResponseWithoutContext) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesConflictResponseWithoutContext) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResponsesConflictResponseWithoutContext) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ResponsesConflictResponseWithoutContext) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponsesConflictResponseWithoutContext) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesConflictResponseWithoutContext) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponsesConflictResponseWithoutContext) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponsesConflictResponseWithoutContext) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponsesConflictResponseWithoutContext) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesConflictResponseWithoutContext) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponsesConflictResponseWithoutContext) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ResponsesConflictResponseWithoutContext) SetStatus(v int32) {
	o.Status = &v
}

func (o ResponsesConflictResponseWithoutContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesConflictResponseWithoutContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResponsesConflictResponseWithoutContext struct {
	value *ResponsesConflictResponseWithoutContext
	isSet bool
}

func (v NullableResponsesConflictResponseWithoutContext) Get() *ResponsesConflictResponseWithoutContext {
	return v.value
}

func (v *NullableResponsesConflictResponseWithoutContext) Set(val *ResponsesConflictResponseWithoutContext) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesConflictResponseWithoutContext) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesConflictResponseWithoutContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesConflictResponseWithoutContext(val *ResponsesConflictResponseWithoutContext) *NullableResponsesConflictResponseWithoutContext {
	return &NullableResponsesConflictResponseWithoutContext{value: val, isSet: true}
}

func (v NullableResponsesConflictResponseWithoutContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesConflictResponseWithoutContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


