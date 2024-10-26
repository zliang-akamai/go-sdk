# GetLkeClusterPools200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoscaler** | Pointer to [**GetLkeClusterPools200ResponseDataInnerAutoscaler**](GetLkeClusterPools200ResponseDataInnerAutoscaler.md) |  | [optional] 
**Count** | Pointer to **int32** | The number of nodes in the Node Pool. | [optional] 
**DiskEncryption** | Pointer to **string** | For new LKE node pools, &#x60;disk_encryption&#x60; is automatically &#x60;enabled&#x60; where disk encryption is supported. It can&#39;t be &#x60;disabled&#x60;. For existing LKE node pools, this derives from the Linode&#39;s &#x60;disk_encryption&#x60; setting. If a Linode&#39;s node pool is not encrypted and you want an encrypted node pool, delete the node pool and create a new node pool. | [optional] 
**Disks** | Pointer to [**[]PostLkeClusterRequestNodePoolsInnerDisksInner**](PostLkeClusterRequestNodePoolsInnerDisksInner.md) | This Node Pool&#39;s custom disk layout. | [optional] 
**Id** | Pointer to **int32** | This Node Pool&#39;s unique ID. | [optional] 
**Labels** | Pointer to [**GetLkeClusterPools200ResponseDataInnerLabels**](GetLkeClusterPools200ResponseDataInnerLabels.md) |  | [optional] 
**Nodes** | Pointer to [**[]GetLkeClusterPools200ResponseDataInnerNodesInner**](GetLkeClusterPools200ResponseDataInnerNodesInner.md) | Status information for the Nodes which are members of this Node Pool. If a Linode has not been provisioned for a given Node slot, the instance_id will be returned as null. | [optional] 
**Tags** | Pointer to **[]string** | An array of tags applied to this object. Tags are for organizational purposes only. | [optional] 
**Taints** | Pointer to [**[]GetLkeClusterPools200ResponseDataInnerTaintsInner**](GetLkeClusterPools200ResponseDataInnerTaintsInner.md) | Kubernetes taints added to nodes in the node pool. Taints help control how pods are scheduled onto nodes, specifically allowing them to repel certain pods. | [optional] 
**Type** | Pointer to **string** | The Linode Type for all of the nodes in the Node Pool. | [optional] 

## Methods

### NewGetLkeClusterPools200ResponseDataInner

`func NewGetLkeClusterPools200ResponseDataInner() *GetLkeClusterPools200ResponseDataInner`

NewGetLkeClusterPools200ResponseDataInner instantiates a new GetLkeClusterPools200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLkeClusterPools200ResponseDataInnerWithDefaults

`func NewGetLkeClusterPools200ResponseDataInnerWithDefaults() *GetLkeClusterPools200ResponseDataInner`

NewGetLkeClusterPools200ResponseDataInnerWithDefaults instantiates a new GetLkeClusterPools200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscaler

`func (o *GetLkeClusterPools200ResponseDataInner) GetAutoscaler() GetLkeClusterPools200ResponseDataInnerAutoscaler`

GetAutoscaler returns the Autoscaler field if non-nil, zero value otherwise.

### GetAutoscalerOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetAutoscalerOk() (*GetLkeClusterPools200ResponseDataInnerAutoscaler, bool)`

GetAutoscalerOk returns a tuple with the Autoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaler

`func (o *GetLkeClusterPools200ResponseDataInner) SetAutoscaler(v GetLkeClusterPools200ResponseDataInnerAutoscaler)`

SetAutoscaler sets Autoscaler field to given value.

### HasAutoscaler

`func (o *GetLkeClusterPools200ResponseDataInner) HasAutoscaler() bool`

HasAutoscaler returns a boolean if a field has been set.

### GetCount

`func (o *GetLkeClusterPools200ResponseDataInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetLkeClusterPools200ResponseDataInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetLkeClusterPools200ResponseDataInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *GetLkeClusterPools200ResponseDataInner) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *GetLkeClusterPools200ResponseDataInner) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *GetLkeClusterPools200ResponseDataInner) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetDisks

`func (o *GetLkeClusterPools200ResponseDataInner) GetDisks() []PostLkeClusterRequestNodePoolsInnerDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetDisksOk() (*[]PostLkeClusterRequestNodePoolsInnerDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *GetLkeClusterPools200ResponseDataInner) SetDisks(v []PostLkeClusterRequestNodePoolsInnerDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *GetLkeClusterPools200ResponseDataInner) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetId

`func (o *GetLkeClusterPools200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLkeClusterPools200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetLkeClusterPools200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *GetLkeClusterPools200ResponseDataInner) GetLabels() GetLkeClusterPools200ResponseDataInnerLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetLabelsOk() (*GetLkeClusterPools200ResponseDataInnerLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetLkeClusterPools200ResponseDataInner) SetLabels(v GetLkeClusterPools200ResponseDataInnerLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetLkeClusterPools200ResponseDataInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNodes

`func (o *GetLkeClusterPools200ResponseDataInner) GetNodes() []GetLkeClusterPools200ResponseDataInnerNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetNodesOk() (*[]GetLkeClusterPools200ResponseDataInnerNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetLkeClusterPools200ResponseDataInner) SetNodes(v []GetLkeClusterPools200ResponseDataInnerNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GetLkeClusterPools200ResponseDataInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTags

`func (o *GetLkeClusterPools200ResponseDataInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetLkeClusterPools200ResponseDataInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetLkeClusterPools200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaints

`func (o *GetLkeClusterPools200ResponseDataInner) GetTaints() []GetLkeClusterPools200ResponseDataInnerTaintsInner`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetTaintsOk() (*[]GetLkeClusterPools200ResponseDataInnerTaintsInner, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *GetLkeClusterPools200ResponseDataInner) SetTaints(v []GetLkeClusterPools200ResponseDataInnerTaintsInner)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *GetLkeClusterPools200ResponseDataInner) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### GetType

`func (o *GetLkeClusterPools200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLkeClusterPools200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLkeClusterPools200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetLkeClusterPools200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


