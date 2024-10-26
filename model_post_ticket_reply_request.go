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

// checks if the PostTicketReplyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTicketReplyRequest{}

// PostTicketReplyRequest struct for PostTicketReplyRequest
type PostTicketReplyRequest struct {
	// The content of your reply.
	Description string `json:"description"`
}

type _PostTicketReplyRequest PostTicketReplyRequest

// NewPostTicketReplyRequest instantiates a new PostTicketReplyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTicketReplyRequest(description string) *PostTicketReplyRequest {
	this := PostTicketReplyRequest{}
	this.Description = description
	return &this
}

// NewPostTicketReplyRequestWithDefaults instantiates a new PostTicketReplyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTicketReplyRequestWithDefaults() *PostTicketReplyRequest {
	this := PostTicketReplyRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *PostTicketReplyRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PostTicketReplyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PostTicketReplyRequest) SetDescription(v string) {
	o.Description = v
}

func (o PostTicketReplyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTicketReplyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *PostTicketReplyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
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

	varPostTicketReplyRequest := _PostTicketReplyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostTicketReplyRequest)

	if err != nil {
		return err
	}

	*o = PostTicketReplyRequest(varPostTicketReplyRequest)

	return err
}

type NullablePostTicketReplyRequest struct {
	value *PostTicketReplyRequest
	isSet bool
}

func (v NullablePostTicketReplyRequest) Get() *PostTicketReplyRequest {
	return v.value
}

func (v *NullablePostTicketReplyRequest) Set(val *PostTicketReplyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTicketReplyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTicketReplyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTicketReplyRequest(val *PostTicketReplyRequest) *NullablePostTicketReplyRequest {
	return &NullablePostTicketReplyRequest{value: val, isSet: true}
}

func (v NullablePostTicketReplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTicketReplyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

