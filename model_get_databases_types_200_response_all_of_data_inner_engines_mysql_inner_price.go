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

// checks if the GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice{}

// GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice Cost in US dollars, broken down into hourly and monthly charges.
type GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice struct {
	// Cost (in US dollars) per hour for this subscription tier.
	Hourly *float32 `json:"hourly,omitempty"`
	// Maximum cost (in US dollars) per month for this subscription tier.
	Monthly *float32 `json:"monthly,omitempty"`
}

// NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice instantiates a new GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice() *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice {
	this := GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice{}
	return &this
}

// NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPriceWithDefaults instantiates a new GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPriceWithDefaults() *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice {
	this := GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice{}
	return &this
}

// GetHourly returns the Hourly field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) GetHourly() float32 {
	if o == nil || IsNil(o.Hourly) {
		var ret float32
		return ret
	}
	return *o.Hourly
}

// GetHourlyOk returns a tuple with the Hourly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) GetHourlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Hourly) {
		return nil, false
	}
	return o.Hourly, true
}

// HasHourly returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) HasHourly() bool {
	if o != nil && !IsNil(o.Hourly) {
		return true
	}

	return false
}

// SetHourly gets a reference to the given float32 and assigns it to the Hourly field.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) SetHourly(v float32) {
	o.Hourly = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) GetMonthly() float32 {
	if o == nil || IsNil(o.Monthly) {
		var ret float32
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) GetMonthlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Monthly) {
		return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) HasMonthly() bool {
	if o != nil && !IsNil(o.Monthly) {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given float32 and assigns it to the Monthly field.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) SetMonthly(v float32) {
	o.Monthly = &v
}

func (o GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hourly) {
		toSerialize["hourly"] = o.Hourly
	}
	if !IsNil(o.Monthly) {
		toSerialize["monthly"] = o.Monthly
	}
	return toSerialize, nil
}

type NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice struct {
	value *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice
	isSet bool
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) Get() *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice {
	return v.value
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) Set(val *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice(val *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) *NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice {
	return &NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice{value: val, isSet: true}
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


