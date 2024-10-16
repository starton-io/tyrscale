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

// checks if the Rpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Rpc{}

// Rpc struct for Rpc
type Rpc struct {
	ChainId *int32 `json:"chain_id,omitempty"`
	Collectors []string `json:"collectors,omitempty"`
	NetworkName string `json:"network_name"`
	Provider string `json:"provider"`
	Type RPCType `json:"type"`
	Url string `json:"url"`
	Uuid *string `json:"uuid,omitempty"`
}

type _Rpc Rpc

// NewRpc instantiates a new Rpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpc(networkName string, provider string, type_ RPCType, url string) *Rpc {
	this := Rpc{}
	this.NetworkName = networkName
	this.Provider = provider
	this.Type = type_
	this.Url = url
	return &this
}

// NewRpcWithDefaults instantiates a new Rpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpcWithDefaults() *Rpc {
	this := Rpc{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *Rpc) GetChainId() int32 {
	if o == nil || IsNil(o.ChainId) {
		var ret int32
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rpc) GetChainIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *Rpc) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given int32 and assigns it to the ChainId field.
func (o *Rpc) SetChainId(v int32) {
	o.ChainId = &v
}

// GetCollectors returns the Collectors field value if set, zero value otherwise.
func (o *Rpc) GetCollectors() []string {
	if o == nil || IsNil(o.Collectors) {
		var ret []string
		return ret
	}
	return o.Collectors
}

// GetCollectorsOk returns a tuple with the Collectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rpc) GetCollectorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Collectors) {
		return nil, false
	}
	return o.Collectors, true
}

// HasCollectors returns a boolean if a field has been set.
func (o *Rpc) HasCollectors() bool {
	if o != nil && !IsNil(o.Collectors) {
		return true
	}

	return false
}

// SetCollectors gets a reference to the given []string and assigns it to the Collectors field.
func (o *Rpc) SetCollectors(v []string) {
	o.Collectors = v
}

// GetNetworkName returns the NetworkName field value
func (o *Rpc) GetNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value
// and a boolean to check if the value has been set.
func (o *Rpc) GetNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkName, true
}

// SetNetworkName sets field value
func (o *Rpc) SetNetworkName(v string) {
	o.NetworkName = v
}

// GetProvider returns the Provider field value
func (o *Rpc) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Rpc) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *Rpc) SetProvider(v string) {
	o.Provider = v
}

// GetType returns the Type field value
func (o *Rpc) GetType() RPCType {
	if o == nil {
		var ret RPCType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Rpc) GetTypeOk() (*RPCType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Rpc) SetType(v RPCType) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *Rpc) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Rpc) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Rpc) SetUrl(v string) {
	o.Url = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Rpc) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rpc) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Rpc) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Rpc) SetUuid(v string) {
	o.Uuid = &v
}

func (o Rpc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Rpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChainId) {
		toSerialize["chain_id"] = o.ChainId
	}
	if !IsNil(o.Collectors) {
		toSerialize["collectors"] = o.Collectors
	}
	toSerialize["network_name"] = o.NetworkName
	toSerialize["provider"] = o.Provider
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

func (o *Rpc) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_name",
		"provider",
		"type",
		"url",
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

	varRpc := _Rpc{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRpc)

	if err != nil {
		return err
	}

	*o = Rpc(varRpc)

	return err
}

type NullableRpc struct {
	value *Rpc
	isSet bool
}

func (v NullableRpc) Get() *Rpc {
	return v.value
}

func (v *NullableRpc) Set(val *Rpc) {
	v.value = val
	v.isSet = true
}

func (v NullableRpc) IsSet() bool {
	return v.isSet
}

func (v *NullableRpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpc(val *Rpc) *NullableRpc {
	return &NullableRpc{value: val, isSet: true}
}

func (v NullableRpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


