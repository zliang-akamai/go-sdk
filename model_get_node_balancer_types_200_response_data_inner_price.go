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

// checks if the GetNodeBalancerTypes200ResponseDataInnerPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodeBalancerTypes200ResponseDataInnerPrice{}

// GetNodeBalancerTypes200ResponseDataInnerPrice The default cost of this NodeBalancer type. Prices are in US dollars, broken down into hourly and monthly charges.  Certain regions have different prices from the default. For region-specific prices, see `region_prices`.
type GetNodeBalancerTypes200ResponseDataInnerPrice struct {
	// Cost (in US dollars) per hour.
	Hourly *float32 `json:"hourly,omitempty"`
	// Cost per month, in US dollars.
	Monthly *float32 `json:"monthly,omitempty"`
}

// NewGetNodeBalancerTypes200ResponseDataInnerPrice instantiates a new GetNodeBalancerTypes200ResponseDataInnerPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodeBalancerTypes200ResponseDataInnerPrice() *GetNodeBalancerTypes200ResponseDataInnerPrice {
	this := GetNodeBalancerTypes200ResponseDataInnerPrice{}
	return &this
}

// NewGetNodeBalancerTypes200ResponseDataInnerPriceWithDefaults instantiates a new GetNodeBalancerTypes200ResponseDataInnerPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodeBalancerTypes200ResponseDataInnerPriceWithDefaults() *GetNodeBalancerTypes200ResponseDataInnerPrice {
	this := GetNodeBalancerTypes200ResponseDataInnerPrice{}
	return &this
}

// GetHourly returns the Hourly field value if set, zero value otherwise.
func (o *GetNodeBalancerTypes200ResponseDataInnerPrice) GetHourly() float32 {
	if o == nil || IsNil(o.Hourly) {
		var ret float32
		return ret
	}
	return *o.Hourly
}

// GetHourlyOk returns a tuple with the Hourly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerTypes200ResponseDataInnerPrice) GetHourlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Hourly) {
		return nil, false
	}
	return o.Hourly, true
}

// HasHourly returns a boolean if a field has been set.
func (o *GetNodeBalancerTypes200ResponseDataInnerPrice) HasHourly() bool {
	if o != nil && !IsNil(o.Hourly) {
		return true
	}

	return false
}

// SetHourly gets a reference to the given float32 and assigns it to the Hourly field.
func (o *GetNodeBalancerTypes200ResponseDataInnerPrice) SetHourly(v float32) {
	o.Hourly = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *GetNodeBalancerTypes200ResponseDataInnerPrice) GetMonthly() float32 {
	if o == nil || IsNil(o.Monthly) {
		var ret float32
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerTypes200ResponseDataInnerPrice) GetMonthlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Monthly) {
		return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *GetNodeBalancerTypes200ResponseDataInnerPrice) HasMonthly() bool {
	if o != nil && !IsNil(o.Monthly) {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given float32 and assigns it to the Monthly field.
func (o *GetNodeBalancerTypes200ResponseDataInnerPrice) SetMonthly(v float32) {
	o.Monthly = &v
}

func (o GetNodeBalancerTypes200ResponseDataInnerPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodeBalancerTypes200ResponseDataInnerPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hourly) {
		toSerialize["hourly"] = o.Hourly
	}
	if !IsNil(o.Monthly) {
		toSerialize["monthly"] = o.Monthly
	}
	return toSerialize, nil
}

type NullableGetNodeBalancerTypes200ResponseDataInnerPrice struct {
	value *GetNodeBalancerTypes200ResponseDataInnerPrice
	isSet bool
}

func (v NullableGetNodeBalancerTypes200ResponseDataInnerPrice) Get() *GetNodeBalancerTypes200ResponseDataInnerPrice {
	return v.value
}

func (v *NullableGetNodeBalancerTypes200ResponseDataInnerPrice) Set(val *GetNodeBalancerTypes200ResponseDataInnerPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodeBalancerTypes200ResponseDataInnerPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodeBalancerTypes200ResponseDataInnerPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodeBalancerTypes200ResponseDataInnerPrice(val *GetNodeBalancerTypes200ResponseDataInnerPrice) *NullableGetNodeBalancerTypes200ResponseDataInnerPrice {
	return &NullableGetNodeBalancerTypes200ResponseDataInnerPrice{value: val, isSet: true}
}

func (v NullableGetNodeBalancerTypes200ResponseDataInnerPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodeBalancerTypes200ResponseDataInnerPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

