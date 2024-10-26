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

// checks if the GetLongviewPlan200ResponsePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLongviewPlan200ResponsePrice{}

// GetLongviewPlan200ResponsePrice Pricing information about this Subscription tier.
type GetLongviewPlan200ResponsePrice struct {
	// The hourly price, in US dollars, for this Subscription tier.
	Hourly *float32 `json:"hourly,omitempty"`
	// The maximum monthly price in US Dollars for this Subscription tier. You will never be charged more than this amount per month for this subscription.
	Monthly *float32 `json:"monthly,omitempty"`
}

// NewGetLongviewPlan200ResponsePrice instantiates a new GetLongviewPlan200ResponsePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLongviewPlan200ResponsePrice() *GetLongviewPlan200ResponsePrice {
	this := GetLongviewPlan200ResponsePrice{}
	return &this
}

// NewGetLongviewPlan200ResponsePriceWithDefaults instantiates a new GetLongviewPlan200ResponsePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLongviewPlan200ResponsePriceWithDefaults() *GetLongviewPlan200ResponsePrice {
	this := GetLongviewPlan200ResponsePrice{}
	return &this
}

// GetHourly returns the Hourly field value if set, zero value otherwise.
func (o *GetLongviewPlan200ResponsePrice) GetHourly() float32 {
	if o == nil || IsNil(o.Hourly) {
		var ret float32
		return ret
	}
	return *o.Hourly
}

// GetHourlyOk returns a tuple with the Hourly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewPlan200ResponsePrice) GetHourlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Hourly) {
		return nil, false
	}
	return o.Hourly, true
}

// HasHourly returns a boolean if a field has been set.
func (o *GetLongviewPlan200ResponsePrice) HasHourly() bool {
	if o != nil && !IsNil(o.Hourly) {
		return true
	}

	return false
}

// SetHourly gets a reference to the given float32 and assigns it to the Hourly field.
func (o *GetLongviewPlan200ResponsePrice) SetHourly(v float32) {
	o.Hourly = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *GetLongviewPlan200ResponsePrice) GetMonthly() float32 {
	if o == nil || IsNil(o.Monthly) {
		var ret float32
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewPlan200ResponsePrice) GetMonthlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Monthly) {
		return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *GetLongviewPlan200ResponsePrice) HasMonthly() bool {
	if o != nil && !IsNil(o.Monthly) {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given float32 and assigns it to the Monthly field.
func (o *GetLongviewPlan200ResponsePrice) SetMonthly(v float32) {
	o.Monthly = &v
}

func (o GetLongviewPlan200ResponsePrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLongviewPlan200ResponsePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hourly) {
		toSerialize["hourly"] = o.Hourly
	}
	if !IsNil(o.Monthly) {
		toSerialize["monthly"] = o.Monthly
	}
	return toSerialize, nil
}

type NullableGetLongviewPlan200ResponsePrice struct {
	value *GetLongviewPlan200ResponsePrice
	isSet bool
}

func (v NullableGetLongviewPlan200ResponsePrice) Get() *GetLongviewPlan200ResponsePrice {
	return v.value
}

func (v *NullableGetLongviewPlan200ResponsePrice) Set(val *GetLongviewPlan200ResponsePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLongviewPlan200ResponsePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLongviewPlan200ResponsePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLongviewPlan200ResponsePrice(val *GetLongviewPlan200ResponsePrice) *NullableGetLongviewPlan200ResponsePrice {
	return &NullableGetLongviewPlan200ResponsePrice{value: val, isSet: true}
}

func (v NullableGetLongviewPlan200ResponsePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLongviewPlan200ResponsePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

