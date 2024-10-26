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

// checks if the PostAddLinodeIpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAddLinodeIpRequest{}

// PostAddLinodeIpRequest struct for PostAddLinodeIpRequest
type PostAddLinodeIpRequest struct {
	// Whether to create a public or private IPv4 address.
	Public bool `json:"public"`
	// The type of address you are allocating. Only IPv4 addresses may be allocated through this operation.
	Type string `json:"type"`
}

type _PostAddLinodeIpRequest PostAddLinodeIpRequest

// NewPostAddLinodeIpRequest instantiates a new PostAddLinodeIpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAddLinodeIpRequest(public bool, type_ string) *PostAddLinodeIpRequest {
	this := PostAddLinodeIpRequest{}
	this.Public = public
	this.Type = type_
	return &this
}

// NewPostAddLinodeIpRequestWithDefaults instantiates a new PostAddLinodeIpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAddLinodeIpRequestWithDefaults() *PostAddLinodeIpRequest {
	this := PostAddLinodeIpRequest{}
	return &this
}

// GetPublic returns the Public field value
func (o *PostAddLinodeIpRequest) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *PostAddLinodeIpRequest) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *PostAddLinodeIpRequest) SetPublic(v bool) {
	o.Public = v
}

// GetType returns the Type field value
func (o *PostAddLinodeIpRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostAddLinodeIpRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostAddLinodeIpRequest) SetType(v string) {
	o.Type = v
}

func (o PostAddLinodeIpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAddLinodeIpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["public"] = o.Public
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *PostAddLinodeIpRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"public",
		"type",
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

	varPostAddLinodeIpRequest := _PostAddLinodeIpRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostAddLinodeIpRequest)

	if err != nil {
		return err
	}

	*o = PostAddLinodeIpRequest(varPostAddLinodeIpRequest)

	return err
}

type NullablePostAddLinodeIpRequest struct {
	value *PostAddLinodeIpRequest
	isSet bool
}

func (v NullablePostAddLinodeIpRequest) Get() *PostAddLinodeIpRequest {
	return v.value
}

func (v *NullablePostAddLinodeIpRequest) Set(val *PostAddLinodeIpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAddLinodeIpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAddLinodeIpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAddLinodeIpRequest(val *PostAddLinodeIpRequest) *NullablePostAddLinodeIpRequest {
	return &NullablePostAddLinodeIpRequest{value: val, isSet: true}
}

func (v NullablePostAddLinodeIpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAddLinodeIpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

