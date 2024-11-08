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

// checks if the WarningObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarningObject{}

// WarningObject An object for describing a single warning associated with a response.
type WarningObject struct {
	// Specific information related to the warning.
	Details *string `json:"details,omitempty"`
	// The general warning message.
	Title *string `json:"title,omitempty"`
}

// NewWarningObject instantiates a new WarningObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarningObject() *WarningObject {
	this := WarningObject{}
	return &this
}

// NewWarningObjectWithDefaults instantiates a new WarningObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningObjectWithDefaults() *WarningObject {
	this := WarningObject{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *WarningObject) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningObject) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *WarningObject) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *WarningObject) SetDetails(v string) {
	o.Details = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WarningObject) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningObject) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WarningObject) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WarningObject) SetTitle(v string) {
	o.Title = &v
}

func (o WarningObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarningObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableWarningObject struct {
	value *WarningObject
	isSet bool
}

func (v NullableWarningObject) Get() *WarningObject {
	return v.value
}

func (v *NullableWarningObject) Set(val *WarningObject) {
	v.value = val
	v.isSet = true
}

func (v NullableWarningObject) IsSet() bool {
	return v.isSet
}

func (v *NullableWarningObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarningObject(val *WarningObject) *NullableWarningObject {
	return &NullableWarningObject{value: val, isSet: true}
}

func (v NullableWarningObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarningObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


