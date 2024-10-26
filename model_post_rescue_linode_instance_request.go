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

// checks if the PostRescueLinodeInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostRescueLinodeInstanceRequest{}

// PostRescueLinodeInstanceRequest struct for PostRescueLinodeInstanceRequest
type PostRescueLinodeInstanceRequest struct {
	Devices *PostRescueLinodeInstanceRequestDevices `json:"devices,omitempty"`
}

// NewPostRescueLinodeInstanceRequest instantiates a new PostRescueLinodeInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRescueLinodeInstanceRequest() *PostRescueLinodeInstanceRequest {
	this := PostRescueLinodeInstanceRequest{}
	return &this
}

// NewPostRescueLinodeInstanceRequestWithDefaults instantiates a new PostRescueLinodeInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRescueLinodeInstanceRequestWithDefaults() *PostRescueLinodeInstanceRequest {
	this := PostRescueLinodeInstanceRequest{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *PostRescueLinodeInstanceRequest) GetDevices() PostRescueLinodeInstanceRequestDevices {
	if o == nil || IsNil(o.Devices) {
		var ret PostRescueLinodeInstanceRequestDevices
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRescueLinodeInstanceRequest) GetDevicesOk() (*PostRescueLinodeInstanceRequestDevices, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *PostRescueLinodeInstanceRequest) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given PostRescueLinodeInstanceRequestDevices and assigns it to the Devices field.
func (o *PostRescueLinodeInstanceRequest) SetDevices(v PostRescueLinodeInstanceRequestDevices) {
	o.Devices = &v
}

func (o PostRescueLinodeInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRescueLinodeInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	return toSerialize, nil
}

type NullablePostRescueLinodeInstanceRequest struct {
	value *PostRescueLinodeInstanceRequest
	isSet bool
}

func (v NullablePostRescueLinodeInstanceRequest) Get() *PostRescueLinodeInstanceRequest {
	return v.value
}

func (v *NullablePostRescueLinodeInstanceRequest) Set(val *PostRescueLinodeInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRescueLinodeInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRescueLinodeInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRescueLinodeInstanceRequest(val *PostRescueLinodeInstanceRequest) *NullablePostRescueLinodeInstanceRequest {
	return &NullablePostRescueLinodeInstanceRequest{value: val, isSet: true}
}

func (v NullablePostRescueLinodeInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRescueLinodeInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


