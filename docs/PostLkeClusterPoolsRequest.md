# PostLkeClusterPoolsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoscaler** | Pointer to [**PostLkeClusterRequestNodePoolsInnerAutoscaler**](PostLkeClusterRequestNodePoolsInnerAutoscaler.md) |  | [optional] 
**Count** | **int32** | The number of nodes in the Node Pool. | 
**Disks** | Pointer to [**[]PostLkeClusterRequestNodePoolsInnerDisksInner**](PostLkeClusterRequestNodePoolsInnerDisksInner.md) | This node pool&#39;s custom disk layout. Each item in this array will create a new disk partition for each node in this node pool.  &gt; 📘 &gt; &gt; Omit this field, except for special use cases. The disks specified here are partitions in _addition_ to the primary partition and reduce the size of the primary partition. This can lead to stability problems for the node.    - The custom disk layout is applied to each node in this node pool.    - The maximum number of custom disk partitions that can be configured is 7.    - Once the requested disk partitions are allocated, the remaining disk space is allocated to the node&#39;s boot disk.    - A node pool&#39;s custom disk layout is immutable over the lifetime of the node pool. | [optional] 
**Labels** | Pointer to  | Key-value pairs added as labels to nodes in the node pool. Labels help classify your nodes and easily select subsets of objects. To learn more, review [Add Labels and Taints to your LKE node pools](https://www.linode.com/docs/products/compute/kubernetes/guides/deploy-and-manage-cluster-with-the-linode-api/#add-labels-and-taints-to-your-lke-node-pools).  Specifying an empty dictionary value will remove all previously set labels. | [optional] 
**Tags** | Pointer to **[]string** | An array of tags applied to this object. Tags are for organizational purposes only. | [optional] 
**Taints** | Pointer to [**[]PostLkeClusterRequestNodePoolsInnerTaintsInner**](PostLkeClusterRequestNodePoolsInnerTaintsInner.md) | Kubernetes taints to add to node pool nodes. Taints help control how pods are scheduled onto nodes, specifically allowing them to repel certain pods. To learn more, review [Add labels and taints to your LKE node pools](https://www.linode.com/docs/products/compute/kubernetes/guides/deploy-and-manage-cluster-with-the-linode-api/#add-labels-and-taints-to-your-lke-node-pools).  Specifying an empty array (&#x60;[]&#x60;) removes all previously set taints. | [optional] 
**Type** | **string** | The Linode Type for all of the nodes in the Node Pool. | 

## Methods

### NewPostLkeClusterPoolsRequest

`func NewPostLkeClusterPoolsRequest(count int32, type_ string, ) *PostLkeClusterPoolsRequest`

NewPostLkeClusterPoolsRequest instantiates a new PostLkeClusterPoolsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLkeClusterPoolsRequestWithDefaults

`func NewPostLkeClusterPoolsRequestWithDefaults() *PostLkeClusterPoolsRequest`

NewPostLkeClusterPoolsRequestWithDefaults instantiates a new PostLkeClusterPoolsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscaler

`func (o *PostLkeClusterPoolsRequest) GetAutoscaler() PostLkeClusterRequestNodePoolsInnerAutoscaler`

GetAutoscaler returns the Autoscaler field if non-nil, zero value otherwise.

### GetAutoscalerOk

`func (o *PostLkeClusterPoolsRequest) GetAutoscalerOk() (*PostLkeClusterRequestNodePoolsInnerAutoscaler, bool)`

GetAutoscalerOk returns a tuple with the Autoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaler

`func (o *PostLkeClusterPoolsRequest) SetAutoscaler(v PostLkeClusterRequestNodePoolsInnerAutoscaler)`

SetAutoscaler sets Autoscaler field to given value.

### HasAutoscaler

`func (o *PostLkeClusterPoolsRequest) HasAutoscaler() bool`

HasAutoscaler returns a boolean if a field has been set.

### GetCount

`func (o *PostLkeClusterPoolsRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PostLkeClusterPoolsRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PostLkeClusterPoolsRequest) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDisks

`func (o *PostLkeClusterPoolsRequest) GetDisks() []PostLkeClusterRequestNodePoolsInnerDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *PostLkeClusterPoolsRequest) GetDisksOk() (*[]PostLkeClusterRequestNodePoolsInnerDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *PostLkeClusterPoolsRequest) SetDisks(v []PostLkeClusterRequestNodePoolsInnerDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *PostLkeClusterPoolsRequest) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetLabels

`func (o *PostLkeClusterPoolsRequest) GetLabels() []interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PostLkeClusterPoolsRequest) GetLabelsOk() (*[]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PostLkeClusterPoolsRequest) SetLabels(v []interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PostLkeClusterPoolsRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *PostLkeClusterPoolsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostLkeClusterPoolsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostLkeClusterPoolsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostLkeClusterPoolsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaints

`func (o *PostLkeClusterPoolsRequest) GetTaints() []PostLkeClusterRequestNodePoolsInnerTaintsInner`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *PostLkeClusterPoolsRequest) GetTaintsOk() (*[]PostLkeClusterRequestNodePoolsInnerTaintsInner, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *PostLkeClusterPoolsRequest) SetTaints(v []PostLkeClusterRequestNodePoolsInnerTaintsInner)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *PostLkeClusterPoolsRequest) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### GetType

`func (o *PostLkeClusterPoolsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostLkeClusterPoolsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostLkeClusterPoolsRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


