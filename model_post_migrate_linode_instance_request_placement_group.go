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
	"bytes"
	"fmt"
)

// checks if the PostMigrateLinodeInstanceRequestPlacementGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostMigrateLinodeInstanceRequestPlacementGroup{}

// PostMigrateLinodeInstanceRequestPlacementGroup Include this to assign this Linode to an existing [placement group](https://www.linode.com/docs/products/compute/compute-instances/guides/placement-groups/) in the data center you're migrating to. These constraints apply:  - If the target Linode is in a placement group, it will be removed from it when migrating. - The target placement group needs to be in the same `region` you're migrating to. - The target placement group needs to have capacity. Run the [Get a region](https://techdocs.akamai.com/linode-api/reference/get-region) operation for the region you want to migrate to, and store the `maximum_linodes_per_pg` value. Run the [Get a placement group](https://techdocs.akamai.com/linode-api/reference/get-placement-group) operation for that same region to review how many Linodes are in that group.
type PostMigrateLinodeInstanceRequestPlacementGroup struct {
	// The placement group's ID. You need to provide it for all operations that affect it.
	Id int32 `json:"id"`
}

type _PostMigrateLinodeInstanceRequestPlacementGroup PostMigrateLinodeInstanceRequestPlacementGroup

// NewPostMigrateLinodeInstanceRequestPlacementGroup instantiates a new PostMigrateLinodeInstanceRequestPlacementGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostMigrateLinodeInstanceRequestPlacementGroup(id int32) *PostMigrateLinodeInstanceRequestPlacementGroup {
	this := PostMigrateLinodeInstanceRequestPlacementGroup{}
	this.Id = id
	return &this
}

// NewPostMigrateLinodeInstanceRequestPlacementGroupWithDefaults instantiates a new PostMigrateLinodeInstanceRequestPlacementGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostMigrateLinodeInstanceRequestPlacementGroupWithDefaults() *PostMigrateLinodeInstanceRequestPlacementGroup {
	this := PostMigrateLinodeInstanceRequestPlacementGroup{}
	return &this
}

// GetId returns the Id field value
func (o *PostMigrateLinodeInstanceRequestPlacementGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostMigrateLinodeInstanceRequestPlacementGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostMigrateLinodeInstanceRequestPlacementGroup) SetId(v int32) {
	o.Id = v
}

func (o PostMigrateLinodeInstanceRequestPlacementGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostMigrateLinodeInstanceRequestPlacementGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *PostMigrateLinodeInstanceRequestPlacementGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostMigrateLinodeInstanceRequestPlacementGroup := _PostMigrateLinodeInstanceRequestPlacementGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostMigrateLinodeInstanceRequestPlacementGroup)

	if err != nil {
		return err
	}

	*o = PostMigrateLinodeInstanceRequestPlacementGroup(varPostMigrateLinodeInstanceRequestPlacementGroup)

	return err
}

type NullablePostMigrateLinodeInstanceRequestPlacementGroup struct {
	value *PostMigrateLinodeInstanceRequestPlacementGroup
	isSet bool
}

func (v NullablePostMigrateLinodeInstanceRequestPlacementGroup) Get() *PostMigrateLinodeInstanceRequestPlacementGroup {
	return v.value
}

func (v *NullablePostMigrateLinodeInstanceRequestPlacementGroup) Set(val *PostMigrateLinodeInstanceRequestPlacementGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePostMigrateLinodeInstanceRequestPlacementGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePostMigrateLinodeInstanceRequestPlacementGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostMigrateLinodeInstanceRequestPlacementGroup(val *PostMigrateLinodeInstanceRequestPlacementGroup) *NullablePostMigrateLinodeInstanceRequestPlacementGroup {
	return &NullablePostMigrateLinodeInstanceRequestPlacementGroup{value: val, isSet: true}
}

func (v NullablePostMigrateLinodeInstanceRequestPlacementGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostMigrateLinodeInstanceRequestPlacementGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


