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

// checks if the StatsAvailableCpuInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatsAvailableCpuInner{}

// StatsAvailableCpuInner A stat data point.
type StatsAvailableCpuInner struct {
	// A stats graph data point.
	X *int32 `json:"x,omitempty"`
	// A stats graph data point.
	Y *int32 `json:"y,omitempty"`
}

// NewStatsAvailableCpuInner instantiates a new StatsAvailableCpuInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsAvailableCpuInner() *StatsAvailableCpuInner {
	this := StatsAvailableCpuInner{}
	return &this
}

// NewStatsAvailableCpuInnerWithDefaults instantiates a new StatsAvailableCpuInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsAvailableCpuInnerWithDefaults() *StatsAvailableCpuInner {
	this := StatsAvailableCpuInner{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise.
func (o *StatsAvailableCpuInner) GetX() int32 {
	if o == nil || IsNil(o.X) {
		var ret int32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAvailableCpuInner) GetXOk() (*int32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *StatsAvailableCpuInner) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given int32 and assigns it to the X field.
func (o *StatsAvailableCpuInner) SetX(v int32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *StatsAvailableCpuInner) GetY() int32 {
	if o == nil || IsNil(o.Y) {
		var ret int32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAvailableCpuInner) GetYOk() (*int32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *StatsAvailableCpuInner) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given int32 and assigns it to the Y field.
func (o *StatsAvailableCpuInner) SetY(v int32) {
	o.Y = &v
}

func (o StatsAvailableCpuInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatsAvailableCpuInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return toSerialize, nil
}

type NullableStatsAvailableCpuInner struct {
	value *StatsAvailableCpuInner
	isSet bool
}

func (v NullableStatsAvailableCpuInner) Get() *StatsAvailableCpuInner {
	return v.value
}

func (v *NullableStatsAvailableCpuInner) Set(val *StatsAvailableCpuInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsAvailableCpuInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsAvailableCpuInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsAvailableCpuInner(val *StatsAvailableCpuInner) *NullableStatsAvailableCpuInner {
	return &NullableStatsAvailableCpuInner{value: val, isSet: true}
}

func (v NullableStatsAvailableCpuInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsAvailableCpuInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


