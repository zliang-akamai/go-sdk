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

// checks if the GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner{}

// GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner struct for GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner
type GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner struct {
	// Returns `true` if the Interface is in use, meaning that the Linode was powered on using the Configuration Profile to which the Interface belongs. Otherwise returns `false`.
	Active *bool `json:"active,omitempty"`
	// ID of the interface.
	Id *int32 `json:"id,omitempty"`
}

// NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner instantiates a new GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner() *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner {
	this := GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner{}
	return &this
}

// NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInnerWithDefaults instantiates a new GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInnerWithDefaults() *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner {
	this := GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) SetActive(v bool) {
	o.Active = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) SetId(v int32) {
	o.Id = &v
}

func (o GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner struct {
	value *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner
	isSet bool
}

func (v NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) Get() *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner {
	return v.value
}

func (v *NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) Set(val *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner(val *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) *NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner {
	return &NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner{value: val, isSet: true}
}

func (v NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


