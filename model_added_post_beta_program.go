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

// checks if the AddedPostBetaProgram type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedPostBetaProgram{}

// AddedPostBetaProgram The Beta Program ID to enroll in for your Account.
type AddedPostBetaProgram struct {
	// The unique identifier of the Beta Program.
	Id string `json:"id"`
}

type _AddedPostBetaProgram AddedPostBetaProgram

// NewAddedPostBetaProgram instantiates a new AddedPostBetaProgram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedPostBetaProgram(id string) *AddedPostBetaProgram {
	this := AddedPostBetaProgram{}
	this.Id = id
	return &this
}

// NewAddedPostBetaProgramWithDefaults instantiates a new AddedPostBetaProgram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedPostBetaProgramWithDefaults() *AddedPostBetaProgram {
	this := AddedPostBetaProgram{}
	return &this
}

// GetId returns the Id field value
func (o *AddedPostBetaProgram) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddedPostBetaProgram) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddedPostBetaProgram) SetId(v string) {
	o.Id = v
}

func (o AddedPostBetaProgram) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedPostBetaProgram) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AddedPostBetaProgram) UnmarshalJSON(data []byte) (err error) {
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

	varAddedPostBetaProgram := _AddedPostBetaProgram{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddedPostBetaProgram)

	if err != nil {
		return err
	}

	*o = AddedPostBetaProgram(varAddedPostBetaProgram)

	return err
}

type NullableAddedPostBetaProgram struct {
	value *AddedPostBetaProgram
	isSet bool
}

func (v NullableAddedPostBetaProgram) Get() *AddedPostBetaProgram {
	return v.value
}

func (v *NullableAddedPostBetaProgram) Set(val *AddedPostBetaProgram) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedPostBetaProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedPostBetaProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedPostBetaProgram(val *AddedPostBetaProgram) *NullableAddedPostBetaProgram {
	return &NullableAddedPostBetaProgram{value: val, isSet: true}
}

func (v NullableAddedPostBetaProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedPostBetaProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

