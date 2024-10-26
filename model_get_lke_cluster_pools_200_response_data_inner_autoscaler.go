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

// checks if the GetLkeClusterPools200ResponseDataInnerAutoscaler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLkeClusterPools200ResponseDataInnerAutoscaler{}

// GetLkeClusterPools200ResponseDataInnerAutoscaler When enabled, the number of nodes autoscales within the defined minimum and maximum values.
type GetLkeClusterPools200ResponseDataInnerAutoscaler struct {
	// Whether autoscaling is enabled for this Node Pool. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty"`
	// The maximum number of nodes to autoscale to. Defaults to the Node Pool's `count`.
	Max *int32 `json:"max,omitempty"`
	// The minimum number of nodes to autoscale to. Defaults to the Node Pool's `count`.
	Min *int32 `json:"min,omitempty"`
}

// NewGetLkeClusterPools200ResponseDataInnerAutoscaler instantiates a new GetLkeClusterPools200ResponseDataInnerAutoscaler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLkeClusterPools200ResponseDataInnerAutoscaler() *GetLkeClusterPools200ResponseDataInnerAutoscaler {
	this := GetLkeClusterPools200ResponseDataInnerAutoscaler{}
	return &this
}

// NewGetLkeClusterPools200ResponseDataInnerAutoscalerWithDefaults instantiates a new GetLkeClusterPools200ResponseDataInnerAutoscaler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLkeClusterPools200ResponseDataInnerAutoscalerWithDefaults() *GetLkeClusterPools200ResponseDataInnerAutoscaler {
	this := GetLkeClusterPools200ResponseDataInnerAutoscaler{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) SetMax(v int32) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) SetMin(v int32) {
	o.Min = &v
}

func (o GetLkeClusterPools200ResponseDataInnerAutoscaler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLkeClusterPools200ResponseDataInnerAutoscaler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	return toSerialize, nil
}

type NullableGetLkeClusterPools200ResponseDataInnerAutoscaler struct {
	value *GetLkeClusterPools200ResponseDataInnerAutoscaler
	isSet bool
}

func (v NullableGetLkeClusterPools200ResponseDataInnerAutoscaler) Get() *GetLkeClusterPools200ResponseDataInnerAutoscaler {
	return v.value
}

func (v *NullableGetLkeClusterPools200ResponseDataInnerAutoscaler) Set(val *GetLkeClusterPools200ResponseDataInnerAutoscaler) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLkeClusterPools200ResponseDataInnerAutoscaler) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLkeClusterPools200ResponseDataInnerAutoscaler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLkeClusterPools200ResponseDataInnerAutoscaler(val *GetLkeClusterPools200ResponseDataInnerAutoscaler) *NullableGetLkeClusterPools200ResponseDataInnerAutoscaler {
	return &NullableGetLkeClusterPools200ResponseDataInnerAutoscaler{value: val, isSet: true}
}

func (v NullableGetLkeClusterPools200ResponseDataInnerAutoscaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLkeClusterPools200ResponseDataInnerAutoscaler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

