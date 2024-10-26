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

// checks if the PostBetaProgramRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBetaProgramRequest{}

// PostBetaProgramRequest The Beta Program ID to enroll in for your Account.
type PostBetaProgramRequest struct {
	// The unique identifier of the Beta Program.
	Id string `json:"id"`
}

type _PostBetaProgramRequest PostBetaProgramRequest

// NewPostBetaProgramRequest instantiates a new PostBetaProgramRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBetaProgramRequest(id string) *PostBetaProgramRequest {
	this := PostBetaProgramRequest{}
	this.Id = id
	return &this
}

// NewPostBetaProgramRequestWithDefaults instantiates a new PostBetaProgramRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBetaProgramRequestWithDefaults() *PostBetaProgramRequest {
	this := PostBetaProgramRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PostBetaProgramRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostBetaProgramRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostBetaProgramRequest) SetId(v string) {
	o.Id = v
}

func (o PostBetaProgramRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBetaProgramRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *PostBetaProgramRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPostBetaProgramRequest := _PostBetaProgramRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostBetaProgramRequest)

	if err != nil {
		return err
	}

	*o = PostBetaProgramRequest(varPostBetaProgramRequest)

	return err
}

type NullablePostBetaProgramRequest struct {
	value *PostBetaProgramRequest
	isSet bool
}

func (v NullablePostBetaProgramRequest) Get() *PostBetaProgramRequest {
	return v.value
}

func (v *NullablePostBetaProgramRequest) Set(val *PostBetaProgramRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBetaProgramRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBetaProgramRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBetaProgramRequest(val *PostBetaProgramRequest) *NullablePostBetaProgramRequest {
	return &NullablePostBetaProgramRequest{value: val, isSet: true}
}

func (v NullablePostBetaProgramRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBetaProgramRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

