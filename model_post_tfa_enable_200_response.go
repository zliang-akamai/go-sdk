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

// checks if the PostTfaEnable200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTfaEnable200Response{}

// PostTfaEnable200Response struct for PostTfaEnable200Response
type PostTfaEnable200Response struct {
	// When this Two Factor secret expires.
	Expiry *time.Time `json:"expiry,omitempty"`
	// Your Two Factor secret. This is used to generate time-based two factor codes required for logging in. Doing this will be required to confirm TFA an actually enable it.
	Secret *string `json:"secret,omitempty"`
}

// NewPostTfaEnable200Response instantiates a new PostTfaEnable200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTfaEnable200Response() *PostTfaEnable200Response {
	this := PostTfaEnable200Response{}
	return &this
}

// NewPostTfaEnable200ResponseWithDefaults instantiates a new PostTfaEnable200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTfaEnable200ResponseWithDefaults() *PostTfaEnable200Response {
	this := PostTfaEnable200Response{}
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *PostTfaEnable200Response) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTfaEnable200Response) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *PostTfaEnable200Response) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *PostTfaEnable200Response) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PostTfaEnable200Response) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTfaEnable200Response) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PostTfaEnable200Response) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PostTfaEnable200Response) SetSecret(v string) {
	o.Secret = &v
}

func (o PostTfaEnable200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTfaEnable200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullablePostTfaEnable200Response struct {
	value *PostTfaEnable200Response
	isSet bool
}

func (v NullablePostTfaEnable200Response) Get() *PostTfaEnable200Response {
	return v.value
}

func (v *NullablePostTfaEnable200Response) Set(val *PostTfaEnable200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTfaEnable200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTfaEnable200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTfaEnable200Response(val *PostTfaEnable200Response) *NullablePostTfaEnable200Response {
	return &NullablePostTfaEnable200Response{value: val, isSet: true}
}

func (v NullablePostTfaEnable200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTfaEnable200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


