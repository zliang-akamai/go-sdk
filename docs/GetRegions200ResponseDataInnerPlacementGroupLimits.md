# GetRegions200ResponseDataInnerPlacementGroupLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumLinodesPerPg** | Pointer to **int32** | The maximum number of compute instances you can include in a single placement group in this region. | [optional] [readonly] 
**MaximumPgsPerCustomer** | Pointer to **NullableInt32** | The maximum number of placement groups you can have in this region. Displayed as &#x60;null&#x60; if you don&#39;t have a limit. | [optional] [readonly] 

## Methods

### NewGetRegions200ResponseDataInnerPlacementGroupLimits

`func NewGetRegions200ResponseDataInnerPlacementGroupLimits() *GetRegions200ResponseDataInnerPlacementGroupLimits`

NewGetRegions200ResponseDataInnerPlacementGroupLimits instantiates a new GetRegions200ResponseDataInnerPlacementGroupLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegions200ResponseDataInnerPlacementGroupLimitsWithDefaults

`func NewGetRegions200ResponseDataInnerPlacementGroupLimitsWithDefaults() *GetRegions200ResponseDataInnerPlacementGroupLimits`

NewGetRegions200ResponseDataInnerPlacementGroupLimitsWithDefaults instantiates a new GetRegions200ResponseDataInnerPlacementGroupLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumLinodesPerPg

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) GetMaximumLinodesPerPg() int32`

GetMaximumLinodesPerPg returns the MaximumLinodesPerPg field if non-nil, zero value otherwise.

### GetMaximumLinodesPerPgOk

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) GetMaximumLinodesPerPgOk() (*int32, bool)`

GetMaximumLinodesPerPgOk returns a tuple with the MaximumLinodesPerPg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLinodesPerPg

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) SetMaximumLinodesPerPg(v int32)`

SetMaximumLinodesPerPg sets MaximumLinodesPerPg field to given value.

### HasMaximumLinodesPerPg

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) HasMaximumLinodesPerPg() bool`

HasMaximumLinodesPerPg returns a boolean if a field has been set.

### GetMaximumPgsPerCustomer

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) GetMaximumPgsPerCustomer() int32`

GetMaximumPgsPerCustomer returns the MaximumPgsPerCustomer field if non-nil, zero value otherwise.

### GetMaximumPgsPerCustomerOk

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) GetMaximumPgsPerCustomerOk() (*int32, bool)`

GetMaximumPgsPerCustomerOk returns a tuple with the MaximumPgsPerCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPgsPerCustomer

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) SetMaximumPgsPerCustomer(v int32)`

SetMaximumPgsPerCustomer sets MaximumPgsPerCustomer field to given value.

### HasMaximumPgsPerCustomer

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) HasMaximumPgsPerCustomer() bool`

HasMaximumPgsPerCustomer returns a boolean if a field has been set.

### SetMaximumPgsPerCustomerNil

`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) SetMaximumPgsPerCustomerNil(b bool)`

 SetMaximumPgsPerCustomerNil sets the value for MaximumPgsPerCustomer to be an explicit nil

### UnsetMaximumPgsPerCustomer
`func (o *GetRegions200ResponseDataInnerPlacementGroupLimits) UnsetMaximumPgsPerCustomer()`

UnsetMaximumPgsPerCustomer ensures that no value is present for MaximumPgsPerCustomer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


