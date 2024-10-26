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

// checks if the GetObjectStorageTypes200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetObjectStorageTypes200ResponseDataInner{}

// GetObjectStorageTypes200ResponseDataInner Returns collection of Object Storage types and prices, including any region-specific rates.
type GetObjectStorageTypes200ResponseDataInner struct {
	// The ID representing the Object Storage type.
	Id *string `json:"id,omitempty"`
	// The Object Storage type label is for display purposes only.
	Label *string `json:"label,omitempty"`
	Price *GetObjectStorageTypes200ResponseDataInnerPrice `json:"price,omitempty"`
	RegionPrices []GetLkeTypes200ResponseDataInnerRegionPricesInner `json:"region_prices,omitempty"`
	// The monthly outbound transfer amount, in MB.
	Transfer *int32 `json:"transfer,omitempty"`
}

// NewGetObjectStorageTypes200ResponseDataInner instantiates a new GetObjectStorageTypes200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetObjectStorageTypes200ResponseDataInner() *GetObjectStorageTypes200ResponseDataInner {
	this := GetObjectStorageTypes200ResponseDataInner{}
	return &this
}

// NewGetObjectStorageTypes200ResponseDataInnerWithDefaults instantiates a new GetObjectStorageTypes200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetObjectStorageTypes200ResponseDataInnerWithDefaults() *GetObjectStorageTypes200ResponseDataInner {
	this := GetObjectStorageTypes200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetObjectStorageTypes200ResponseDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetObjectStorageTypes200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetObjectStorageTypes200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetObjectStorageTypes200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetObjectStorageTypes200ResponseDataInner) GetPrice() GetObjectStorageTypes200ResponseDataInnerPrice {
	if o == nil || IsNil(o.Price) {
		var ret GetObjectStorageTypes200ResponseDataInnerPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) GetPriceOk() (*GetObjectStorageTypes200ResponseDataInnerPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given GetObjectStorageTypes200ResponseDataInnerPrice and assigns it to the Price field.
func (o *GetObjectStorageTypes200ResponseDataInner) SetPrice(v GetObjectStorageTypes200ResponseDataInnerPrice) {
	o.Price = &v
}

// GetRegionPrices returns the RegionPrices field value if set, zero value otherwise.
func (o *GetObjectStorageTypes200ResponseDataInner) GetRegionPrices() []GetLkeTypes200ResponseDataInnerRegionPricesInner {
	if o == nil || IsNil(o.RegionPrices) {
		var ret []GetLkeTypes200ResponseDataInnerRegionPricesInner
		return ret
	}
	return o.RegionPrices
}

// GetRegionPricesOk returns a tuple with the RegionPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) GetRegionPricesOk() ([]GetLkeTypes200ResponseDataInnerRegionPricesInner, bool) {
	if o == nil || IsNil(o.RegionPrices) {
		return nil, false
	}
	return o.RegionPrices, true
}

// HasRegionPrices returns a boolean if a field has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) HasRegionPrices() bool {
	if o != nil && !IsNil(o.RegionPrices) {
		return true
	}

	return false
}

// SetRegionPrices gets a reference to the given []GetLkeTypes200ResponseDataInnerRegionPricesInner and assigns it to the RegionPrices field.
func (o *GetObjectStorageTypes200ResponseDataInner) SetRegionPrices(v []GetLkeTypes200ResponseDataInnerRegionPricesInner) {
	o.RegionPrices = v
}

// GetTransfer returns the Transfer field value if set, zero value otherwise.
func (o *GetObjectStorageTypes200ResponseDataInner) GetTransfer() int32 {
	if o == nil || IsNil(o.Transfer) {
		var ret int32
		return ret
	}
	return *o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) GetTransferOk() (*int32, bool) {
	if o == nil || IsNil(o.Transfer) {
		return nil, false
	}
	return o.Transfer, true
}

// HasTransfer returns a boolean if a field has been set.
func (o *GetObjectStorageTypes200ResponseDataInner) HasTransfer() bool {
	if o != nil && !IsNil(o.Transfer) {
		return true
	}

	return false
}

// SetTransfer gets a reference to the given int32 and assigns it to the Transfer field.
func (o *GetObjectStorageTypes200ResponseDataInner) SetTransfer(v int32) {
	o.Transfer = &v
}

func (o GetObjectStorageTypes200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetObjectStorageTypes200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.RegionPrices) {
		toSerialize["region_prices"] = o.RegionPrices
	}
	if !IsNil(o.Transfer) {
		toSerialize["transfer"] = o.Transfer
	}
	return toSerialize, nil
}

type NullableGetObjectStorageTypes200ResponseDataInner struct {
	value *GetObjectStorageTypes200ResponseDataInner
	isSet bool
}

func (v NullableGetObjectStorageTypes200ResponseDataInner) Get() *GetObjectStorageTypes200ResponseDataInner {
	return v.value
}

func (v *NullableGetObjectStorageTypes200ResponseDataInner) Set(val *GetObjectStorageTypes200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObjectStorageTypes200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObjectStorageTypes200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObjectStorageTypes200ResponseDataInner(val *GetObjectStorageTypes200ResponseDataInner) *NullableGetObjectStorageTypes200ResponseDataInner {
	return &NullableGetObjectStorageTypes200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetObjectStorageTypes200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetObjectStorageTypes200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


