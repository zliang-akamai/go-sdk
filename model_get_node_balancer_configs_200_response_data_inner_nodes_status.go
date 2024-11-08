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

// checks if the GetNodeBalancerConfigs200ResponseDataInnerNodesStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodeBalancerConfigs200ResponseDataInnerNodesStatus{}

// GetNodeBalancerConfigs200ResponseDataInnerNodesStatus A structure containing information about the health of the backends for this port.  This information is updated periodically as checks are performed against backends.
type GetNodeBalancerConfigs200ResponseDataInnerNodesStatus struct {
	// The number of backends considered to be \"DOWN\" and unhealthy.  These are not in rotation, and not serving requests.
	Down *int32 `json:"down,omitempty"`
	// The number of backends considered to be \"UP\" and healthy, and that are serving requests.
	Up *int32 `json:"up,omitempty"`
}

// NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatus instantiates a new GetNodeBalancerConfigs200ResponseDataInnerNodesStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatus() *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus {
	this := GetNodeBalancerConfigs200ResponseDataInnerNodesStatus{}
	return &this
}

// NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatusWithDefaults instantiates a new GetNodeBalancerConfigs200ResponseDataInnerNodesStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatusWithDefaults() *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus {
	this := GetNodeBalancerConfigs200ResponseDataInnerNodesStatus{}
	return &this
}

// GetDown returns the Down field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) GetDown() int32 {
	if o == nil || IsNil(o.Down) {
		var ret int32
		return ret
	}
	return *o.Down
}

// GetDownOk returns a tuple with the Down field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) GetDownOk() (*int32, bool) {
	if o == nil || IsNil(o.Down) {
		return nil, false
	}
	return o.Down, true
}

// HasDown returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) HasDown() bool {
	if o != nil && !IsNil(o.Down) {
		return true
	}

	return false
}

// SetDown gets a reference to the given int32 and assigns it to the Down field.
func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) SetDown(v int32) {
	o.Down = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) GetUp() int32 {
	if o == nil || IsNil(o.Up) {
		var ret int32
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) GetUpOk() (*int32, bool) {
	if o == nil || IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) HasUp() bool {
	if o != nil && !IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given int32 and assigns it to the Up field.
func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) SetUp(v int32) {
	o.Up = &v
}

func (o GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Down) {
		toSerialize["down"] = o.Down
	}
	if !IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}
	return toSerialize, nil
}

type NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus struct {
	value *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus
	isSet bool
}

func (v NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus) Get() *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus {
	return v.value
}

func (v *NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus) Set(val *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus(val *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) *NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus {
	return &NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus{value: val, isSet: true}
}

func (v NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodeBalancerConfigs200ResponseDataInnerNodesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


