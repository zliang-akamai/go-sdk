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

// checks if the PutSshKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutSshKeyRequest{}

// PutSshKeyRequest struct for PutSshKeyRequest
type PutSshKeyRequest struct {
	// A label for the SSH Key.
	Label *string `json:"label,omitempty"`
}

// NewPutSshKeyRequest instantiates a new PutSshKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutSshKeyRequest() *PutSshKeyRequest {
	this := PutSshKeyRequest{}
	return &this
}

// NewPutSshKeyRequestWithDefaults instantiates a new PutSshKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutSshKeyRequestWithDefaults() *PutSshKeyRequest {
	this := PutSshKeyRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PutSshKeyRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutSshKeyRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PutSshKeyRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PutSshKeyRequest) SetLabel(v string) {
	o.Label = &v
}

func (o PutSshKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutSshKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullablePutSshKeyRequest struct {
	value *PutSshKeyRequest
	isSet bool
}

func (v NullablePutSshKeyRequest) Get() *PutSshKeyRequest {
	return v.value
}

func (v *NullablePutSshKeyRequest) Set(val *PutSshKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutSshKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutSshKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutSshKeyRequest(val *PutSshKeyRequest) *NullablePutSshKeyRequest {
	return &NullablePutSshKeyRequest{value: val, isSet: true}
}

func (v NullablePutSshKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutSshKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


