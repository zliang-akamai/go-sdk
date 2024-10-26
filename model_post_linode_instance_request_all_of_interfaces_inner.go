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
	"bytes"
	"fmt"
)

// checks if the PostLinodeInstanceRequestAllOfInterfacesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLinodeInstanceRequestAllOfInterfacesInner{}

// PostLinodeInstanceRequestAllOfInterfacesInner The Network Interface to apply to this Linode's configuration profile.
type PostLinodeInstanceRequestAllOfInterfacesInner struct {
	// Returns `true` if the Interface is in use, meaning that Compute Instance has been booted using the Configuration Profile to which the Interface belongs. Otherwise returns `false`.
	Active *bool `json:"active,omitempty"`
	// The unique ID representing this Interface.
	Id *int32 `json:"id,omitempty"`
	// An array of IPv4 CIDR VPC Subnet ranges that are routed to this Interface.  - Array items are only allowed for `vpc` type Interfaces. - This must be empty for non-`vpc` type Interfaces.  For requests:  - Addresses in submitted ranges must not already be actively assigned. - Submitting values replaces any existing values. - Submitting an empty array removes any existing values. - Omitting this property results in no change to existing values.
	IpRanges []string `json:"ip_ranges,omitempty"`
	// This Network Interface's private IP address in Classless Inter-Domain Routing (CIDR) notation.  For `vlan` purpose Interfaces:  - Must be unique among the Linode's Interfaces to avoid conflicting addresses. - Should be unique among devices attached to the VLAN to avoid conflict. - The Linode is configured to use this address for the associated Interface upon reboot if Network Helper is enabled. If Network Helper is disabled, the address can be enabled with [manual static IP configuration](https://www.linode.com/docs/guides/manual-network-configuration/).  For `public` purpose Interfaces:  - In requests, must be an empty string (`\"\"`) or `null` if included. - In responses, always returns `null`.  For `vpc` purpose Interfaces:  - In requests, must be an empty string (`\"\"`) or `null` if included. - In responses, always returns `null`.
	IpamAddress NullableString `json:"ipam_address,omitempty"`
	Ipv4 *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 `json:"ipv4,omitempty"`
	// The name of this Interface.  For `vlan` purpose Interfaces:  - Required. - Must be unique among the Linode's Interfaces (a Linode cannot be attached to the same VLAN multiple times). - Can only contain ASCII letters, numbers, and hyphens (`-`). You can't use two consecutive hyphens (`--`). - If the VLAN label is new, a VLAN is created. Up to 10 VLANs can be created in each data center region. To view your active VLANs, run the [List VLANs](https://techdocs.akamai.com/linode-api/reference/get-vlans) operation.  For `public` purpose Interfaces:  - In requests, must be an empty string (`\"\"`) or `null` if included. - In responses, always returns `null`.  For `vpc` purpose Interfaces:  - In requests, must be an empty string (`\"\"`) or `null` if included. - In responses, always returns `null`.
	Label NullableString `json:"label,omitempty" validate:"regexp=[a-zA-Z0-9-]+"`
	// The primary Interface is configured as the default route to the Linode.  Each Configuration Profile can have up to one `\"primary\": true` Interface at a time.  Must be `false` for `vlan` type Interfaces.  If no Interface is configured as the primary, the first non-`vlan` type Interface in the `interfaces` array is automatically treated as the primary Interface.
	Primary *bool `json:"primary,omitempty"`
	// The type of Interface.  - `public`   - Only one `public` Interface per Linode can be defined.   - The Linode's default public IPv4 address is assigned to the `public` Interface.   - A Linode must have a public Interface in the first/eth0 position to be reachable via the public internet upon boot without additional system configuration. If no `public` Interface is configured, the Linode is not directly reachable via the public internet. In this case, access can only be established via [LISH](https://www.linode.com/docs/products/compute/compute-instances/guides/lish/) or other Linodes connected to the same VLAN or VPC.  - `vlan`   - Configuring a `vlan` purpose Interface attaches this Linode to the VLAN with the specified `label`.   - The Linode is configured to use the specified `ipam_address`, if any.  - `vpc`   - Configuring a `vpc` purpose Interface attaches this Linode to the existing VPC Subnet with the specified `subnet_id`.   - When the Interface is activated, the Linode is configured to use an IP address from the range in the assigned VPC Subnet. See `ipv4.vpc` for more information.
	Purpose string `json:"purpose"`
	// The `id` of the VPC Subnet for this Interface.  In requests, this value is used to assign a Linode to a VPC Subnet.  - Required for `vpc` type Interfaces. - Returns `null` for non-`vpc` type Interfaces. - Once a VPC Subnet is assigned to an Interface, it cannot be updated. - The Linode must be rebooted with the Interface's Configuration Profile to complete assignment to a VPC Subnet.
	SubnetId NullableInt32 `json:"subnet_id,omitempty"`
	// The `id` of the VPC configured for this Interface. Returns `null` for non-`vpc` type Interfaces.
	VpcId NullableInt32 `json:"vpc_id,omitempty"`
}

type _PostLinodeInstanceRequestAllOfInterfacesInner PostLinodeInstanceRequestAllOfInterfacesInner

// NewPostLinodeInstanceRequestAllOfInterfacesInner instantiates a new PostLinodeInstanceRequestAllOfInterfacesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLinodeInstanceRequestAllOfInterfacesInner(purpose string) *PostLinodeInstanceRequestAllOfInterfacesInner {
	this := PostLinodeInstanceRequestAllOfInterfacesInner{}
	this.Purpose = purpose
	return &this
}

