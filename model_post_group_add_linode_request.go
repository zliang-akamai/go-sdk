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

// checks if the PostGroupAddLinodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostGroupAddLinodeRequest{}

// PostGroupAddLinodeRequest The compute instances included in a placement group.
type PostGroupAddLinodeRequest struct {
	// The `linodeId` values for individual compute instances included in the placement group.
	Linodes []int32 `json:"linodes,omitempty"`
}

// NewPostGroupAddLinodeRequest instantiates a new PostGroupAddLinodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostGroupAddLinodeRequest() *PostGroupAddLinodeRequest {
	this := PostGroupAddLinodeRequest{}
	return &this
}

// NewPostGroupAddLinodeRequestWithDefaults instantiates a new PostGroupAddLinodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostGroupAddLinodeRequestWithDefaults() *PostGroupAddLinodeRequest {
	this := PostGroupAddLinodeRequest{}
	return &this
}

// GetLinodes returns the Linodes field value if set, zero value otherwise.
func (o *PostGroupAddLinodeRequest) GetLinodes() []int32 {
	if o == nil || IsNil(o.Linodes) {
		var ret []int32
		return ret
	}
	return o.Linodes
}

// GetLinodesOk returns a tuple with the Linodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGroupAddLinodeRequest) GetLinodesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Linodes) {
		return nil, false
	}
	return o.Linodes, true
}

// HasLinodes returns a boolean if a field has been set.
func (o *PostGroupAddLinodeRequest) HasLinodes() bool {
	if o != nil && !IsNil(o.Linodes) {
		return true
	}

	return false
}

// SetLinodes gets a reference to the given []int32 and assigns it to the Linodes field.
func (o *PostGroupAddLinodeRequest) SetLinodes(v []int32) {
	o.Linodes = v
}

func (o PostGroupAddLinodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostGroupAddLinodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Linodes) {
		toSerialize["linodes"] = o.Linodes
	}
	return toSerialize, nil
}

type NullablePostGroupAddLinodeRequest struct {
	value *PostGroupAddLinodeRequest
	isSet bool
}

func (v NullablePostGroupAddLinodeRequest) Get() *PostGroupAddLinodeRequest {
	return v.value
}

func (v *NullablePostGroupAddLinodeRequest) Set(val *PostGroupAddLinodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostGroupAddLinodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostGroupAddLinodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostGroupAddLinodeRequest(val *PostGroupAddLinodeRequest) *NullablePostGroupAddLinodeRequest {
	return &NullablePostGroupAddLinodeRequest{value: val, isSet: true}
}

func (v NullablePostGroupAddLinodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostGroupAddLinodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


