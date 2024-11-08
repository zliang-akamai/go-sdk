/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetLinodeStats200ResponseIo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLinodeStats200ResponseIo{}

// GetLinodeStats200ResponseIo Input/Output statistics.
type GetLinodeStats200ResponseIo struct {
	// Block/s written.
	Io [][]float32 `json:"io,omitempty"`
	// Block/s written.
	Swap [][]float32 `json:"swap,omitempty"`
}

// NewGetLinodeStats200ResponseIo instantiates a new GetLinodeStats200ResponseIo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLinodeStats200ResponseIo() *GetLinodeStats200ResponseIo {
	this := GetLinodeStats200ResponseIo{}
	return &this
}

// NewGetLinodeStats200ResponseIoWithDefaults instantiates a new GetLinodeStats200ResponseIo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLinodeStats200ResponseIoWithDefaults() *GetLinodeStats200ResponseIo {
	this := GetLinodeStats200ResponseIo{}
	return &this
}

// GetIo returns the Io field value if set, zero value otherwise.
func (o *GetLinodeStats200ResponseIo) GetIo() [][]float32 {
	if o == nil || IsNil(o.Io) {
		var ret [][]float32
		return ret
	}
	return o.Io
}

// GetIoOk returns a tuple with the Io field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeStats200ResponseIo) GetIoOk() ([][]float32, bool) {
	if o == nil || IsNil(o.Io) {
		return nil, false
	}
	return o.Io, true
}

// HasIo returns a boolean if a field has been set.
func (o *GetLinodeStats200ResponseIo) HasIo() bool {
	if o != nil && !IsNil(o.Io) {
		return true
	}

	return false
}

// SetIo gets a reference to the given [][]float32 and assigns it to the Io field.
func (o *GetLinodeStats200ResponseIo) SetIo(v [][]float32) {
	o.Io = v
}

// GetSwap returns the Swap field value if set, zero value otherwise.
func (o *GetLinodeStats200ResponseIo) GetSwap() [][]float32 {
	if o == nil || IsNil(o.Swap) {
		var ret [][]float32
		return ret
	}
	return o.Swap
}

// GetSwapOk returns a tuple with the Swap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeStats200ResponseIo) GetSwapOk() ([][]float32, bool) {
	if o == nil || IsNil(o.Swap) {
		return nil, false
	}
	return o.Swap, true
}

// HasSwap returns a boolean if a field has been set.
func (o *GetLinodeStats200ResponseIo) HasSwap() bool {
	if o != nil && !IsNil(o.Swap) {
		return true
	}

	return false
}

// SetSwap gets a reference to the given [][]float32 and assigns it to the Swap field.
func (o *GetLinodeStats200ResponseIo) SetSwap(v [][]float32) {
	o.Swap = v
}

func (o GetLinodeStats200ResponseIo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLinodeStats200ResponseIo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Io) {
		toSerialize["io"] = o.Io
	}
	if !IsNil(o.Swap) {
		toSerialize["swap"] = o.Swap
	}
	return toSerialize, nil
}

type NullableGetLinodeStats200ResponseIo struct {
	value *GetLinodeStats200ResponseIo
	isSet bool
}

func (v NullableGetLinodeStats200ResponseIo) Get() *GetLinodeStats200ResponseIo {
	return v.value
}

func (v *NullableGetLinodeStats200ResponseIo) Set(val *GetLinodeStats200ResponseIo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLinodeStats200ResponseIo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLinodeStats200ResponseIo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLinodeStats200ResponseIo(val *GetLinodeStats200ResponseIo) *NullableGetLinodeStats200ResponseIo {
	return &NullableGetLinodeStats200ResponseIo{value: val, isSet: true}
}

func (v NullableGetLinodeStats200ResponseIo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLinodeStats200ResponseIo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


