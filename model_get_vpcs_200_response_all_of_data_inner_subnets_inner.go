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
	"time"
)

// checks if the GetVpcs200ResponseAllOfDataInnerSubnetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVpcs200ResponseAllOfDataInnerSubnetsInner{}

// GetVpcs200ResponseAllOfDataInnerSubnetsInner An object describing a VPC Subnet.
type GetVpcs200ResponseAllOfDataInnerSubnetsInner struct {
	// The date-time of VPC Subnet creation.
	Created *time.Time `json:"created,omitempty"`
	// The unique ID of the VPC Subnet.
	Id *int32 `json:"id,omitempty"`
	// IPv4 range in CIDR canonical form.  - The range must belong to a private address space as defined in [RFC1918](https://datatracker.ietf.org/doc/html/rfc1918). - Allowed prefix lengths: 1-29. - The range must not overlap with 192.168.128.0/17. - The range must not overlap with other Subnets on the same VPC.
	Ipv4 *string `json:"ipv4,omitempty"`
	// The VPC Subnet's label, for display purposes only.  - Must be unique among the VPC's Subnets. - Can only contain ASCII letters, numbers, and hyphens (`-`). You can't use two consecutive hyphens (`--`).
	Label *string `json:"label,omitempty"`
	// An array of Linode IDs assigned to the VPC Subnet.  A Linode is assigned to a VPC Subnet if it has a Configuration Profile with a `vpc` purpose interface with the subnet's `subnet_id`. Even if the Configuration Profile is not active, meaning the Linode does not have access to the Subnet, the Linode still appears in this array.
	Linodes []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner `json:"linodes,omitempty"`
	// The date-time of the most recent VPC Subnet update.
	Updated NullableTime `json:"updated,omitempty"`
}

// NewGetVpcs200ResponseAllOfDataInnerSubnetsInner instantiates a new GetVpcs200ResponseAllOfDataInnerSubnetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVpcs200ResponseAllOfDataInnerSubnetsInner() *GetVpcs200ResponseAllOfDataInnerSubnetsInner {
	this := GetVpcs200ResponseAllOfDataInnerSubnetsInner{}
	return &this
}

// NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerWithDefaults instantiates a new GetVpcs200ResponseAllOfDataInnerSubnetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerWithDefaults() *GetVpcs200ResponseAllOfDataInnerSubnetsInner {
	this := GetVpcs200ResponseAllOfDataInnerSubnetsInner{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetId(v int32) {
	o.Id = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetIpv4() string {
	if o == nil || IsNil(o.Ipv4) {
		var ret string
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetIpv4Ok() (*string, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given string and assigns it to the Ipv4 field.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetIpv4(v string) {
	o.Ipv4 = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetLabel(v string) {
	o.Label = &v
}

// GetLinodes returns the Linodes field value if set, zero value otherwise.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetLinodes() []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner {
	if o == nil || IsNil(o.Linodes) {
		var ret []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner
		return ret
	}
	return o.Linodes
}

// GetLinodesOk returns a tuple with the Linodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetLinodesOk() ([]GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner, bool) {
	if o == nil || IsNil(o.Linodes) {
		return nil, false
	}
	return o.Linodes, true
}

// HasLinodes returns a boolean if a field has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasLinodes() bool {
	if o != nil && !IsNil(o.Linodes) {
		return true
	}

	return false
}

// SetLinodes gets a reference to the given []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner and assigns it to the Linodes field.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetLinodes(v []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) {
	o.Linodes = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// HasUpdated returns a boolean if a field has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasUpdated() bool {
	if o != nil && o.Updated.IsSet() {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given NullableTime and assigns it to the Updated field.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetUpdated(v time.Time) {
	o.Updated.Set(&v)
}
// SetUpdatedNil sets the value for Updated to be an explicit nil
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetUpdatedNil() {
	o.Updated.Set(nil)
}

// UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) UnsetUpdated() {
	o.Updated.Unset()
}

func (o GetVpcs200ResponseAllOfDataInnerSubnetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVpcs200ResponseAllOfDataInnerSubnetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Linodes) {
		toSerialize["linodes"] = o.Linodes
	}
	if o.Updated.IsSet() {
		toSerialize["updated"] = o.Updated.Get()
	}
	return toSerialize, nil
}

type NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner struct {
	value *GetVpcs200ResponseAllOfDataInnerSubnetsInner
	isSet bool
}

func (v NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner) Get() *GetVpcs200ResponseAllOfDataInnerSubnetsInner {
	return v.value
}

func (v *NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner) Set(val *GetVpcs200ResponseAllOfDataInnerSubnetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVpcs200ResponseAllOfDataInnerSubnetsInner(val *GetVpcs200ResponseAllOfDataInnerSubnetsInner) *NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner {
	return &NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner{value: val, isSet: true}
}

func (v NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVpcs200ResponseAllOfDataInnerSubnetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


