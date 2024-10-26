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

// checks if the GetLinodeTypes200ResponseDataInnerAddonsBackups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLinodeTypes200ResponseDataInnerAddonsBackups{}

// GetLinodeTypes200ResponseDataInnerAddonsBackups Information about the optional Backup service offered for Linodes.
type GetLinodeTypes200ResponseDataInnerAddonsBackups struct {
	Price *GetLinodeTypes200ResponseDataInnerAddonsBackupsPrice `json:"price,omitempty"`
	RegionPrices []GetLinodeTypes200ResponseDataInnerAddonsBackupsRegionPricesInner `json:"region_prices,omitempty"`
}

// NewGetLinodeTypes200ResponseDataInnerAddonsBackups instantiates a new GetLinodeTypes200ResponseDataInnerAddonsBackups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLinodeTypes200ResponseDataInnerAddonsBackups() *GetLinodeTypes200ResponseDataInnerAddonsBackups {
	this := GetLinodeTypes200ResponseDataInnerAddonsBackups{}
	return &this
}

// NewGetLinodeTypes200ResponseDataInnerAddonsBackupsWithDefaults instantiates a new GetLinodeTypes200ResponseDataInnerAddonsBackups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLinodeTypes200ResponseDataInnerAddonsBackupsWithDefaults() *GetLinodeTypes200ResponseDataInnerAddonsBackups {
	this := GetLinodeTypes200ResponseDataInnerAddonsBackups{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetLinodeTypes200ResponseDataInnerAddonsBackups) GetPrice() GetLinodeTypes200ResponseDataInnerAddonsBackupsPrice {
	if o == nil || IsNil(o.Price) {
		var ret GetLinodeTypes200ResponseDataInnerAddonsBackupsPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeTypes200ResponseDataInnerAddonsBackups) GetPriceOk() (*GetLinodeTypes200ResponseDataInnerAddonsBackupsPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetLinodeTypes200ResponseDataInnerAddonsBackups) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given GetLinodeTypes200ResponseDataInnerAddonsBackupsPrice and assigns it to the Price field.
func (o *GetLinodeTypes200ResponseDataInnerAddonsBackups) SetPrice(v GetLinodeTypes200ResponseDataInnerAddonsBackupsPrice) {
	o.Price = &v
}

// GetRegionPrices returns the RegionPrices field value if set, zero value otherwise.
func (o *GetLinodeTypes200ResponseDataInnerAddonsBackups) GetRegionPrices() []GetLinodeTypes200ResponseDataInnerAddonsBackupsRegionPricesInner {
	if o == nil || IsNil(o.RegionPrices) {
		var ret []GetLinodeTypes200ResponseDataInnerAddonsBackupsRegionPricesInner
		return ret
	}
	return o.RegionPrices
}

// GetRegionPricesOk returns a tuple with the RegionPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeTypes200ResponseDataInnerAddonsBackups) GetRegionPricesOk() ([]GetLinodeTypes200ResponseDataInnerAddonsBackupsRegionPricesInner, bool) {
	if o == nil || IsNil(o.RegionPrices) {
		return nil, false
	}
	return o.RegionPrices, true
}

// HasRegionPrices returns a boolean if a field has been set.
func (o *GetLinodeTypes200ResponseDataInnerAddonsBackups) HasRegionPrices() bool {
	if o != nil && !IsNil(o.RegionPrices) {
		return true
	}

	return false
}

// SetRegionPrices gets a reference to the given []GetLinodeTypes200ResponseDataInnerAddonsBackupsRegionPricesInner and assigns it to the RegionPrices field.
func (o *GetLinodeTypes200ResponseDataInnerAddonsBackups) SetRegionPrices(v []GetLinodeTypes200ResponseDataInnerAddonsBackupsRegionPricesInner) {
	o.RegionPrices = v
}

func (o GetLinodeTypes200ResponseDataInnerAddonsBackups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLinodeTypes200ResponseDataInnerAddonsBackups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.RegionPrices) {
		toSerialize["region_prices"] = o.RegionPrices
	}
	return toSerialize, nil
}

type NullableGetLinodeTypes200ResponseDataInnerAddonsBackups struct {
	value *GetLinodeTypes200ResponseDataInnerAddonsBackups
	isSet bool
}

func (v NullableGetLinodeTypes200ResponseDataInnerAddonsBackups) Get() *GetLinodeTypes200ResponseDataInnerAddonsBackups {
	return v.value
}

func (v *NullableGetLinodeTypes200ResponseDataInnerAddonsBackups) Set(val *GetLinodeTypes200ResponseDataInnerAddonsBackups) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLinodeTypes200ResponseDataInnerAddonsBackups) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLinodeTypes200ResponseDataInnerAddonsBackups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLinodeTypes200ResponseDataInnerAddonsBackups(val *GetLinodeTypes200ResponseDataInnerAddonsBackups) *NullableGetLinodeTypes200ResponseDataInnerAddonsBackups {
	return &NullableGetLinodeTypes200ResponseDataInnerAddonsBackups{value: val, isSet: true}
}

func (v NullableGetLinodeTypes200ResponseDataInnerAddonsBackups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLinodeTypes200ResponseDataInnerAddonsBackups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


