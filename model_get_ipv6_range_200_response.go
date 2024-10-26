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

// checks if the GetIpv6Range200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIpv6Range200Response{}

// GetIpv6Range200Response An object representing an IPv6 range.
type GetIpv6Range200Response struct {
	// Whether this IPv6 range is shared.
	IsBgp *bool `json:"is_bgp,omitempty"`
	// A list of Linodes targeted by this IPv6 range. Includes Linodes with IP sharing.
	Linodes []int32 `json:"linodes,omitempty"`
	// The prefix length of the address. The total number of addresses that can be assigned from this range is calculated as 2<sup>(128 - prefix length)</sup>.
	Prefix *int32 `json:"prefix,omitempty"`
	// The IPv6 address of this range.
	Range *string `json:"range,omitempty"`
	// The region for this range of IPv6 addresses.
	Region *string `json:"region,omitempty"`
}

// NewGetIpv6Range200Response instantiates a new GetIpv6Range200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIpv6Range200Response() *GetIpv6Range200Response {
	this := GetIpv6Range200Response{}
	return &this
}

// NewGetIpv6Range200ResponseWithDefaults instantiates a new GetIpv6Range200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIpv6Range200ResponseWithDefaults() *GetIpv6Range200Response {
	this := GetIpv6Range200Response{}
	return &this
}

// GetIsBgp returns the IsBgp field value if set, zero value otherwise.
func (o *GetIpv6Range200Response) GetIsBgp() bool {
	if o == nil || IsNil(o.IsBgp) {
		var ret bool
		return ret
	}
	return *o.IsBgp
}

// GetIsBgpOk returns a tuple with the IsBgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Range200Response) GetIsBgpOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBgp) {
		return nil, false
	}
	return o.IsBgp, true
}

// HasIsBgp returns a boolean if a field has been set.
func (o *GetIpv6Range200Response) HasIsBgp() bool {
	if o != nil && !IsNil(o.IsBgp) {
		return true
	}

	return false
}

// SetIsBgp gets a reference to the given bool and assigns it to the IsBgp field.
func (o *GetIpv6Range200Response) SetIsBgp(v bool) {
	o.IsBgp = &v
}

// GetLinodes returns the Linodes field value if set, zero value otherwise.
func (o *GetIpv6Range200Response) GetLinodes() []int32 {
	if o == nil || IsNil(o.Linodes) {
		var ret []int32
		return ret
	}
	return o.Linodes
}

// GetLinodesOk returns a tuple with the Linodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Range200Response) GetLinodesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Linodes) {
		return nil, false
	}
	return o.Linodes, true
}

// HasLinodes returns a boolean if a field has been set.
func (o *GetIpv6Range200Response) HasLinodes() bool {
	if o != nil && !IsNil(o.Linodes) {
		return true
	}

	return false
}

// SetLinodes gets a reference to the given []int32 and assigns it to the Linodes field.
func (o *GetIpv6Range200Response) SetLinodes(v []int32) {
	o.Linodes = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *GetIpv6Range200Response) GetPrefix() int32 {
	if o == nil || IsNil(o.Prefix) {
		var ret int32
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Range200Response) GetPrefixOk() (*int32, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *GetIpv6Range200Response) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given int32 and assigns it to the Prefix field.
func (o *GetIpv6Range200Response) SetPrefix(v int32) {
	o.Prefix = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *GetIpv6Range200Response) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Range200Response) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *GetIpv6Range200Response) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *GetIpv6Range200Response) SetRange(v string) {
	o.Range = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetIpv6Range200Response) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Range200Response) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetIpv6Range200Response) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetIpv6Range200Response) SetRegion(v string) {
	o.Region = &v
}

func (o GetIpv6Range200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIpv6Range200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsBgp) {
		toSerialize["is_bgp"] = o.IsBgp
	}
	if !IsNil(o.Linodes) {
		toSerialize["linodes"] = o.Linodes
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableGetIpv6Range200Response struct {
	value *GetIpv6Range200Response
	isSet bool
}

func (v NullableGetIpv6Range200Response) Get() *GetIpv6Range200Response {
	return v.value
}

func (v *NullableGetIpv6Range200Response) Set(val *GetIpv6Range200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIpv6Range200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIpv6Range200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIpv6Range200Response(val *GetIpv6Range200Response) *NullableGetIpv6Range200Response {
	return &NullableGetIpv6Range200Response{value: val, isSet: true}
}

func (v NullableGetIpv6Range200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIpv6Range200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


