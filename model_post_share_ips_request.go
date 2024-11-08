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
	"bytes"
	"fmt"
)

// checks if the PostShareIpsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostShareIpsRequest{}

// PostShareIpsRequest A request object IP Addresses Share (POST /networking/ips/share).
type PostShareIpsRequest struct {
	// A list of secondary Linode IPs to share with the primary Linode.  - Can include both IPv4 addresses and IPv6 ranges (omit /56 and /64 prefix lengths) - Can include both private and public IPv4 addresses. - You must have access to all of these addresses and they must be in the same Region as the primary Linode. - Enter an empty array to remove all shared IP addresses.
	Ips []string `json:"ips"`
	// The ID of the primary Linode that the addresses will be shared with.
	LinodeId int32 `json:"linode_id"`
}

type _PostShareIpsRequest PostShareIpsRequest

// NewPostShareIpsRequest instantiates a new PostShareIpsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostShareIpsRequest(ips []string, linodeId int32) *PostShareIpsRequest {
	this := PostShareIpsRequest{}
	this.Ips = ips
	this.LinodeId = linodeId
	return &this
}

// NewPostShareIpsRequestWithDefaults instantiates a new PostShareIpsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostShareIpsRequestWithDefaults() *PostShareIpsRequest {
	this := PostShareIpsRequest{}
	return &this
}

// GetIps returns the Ips field value
func (o *PostShareIpsRequest) GetIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *PostShareIpsRequest) GetIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *PostShareIpsRequest) SetIps(v []string) {
	o.Ips = v
}

// GetLinodeId returns the LinodeId field value
func (o *PostShareIpsRequest) GetLinodeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LinodeId
}

// GetLinodeIdOk returns a tuple with the LinodeId field value
// and a boolean to check if the value has been set.
func (o *PostShareIpsRequest) GetLinodeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinodeId, true
}

// SetLinodeId sets field value
func (o *PostShareIpsRequest) SetLinodeId(v int32) {
	o.LinodeId = v
}

func (o PostShareIpsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostShareIpsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ips"] = o.Ips
	toSerialize["linode_id"] = o.LinodeId
	return toSerialize, nil
}

func (o *PostShareIpsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ips",
		"linode_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostShareIpsRequest := _PostShareIpsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostShareIpsRequest)

	if err != nil {
		return err
	}

	*o = PostShareIpsRequest(varPostShareIpsRequest)

	return err
}

type NullablePostShareIpsRequest struct {
	value *PostShareIpsRequest
	isSet bool
}

func (v NullablePostShareIpsRequest) Get() *PostShareIpsRequest {
	return v.value
}

func (v *NullablePostShareIpsRequest) Set(val *PostShareIpsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostShareIpsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostShareIpsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostShareIpsRequest(val *PostShareIpsRequest) *NullablePostShareIpsRequest {
	return &NullablePostShareIpsRequest{value: val, isSet: true}
}

func (v NullablePostShareIpsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostShareIpsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


