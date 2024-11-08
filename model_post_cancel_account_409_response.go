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

// checks if the PostCancelAccount409Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCancelAccount409Response{}

// PostCancelAccount409Response struct for PostCancelAccount409Response
type PostCancelAccount409Response struct {
	Errors []PostCancelAccount409ResponseErrorsInner `json:"errors,omitempty"`
}

// NewPostCancelAccount409Response instantiates a new PostCancelAccount409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCancelAccount409Response() *PostCancelAccount409Response {
	this := PostCancelAccount409Response{}
	return &this
}

// NewPostCancelAccount409ResponseWithDefaults instantiates a new PostCancelAccount409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCancelAccount409ResponseWithDefaults() *PostCancelAccount409Response {
	this := PostCancelAccount409Response{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PostCancelAccount409Response) GetErrors() []PostCancelAccount409ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []PostCancelAccount409ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCancelAccount409Response) GetErrorsOk() ([]PostCancelAccount409ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PostCancelAccount409Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []PostCancelAccount409ResponseErrorsInner and assigns it to the Errors field.
func (o *PostCancelAccount409Response) SetErrors(v []PostCancelAccount409ResponseErrorsInner) {
	o.Errors = v
}

func (o PostCancelAccount409Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCancelAccount409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullablePostCancelAccount409Response struct {
	value *PostCancelAccount409Response
	isSet bool
}

func (v NullablePostCancelAccount409Response) Get() *PostCancelAccount409Response {
	return v.value
}

func (v *NullablePostCancelAccount409Response) Set(val *PostCancelAccount409Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCancelAccount409Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCancelAccount409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCancelAccount409Response(val *PostCancelAccount409Response) *NullablePostCancelAccount409Response {
	return &NullablePostCancelAccount409Response{value: val, isSet: true}
}

func (v NullablePostCancelAccount409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCancelAccount409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