// NewPostLinodeInstanceRequestAllOfInterfacesInnerWithDefaults instantiates a new PostLinodeInstanceRequestAllOfInterfacesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLinodeInstanceRequestAllOfInterfacesInnerWithDefaults() *PostLinodeInstanceRequestAllOfInterfacesInner {
	this := PostLinodeInstanceRequestAllOfInterfacesInner{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetActive(v bool) {
	o.Active = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetId(v int32) {
	o.Id = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpRanges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpRanges) {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasIpRanges() bool {
	if o != nil && !IsNil(o.IpRanges) {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetIpamAddress returns the IpamAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpamAddress() string {
	if o == nil || IsNil(o.IpamAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpamAddress.Get()
}

// GetIpamAddressOk returns a tuple with the IpamAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpamAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpamAddress.Get(), o.IpamAddress.IsSet()
}

// HasIpamAddress returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasIpamAddress() bool {
	if o != nil && o.IpamAddress.IsSet() {
		return true
	}

	return false
}

// SetIpamAddress gets a reference to the given NullableString and assigns it to the IpamAddress field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpamAddress(v string) {
	o.IpamAddress.Set(&v)
}
// SetIpamAddressNil sets the value for IpamAddress to be an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpamAddressNil() {
	o.IpamAddress.Set(nil)
}

// UnsetIpamAddress ensures that no value is present for IpamAddress, not even an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetIpamAddress() {
	o.IpamAddress.Unset()
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpv4() PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 {
	if o == nil || IsNil(o.Ipv4) {
		var ret PostLinodeInstanceRequestAllOfInterfacesInnerIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpv4Ok() (*PostLinodeInstanceRequestAllOfInterfacesInnerIpv4, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 and assigns it to the Ipv4 field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpv4(v PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) {
	o.Ipv4 = &v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetLabel() {
	o.Label.Unset()
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetPrimary(v bool) {
	o.Primary = &v
}

// GetPurpose returns the Purpose field value
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetPurpose(v string) {
	o.Purpose = v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetSubnetId() int32 {
	if o == nil || IsNil(o.SubnetId.Get()) {
		var ret int32
		return ret
	}
	return *o.SubnetId.Get()
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetSubnetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubnetId.Get(), o.SubnetId.IsSet()
}

// HasSubnetId returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasSubnetId() bool {
	if o != nil && o.SubnetId.IsSet() {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given NullableInt32 and assigns it to the SubnetId field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetSubnetId(v int32) {
	o.SubnetId.Set(&v)
}
// SetSubnetIdNil sets the value for SubnetId to be an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetSubnetIdNil() {
	o.SubnetId.Set(nil)
}

// UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetSubnetId() {
	o.SubnetId.Unset()
}

// GetVpcId returns the VpcId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetVpcId() int32 {
	if o == nil || IsNil(o.VpcId.Get()) {
		var ret int32
		return ret
	}
	return *o.VpcId.Get()
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetVpcIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VpcId.Get(), o.VpcId.IsSet()
}

// HasVpcId returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasVpcId() bool {
	if o != nil && o.VpcId.IsSet() {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given NullableInt32 and assigns it to the VpcId field.
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetVpcId(v int32) {
	o.VpcId.Set(&v)
}
// SetVpcIdNil sets the value for VpcId to be an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetVpcIdNil() {
	o.VpcId.Set(nil)
}

// UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetVpcId() {
	o.VpcId.Unset()
}

func (o PostLinodeInstanceRequestAllOfInterfacesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLinodeInstanceRequestAllOfInterfacesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.IpRanges != nil {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if o.IpamAddress.IsSet() {
		toSerialize["ipam_address"] = o.IpamAddress.Get()
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	toSerialize["purpose"] = o.Purpose
	if o.SubnetId.IsSet() {
		toSerialize["subnet_id"] = o.SubnetId.Get()
	}
	if o.VpcId.IsSet() {
		toSerialize["vpc_id"] = o.VpcId.Get()
	}
	return toSerialize, nil
}

func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"purpose",
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

	varPostLinodeInstanceRequestAllOfInterfacesInner := _PostLinodeInstanceRequestAllOfInterfacesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostLinodeInstanceRequestAllOfInterfacesInner)

	if err != nil {
		return err
	}

	*o = PostLinodeInstanceRequestAllOfInterfacesInner(varPostLinodeInstanceRequestAllOfInterfacesInner)

	return err
}

type NullablePostLinodeInstanceRequestAllOfInterfacesInner struct {
	value *PostLinodeInstanceRequestAllOfInterfacesInner
	isSet bool
}

func (v NullablePostLinodeInstanceRequestAllOfInterfacesInner) Get() *PostLinodeInstanceRequestAllOfInterfacesInner {
	return v.value
}

func (v *NullablePostLinodeInstanceRequestAllOfInterfacesInner) Set(val *PostLinodeInstanceRequestAllOfInterfacesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLinodeInstanceRequestAllOfInterfacesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLinodeInstanceRequestAllOfInterfacesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLinodeInstanceRequestAllOfInterfacesInner(val *PostLinodeInstanceRequestAllOfInterfacesInner) *NullablePostLinodeInstanceRequestAllOfInterfacesInner {
	return &NullablePostLinodeInstanceRequestAllOfInterfacesInner{value: val, isSet: true}
}

func (v NullablePostLinodeInstanceRequestAllOfInterfacesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLinodeInstanceRequestAllOfInterfacesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


