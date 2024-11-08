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

// checks if the GetIps200ResponseAllOfDataInnerVpcNat11 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIps200ResponseAllOfDataInnerVpcNat11{}

// GetIps200ResponseAllOfDataInnerVpcNat11 IPv4 address configured as a 1:1 NAT for this Interface. Empty if no address is configured as a 1:1 NAT.  __Note__. Only allowed for `vpc` type Interfaces.
type GetIps200ResponseAllOfDataInnerVpcNat11 struct {
	// The IPv4 address that is configured as a 1:1 NAT for this VPC interface.
	Address *string `json:"address,omitempty"`
	// The `id` of the VPC Subnet for this Interface.
	SubnetId *int32 `json:"subnet_id,omitempty"`
	// The `id` of the VPC configured for this Interface.
	VpcId *int32 `json:"vpc_id,omitempty"`
}

// NewGetIps200ResponseAllOfDataInnerVpcNat11 instantiates a new GetIps200ResponseAllOfDataInnerVpcNat11 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIps200ResponseAllOfDataInnerVpcNat11() *GetIps200ResponseAllOfDataInnerVpcNat11 {
	this := GetIps200ResponseAllOfDataInnerVpcNat11{}
	return &this
}

// NewGetIps200ResponseAllOfDataInnerVpcNat11WithDefaults instantiates a new GetIps200ResponseAllOfDataInnerVpcNat11 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIps200ResponseAllOfDataInnerVpcNat11WithDefaults() *GetIps200ResponseAllOfDataInnerVpcNat11 {
	this := GetIps200ResponseAllOfDataInnerVpcNat11{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) SetAddress(v string) {
	o.Address = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) GetSubnetId() int32 {
	if o == nil || IsNil(o.SubnetId) {
		var ret int32
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) GetSubnetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SubnetId) {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) HasSubnetId() bool {
	if o != nil && !IsNil(o.SubnetId) {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given int32 and assigns it to the SubnetId field.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) SetSubnetId(v int32) {
	o.SubnetId = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) GetVpcId() int32 {
	if o == nil || IsNil(o.VpcId) {
		var ret int32
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) GetVpcIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given int32 and assigns it to the VpcId field.
func (o *GetIps200ResponseAllOfDataInnerVpcNat11) SetVpcId(v int32) {
	o.VpcId = &v
}

func (o GetIps200ResponseAllOfDataInnerVpcNat11) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIps200ResponseAllOfDataInnerVpcNat11) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.SubnetId) {
		toSerialize["subnet_id"] = o.SubnetId
	}
	if !IsNil(o.VpcId) {
		toSerialize["vpc_id"] = o.VpcId
	}
	return toSerialize, nil
}

type NullableGetIps200ResponseAllOfDataInnerVpcNat11 struct {
	value *GetIps200ResponseAllOfDataInnerVpcNat11
	isSet bool
}

func (v NullableGetIps200ResponseAllOfDataInnerVpcNat11) Get() *GetIps200ResponseAllOfDataInnerVpcNat11 {
	return v.value
}

func (v *NullableGetIps200ResponseAllOfDataInnerVpcNat11) Set(val *GetIps200ResponseAllOfDataInnerVpcNat11) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIps200ResponseAllOfDataInnerVpcNat11) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIps200ResponseAllOfDataInnerVpcNat11) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIps200ResponseAllOfDataInnerVpcNat11(val *GetIps200ResponseAllOfDataInnerVpcNat11) *NullableGetIps200ResponseAllOfDataInnerVpcNat11 {
	return &NullableGetIps200ResponseAllOfDataInnerVpcNat11{value: val, isSet: true}
}

func (v NullableGetIps200ResponseAllOfDataInnerVpcNat11) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIps200ResponseAllOfDataInnerVpcNat11) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


