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

// checks if the PostIpv6RangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostIpv6RangeRequest{}

// PostIpv6RangeRequest struct for PostIpv6RangeRequest
type PostIpv6RangeRequest struct {
	// The ID of the Linode to assign this range to. The SLAAC address for the provided Linode is used as the range's `route_target`.  - __Required__ if `route_target` is omitted from the request.  - Mutually exclusive with `route_target`. Submitting values for both properties in a request results in an error.
	LinodeId *int32 `json:"linode_id,omitempty"`
	// The prefix length of the IPv6 range.
	PrefixLength int32 `json:"prefix_length"`
	// The IPv6 SLAAC address to assign this range to.  - __Required__ if `linode_id` is omitted from the request.  - Mutually exclusive with `linode_id`. Submitting values for both properties in a request results in an error.  - __Note__. Omit the `/128` prefix length of the SLAAC address when using this property.
	RouteTarget *string `json:"route_target,omitempty"`
}

type _PostIpv6RangeRequest PostIpv6RangeRequest

// NewPostIpv6RangeRequest instantiates a new PostIpv6RangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostIpv6RangeRequest(prefixLength int32) *PostIpv6RangeRequest {
	this := PostIpv6RangeRequest{}
	this.PrefixLength = prefixLength
	return &this
}

// NewPostIpv6RangeRequestWithDefaults instantiates a new PostIpv6RangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostIpv6RangeRequestWithDefaults() *PostIpv6RangeRequest {
	this := PostIpv6RangeRequest{}
	return &this
}

// GetLinodeId returns the LinodeId field value if set, zero value otherwise.
func (o *PostIpv6RangeRequest) GetLinodeId() int32 {
	if o == nil || IsNil(o.LinodeId) {
		var ret int32
		return ret
	}
	return *o.LinodeId
}

// GetLinodeIdOk returns a tuple with the LinodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostIpv6RangeRequest) GetLinodeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LinodeId) {
		return nil, false
	}
	return o.LinodeId, true
}

// HasLinodeId returns a boolean if a field has been set.
func (o *PostIpv6RangeRequest) HasLinodeId() bool {
	if o != nil && !IsNil(o.LinodeId) {
		return true
	}

	return false
}

// SetLinodeId gets a reference to the given int32 and assigns it to the LinodeId field.
func (o *PostIpv6RangeRequest) SetLinodeId(v int32) {
	o.LinodeId = &v
}

// GetPrefixLength returns the PrefixLength field value
func (o *PostIpv6RangeRequest) GetPrefixLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *PostIpv6RangeRequest) GetPrefixLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *PostIpv6RangeRequest) SetPrefixLength(v int32) {
	o.PrefixLength = v
}

// GetRouteTarget returns the RouteTarget field value if set, zero value otherwise.
func (o *PostIpv6RangeRequest) GetRouteTarget() string {
	if o == nil || IsNil(o.RouteTarget) {
		var ret string
		return ret
	}
	return *o.RouteTarget
}

// GetRouteTargetOk returns a tuple with the RouteTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostIpv6RangeRequest) GetRouteTargetOk() (*string, bool) {
	if o == nil || IsNil(o.RouteTarget) {
		return nil, false
	}
	return o.RouteTarget, true
}

// HasRouteTarget returns a boolean if a field has been set.
func (o *PostIpv6RangeRequest) HasRouteTarget() bool {
	if o != nil && !IsNil(o.RouteTarget) {
		return true
	}

	return false
}

// SetRouteTarget gets a reference to the given string and assigns it to the RouteTarget field.
func (o *PostIpv6RangeRequest) SetRouteTarget(v string) {
	o.RouteTarget = &v
}

func (o PostIpv6RangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostIpv6RangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinodeId) {
		toSerialize["linode_id"] = o.LinodeId
	}
	toSerialize["prefix_length"] = o.PrefixLength
	if !IsNil(o.RouteTarget) {
		toSerialize["route_target"] = o.RouteTarget
	}
	return toSerialize, nil
}

func (o *PostIpv6RangeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prefix_length",
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

	varPostIpv6RangeRequest := _PostIpv6RangeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostIpv6RangeRequest)

	if err != nil {
		return err
	}

	*o = PostIpv6RangeRequest(varPostIpv6RangeRequest)

	return err
}

type NullablePostIpv6RangeRequest struct {
	value *PostIpv6RangeRequest
	isSet bool
}

func (v NullablePostIpv6RangeRequest) Get() *PostIpv6RangeRequest {
	return v.value
}

func (v *NullablePostIpv6RangeRequest) Set(val *PostIpv6RangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostIpv6RangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostIpv6RangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostIpv6RangeRequest(val *PostIpv6RangeRequest) *NullablePostIpv6RangeRequest {
	return &NullablePostIpv6RangeRequest{value: val, isSet: true}
}

func (v NullablePostIpv6RangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostIpv6RangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


