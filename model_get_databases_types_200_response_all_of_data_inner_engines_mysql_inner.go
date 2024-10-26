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

// checks if the GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner{}

// GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner struct for GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner
type GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner struct {
	Price *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice `json:"price,omitempty"`
	// The number of nodes for the Managed Database cluster for this subscription tier.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner instantiates a new GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner() *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner {
	this := GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner{}
	return &this
}

// NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerWithDefaults instantiates a new GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerWithDefaults() *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner {
	this := GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) GetPrice() GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice {
	if o == nil || IsNil(o.Price) {
		var ret GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) GetPriceOk() (*GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice and assigns it to the Price field.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) SetPrice(v GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInnerPrice) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner struct {
	value *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner
	isSet bool
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) Get() *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner {
	return v.value
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) Set(val *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner(val *GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) *NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner {
	return &NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner{value: val, isSet: true}
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

