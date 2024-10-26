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

// checks if the AddedGetEntityTransfers200AllOfEntities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedGetEntityTransfers200AllOfEntities{}

// AddedGetEntityTransfers200AllOfEntities A collection of the entities to include in this transfer request, separated by type.
type AddedGetEntityTransfers200AllOfEntities struct {
	// An array containing the IDs of each of the Linodes included in this transfer.
	Linodes []int32 `json:"linodes,omitempty"`
}

// NewAddedGetEntityTransfers200AllOfEntities instantiates a new AddedGetEntityTransfers200AllOfEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedGetEntityTransfers200AllOfEntities() *AddedGetEntityTransfers200AllOfEntities {
	this := AddedGetEntityTransfers200AllOfEntities{}
	return &this
}

// NewAddedGetEntityTransfers200AllOfEntitiesWithDefaults instantiates a new AddedGetEntityTransfers200AllOfEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedGetEntityTransfers200AllOfEntitiesWithDefaults() *AddedGetEntityTransfers200AllOfEntities {
	this := AddedGetEntityTransfers200AllOfEntities{}
	return &this
}

// GetLinodes returns the Linodes field value if set, zero value otherwise.
func (o *AddedGetEntityTransfers200AllOfEntities) GetLinodes() []int32 {
	if o == nil || IsNil(o.Linodes) {
		var ret []int32
		return ret
	}
	return o.Linodes
}

// GetLinodesOk returns a tuple with the Linodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetEntityTransfers200AllOfEntities) GetLinodesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Linodes) {
		return nil, false
	}
	return o.Linodes, true
}

// HasLinodes returns a boolean if a field has been set.
func (o *AddedGetEntityTransfers200AllOfEntities) HasLinodes() bool {
	if o != nil && !IsNil(o.Linodes) {
		return true
	}

	return false
}

// SetLinodes gets a reference to the given []int32 and assigns it to the Linodes field.
func (o *AddedGetEntityTransfers200AllOfEntities) SetLinodes(v []int32) {
	o.Linodes = v
}

func (o AddedGetEntityTransfers200AllOfEntities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedGetEntityTransfers200AllOfEntities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Linodes) {
		toSerialize["linodes"] = o.Linodes
	}
	return toSerialize, nil
}

type NullableAddedGetEntityTransfers200AllOfEntities struct {
	value *AddedGetEntityTransfers200AllOfEntities
	isSet bool
}

func (v NullableAddedGetEntityTransfers200AllOfEntities) Get() *AddedGetEntityTransfers200AllOfEntities {
	return v.value
}

func (v *NullableAddedGetEntityTransfers200AllOfEntities) Set(val *AddedGetEntityTransfers200AllOfEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedGetEntityTransfers200AllOfEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedGetEntityTransfers200AllOfEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedGetEntityTransfers200AllOfEntities(val *AddedGetEntityTransfers200AllOfEntities) *NullableAddedGetEntityTransfers200AllOfEntities {
	return &NullableAddedGetEntityTransfers200AllOfEntities{value: val, isSet: true}
}

func (v NullableAddedGetEntityTransfers200AllOfEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedGetEntityTransfers200AllOfEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


