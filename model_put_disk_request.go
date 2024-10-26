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

// checks if the PutDiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutDiskRequest{}

// PutDiskRequest struct for PutDiskRequest
type PutDiskRequest struct {
	// The name of the disk. This is for display purposes only.
	Label *string `json:"label,omitempty"`
}

// NewPutDiskRequest instantiates a new PutDiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutDiskRequest() *PutDiskRequest {
	this := PutDiskRequest{}
	return &this
}

// NewPutDiskRequestWithDefaults instantiates a new PutDiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutDiskRequestWithDefaults() *PutDiskRequest {
	this := PutDiskRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PutDiskRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutDiskRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PutDiskRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PutDiskRequest) SetLabel(v string) {
	o.Label = &v
}

func (o PutDiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutDiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullablePutDiskRequest struct {
	value *PutDiskRequest
	isSet bool
}

func (v NullablePutDiskRequest) Get() *PutDiskRequest {
	return v.value
}

func (v *NullablePutDiskRequest) Set(val *PutDiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutDiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutDiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutDiskRequest(val *PutDiskRequest) *NullablePutDiskRequest {
	return &NullablePutDiskRequest{value: val, isSet: true}
}

func (v NullablePutDiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutDiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


