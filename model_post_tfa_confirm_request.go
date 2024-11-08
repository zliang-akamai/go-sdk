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

// checks if the PostTfaConfirmRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTfaConfirmRequest{}

// PostTfaConfirmRequest struct for PostTfaConfirmRequest
type PostTfaConfirmRequest struct {
	// The Two Factor code you generated with your Two Factor secret. These codes are time-based, so be sure it is current.
	TfaCode *string `json:"tfa_code,omitempty"`
}

// NewPostTfaConfirmRequest instantiates a new PostTfaConfirmRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTfaConfirmRequest() *PostTfaConfirmRequest {
	this := PostTfaConfirmRequest{}
	return &this
}

// NewPostTfaConfirmRequestWithDefaults instantiates a new PostTfaConfirmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTfaConfirmRequestWithDefaults() *PostTfaConfirmRequest {
	this := PostTfaConfirmRequest{}
	return &this
}

// GetTfaCode returns the TfaCode field value if set, zero value otherwise.
func (o *PostTfaConfirmRequest) GetTfaCode() string {
	if o == nil || IsNil(o.TfaCode) {
		var ret string
		return ret
	}
	return *o.TfaCode
}

// GetTfaCodeOk returns a tuple with the TfaCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTfaConfirmRequest) GetTfaCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TfaCode) {
		return nil, false
	}
	return o.TfaCode, true
}

// HasTfaCode returns a boolean if a field has been set.
func (o *PostTfaConfirmRequest) HasTfaCode() bool {
	if o != nil && !IsNil(o.TfaCode) {
		return true
	}

	return false
}

// SetTfaCode gets a reference to the given string and assigns it to the TfaCode field.
func (o *PostTfaConfirmRequest) SetTfaCode(v string) {
	o.TfaCode = &v
}

func (o PostTfaConfirmRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTfaConfirmRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TfaCode) {
		toSerialize["tfa_code"] = o.TfaCode
	}
	return toSerialize, nil
}

type NullablePostTfaConfirmRequest struct {
	value *PostTfaConfirmRequest
	isSet bool
}

func (v NullablePostTfaConfirmRequest) Get() *PostTfaConfirmRequest {
	return v.value
}

func (v *NullablePostTfaConfirmRequest) Set(val *PostTfaConfirmRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTfaConfirmRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTfaConfirmRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTfaConfirmRequest(val *PostTfaConfirmRequest) *NullablePostTfaConfirmRequest {
	return &NullablePostTfaConfirmRequest{value: val, isSet: true}
}

func (v NullablePostTfaConfirmRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTfaConfirmRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


