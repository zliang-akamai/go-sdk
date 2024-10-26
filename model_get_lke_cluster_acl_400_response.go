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

// checks if the GetLkeClusterAcl400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLkeClusterAcl400Response{}

// GetLkeClusterAcl400Response struct for GetLkeClusterAcl400Response
type GetLkeClusterAcl400Response struct {
	Errors []GetLkeClusterAcl400ResponseErrorsInner `json:"errors,omitempty"`
}

// NewGetLkeClusterAcl400Response instantiates a new GetLkeClusterAcl400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLkeClusterAcl400Response() *GetLkeClusterAcl400Response {
	this := GetLkeClusterAcl400Response{}
	return &this
}

// NewGetLkeClusterAcl400ResponseWithDefaults instantiates a new GetLkeClusterAcl400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLkeClusterAcl400ResponseWithDefaults() *GetLkeClusterAcl400Response {
	this := GetLkeClusterAcl400Response{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetLkeClusterAcl400Response) GetErrors() []GetLkeClusterAcl400ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []GetLkeClusterAcl400ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeClusterAcl400Response) GetErrorsOk() ([]GetLkeClusterAcl400ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetLkeClusterAcl400Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GetLkeClusterAcl400ResponseErrorsInner and assigns it to the Errors field.
func (o *GetLkeClusterAcl400Response) SetErrors(v []GetLkeClusterAcl400ResponseErrorsInner) {
	o.Errors = v
}

func (o GetLkeClusterAcl400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLkeClusterAcl400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetLkeClusterAcl400Response struct {
	value *GetLkeClusterAcl400Response
	isSet bool
}

func (v NullableGetLkeClusterAcl400Response) Get() *GetLkeClusterAcl400Response {
	return v.value
}

func (v *NullableGetLkeClusterAcl400Response) Set(val *GetLkeClusterAcl400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLkeClusterAcl400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLkeClusterAcl400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLkeClusterAcl400Response(val *GetLkeClusterAcl400Response) *NullableGetLkeClusterAcl400Response {
	return &NullableGetLkeClusterAcl400Response{value: val, isSet: true}
}

func (v NullableGetLkeClusterAcl400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLkeClusterAcl400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

