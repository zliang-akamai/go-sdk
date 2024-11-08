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

// checks if the PostPromoCreditRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPromoCreditRequest{}

// PostPromoCreditRequest struct for PostPromoCreditRequest
type PostPromoCreditRequest struct {
	// The Promo Code.
	PromoCode string `json:"promo_code"`
}

type _PostPromoCreditRequest PostPromoCreditRequest

// NewPostPromoCreditRequest instantiates a new PostPromoCreditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPromoCreditRequest(promoCode string) *PostPromoCreditRequest {
	this := PostPromoCreditRequest{}
	this.PromoCode = promoCode
	return &this
}

// NewPostPromoCreditRequestWithDefaults instantiates a new PostPromoCreditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPromoCreditRequestWithDefaults() *PostPromoCreditRequest {
	this := PostPromoCreditRequest{}
	return &this
}

// GetPromoCode returns the PromoCode field value
func (o *PostPromoCreditRequest) GetPromoCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PromoCode
}

// GetPromoCodeOk returns a tuple with the PromoCode field value
// and a boolean to check if the value has been set.
func (o *PostPromoCreditRequest) GetPromoCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromoCode, true
}

// SetPromoCode sets field value
func (o *PostPromoCreditRequest) SetPromoCode(v string) {
	o.PromoCode = v
}

func (o PostPromoCreditRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPromoCreditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["promo_code"] = o.PromoCode
	return toSerialize, nil
}

func (o *PostPromoCreditRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"promo_code",
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

	varPostPromoCreditRequest := _PostPromoCreditRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostPromoCreditRequest)

	if err != nil {
		return err
	}

	*o = PostPromoCreditRequest(varPostPromoCreditRequest)

	return err
}

type NullablePostPromoCreditRequest struct {
	value *PostPromoCreditRequest
	isSet bool
}

func (v NullablePostPromoCreditRequest) Get() *PostPromoCreditRequest {
	return v.value
}

func (v *NullablePostPromoCreditRequest) Set(val *PostPromoCreditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPromoCreditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPromoCreditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPromoCreditRequest(val *PostPromoCreditRequest) *NullablePostPromoCreditRequest {
	return &NullablePostPromoCreditRequest{value: val, isSet: true}
}

func (v NullablePostPromoCreditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPromoCreditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


