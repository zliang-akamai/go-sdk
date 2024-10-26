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

// checks if the PostPayment202ResponseWarningsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPayment202ResponseWarningsInner{}

// PostPayment202ResponseWarningsInner An object for describing a single warning associated with a response.
type PostPayment202ResponseWarningsInner struct {
	// Specific information related to the warning.
	Details *string `json:"details,omitempty"`
	// The general warning message.
	Title *string `json:"title,omitempty"`
}

// NewPostPayment202ResponseWarningsInner instantiates a new PostPayment202ResponseWarningsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPayment202ResponseWarningsInner() *PostPayment202ResponseWarningsInner {
	this := PostPayment202ResponseWarningsInner{}
	return &this
}

// NewPostPayment202ResponseWarningsInnerWithDefaults instantiates a new PostPayment202ResponseWarningsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPayment202ResponseWarningsInnerWithDefaults() *PostPayment202ResponseWarningsInner {
	this := PostPayment202ResponseWarningsInner{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *PostPayment202ResponseWarningsInner) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPayment202ResponseWarningsInner) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PostPayment202ResponseWarningsInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *PostPayment202ResponseWarningsInner) SetDetails(v string) {
	o.Details = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PostPayment202ResponseWarningsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPayment202ResponseWarningsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PostPayment202ResponseWarningsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PostPayment202ResponseWarningsInner) SetTitle(v string) {
	o.Title = &v
}

func (o PostPayment202ResponseWarningsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPayment202ResponseWarningsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullablePostPayment202ResponseWarningsInner struct {
	value *PostPayment202ResponseWarningsInner
	isSet bool
}

func (v NullablePostPayment202ResponseWarningsInner) Get() *PostPayment202ResponseWarningsInner {
	return v.value
}

func (v *NullablePostPayment202ResponseWarningsInner) Set(val *PostPayment202ResponseWarningsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPayment202ResponseWarningsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPayment202ResponseWarningsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPayment202ResponseWarningsInner(val *PostPayment202ResponseWarningsInner) *NullablePostPayment202ResponseWarningsInner {
	return &NullablePostPayment202ResponseWarningsInner{value: val, isSet: true}
}

func (v NullablePostPayment202ResponseWarningsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPayment202ResponseWarningsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

