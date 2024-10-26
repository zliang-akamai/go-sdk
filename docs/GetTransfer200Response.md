# GetTransfer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billable** | Pointer to **int32** | The amount of your transfer pool that is billable this billing cycle. | [optional] [readonly] 
**Quota** | Pointer to **int32** | The amount of network usage allowed this billing cycle. | [optional] [readonly] 
**RegionTransfers** | Pointer to [**[]GetTransfer200ResponseRegionTransfersInner**](GetTransfer200ResponseRegionTransfersInner.md) |  | [optional] 
**Used** | Pointer to **int32** | The amount of network usage you have used this billing cycle. | [optional] [readonly] 

## Methods

### NewGetTransfer200Response

`func NewGetTransfer200Response() *GetTransfer200Response`

NewGetTransfer200Response instantiates a new GetTransfer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransfer200ResponseWithDefaults

`func NewGetTransfer200ResponseWithDefaults() *GetTransfer200Response`

NewGetTransfer200ResponseWithDefaults instantiates a new GetTransfer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillable

`func (o *GetTransfer200Response) GetBillable() int32`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *GetTransfer200Response) GetBillableOk() (*int32, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *GetTransfer200Response) SetBillable(v int32)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *GetTransfer200Response) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### GetQuota

`func (o *GetTransfer200Response) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetTransfer200Response) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetTransfer200Response) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *GetTransfer200Response) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetRegionTransfers

`func (o *GetTransfer200Response) GetRegionTransfers() []GetTransfer200ResponseRegionTransfersInner`

GetRegionTransfers returns the RegionTransfers field if non-nil, zero value otherwise.

### GetRegionTransfersOk

`func (o *GetTransfer200Response) GetRegionTransfersOk() (*[]GetTransfer200ResponseRegionTransfersInner, bool)`

GetRegionTransfersOk returns a tuple with the RegionTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionTransfers

`func (o *GetTransfer200Response) SetRegionTransfers(v []GetTransfer200ResponseRegionTransfersInner)`

SetRegionTransfers sets RegionTransfers field to given value.

### HasRegionTransfers

`func (o *GetTransfer200Response) HasRegionTransfers() bool`

HasRegionTransfers returns a boolean if a field has been set.

### GetUsed

`func (o *GetTransfer200Response) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *GetTransfer200Response) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *GetTransfer200Response) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *GetTransfer200Response) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


