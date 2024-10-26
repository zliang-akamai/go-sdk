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

// checks if the GetNodeBalancerStats200ResponseDataTraffic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodeBalancerStats200ResponseDataTraffic{}

// GetNodeBalancerStats200ResponseDataTraffic Traffic statistics for this NodeBalancer.
type GetNodeBalancerStats200ResponseDataTraffic struct {
	// An array of key/value pairs representing unix timestamp and reading for inbound traffic.
	In []float32 `json:"in,omitempty"`
	// An array of key/value pairs representing unix timestamp and reading for outbound traffic.
	Out []float32 `json:"out,omitempty"`
}

// NewGetNodeBalancerStats200ResponseDataTraffic instantiates a new GetNodeBalancerStats200ResponseDataTraffic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodeBalancerStats200ResponseDataTraffic() *GetNodeBalancerStats200ResponseDataTraffic {
	this := GetNodeBalancerStats200ResponseDataTraffic{}
	return &this
}

// NewGetNodeBalancerStats200ResponseDataTrafficWithDefaults instantiates a new GetNodeBalancerStats200ResponseDataTraffic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodeBalancerStats200ResponseDataTrafficWithDefaults() *GetNodeBalancerStats200ResponseDataTraffic {
	this := GetNodeBalancerStats200ResponseDataTraffic{}
	return &this
}

// GetIn returns the In field value if set, zero value otherwise.
func (o *GetNodeBalancerStats200ResponseDataTraffic) GetIn() []float32 {
	if o == nil || IsNil(o.In) {
		var ret []float32
		return ret
	}
	return o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerStats200ResponseDataTraffic) GetInOk() ([]float32, bool) {
	if o == nil || IsNil(o.In) {
		return nil, false
	}
	return o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *GetNodeBalancerStats200ResponseDataTraffic) HasIn() bool {
	if o != nil && !IsNil(o.In) {
		return true
	}

	return false
}

// SetIn gets a reference to the given []float32 and assigns it to the In field.
func (o *GetNodeBalancerStats200ResponseDataTraffic) SetIn(v []float32) {
	o.In = v
}

// GetOut returns the Out field value if set, zero value otherwise.
func (o *GetNodeBalancerStats200ResponseDataTraffic) GetOut() []float32 {
	if o == nil || IsNil(o.Out) {
		var ret []float32
		return ret
	}
	return o.Out
}

// GetOutOk returns a tuple with the Out field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerStats200ResponseDataTraffic) GetOutOk() ([]float32, bool) {
	if o == nil || IsNil(o.Out) {
		return nil, false
	}
	return o.Out, true
}

// HasOut returns a boolean if a field has been set.
func (o *GetNodeBalancerStats200ResponseDataTraffic) HasOut() bool {
	if o != nil && !IsNil(o.Out) {
		return true
	}

	return false
}

// SetOut gets a reference to the given []float32 and assigns it to the Out field.
func (o *GetNodeBalancerStats200ResponseDataTraffic) SetOut(v []float32) {
	o.Out = v
}

func (o GetNodeBalancerStats200ResponseDataTraffic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodeBalancerStats200ResponseDataTraffic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.In) {
		toSerialize["in"] = o.In
	}
	if !IsNil(o.Out) {
		toSerialize["out"] = o.Out
	}
	return toSerialize, nil
}

type NullableGetNodeBalancerStats200ResponseDataTraffic struct {
	value *GetNodeBalancerStats200ResponseDataTraffic
	isSet bool
}

func (v NullableGetNodeBalancerStats200ResponseDataTraffic) Get() *GetNodeBalancerStats200ResponseDataTraffic {
	return v.value
}

func (v *NullableGetNodeBalancerStats200ResponseDataTraffic) Set(val *GetNodeBalancerStats200ResponseDataTraffic) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodeBalancerStats200ResponseDataTraffic) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodeBalancerStats200ResponseDataTraffic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodeBalancerStats200ResponseDataTraffic(val *GetNodeBalancerStats200ResponseDataTraffic) *NullableGetNodeBalancerStats200ResponseDataTraffic {
	return &NullableGetNodeBalancerStats200ResponseDataTraffic{value: val, isSet: true}
}

func (v NullableGetNodeBalancerStats200ResponseDataTraffic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodeBalancerStats200ResponseDataTraffic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

