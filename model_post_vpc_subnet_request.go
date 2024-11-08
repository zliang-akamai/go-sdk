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

// checks if the PostVpcSubnetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostVpcSubnetRequest{}

// PostVpcSubnetRequest VPC Subnet Create request object.
type PostVpcSubnetRequest struct {
	// IPv4 range in CIDR canonical form.  - The range must belong to a private address space as defined in [RFC1918](https://datatracker.ietf.org/doc/html/rfc1918). - Allowed prefix lengths: 1-29. - The range must not overlap with 192.168.128.0/17. - The range must not overlap with other Subnets on the same VPC.
	Ipv4 string `json:"ipv4"`
	// The VPC Subnet's label, for display purposes only.  - Must be unique among the VPC's Subnets. - Can only contain ASCII letters, numbers, and hyphens (`-`). You can't use two consecutive hyphens (`--`).
	Label string `json:"label"`
}

type _PostVpcSubnetRequest PostVpcSubnetRequest

// NewPostVpcSubnetRequest instantiates a new PostVpcSubnetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostVpcSubnetRequest(ipv4 string, label string) *PostVpcSubnetRequest {
	this := PostVpcSubnetRequest{}
	this.Ipv4 = ipv4
	this.Label = label
	return &this
}

// NewPostVpcSubnetRequestWithDefaults instantiates a new PostVpcSubnetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostVpcSubnetRequestWithDefaults() *PostVpcSubnetRequest {
	this := PostVpcSubnetRequest{}
	return &this
}

// GetIpv4 returns the Ipv4 field value
func (o *PostVpcSubnetRequest) GetIpv4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value
// and a boolean to check if the value has been set.
func (o *PostVpcSubnetRequest) GetIpv4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ipv4, true
}

// SetIpv4 sets field value
func (o *PostVpcSubnetRequest) SetIpv4(v string) {
	o.Ipv4 = v
}

// GetLabel returns the Label field value
func (o *PostVpcSubnetRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PostVpcSubnetRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PostVpcSubnetRequest) SetLabel(v string) {
	o.Label = v
}

func (o PostVpcSubnetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostVpcSubnetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ipv4"] = o.Ipv4
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

func (o *PostVpcSubnetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ipv4",
		"label",
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

	varPostVpcSubnetRequest := _PostVpcSubnetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostVpcSubnetRequest)

	if err != nil {
		return err
	}

	*o = PostVpcSubnetRequest(varPostVpcSubnetRequest)

	return err
}

type NullablePostVpcSubnetRequest struct {
	value *PostVpcSubnetRequest
	isSet bool
}

func (v NullablePostVpcSubnetRequest) Get() *PostVpcSubnetRequest {
	return v.value
}

func (v *NullablePostVpcSubnetRequest) Set(val *PostVpcSubnetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostVpcSubnetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostVpcSubnetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostVpcSubnetRequest(val *PostVpcSubnetRequest) *NullablePostVpcSubnetRequest {
	return &NullablePostVpcSubnetRequest{value: val, isSet: true}
}

func (v NullablePostVpcSubnetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostVpcSubnetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


