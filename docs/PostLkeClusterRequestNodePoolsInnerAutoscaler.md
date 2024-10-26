# PostLkeClusterRequestNodePoolsInnerAutoscaler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether automatic scaling is enabled for this node pool. Defaults to &#x60;false&#x60;. | [optional] 
**Max** | Pointer to **int32** | The maximum number of nodes to automatically scale to. Defaults to the value provided by the &#x60;count&#x60; field. | [optional] 
**Min** | Pointer to **int32** | The minimum number of nodes to automatically scale to. Defaults to the node pool&#39;s &#x60;count&#x60;. | [optional] 

## Methods

### NewPostLkeClusterRequestNodePoolsInnerAutoscaler

`func NewPostLkeClusterRequestNodePoolsInnerAutoscaler() *PostLkeClusterRequestNodePoolsInnerAutoscaler`

NewPostLkeClusterRequestNodePoolsInnerAutoscaler instantiates a new PostLkeClusterRequestNodePoolsInnerAutoscaler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLkeClusterRequestNodePoolsInnerAutoscalerWithDefaults

`func NewPostLkeClusterRequestNodePoolsInnerAutoscalerWithDefaults() *PostLkeClusterRequestNodePoolsInnerAutoscaler`

NewPostLkeClusterRequestNodePoolsInnerAutoscalerWithDefaults instantiates a new PostLkeClusterRequestNodePoolsInnerAutoscaler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMax

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *PostLkeClusterRequestNodePoolsInnerAutoscaler) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


