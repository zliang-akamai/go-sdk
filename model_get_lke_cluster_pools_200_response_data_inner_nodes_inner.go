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

// checks if the GetLkeClusterPools200ResponseDataInnerNodesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLkeClusterPools200ResponseDataInnerNodesInner{}

// GetLkeClusterPools200ResponseDataInnerNodesInner Status information for a Node which is a member of a Kubernetes cluster.
type GetLkeClusterPools200ResponseDataInnerNodesInner struct {
	// The Node's ID.
	Id *string `json:"id,omitempty"`
	// The Linode's ID. When no Linode is currently provisioned for this Node, this will be null.
	InstanceId *string `json:"instance_id,omitempty"`
	// The Node's status as it pertains to being a Kubernetes node.
	Status *string `json:"status,omitempty"`
}

// NewGetLkeClusterPools200ResponseDataInnerNodesInner instantiates a new GetLkeClusterPools200ResponseDataInnerNodesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLkeClusterPools200ResponseDataInnerNodesInner() *GetLkeClusterPools200ResponseDataInnerNodesInner {
	this := GetLkeClusterPools200ResponseDataInnerNodesInner{}
	return &this
}

// NewGetLkeClusterPools200ResponseDataInnerNodesInnerWithDefaults instantiates a new GetLkeClusterPools200ResponseDataInnerNodesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLkeClusterPools200ResponseDataInnerNodesInnerWithDefaults() *GetLkeClusterPools200ResponseDataInnerNodesInner {
	this := GetLkeClusterPools200ResponseDataInnerNodesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) SetId(v string) {
	o.Id = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetLkeClusterPools200ResponseDataInnerNodesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLkeClusterPools200ResponseDataInnerNodesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instance_id"] = o.InstanceId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGetLkeClusterPools200ResponseDataInnerNodesInner struct {
	value *GetLkeClusterPools200ResponseDataInnerNodesInner
	isSet bool
}

func (v NullableGetLkeClusterPools200ResponseDataInnerNodesInner) Get() *GetLkeClusterPools200ResponseDataInnerNodesInner {
	return v.value
}

func (v *NullableGetLkeClusterPools200ResponseDataInnerNodesInner) Set(val *GetLkeClusterPools200ResponseDataInnerNodesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLkeClusterPools200ResponseDataInnerNodesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLkeClusterPools200ResponseDataInnerNodesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLkeClusterPools200ResponseDataInnerNodesInner(val *GetLkeClusterPools200ResponseDataInnerNodesInner) *NullableGetLkeClusterPools200ResponseDataInnerNodesInner {
	return &NullableGetLkeClusterPools200ResponseDataInnerNodesInner{value: val, isSet: true}
}

func (v NullableGetLkeClusterPools200ResponseDataInnerNodesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLkeClusterPools200ResponseDataInnerNodesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


