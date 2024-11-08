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

// checks if the GetNodeBalancerStats200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodeBalancerStats200ResponseData{}

// GetNodeBalancerStats200ResponseData The data returned about this NodeBalancers.
type GetNodeBalancerStats200ResponseData struct {
	// An array of key/value pairs representing unix timestamp and reading for connections to this NodeBalancer.
	Connections []float32 `json:"connections,omitempty"`
	Traffic *GetNodeBalancerStats200ResponseDataTraffic `json:"traffic,omitempty"`
}

// NewGetNodeBalancerStats200ResponseData instantiates a new GetNodeBalancerStats200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodeBalancerStats200ResponseData() *GetNodeBalancerStats200ResponseData {
	this := GetNodeBalancerStats200ResponseData{}
	return &this
}

// NewGetNodeBalancerStats200ResponseDataWithDefaults instantiates a new GetNodeBalancerStats200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodeBalancerStats200ResponseDataWithDefaults() *GetNodeBalancerStats200ResponseData {
	this := GetNodeBalancerStats200ResponseData{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *GetNodeBalancerStats200ResponseData) GetConnections() []float32 {
	if o == nil || IsNil(o.Connections) {
		var ret []float32
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerStats200ResponseData) GetConnectionsOk() ([]float32, bool) {
	if o == nil || IsNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *GetNodeBalancerStats200ResponseData) HasConnections() bool {
	if o != nil && !IsNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []float32 and assigns it to the Connections field.
func (o *GetNodeBalancerStats200ResponseData) SetConnections(v []float32) {
	o.Connections = v
}

// GetTraffic returns the Traffic field value if set, zero value otherwise.
func (o *GetNodeBalancerStats200ResponseData) GetTraffic() GetNodeBalancerStats200ResponseDataTraffic {
	if o == nil || IsNil(o.Traffic) {
		var ret GetNodeBalancerStats200ResponseDataTraffic
		return ret
	}
	return *o.Traffic
}

// GetTrafficOk returns a tuple with the Traffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerStats200ResponseData) GetTrafficOk() (*GetNodeBalancerStats200ResponseDataTraffic, bool) {
	if o == nil || IsNil(o.Traffic) {
		return nil, false
	}
	return o.Traffic, true
}

// HasTraffic returns a boolean if a field has been set.
func (o *GetNodeBalancerStats200ResponseData) HasTraffic() bool {
	if o != nil && !IsNil(o.Traffic) {
		return true
	}

	return false
}

// SetTraffic gets a reference to the given GetNodeBalancerStats200ResponseDataTraffic and assigns it to the Traffic field.
func (o *GetNodeBalancerStats200ResponseData) SetTraffic(v GetNodeBalancerStats200ResponseDataTraffic) {
	o.Traffic = &v
}

func (o GetNodeBalancerStats200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodeBalancerStats200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	if !IsNil(o.Traffic) {
		toSerialize["traffic"] = o.Traffic
	}
	return toSerialize, nil
}

type NullableGetNodeBalancerStats200ResponseData struct {
	value *GetNodeBalancerStats200ResponseData
	isSet bool
}

func (v NullableGetNodeBalancerStats200ResponseData) Get() *GetNodeBalancerStats200ResponseData {
	return v.value
}

func (v *NullableGetNodeBalancerStats200ResponseData) Set(val *GetNodeBalancerStats200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodeBalancerStats200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodeBalancerStats200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodeBalancerStats200ResponseData(val *GetNodeBalancerStats200ResponseData) *NullableGetNodeBalancerStats200ResponseData {
	return &NullableGetNodeBalancerStats200ResponseData{value: val, isSet: true}
}

func (v NullableGetNodeBalancerStats200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodeBalancerStats200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


