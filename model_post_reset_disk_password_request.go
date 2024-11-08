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

// checks if the PostResetDiskPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostResetDiskPasswordRequest{}

// PostResetDiskPasswordRequest struct for PostResetDiskPasswordRequest
type PostResetDiskPasswordRequest struct {
	// The new root password for the OS installed on this Disk. The password must meet the complexity strength validation requirements for a strong password.
	Password string `json:"password"`
}

type _PostResetDiskPasswordRequest PostResetDiskPasswordRequest

// NewPostResetDiskPasswordRequest instantiates a new PostResetDiskPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostResetDiskPasswordRequest(password string) *PostResetDiskPasswordRequest {
	this := PostResetDiskPasswordRequest{}
	this.Password = password
	return &this
}

// NewPostResetDiskPasswordRequestWithDefaults instantiates a new PostResetDiskPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostResetDiskPasswordRequestWithDefaults() *PostResetDiskPasswordRequest {
	this := PostResetDiskPasswordRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *PostResetDiskPasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PostResetDiskPasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PostResetDiskPasswordRequest) SetPassword(v string) {
	o.Password = v
}

func (o PostResetDiskPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostResetDiskPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *PostResetDiskPasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
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

	varPostResetDiskPasswordRequest := _PostResetDiskPasswordRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostResetDiskPasswordRequest)

	if err != nil {
		return err
	}

	*o = PostResetDiskPasswordRequest(varPostResetDiskPasswordRequest)

	return err
}

type NullablePostResetDiskPasswordRequest struct {
	value *PostResetDiskPasswordRequest
	isSet bool
}

func (v NullablePostResetDiskPasswordRequest) Get() *PostResetDiskPasswordRequest {
	return v.value
}

func (v *NullablePostResetDiskPasswordRequest) Set(val *PostResetDiskPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostResetDiskPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostResetDiskPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostResetDiskPasswordRequest(val *PostResetDiskPasswordRequest) *NullablePostResetDiskPasswordRequest {
	return &NullablePostResetDiskPasswordRequest{value: val, isSet: true}
}

func (v NullablePostResetDiskPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostResetDiskPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


