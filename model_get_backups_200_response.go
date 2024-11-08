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

// checks if the GetBackups200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBackups200Response{}

// GetBackups200Response struct for GetBackups200Response
type GetBackups200Response struct {
	Automatic []GetBackups200ResponseAutomaticInner `json:"automatic,omitempty"`
	Snapshot *GetBackups200ResponseSnapshot `json:"snapshot,omitempty"`
}

// NewGetBackups200Response instantiates a new GetBackups200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBackups200Response() *GetBackups200Response {
	this := GetBackups200Response{}
	return &this
}

// NewGetBackups200ResponseWithDefaults instantiates a new GetBackups200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBackups200ResponseWithDefaults() *GetBackups200Response {
	this := GetBackups200Response{}
	return &this
}

// GetAutomatic returns the Automatic field value if set, zero value otherwise.
func (o *GetBackups200Response) GetAutomatic() []GetBackups200ResponseAutomaticInner {
	if o == nil || IsNil(o.Automatic) {
		var ret []GetBackups200ResponseAutomaticInner
		return ret
	}
	return o.Automatic
}

// GetAutomaticOk returns a tuple with the Automatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBackups200Response) GetAutomaticOk() ([]GetBackups200ResponseAutomaticInner, bool) {
	if o == nil || IsNil(o.Automatic) {
		return nil, false
	}
	return o.Automatic, true
}

// HasAutomatic returns a boolean if a field has been set.
func (o *GetBackups200Response) HasAutomatic() bool {
	if o != nil && !IsNil(o.Automatic) {
		return true
	}

	return false
}

// SetAutomatic gets a reference to the given []GetBackups200ResponseAutomaticInner and assigns it to the Automatic field.
func (o *GetBackups200Response) SetAutomatic(v []GetBackups200ResponseAutomaticInner) {
	o.Automatic = v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *GetBackups200Response) GetSnapshot() GetBackups200ResponseSnapshot {
	if o == nil || IsNil(o.Snapshot) {
		var ret GetBackups200ResponseSnapshot
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBackups200Response) GetSnapshotOk() (*GetBackups200ResponseSnapshot, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *GetBackups200Response) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given GetBackups200ResponseSnapshot and assigns it to the Snapshot field.
func (o *GetBackups200Response) SetSnapshot(v GetBackups200ResponseSnapshot) {
	o.Snapshot = &v
}

func (o GetBackups200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBackups200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Automatic) {
		toSerialize["automatic"] = o.Automatic
	}
	if !IsNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	return toSerialize, nil
}

type NullableGetBackups200Response struct {
	value *GetBackups200Response
	isSet bool
}

func (v NullableGetBackups200Response) Get() *GetBackups200Response {
	return v.value
}

func (v *NullableGetBackups200Response) Set(val *GetBackups200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBackups200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBackups200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBackups200Response(val *GetBackups200Response) *NullableGetBackups200Response {
	return &NullableGetBackups200Response{value: val, isSet: true}
}

func (v NullableGetBackups200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBackups200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


