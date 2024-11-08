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

// checks if the PostLinodeInstanceRequestAllOfPlacementGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLinodeInstanceRequestAllOfPlacementGroup{}

// PostLinodeInstanceRequestAllOfPlacementGroup Include this to assign this Linode to an existing [placement group](https://www.linode.com/docs/products/compute/compute-instances/guides/placement-groups/). These constraints apply:  - The target placement group needs to be in the same `region` set for this Linode. - The placement group needs to have capacity. Run the [Get a region](https://techdocs.akamai.com/linode-api/reference/get-region) operation and store the `maximum_linodes_per_pg` value to know the Linode limit per placement group. You can then run the [Get a placement group](https://techdocs.akamai.com/linode-api/reference/get-placement-group) operation to review the Linodes in that group.
type PostLinodeInstanceRequestAllOfPlacementGroup struct {
	// The placement group's ID. You need to provide it for all operations that affect it.
	Id *int32 `json:"id,omitempty"`
}

// NewPostLinodeInstanceRequestAllOfPlacementGroup instantiates a new PostLinodeInstanceRequestAllOfPlacementGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLinodeInstanceRequestAllOfPlacementGroup() *PostLinodeInstanceRequestAllOfPlacementGroup {
	this := PostLinodeInstanceRequestAllOfPlacementGroup{}
	return &this
}

// NewPostLinodeInstanceRequestAllOfPlacementGroupWithDefaults instantiates a new PostLinodeInstanceRequestAllOfPlacementGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLinodeInstanceRequestAllOfPlacementGroupWithDefaults() *PostLinodeInstanceRequestAllOfPlacementGroup {
	this := PostLinodeInstanceRequestAllOfPlacementGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequestAllOfPlacementGroup) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequestAllOfPlacementGroup) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequestAllOfPlacementGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PostLinodeInstanceRequestAllOfPlacementGroup) SetId(v int32) {
	o.Id = &v
}

func (o PostLinodeInstanceRequestAllOfPlacementGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLinodeInstanceRequestAllOfPlacementGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePostLinodeInstanceRequestAllOfPlacementGroup struct {
	value *PostLinodeInstanceRequestAllOfPlacementGroup
	isSet bool
}

func (v NullablePostLinodeInstanceRequestAllOfPlacementGroup) Get() *PostLinodeInstanceRequestAllOfPlacementGroup {
	return v.value
}

func (v *NullablePostLinodeInstanceRequestAllOfPlacementGroup) Set(val *PostLinodeInstanceRequestAllOfPlacementGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLinodeInstanceRequestAllOfPlacementGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLinodeInstanceRequestAllOfPlacementGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLinodeInstanceRequestAllOfPlacementGroup(val *PostLinodeInstanceRequestAllOfPlacementGroup) *NullablePostLinodeInstanceRequestAllOfPlacementGroup {
	return &NullablePostLinodeInstanceRequestAllOfPlacementGroup{value: val, isSet: true}
}

func (v NullablePostLinodeInstanceRequestAllOfPlacementGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLinodeInstanceRequestAllOfPlacementGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


