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

// checks if the GetTags200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTags200ResponseDataInner{}

// GetTags200ResponseDataInner A tag that has been applied to an object on your Account. Tags are currently for organizational purposes only.
type GetTags200ResponseDataInner struct {
	// A Label used for organization of objects on your Account.
	Label *string `json:"label,omitempty"`
}

// NewGetTags200ResponseDataInner instantiates a new GetTags200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTags200ResponseDataInner() *GetTags200ResponseDataInner {
	this := GetTags200ResponseDataInner{}
	return &this
}

// NewGetTags200ResponseDataInnerWithDefaults instantiates a new GetTags200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTags200ResponseDataInnerWithDefaults() *GetTags200ResponseDataInner {
	this := GetTags200ResponseDataInner{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetTags200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTags200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetTags200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetTags200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

func (o GetTags200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTags200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableGetTags200ResponseDataInner struct {
	value *GetTags200ResponseDataInner
	isSet bool
}

func (v NullableGetTags200ResponseDataInner) Get() *GetTags200ResponseDataInner {
	return v.value
}

func (v *NullableGetTags200ResponseDataInner) Set(val *GetTags200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTags200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTags200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTags200ResponseDataInner(val *GetTags200ResponseDataInner) *NullableGetTags200ResponseDataInner {
	return &NullableGetTags200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetTags200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTags200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


