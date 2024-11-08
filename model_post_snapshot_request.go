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

// checks if the PostSnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSnapshotRequest{}

// PostSnapshotRequest struct for PostSnapshotRequest
type PostSnapshotRequest struct {
	// The label for the new snapshot.
	Label string `json:"label"`
}

type _PostSnapshotRequest PostSnapshotRequest

// NewPostSnapshotRequest instantiates a new PostSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSnapshotRequest(label string) *PostSnapshotRequest {
	this := PostSnapshotRequest{}
	this.Label = label
	return &this
}

// NewPostSnapshotRequestWithDefaults instantiates a new PostSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSnapshotRequestWithDefaults() *PostSnapshotRequest {
	this := PostSnapshotRequest{}
	return &this
}

// GetLabel returns the Label field value
func (o *PostSnapshotRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PostSnapshotRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PostSnapshotRequest) SetLabel(v string) {
	o.Label = v
}

func (o PostSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

func (o *PostSnapshotRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
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

	varPostSnapshotRequest := _PostSnapshotRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSnapshotRequest)

	if err != nil {
		return err
	}

	*o = PostSnapshotRequest(varPostSnapshotRequest)

	return err
}

type NullablePostSnapshotRequest struct {
	value *PostSnapshotRequest
	isSet bool
}

func (v NullablePostSnapshotRequest) Get() *PostSnapshotRequest {
	return v.value
}

func (v *NullablePostSnapshotRequest) Set(val *PostSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSnapshotRequest(val *PostSnapshotRequest) *NullablePostSnapshotRequest {
	return &NullablePostSnapshotRequest{value: val, isSet: true}
}

func (v NullablePostSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


