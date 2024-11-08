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

// checks if the GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner{}

// GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner The individual compute instances the system is migrating out of the placement group.
type GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner struct {
	// The unique identifier for a compute instance being migrated out of the placement group.
	LinodeId *int32 `json:"linode_id,omitempty"`
}

// NewGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner instantiates a new GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner() *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner {
	this := GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner{}
	return &this
}

// NewGetPlacementGroups200ResponseDataInnerMigrationsOutboundInnerWithDefaults instantiates a new GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPlacementGroups200ResponseDataInnerMigrationsOutboundInnerWithDefaults() *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner {
	this := GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner{}
	return &this
}

// GetLinodeId returns the LinodeId field value if set, zero value otherwise.
func (o *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) GetLinodeId() int32 {
	if o == nil || IsNil(o.LinodeId) {
		var ret int32
		return ret
	}
	return *o.LinodeId
}

// GetLinodeIdOk returns a tuple with the LinodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) GetLinodeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LinodeId) {
		return nil, false
	}
	return o.LinodeId, true
}

// HasLinodeId returns a boolean if a field has been set.
func (o *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) HasLinodeId() bool {
	if o != nil && !IsNil(o.LinodeId) {
		return true
	}

	return false
}

// SetLinodeId gets a reference to the given int32 and assigns it to the LinodeId field.
func (o *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) SetLinodeId(v int32) {
	o.LinodeId = &v
}

func (o GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinodeId) {
		toSerialize["linode_id"] = o.LinodeId
	}
	return toSerialize, nil
}

type NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner struct {
	value *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner
	isSet bool
}

func (v NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) Get() *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner {
	return v.value
}

func (v *NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) Set(val *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner(val *GetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) *NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner {
	return &NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner{value: val, isSet: true}
}

func (v NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPlacementGroups200ResponseDataInnerMigrationsOutboundInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


