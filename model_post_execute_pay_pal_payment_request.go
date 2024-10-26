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

// checks if the PostExecutePayPalPaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostExecutePayPalPaymentRequest{}

// PostExecutePayPalPaymentRequest An object representing an execution of Payment to PayPal to capture the funds and credit your Linode Account.
type PostExecutePayPalPaymentRequest struct {
	// The PayerID returned by PayPal during the transaction authorization process.
	PayerId string `json:"payer_id"`
	// The PaymentID returned from [Stage a PayPal payment](https://techdocs.akamai.com/linode-api/reference/post-pay-pal-payment) that has been approved with PayPal.
	PaymentId string `json:"payment_id"`
}

type _PostExecutePayPalPaymentRequest PostExecutePayPalPaymentRequest

// NewPostExecutePayPalPaymentRequest instantiates a new PostExecutePayPalPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostExecutePayPalPaymentRequest(payerId string, paymentId string) *PostExecutePayPalPaymentRequest {
	this := PostExecutePayPalPaymentRequest{}
	this.PayerId = payerId
	this.PaymentId = paymentId
	return &this
}

// NewPostExecutePayPalPaymentRequestWithDefaults instantiates a new PostExecutePayPalPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostExecutePayPalPaymentRequestWithDefaults() *PostExecutePayPalPaymentRequest {
	this := PostExecutePayPalPaymentRequest{}
	return &this
}

// GetPayerId returns the PayerId field value
func (o *PostExecutePayPalPaymentRequest) GetPayerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayerId
}

// GetPayerIdOk returns a tuple with the PayerId field value
// and a boolean to check if the value has been set.
func (o *PostExecutePayPalPaymentRequest) GetPayerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerId, true
}

// SetPayerId sets field value
func (o *PostExecutePayPalPaymentRequest) SetPayerId(v string) {
	o.PayerId = v
}

// GetPaymentId returns the PaymentId field value
func (o *PostExecutePayPalPaymentRequest) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PostExecutePayPalPaymentRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PostExecutePayPalPaymentRequest) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o PostExecutePayPalPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostExecutePayPalPaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payer_id"] = o.PayerId
	toSerialize["payment_id"] = o.PaymentId
	return toSerialize, nil
}

func (o *PostExecutePayPalPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payer_id",
		"payment_id",
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

	varPostExecutePayPalPaymentRequest := _PostExecutePayPalPaymentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostExecutePayPalPaymentRequest)

	if err != nil {
		return err
	}

	*o = PostExecutePayPalPaymentRequest(varPostExecutePayPalPaymentRequest)

	return err
}

type NullablePostExecutePayPalPaymentRequest struct {
	value *PostExecutePayPalPaymentRequest
	isSet bool
}

func (v NullablePostExecutePayPalPaymentRequest) Get() *PostExecutePayPalPaymentRequest {
	return v.value
}

func (v *NullablePostExecutePayPalPaymentRequest) Set(val *PostExecutePayPalPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostExecutePayPalPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostExecutePayPalPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostExecutePayPalPaymentRequest(val *PostExecutePayPalPaymentRequest) *NullablePostExecutePayPalPaymentRequest {
	return &NullablePostExecutePayPalPaymentRequest{value: val, isSet: true}
}

func (v NullablePostExecutePayPalPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostExecutePayPalPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

