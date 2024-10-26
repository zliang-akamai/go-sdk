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

// checks if the PropertiesEntities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertiesEntities{}

// PropertiesEntities A collection of the services to include in this transfer request, separated by type.
type PropertiesEntities struct {
	// An array containing the IDs of each of the Linodes included in this transfer.
	Linodes []int32 `json:"linodes,omitempty"`
}

// NewPropertiesEntities instantiates a new PropertiesEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertiesEntities() *PropertiesEntities {
	this := PropertiesEntities{}
	return &this
}

// NewPropertiesEntitiesWithDefaults instantiates a new PropertiesEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertiesEntitiesWithDefaults() *PropertiesEntities {
	this := PropertiesEntities{}
	return &this
}

// GetLinodes returns the Linodes field value if set, zero value otherwise.
func (o *PropertiesEntities) GetLinodes() []int32 {
	if o == nil || IsNil(o.Linodes) {
		var ret []int32
		return ret
	}
	return o.Linodes
}

// GetLinodesOk returns a tuple with the Linodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertiesEntities) GetLinodesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Linodes) {
		return nil, false
	}
	return o.Linodes, true
}

// HasLinodes returns a boolean if a field has been set.
func (o *PropertiesEntities) HasLinodes() bool {
	if o != nil && !IsNil(o.Linodes) {
		return true
	}

	return false
}

// SetLinodes gets a reference to the given []int32 and assigns it to the Linodes field.
func (o *PropertiesEntities) SetLinodes(v []int32) {
	o.Linodes = v
}

func (o PropertiesEntities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertiesEntities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Linodes) {
		toSerialize["linodes"] = o.Linodes
	}
	return toSerialize, nil
}

type NullablePropertiesEntities struct {
	value *PropertiesEntities
	isSet bool
}

func (v NullablePropertiesEntities) Get() *PropertiesEntities {
	return v.value
}

func (v *NullablePropertiesEntities) Set(val *PropertiesEntities) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertiesEntities) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertiesEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertiesEntities(val *PropertiesEntities) *NullablePropertiesEntities {
	return &NullablePropertiesEntities{value: val, isSet: true}
}

func (v NullablePropertiesEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertiesEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

