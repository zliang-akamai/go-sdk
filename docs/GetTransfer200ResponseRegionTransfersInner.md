# GetTransfer200ResponseRegionTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billable** | Pointer to **int32** | The amount of your transfer pool that is billable this billing cycle for this Region. | [optional] [readonly] 
**Id** | Pointer to **string** | The Region ID for this network utilization data. | [optional] 
**Quota** | Pointer to **int32** | The amount of network usage allowed this billing cycle for this Region. | [optional] [readonly] 
**Used** | Pointer to **int32** | The amount of network usage you have used this billing cycle for this Region. | [optional] [readonly] 

## Methods

### NewGetTransfer200ResponseRegionTransfersInner

`func NewGetTransfer200ResponseRegionTransfersInner() *GetTransfer200ResponseRegionTransfersInner`

NewGetTransfer200ResponseRegionTransfersInner instantiates a new GetTransfer200ResponseRegionTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransfer200ResponseRegionTransfersInnerWithDefaults

`func NewGetTransfer200ResponseRegionTransfersInnerWithDefaults() *GetTransfer200ResponseRegionTransfersInner`

NewGetTransfer200ResponseRegionTransfersInnerWithDefaults instantiates a new GetTransfer200ResponseRegionTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillable

`func (o *GetTransfer200ResponseRegionTransfersInner) GetBillable() int32`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *GetTransfer200ResponseRegionTransfersInner) GetBillableOk() (*int32, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *GetTransfer200ResponseRegionTransfersInner) SetBillable(v int32)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *GetTransfer200ResponseRegionTransfersInner) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### GetId

`func (o *GetTransfer200ResponseRegionTransfersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTransfer200ResponseRegionTransfersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTransfer200ResponseRegionTransfersInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetTransfer200ResponseRegionTransfersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuota

`func (o *GetTransfer200ResponseRegionTransfersInner) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetTransfer200ResponseRegionTransfersInner) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetTransfer200ResponseRegionTransfersInner) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *GetTransfer200ResponseRegionTransfersInner) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetUsed

`func (o *GetTransfer200ResponseRegionTransfersInner) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *GetTransfer200ResponseRegionTransfersInner) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *GetTransfer200ResponseRegionTransfersInner) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *GetTransfer200ResponseRegionTransfersInner) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


