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

// checks if the PostLkeClusterRequestNodePoolsInnerAutoscaler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLkeClusterRequestNodePoolsInnerAutoscaler{}

// PostLkeClusterRequestNodePoolsInnerAutoscaler When enabled, the number of nodes automatically scales within the defined minimum and maximum values. When making a request, `max` and `min` require each other.
type PostLkeClusterRequestNodePoolsInnerAutoscaler struct {
	// Whether automatic scaling is enabled for this node pool. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty"`
	// The maximum number of nodes to automatically scale to. Defaults to the value provided by the `count` field.
	Max *int32 `json:"max,omitempty"`
	// The minimum number of nodes to automatically scale to. Defaults to the node pool's `count`.
	Min *int32 `json:"min,omitempty"`
}

// NewPostLkeClusterRequestNodePoolsInnerAutoscaler instantiates a new PostLkeClusterRequestNodePoolsInnerAutoscaler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLkeClusterRequestNodePoolsInnerAutoscaler() *PostLkeClusterRequestNodePoolsInnerAutoscaler {
	this := PostLkeClusterRequestNodePoolsInnerAutoscaler{}
	return &this
}

// NewPostLkeClusterRequestNodePoolsInnerAutoscalerWithDefaults instantiates a new PostLkeClusterRequestNodePoolsInnerAutoscaler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLkeClusterRequestNodePoolsInnerAutoscalerWithDefaults() *PostLkeClusterRequestNodePoolsInnerAutoscaler {
	this := PostLkeClusterRequestNodePoolsInnerAutoscaler{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) SetMax(v int32) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) SetMin(v int32) {
	o.Min = &v
}

func (o PostLkeClusterRequestNodePoolsInnerAutoscaler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLkeClusterRequestNodePoolsInnerAutoscaler) ToMap() (map[string]interface{}, error) {
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

type NullablePostLkeClusterRequestNodePoolsInnerAutoscaler struct {
	value *PostLkeClusterRequestNodePoolsInnerAutoscaler
	isSet bool
}

func (v NullablePostLkeClusterRequestNodePoolsInnerAutoscaler) Get() *PostLkeClusterRequestNodePoolsInnerAutoscaler {
	return v.value
}

func (v *NullablePostLkeClusterRequestNodePoolsInnerAutoscaler) Set(val *PostLkeClusterRequestNodePoolsInnerAutoscaler) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLkeClusterRequestNodePoolsInnerAutoscaler) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLkeClusterRequestNodePoolsInnerAutoscaler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLkeClusterRequestNodePoolsInnerAutoscaler(val *PostLkeClusterRequestNodePoolsInnerAutoscaler) *NullablePostLkeClusterRequestNodePoolsInnerAutoscaler {
	return &NullablePostLkeClusterRequestNodePoolsInnerAutoscaler{value: val, isSet: true}
}

func (v NullablePostLkeClusterRequestNodePoolsInnerAutoscaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLkeClusterRequestNodePoolsInnerAutoscaler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


