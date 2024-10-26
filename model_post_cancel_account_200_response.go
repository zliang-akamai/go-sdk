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

// checks if the PostCancelAccount200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCancelAccount200Response{}

// PostCancelAccount200Response struct for PostCancelAccount200Response
type PostCancelAccount200Response struct {
	// A link to Linode's exit survey.
	SurveyLink *string `json:"survey_link,omitempty"`
}

// NewPostCancelAccount200Response instantiates a new PostCancelAccount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCancelAccount200Response() *PostCancelAccount200Response {
	this := PostCancelAccount200Response{}
	return &this
}

// NewPostCancelAccount200ResponseWithDefaults instantiates a new PostCancelAccount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCancelAccount200ResponseWithDefaults() *PostCancelAccount200Response {
	this := PostCancelAccount200Response{}
	return &this
}

// GetSurveyLink returns the SurveyLink field value if set, zero value otherwise.
func (o *PostCancelAccount200Response) GetSurveyLink() string {
	if o == nil || IsNil(o.SurveyLink) {
		var ret string
		return ret
	}
	return *o.SurveyLink
}

// GetSurveyLinkOk returns a tuple with the SurveyLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCancelAccount200Response) GetSurveyLinkOk() (*string, bool) {
	if o == nil || IsNil(o.SurveyLink) {
		return nil, false
	}
	return o.SurveyLink, true
}

// HasSurveyLink returns a boolean if a field has been set.
func (o *PostCancelAccount200Response) HasSurveyLink() bool {
	if o != nil && !IsNil(o.SurveyLink) {
		return true
	}

	return false
}

// SetSurveyLink gets a reference to the given string and assigns it to the SurveyLink field.
func (o *PostCancelAccount200Response) SetSurveyLink(v string) {
	o.SurveyLink = &v
}

func (o PostCancelAccount200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCancelAccount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SurveyLink) {
		toSerialize["survey_link"] = o.SurveyLink
	}
	return toSerialize, nil
}

type NullablePostCancelAccount200Response struct {
	value *PostCancelAccount200Response
	isSet bool
}

func (v NullablePostCancelAccount200Response) Get() *PostCancelAccount200Response {
	return v.value
}

func (v *NullablePostCancelAccount200Response) Set(val *PostCancelAccount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCancelAccount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCancelAccount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCancelAccount200Response(val *PostCancelAccount200Response) *NullablePostCancelAccount200Response {
	return &NullablePostCancelAccount200Response{value: val, isSet: true}
}

func (v NullablePostCancelAccount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCancelAccount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


