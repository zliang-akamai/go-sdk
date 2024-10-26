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

// checks if the PutLkeNodePoolRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutLkeNodePoolRequest{}

// PutLkeNodePoolRequest struct for PutLkeNodePoolRequest
type PutLkeNodePoolRequest struct {
	Autoscaler *PostLkeClusterRequestNodePoolsInnerAutoscaler `json:"autoscaler,omitempty"`
	// The number of nodes in the Node Pool.
	Count *int32 `json:"count,omitempty"`
	// Key-value pairs added as labels to nodes in the node pool. Labels help classify your nodes and easily select subsets of objects. To learn more, review [Add Labels and Taints to your LKE node pools](https://www.linode.com/docs/products/compute/kubernetes/guides/deploy-and-manage-cluster-with-the-linode-api/#add-labels-and-taints-to-your-lke-node-pools).  Specifying an empty dictionary value will remove all previously set labels.
	Labels *[]interface{} `json:"labels,omitempty"`
	// Kubernetes taints to add to node pool nodes. Taints help control how pods are scheduled onto nodes, specifically allowing them to repel certain pods. To learn more, review [Add labels and taints to your LKE node pools](https://www.linode.com/docs/products/compute/kubernetes/guides/deploy-and-manage-cluster-with-the-linode-api/#add-labels-and-taints-to-your-lke-node-pools).  Specifying an empty array (`[]`) removes all previously set taints.
	Taints []PostLkeClusterRequestNodePoolsInnerTaintsInner `json:"taints,omitempty"`
}

// NewPutLkeNodePoolRequest instantiates a new PutLkeNodePoolRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutLkeNodePoolRequest() *PutLkeNodePoolRequest {
	this := PutLkeNodePoolRequest{}
	return &this
}

// NewPutLkeNodePoolRequestWithDefaults instantiates a new PutLkeNodePoolRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutLkeNodePoolRequestWithDefaults() *PutLkeNodePoolRequest {
	this := PutLkeNodePoolRequest{}
	return &this
}

// GetAutoscaler returns the Autoscaler field value if set, zero value otherwise.
func (o *PutLkeNodePoolRequest) GetAutoscaler() PostLkeClusterRequestNodePoolsInnerAutoscaler {
	if o == nil || IsNil(o.Autoscaler) {
		var ret PostLkeClusterRequestNodePoolsInnerAutoscaler
		return ret
	}
	return *o.Autoscaler
}

// GetAutoscalerOk returns a tuple with the Autoscaler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLkeNodePoolRequest) GetAutoscalerOk() (*PostLkeClusterRequestNodePoolsInnerAutoscaler, bool) {
	if o == nil || IsNil(o.Autoscaler) {
		return nil, false
	}
	return o.Autoscaler, true
}

// HasAutoscaler returns a boolean if a field has been set.
func (o *PutLkeNodePoolRequest) HasAutoscaler() bool {
	if o != nil && !IsNil(o.Autoscaler) {
		return true
	}

	return false
}

// SetAutoscaler gets a reference to the given PostLkeClusterRequestNodePoolsInnerAutoscaler and assigns it to the Autoscaler field.
func (o *PutLkeNodePoolRequest) SetAutoscaler(v PostLkeClusterRequestNodePoolsInnerAutoscaler) {
	o.Autoscaler = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PutLkeNodePoolRequest) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLkeNodePoolRequest) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PutLkeNodePoolRequest) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PutLkeNodePoolRequest) SetCount(v int32) {
	o.Count = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PutLkeNodePoolRequest) GetLabels() []interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret []interface{}
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLkeNodePoolRequest) GetLabelsOk() (*[]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PutLkeNodePoolRequest) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []interface{} and assigns it to the Labels field.
func (o *PutLkeNodePoolRequest) SetLabels(v []interface{}) {
	o.Labels = &v
}

// GetTaints returns the Taints field value if set, zero value otherwise.
func (o *PutLkeNodePoolRequest) GetTaints() []PostLkeClusterRequestNodePoolsInnerTaintsInner {
	if o == nil || IsNil(o.Taints) {
		var ret []PostLkeClusterRequestNodePoolsInnerTaintsInner
		return ret
	}
	return o.Taints
}

// GetTaintsOk returns a tuple with the Taints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLkeNodePoolRequest) GetTaintsOk() ([]PostLkeClusterRequestNodePoolsInnerTaintsInner, bool) {
	if o == nil || IsNil(o.Taints) {
		return nil, false
	}
	return o.Taints, true
}

// HasTaints returns a boolean if a field has been set.
func (o *PutLkeNodePoolRequest) HasTaints() bool {
	if o != nil && !IsNil(o.Taints) {
		return true
	}

	return false
}

// SetTaints gets a reference to the given []PostLkeClusterRequestNodePoolsInnerTaintsInner and assigns it to the Taints field.
func (o *PutLkeNodePoolRequest) SetTaints(v []PostLkeClusterRequestNodePoolsInnerTaintsInner) {
	o.Taints = v
}

func (o PutLkeNodePoolRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutLkeNodePoolRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Autoscaler) {
		toSerialize["autoscaler"] = o.Autoscaler
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Taints) {
		toSerialize["taints"] = o.Taints
	}
	return toSerialize, nil
}

type NullablePutLkeNodePoolRequest struct {
	value *PutLkeNodePoolRequest
	isSet bool
}

func (v NullablePutLkeNodePoolRequest) Get() *PutLkeNodePoolRequest {
	return v.value
}

func (v *NullablePutLkeNodePoolRequest) Set(val *PutLkeNodePoolRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutLkeNodePoolRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutLkeNodePoolRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutLkeNodePoolRequest(val *PutLkeNodePoolRequest) *NullablePutLkeNodePoolRequest {
	return &NullablePutLkeNodePoolRequest{value: val, isSet: true}
}

func (v NullablePutLkeNodePoolRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutLkeNodePoolRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

