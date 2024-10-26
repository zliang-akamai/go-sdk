# GetLkeClusterPools200ResponseDataInnerAutoscaler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether autoscaling is enabled for this Node Pool. Defaults to &#x60;false&#x60;. | [optional] 
**Max** | Pointer to **int32** | The maximum number of nodes to autoscale to. Defaults to the Node Pool&#39;s &#x60;count&#x60;. | [optional] 
**Min** | Pointer to **int32** | The minimum number of nodes to autoscale to. Defaults to the Node Pool&#39;s &#x60;count&#x60;. | [optional] 

## Methods

### NewGetLkeClusterPools200ResponseDataInnerAutoscaler

`func NewGetLkeClusterPools200ResponseDataInnerAutoscaler() *GetLkeClusterPools200ResponseDataInnerAutoscaler`

NewGetLkeClusterPools200ResponseDataInnerAutoscaler instantiates a new GetLkeClusterPools200ResponseDataInnerAutoscaler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLkeClusterPools200ResponseDataInnerAutoscalerWithDefaults

`func NewGetLkeClusterPools200ResponseDataInnerAutoscalerWithDefaults() *GetLkeClusterPools200ResponseDataInnerAutoscaler`

NewGetLkeClusterPools200ResponseDataInnerAutoscalerWithDefaults instantiates a new GetLkeClusterPools200ResponseDataInnerAutoscaler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMax

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *GetLkeClusterPools200ResponseDataInnerAutoscaler) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


