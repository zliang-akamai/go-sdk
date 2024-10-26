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

// checks if the PostPaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPaymentRequest{}

// PostPaymentRequest struct for PostPaymentRequest
type PostPaymentRequest struct {
	// The ID of the Payment Method to apply to the Payment.
	PaymentMethodId *int32 `json:"payment_method_id,omitempty"`
	// The amount in US Dollars of the Payment.  - Can begin with or without `$`. - Commas (`,`) are not accepted. - Must end with a decimal expression, such as `.00` or `.99`. - Minimum: `$5.00` or the Account balance, whichever is lower. - Maximum: `$2000.00` or the Account balance up to `$50000.00`, whichever is greater.
	Usd *string `json:"usd,omitempty" validate:"regexp=^\\\\$?\\\\d+\\\\.\\\\d{2}$"`
}

// NewPostPaymentRequest instantiates a new PostPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPaymentRequest() *PostPaymentRequest {
	this := PostPaymentRequest{}
	return &this
}

// NewPostPaymentRequestWithDefaults instantiates a new PostPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPaymentRequestWithDefaults() *PostPaymentRequest {
	this := PostPaymentRequest{}
	return &this
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *PostPaymentRequest) GetPaymentMethodId() int32 {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret int32
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPaymentRequest) GetPaymentMethodIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *PostPaymentRequest) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given int32 and assigns it to the PaymentMethodId field.
func (o *PostPaymentRequest) SetPaymentMethodId(v int32) {
	o.PaymentMethodId = &v
}

// GetUsd returns the Usd field value if set, zero value otherwise.
func (o *PostPaymentRequest) GetUsd() string {
	if o == nil || IsNil(o.Usd) {
		var ret string
		return ret
	}
	return *o.Usd
}

// GetUsdOk returns a tuple with the Usd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPaymentRequest) GetUsdOk() (*string, bool) {
	if o == nil || IsNil(o.Usd) {
		return nil, false
	}
	return o.Usd, true
}

// HasUsd returns a boolean if a field has been set.
func (o *PostPaymentRequest) HasUsd() bool {
	if o != nil && !IsNil(o.Usd) {
		return true
	}

	return false
}

// SetUsd gets a reference to the given string and assigns it to the Usd field.
func (o *PostPaymentRequest) SetUsd(v string) {
	o.Usd = &v
}

func (o PostPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if !IsNil(o.Usd) {
		toSerialize["usd"] = o.Usd
	}
	return toSerialize, nil
}

type NullablePostPaymentRequest struct {
	value *PostPaymentRequest
	isSet bool
}

func (v NullablePostPaymentRequest) Get() *PostPaymentRequest {
	return v.value
}

func (v *NullablePostPaymentRequest) Set(val *PostPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPaymentRequest(val *PostPaymentRequest) *NullablePostPaymentRequest {
	return &NullablePostPaymentRequest{value: val, isSet: true}
}

func (v NullablePostPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


