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

// checks if the PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLinodeInstanceRequestAllOfInterfacesInnerIpv4{}

// PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 IPv4 addresses configured for this Interface. Only allowed for `vpc` type Interfaces. Returns `null` if no `vpc` Interface is assigned.
type PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 struct {
	// The 1:1 NAT IPv4 address, used to associate a public IPv4 address with the VPC Subnet IPv4 address assigned to this Interface.  - Only allowed for `vpc` type Interfaces. - Returns `null` if no 1:1 NAT is set for a `vpc` type Interface. - Returns an empty string (`\"\"`) for non-`vpc` type Interfaces.  For requests:  - Setting this value to `any` enables the Linode's assigned public IPv4 address on this Interface and establishes a 1:1 NAT between the public IPv4 and VPC Subnet IPv4 addresses. - Setting the value to a specific public IPv4 address that is assigned to the Linode enables a 1:1 NAT between that address and the VPC Subnet IPv4 address. - The public IPv4 address can't be shared with another Linode. - If omitted, set to `null`, or set to an empty string (`\"\"`), no 1:1 NAT is established.  > 📘 > > When creating a new compute instance, you can't set this to a specific IPv4 address. When a new compute instance is created, the network establishes a public IPv4 address for it. Since this address doesn't exist yet, you can't include a custom IPv4 address to change it. Once your compute instance is created, you can [update your configuration profile interface](https://www.linode.com/docs/api/linode-instances/#configuration-profile-interface-update) to change the `nat_1_1` address.
	Nat11 NullableString `json:"nat_1_1,omitempty"`
	// The VPC Subnet IPv4 address for this Interface.  - Only allowed for `vpc` type Interfaces. - Returns an empty string (`\"\"`) for non-`vpc` type Interfaces.  For requests:  - Must not already be actively assigned as an address or within a range to any Linodes. - Must not be the first two or last two addresses in the Subnet IPv4 Range. - If omitted, a valid address within the Subnet IPv4 range is automatically assigned.
	Vpc NullableString `json:"vpc,omitempty"`
}

// NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4 instantiates a new PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4() *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 {
	this := PostLinodeInstanceRequestAllOfInterfacesInnerIpv4{}
	return &this
}

// NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4WithDefaults instantiates a new PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4WithDefaults() *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 {
	this := PostLinodeInstanceRequestAllOfInterfacesInnerIpv4{}
	return &this
}

// GetNat11 returns the Nat11 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) GetNat11() string {
	if o == nil || IsNil(o.Nat11.Get()) {
		var ret string
		return ret
	}
	return *o.Nat11.Get()
}

// GetNat11Ok returns a tuple with the Nat11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) GetNat11Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nat11.Get(), o.Nat11.IsSet()
}

// HasNat11 returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) HasNat11() bool {
	if o != nil && o.Nat11.IsSet() {
		return true
	}

	return false
}

// SetNat11 gets a reference to the given NullableString and assigns it to the Nat11 field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) SetNat11(v string) {
	o.Nat11.Set(&v)
}
// SetNat11Nil sets the value for Nat11 to be an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) SetNat11Nil() {
	o.Nat11.Set(nil)
}

// UnsetNat11 ensures that no value is present for Nat11, not even an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) UnsetNat11() {
	o.Nat11.Unset()
}

// GetVpc returns the Vpc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) GetVpc() string {
	if o == nil || IsNil(o.Vpc.Get()) {
		var ret string
		return ret
	}
	return *o.Vpc.Get()
}

// GetVpcOk returns a tuple with the Vpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) GetVpcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vpc.Get(), o.Vpc.IsSet()
}

// HasVpc returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) HasVpc() bool {
	if o != nil && o.Vpc.IsSet() {
		return true
	}

	return false
}

// SetVpc gets a reference to the given NullableString and assigns it to the Vpc field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) SetVpc(v string) {
	o.Vpc.Set(&v)
}
// SetVpcNil sets the value for Vpc to be an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) SetVpcNil() {
	o.Vpc.Set(nil)
}

// UnsetVpc ensures that no value is present for Vpc, not even an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) UnsetVpc() {
	o.Vpc.Unset()
}

func (o PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Nat11.IsSet() {
		toSerialize["nat_1_1"] = o.Nat11.Get()
	}
	if o.Vpc.IsSet() {
		toSerialize["vpc"] = o.Vpc.Get()
	}
	return toSerialize, nil
}

type NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4 struct {
	value *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4
	isSet bool
}

func (v NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4) Get() *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 {
	return v.value
}

func (v *NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4) Set(val *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4(val *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) *NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4 {
	return &NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4{value: val, isSet: true}
}

func (v NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLinodeInstanceRequestAllOfInterfacesInnerIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


