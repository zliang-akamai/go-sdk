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

// checks if the GetLkeTypes200ResponseDataInnerPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLkeTypes200ResponseDataInnerPrice{}

// GetLkeTypes200ResponseDataInnerPrice The default cost of this Kubernetes type. Prices are in US dollars, broken down into hourly and monthly charges.  Certain regions have different prices from the default. For region-specific prices, see `region_prices`.
type GetLkeTypes200ResponseDataInnerPrice struct {
	// Cost (in US dollars) per hour.
	Hourly *float32 `json:"hourly,omitempty"`
	// Cost per month, in US dollars.
	Monthly *float32 `json:"monthly,omitempty"`
}

// NewGetLkeTypes200ResponseDataInnerPrice instantiates a new GetLkeTypes200ResponseDataInnerPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLkeTypes200ResponseDataInnerPrice() *GetLkeTypes200ResponseDataInnerPrice {
	this := GetLkeTypes200ResponseDataInnerPrice{}
	return &this
}

// NewGetLkeTypes200ResponseDataInnerPriceWithDefaults instantiates a new GetLkeTypes200ResponseDataInnerPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLkeTypes200ResponseDataInnerPriceWithDefaults() *GetLkeTypes200ResponseDataInnerPrice {
	this := GetLkeTypes200ResponseDataInnerPrice{}
	return &this
}

// GetHourly returns the Hourly field value if set, zero value otherwise.
func (o *GetLkeTypes200ResponseDataInnerPrice) GetHourly() float32 {
	if o == nil || IsNil(o.Hourly) {
		var ret float32
		return ret
	}
	return *o.Hourly
}

// GetHourlyOk returns a tuple with the Hourly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeTypes200ResponseDataInnerPrice) GetHourlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Hourly) {
		return nil, false
	}
	return o.Hourly, true
}

// HasHourly returns a boolean if a field has been set.
func (o *GetLkeTypes200ResponseDataInnerPrice) HasHourly() bool {
	if o != nil && !IsNil(o.Hourly) {
		return true
	}

	return false
}

// SetHourly gets a reference to the given float32 and assigns it to the Hourly field.
func (o *GetLkeTypes200ResponseDataInnerPrice) SetHourly(v float32) {
	o.Hourly = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *GetLkeTypes200ResponseDataInnerPrice) GetMonthly() float32 {
	if o == nil || IsNil(o.Monthly) {
		var ret float32
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeTypes200ResponseDataInnerPrice) GetMonthlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Monthly) {
		return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *GetLkeTypes200ResponseDataInnerPrice) HasMonthly() bool {
	if o != nil && !IsNil(o.Monthly) {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given float32 and assigns it to the Monthly field.
func (o *GetLkeTypes200ResponseDataInnerPrice) SetMonthly(v float32) {
	o.Monthly = &v
}

func (o GetLkeTypes200ResponseDataInnerPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLkeTypes200ResponseDataInnerPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hourly) {
		toSerialize["hourly"] = o.Hourly
	}
	if !IsNil(o.Monthly) {
		toSerialize["monthly"] = o.Monthly
	}
	return toSerialize, nil
}

type NullableGetLkeTypes200ResponseDataInnerPrice struct {
	value *GetLkeTypes200ResponseDataInnerPrice
	isSet bool
}

func (v NullableGetLkeTypes200ResponseDataInnerPrice) Get() *GetLkeTypes200ResponseDataInnerPrice {
	return v.value
}

func (v *NullableGetLkeTypes200ResponseDataInnerPrice) Set(val *GetLkeTypes200ResponseDataInnerPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLkeTypes200ResponseDataInnerPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLkeTypes200ResponseDataInnerPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLkeTypes200ResponseDataInnerPrice(val *GetLkeTypes200ResponseDataInnerPrice) *NullableGetLkeTypes200ResponseDataInnerPrice {
	return &NullableGetLkeTypes200ResponseDataInnerPrice{value: val, isSet: true}
}

func (v NullableGetLkeTypes200ResponseDataInnerPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLkeTypes200ResponseDataInnerPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

