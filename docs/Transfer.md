# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billable** | Pointer to **int32** | The amount of your transfer pool that is billable this billing cycle. | [optional] [readonly] 
**Quota** | Pointer to **int32** | The amount of network usage allowed this billing cycle. | [optional] [readonly] 
**RegionTransfers** | Pointer to [**[]GetTransfer200ResponseRegionTransfersInner**](GetTransfer200ResponseRegionTransfersInner.md) |  | [optional] 
**Used** | Pointer to **int32** | The amount of network usage you have used this billing cycle. | [optional] [readonly] 

## Methods

### NewTransfer

`func NewTransfer() *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillable

`func (o *Transfer) GetBillable() int32`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *Transfer) GetBillableOk() (*int32, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *Transfer) SetBillable(v int32)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *Transfer) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### GetQuota

`func (o *Transfer) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Transfer) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Transfer) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Transfer) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetRegionTransfers

`func (o *Transfer) GetRegionTransfers() []GetTransfer200ResponseRegionTransfersInner`

GetRegionTransfers returns the RegionTransfers field if non-nil, zero value otherwise.

### GetRegionTransfersOk

`func (o *Transfer) GetRegionTransfersOk() (*[]GetTransfer200ResponseRegionTransfersInner, bool)`

GetRegionTransfersOk returns a tuple with the RegionTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionTransfers

`func (o *Transfer) SetRegionTransfers(v []GetTransfer200ResponseRegionTransfersInner)`

SetRegionTransfers sets RegionTransfers field to given value.

### HasRegionTransfers

`func (o *Transfer) HasRegionTransfers() bool`

HasRegionTransfers returns a boolean if a field has been set.

### GetUsed

`func (o *Transfer) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Transfer) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Transfer) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Transfer) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


