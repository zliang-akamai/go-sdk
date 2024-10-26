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

// checks if the PutVpcSubnetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutVpcSubnetRequest{}

// PutVpcSubnetRequest A VPC Subnet Update request object.
type PutVpcSubnetRequest struct {
	// The VPC Subnet's label, for display purposes only.  - Must be unique among the VPC's Subnets. - Can only contain ASCII letters, numbers, and hyphens (`-`). You can't use two consecutive hyphens (`--`).
	Label *string `json:"label,omitempty"`
}

// NewPutVpcSubnetRequest instantiates a new PutVpcSubnetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutVpcSubnetRequest() *PutVpcSubnetRequest {
	this := PutVpcSubnetRequest{}
	return &this
}

// NewPutVpcSubnetRequestWithDefaults instantiates a new PutVpcSubnetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutVpcSubnetRequestWithDefaults() *PutVpcSubnetRequest {
	this := PutVpcSubnetRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PutVpcSubnetRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutVpcSubnetRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PutVpcSubnetRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PutVpcSubnetRequest) SetLabel(v string) {
	o.Label = &v
}

func (o PutVpcSubnetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutVpcSubnetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullablePutVpcSubnetRequest struct {
	value *PutVpcSubnetRequest
	isSet bool
}

func (v NullablePutVpcSubnetRequest) Get() *PutVpcSubnetRequest {
	return v.value
}

func (v *NullablePutVpcSubnetRequest) Set(val *PutVpcSubnetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutVpcSubnetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutVpcSubnetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutVpcSubnetRequest(val *PutVpcSubnetRequest) *NullablePutVpcSubnetRequest {
	return &NullablePutVpcSubnetRequest{value: val, isSet: true}
}

func (v NullablePutVpcSubnetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutVpcSubnetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


