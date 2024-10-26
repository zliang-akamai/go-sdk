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

// checks if the GetLkeVersions200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLkeVersions200ResponseDataInner{}

// GetLkeVersions200ResponseDataInner LKE versions.
type GetLkeVersions200ResponseDataInner struct {
	// A Kubernetes version number available for deployment to a Kubernetes cluster in the format of &lt;major&gt;.&lt;minor&gt;, and the latest supported patch version.
	Id *string `json:"id,omitempty"`
}

// NewGetLkeVersions200ResponseDataInner instantiates a new GetLkeVersions200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLkeVersions200ResponseDataInner() *GetLkeVersions200ResponseDataInner {
	this := GetLkeVersions200ResponseDataInner{}
	return &this
}

// NewGetLkeVersions200ResponseDataInnerWithDefaults instantiates a new GetLkeVersions200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLkeVersions200ResponseDataInnerWithDefaults() *GetLkeVersions200ResponseDataInner {
	this := GetLkeVersions200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetLkeVersions200ResponseDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeVersions200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetLkeVersions200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetLkeVersions200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

func (o GetLkeVersions200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLkeVersions200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetLkeVersions200ResponseDataInner struct {
	value *GetLkeVersions200ResponseDataInner
	isSet bool
}

func (v NullableGetLkeVersions200ResponseDataInner) Get() *GetLkeVersions200ResponseDataInner {
	return v.value
}

func (v *NullableGetLkeVersions200ResponseDataInner) Set(val *GetLkeVersions200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLkeVersions200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLkeVersions200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLkeVersions200ResponseDataInner(val *GetLkeVersions200ResponseDataInner) *NullableGetLkeVersions200ResponseDataInner {
	return &NullableGetLkeVersions200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetLkeVersions200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLkeVersions200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

