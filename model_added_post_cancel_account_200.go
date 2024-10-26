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

// checks if the AddedPostCancelAccount200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedPostCancelAccount200{}

// AddedPostCancelAccount200 struct for AddedPostCancelAccount200
type AddedPostCancelAccount200 struct {
	// A link to Linode's exit survey.
	SurveyLink *string `json:"survey_link,omitempty"`
}

// NewAddedPostCancelAccount200 instantiates a new AddedPostCancelAccount200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedPostCancelAccount200() *AddedPostCancelAccount200 {
	this := AddedPostCancelAccount200{}
	return &this
}

// NewAddedPostCancelAccount200WithDefaults instantiates a new AddedPostCancelAccount200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedPostCancelAccount200WithDefaults() *AddedPostCancelAccount200 {
	this := AddedPostCancelAccount200{}
	return &this
}

// GetSurveyLink returns the SurveyLink field value if set, zero value otherwise.
func (o *AddedPostCancelAccount200) GetSurveyLink() string {
	if o == nil || IsNil(o.SurveyLink) {
		var ret string
		return ret
	}
	return *o.SurveyLink
}

// GetSurveyLinkOk returns a tuple with the SurveyLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedPostCancelAccount200) GetSurveyLinkOk() (*string, bool) {
	if o == nil || IsNil(o.SurveyLink) {
		return nil, false
	}
	return o.SurveyLink, true
}

// HasSurveyLink returns a boolean if a field has been set.
func (o *AddedPostCancelAccount200) HasSurveyLink() bool {
	if o != nil && !IsNil(o.SurveyLink) {
		return true
	}

	return false
}

// SetSurveyLink gets a reference to the given string and assigns it to the SurveyLink field.
func (o *AddedPostCancelAccount200) SetSurveyLink(v string) {
	o.SurveyLink = &v
}

func (o AddedPostCancelAccount200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedPostCancelAccount200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SurveyLink) {
		toSerialize["survey_link"] = o.SurveyLink
	}
	return toSerialize, nil
}

type NullableAddedPostCancelAccount200 struct {
	value *AddedPostCancelAccount200
	isSet bool
}

func (v NullableAddedPostCancelAccount200) Get() *AddedPostCancelAccount200 {
	return v.value
}

func (v *NullableAddedPostCancelAccount200) Set(val *AddedPostCancelAccount200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedPostCancelAccount200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedPostCancelAccount200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedPostCancelAccount200(val *AddedPostCancelAccount200) *NullableAddedPostCancelAccount200 {
	return &NullableAddedPostCancelAccount200{value: val, isSet: true}
}

func (v NullableAddedPostCancelAccount200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedPostCancelAccount200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


