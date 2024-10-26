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

// checks if the GetLinodeIps200ResponseIpv6GlobalInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLinodeIps200ResponseIpv6GlobalInner{}

// GetLinodeIps200ResponseIpv6GlobalInner An object representing an IPv6 range.
type GetLinodeIps200ResponseIpv6GlobalInner struct {
	// The prefix length of the address. The total number of addresses that can be assigned from this range is calculated as 2<sup>(128 - prefix length)</sup>.
	Prefix *int32 `json:"prefix,omitempty"`
	// The IPv6 address of this range.
	Range *string `json:"range,omitempty"`
	// The region for this range of IPv6 addresses.
	Region *string `json:"region,omitempty"`
	// The IPv6 SLAAC address.
	RouteTarget *string `json:"route_target,omitempty"`
}

// NewGetLinodeIps200ResponseIpv6GlobalInner instantiates a new GetLinodeIps200ResponseIpv6GlobalInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLinodeIps200ResponseIpv6GlobalInner() *GetLinodeIps200ResponseIpv6GlobalInner {
	this := GetLinodeIps200ResponseIpv6GlobalInner{}
	return &this
}

// NewGetLinodeIps200ResponseIpv6GlobalInnerWithDefaults instantiates a new GetLinodeIps200ResponseIpv6GlobalInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLinodeIps200ResponseIpv6GlobalInnerWithDefaults() *GetLinodeIps200ResponseIpv6GlobalInner {
	this := GetLinodeIps200ResponseIpv6GlobalInner{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetPrefix() int32 {
	if o == nil || IsNil(o.Prefix) {
		var ret int32
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetPrefixOk() (*int32, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given int32 and assigns it to the Prefix field.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) SetPrefix(v int32) {
	o.Prefix = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) SetRange(v string) {
	o.Range = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) SetRegion(v string) {
	o.Region = &v
}

// GetRouteTarget returns the RouteTarget field value if set, zero value otherwise.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRouteTarget() string {
	if o == nil || IsNil(o.RouteTarget) {
		var ret string
		return ret
	}
	return *o.RouteTarget
}

// GetRouteTargetOk returns a tuple with the RouteTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRouteTargetOk() (*string, bool) {
	if o == nil || IsNil(o.RouteTarget) {
		return nil, false
	}
	return o.RouteTarget, true
}

// HasRouteTarget returns a boolean if a field has been set.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) HasRouteTarget() bool {
	if o != nil && !IsNil(o.RouteTarget) {
		return true
	}

	return false
}

// SetRouteTarget gets a reference to the given string and assigns it to the RouteTarget field.
func (o *GetLinodeIps200ResponseIpv6GlobalInner) SetRouteTarget(v string) {
	o.RouteTarget = &v
}

func (o GetLinodeIps200ResponseIpv6GlobalInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLinodeIps200ResponseIpv6GlobalInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.RouteTarget) {
		toSerialize["route_target"] = o.RouteTarget
	}
	return toSerialize, nil
}

type NullableGetLinodeIps200ResponseIpv6GlobalInner struct {
	value *GetLinodeIps200ResponseIpv6GlobalInner
	isSet bool
}

func (v NullableGetLinodeIps200ResponseIpv6GlobalInner) Get() *GetLinodeIps200ResponseIpv6GlobalInner {
	return v.value
}

func (v *NullableGetLinodeIps200ResponseIpv6GlobalInner) Set(val *GetLinodeIps200ResponseIpv6GlobalInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLinodeIps200ResponseIpv6GlobalInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLinodeIps200ResponseIpv6GlobalInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLinodeIps200ResponseIpv6GlobalInner(val *GetLinodeIps200ResponseIpv6GlobalInner) *NullableGetLinodeIps200ResponseIpv6GlobalInner {
	return &NullableGetLinodeIps200ResponseIpv6GlobalInner{value: val, isSet: true}
}

func (v NullableGetLinodeIps200ResponseIpv6GlobalInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLinodeIps200ResponseIpv6GlobalInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

