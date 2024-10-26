# PutLkeNodePoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoscaler** | Pointer to [**PostLkeClusterRequestNodePoolsInnerAutoscaler**](PostLkeClusterRequestNodePoolsInnerAutoscaler.md) |  | [optional] 
**Count** | Pointer to **int32** | The number of nodes in the Node Pool. | [optional] 
**Labels** | Pointer to  | Key-value pairs added as labels to nodes in the node pool. Labels help classify your nodes and easily select subsets of objects. To learn more, review [Add Labels and Taints to your LKE node pools](https://www.linode.com/docs/products/compute/kubernetes/guides/deploy-and-manage-cluster-with-the-linode-api/#add-labels-and-taints-to-your-lke-node-pools).  Specifying an empty dictionary value will remove all previously set labels. | [optional] 
**Taints** | Pointer to [**[]PostLkeClusterRequestNodePoolsInnerTaintsInner**](PostLkeClusterRequestNodePoolsInnerTaintsInner.md) | Kubernetes taints to add to node pool nodes. Taints help control how pods are scheduled onto nodes, specifically allowing them to repel certain pods. To learn more, review [Add labels and taints to your LKE node pools](https://www.linode.com/docs/products/compute/kubernetes/guides/deploy-and-manage-cluster-with-the-linode-api/#add-labels-and-taints-to-your-lke-node-pools).  Specifying an empty array (&#x60;[]&#x60;) removes all previously set taints. | [optional] 

## Methods

### NewPutLkeNodePoolRequest

`func NewPutLkeNodePoolRequest() *PutLkeNodePoolRequest`

NewPutLkeNodePoolRequest instantiates a new PutLkeNodePoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLkeNodePoolRequestWithDefaults

`func NewPutLkeNodePoolRequestWithDefaults() *PutLkeNodePoolRequest`

NewPutLkeNodePoolRequestWithDefaults instantiates a new PutLkeNodePoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscaler

`func (o *PutLkeNodePoolRequest) GetAutoscaler() PostLkeClusterRequestNodePoolsInnerAutoscaler`

GetAutoscaler returns the Autoscaler field if non-nil, zero value otherwise.

### GetAutoscalerOk

`func (o *PutLkeNodePoolRequest) GetAutoscalerOk() (*PostLkeClusterRequestNodePoolsInnerAutoscaler, bool)`

GetAutoscalerOk returns a tuple with the Autoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaler

`func (o *PutLkeNodePoolRequest) SetAutoscaler(v PostLkeClusterRequestNodePoolsInnerAutoscaler)`

SetAutoscaler sets Autoscaler field to given value.

### HasAutoscaler

`func (o *PutLkeNodePoolRequest) HasAutoscaler() bool`

HasAutoscaler returns a boolean if a field has been set.

### GetCount

`func (o *PutLkeNodePoolRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PutLkeNodePoolRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PutLkeNodePoolRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PutLkeNodePoolRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLabels

`func (o *PutLkeNodePoolRequest) GetLabels() []interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PutLkeNodePoolRequest) GetLabelsOk() (*[]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PutLkeNodePoolRequest) SetLabels(v []interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PutLkeNodePoolRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTaints

`func (o *PutLkeNodePoolRequest) GetTaints() []PostLkeClusterRequestNodePoolsInnerTaintsInner`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *PutLkeNodePoolRequest) GetTaintsOk() (*[]PostLkeClusterRequestNodePoolsInnerTaintsInner, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *PutLkeNodePoolRequest) SetTaints(v []PostLkeClusterRequestNodePoolsInnerTaintsInner)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *PutLkeNodePoolRequest) HasTaints() bool`

HasTaints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


