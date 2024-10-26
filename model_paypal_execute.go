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

// checks if the PaypalExecute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaypalExecute{}

// PaypalExecute An object representing an execution of Payment to PayPal to capture the funds and credit your Linode Account.
type PaypalExecute struct {
	// The PayerID returned by PayPal during the transaction authorization process.
	PayerId string `json:"payer_id"`
	// The PaymentID returned from [Stage a PayPal payment](https://techdocs.akamai.com/linode-api/reference/post-pay-pal-payment) that has been approved with PayPal.
	PaymentId string `json:"payment_id"`
}

type _PaypalExecute PaypalExecute

// NewPaypalExecute instantiates a new PaypalExecute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalExecute(payerId string, paymentId string) *PaypalExecute {
	this := PaypalExecute{}
	this.PayerId = payerId
	this.PaymentId = paymentId
	return &this
}

// NewPaypalExecuteWithDefaults instantiates a new PaypalExecute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalExecuteWithDefaults() *PaypalExecute {
	this := PaypalExecute{}
	return &this
}

// GetPayerId returns the PayerId field value
func (o *PaypalExecute) GetPayerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayerId
}

// GetPayerIdOk returns a tuple with the PayerId field value
// and a boolean to check if the value has been set.
func (o *PaypalExecute) GetPayerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerId, true
}

// SetPayerId sets field value
func (o *PaypalExecute) SetPayerId(v string) {
	o.PayerId = v
}

// GetPaymentId returns the PaymentId field value
func (o *PaypalExecute) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaypalExecute) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaypalExecute) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o PaypalExecute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaypalExecute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payer_id"] = o.PayerId
	toSerialize["payment_id"] = o.PaymentId
	return toSerialize, nil
}

func (o *PaypalExecute) UnmarshalJSON(data []byte) (err error) {
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

	varPaypalExecute := _PaypalExecute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaypalExecute)

	if err != nil {
		return err
	}

	*o = PaypalExecute(varPaypalExecute)

	return err
}

type NullablePaypalExecute struct {
	value *PaypalExecute
	isSet bool
}

func (v NullablePaypalExecute) Get() *PaypalExecute {
	return v.value
}

func (v *NullablePaypalExecute) Set(val *PaypalExecute) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalExecute) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalExecute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalExecute(val *PaypalExecute) *NullablePaypalExecute {
	return &NullablePaypalExecute{value: val, isSet: true}
}

func (v NullablePaypalExecute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalExecute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


