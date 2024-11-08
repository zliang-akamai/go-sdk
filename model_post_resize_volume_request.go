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

// checks if the PostResizeVolumeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostResizeVolumeRequest{}

// PostResizeVolumeRequest struct for PostResizeVolumeRequest
type PostResizeVolumeRequest struct {
	// The Volume's size, in GiB.
	Size int32 `json:"size"`
}

type _PostResizeVolumeRequest PostResizeVolumeRequest

// NewPostResizeVolumeRequest instantiates a new PostResizeVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostResizeVolumeRequest(size int32) *PostResizeVolumeRequest {
	this := PostResizeVolumeRequest{}
	this.Size = size
	return &this
}

// NewPostResizeVolumeRequestWithDefaults instantiates a new PostResizeVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostResizeVolumeRequestWithDefaults() *PostResizeVolumeRequest {
	this := PostResizeVolumeRequest{}
	return &this
}

// GetSize returns the Size field value
func (o *PostResizeVolumeRequest) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PostResizeVolumeRequest) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PostResizeVolumeRequest) SetSize(v int32) {
	o.Size = v
}

func (o PostResizeVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostResizeVolumeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

func (o *PostResizeVolumeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
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

	varPostResizeVolumeRequest := _PostResizeVolumeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostResizeVolumeRequest)

	if err != nil {
		return err
	}

	*o = PostResizeVolumeRequest(varPostResizeVolumeRequest)

	return err
}

type NullablePostResizeVolumeRequest struct {
	value *PostResizeVolumeRequest
	isSet bool
}

func (v NullablePostResizeVolumeRequest) Get() *PostResizeVolumeRequest {
	return v.value
}

func (v *NullablePostResizeVolumeRequest) Set(val *PostResizeVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostResizeVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostResizeVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostResizeVolumeRequest(val *PostResizeVolumeRequest) *NullablePostResizeVolumeRequest {
	return &NullablePostResizeVolumeRequest{value: val, isSet: true}
}

func (v NullablePostResizeVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostResizeVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


