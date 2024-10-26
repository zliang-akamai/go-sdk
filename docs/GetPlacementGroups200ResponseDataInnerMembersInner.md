# GetPlacementGroups200ResponseDataInnerMembersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCompliant** | Pointer to **bool** | The compliance status of each individual compute instance in the placement group. | [optional] 
**LinodeId** | Pointer to **int32** | The unique identifier for a compute instance included in the placement group. | [optional] [readonly] 

## Methods

### NewGetPlacementGroups200ResponseDataInnerMembersInner

`func NewGetPlacementGroups200ResponseDataInnerMembersInner() *GetPlacementGroups200ResponseDataInnerMembersInner`

NewGetPlacementGroups200ResponseDataInnerMembersInner instantiates a new GetPlacementGroups200ResponseDataInnerMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlacementGroups200ResponseDataInnerMembersInnerWithDefaults

`func NewGetPlacementGroups200ResponseDataInnerMembersInnerWithDefaults() *GetPlacementGroups200ResponseDataInnerMembersInner`

NewGetPlacementGroups200ResponseDataInnerMembersInnerWithDefaults instantiates a new GetPlacementGroups200ResponseDataInnerMembersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCompliant

`func (o *GetPlacementGroups200ResponseDataInnerMembersInner) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *GetPlacementGroups200ResponseDataInnerMembersInner) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *GetPlacementGroups200ResponseDataInnerMembersInner) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *GetPlacementGroups200ResponseDataInnerMembersInner) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### GetLinodeId

`func (o *GetPlacementGroups200ResponseDataInnerMembersInner) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *GetPlacementGroups200ResponseDataInnerMembersInner) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *GetPlacementGroups200ResponseDataInnerMembersInner) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *GetPlacementGroups200ResponseDataInnerMembersInner) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


