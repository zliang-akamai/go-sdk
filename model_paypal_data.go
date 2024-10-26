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

// checks if the PaypalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaypalData{}

// PaypalData PayPal information.
type PaypalData struct {
	// The email address associated with your PayPal account.
	Email *string `json:"email,omitempty"`
	// PayPal Merchant ID associated with your PayPal account.
	PaypalId *string `json:"paypal_id,omitempty"`
}

// NewPaypalData instantiates a new PaypalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalData() *PaypalData {
	this := PaypalData{}
	return &this
}

// NewPaypalDataWithDefaults instantiates a new PaypalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalDataWithDefaults() *PaypalData {
	this := PaypalData{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PaypalData) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalData) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PaypalData) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PaypalData) SetEmail(v string) {
	o.Email = &v
}

// GetPaypalId returns the PaypalId field value if set, zero value otherwise.
func (o *PaypalData) GetPaypalId() string {
	if o == nil || IsNil(o.PaypalId) {
		var ret string
		return ret
	}
	return *o.PaypalId
}

// GetPaypalIdOk returns a tuple with the PaypalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalData) GetPaypalIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaypalId) {
		return nil, false
	}
	return o.PaypalId, true
}

// HasPaypalId returns a boolean if a field has been set.
func (o *PaypalData) HasPaypalId() bool {
	if o != nil && !IsNil(o.PaypalId) {
		return true
	}

	return false
}

// SetPaypalId gets a reference to the given string and assigns it to the PaypalId field.
func (o *PaypalData) SetPaypalId(v string) {
	o.PaypalId = &v
}

func (o PaypalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaypalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PaypalId) {
		toSerialize["paypal_id"] = o.PaypalId
	}
	return toSerialize, nil
}

type NullablePaypalData struct {
	value *PaypalData
	isSet bool
}

func (v NullablePaypalData) Get() *PaypalData {
	return v.value
}

func (v *NullablePaypalData) Set(val *PaypalData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalData(val *PaypalData) *NullablePaypalData {
	return &NullablePaypalData{value: val, isSet: true}
}

func (v NullablePaypalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


