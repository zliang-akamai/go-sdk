# GetLinodeTransfer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billable** | Pointer to **int32** | The amount of network transfer this Linode has used, in GB, past your monthly quota. | [optional] [readonly] 
**Quota** | Pointer to **int32** | The amount of network transfer this Linode adds to your transfer pool, in GB, for the current month&#39;s billing cycle. | [optional] [readonly] 
**Used** | Pointer to **int32** | The amount of network transfer used by this Linode, in bytes, for the current month&#39;s billing cycle. | [optional] [readonly] 

## Methods

### NewGetLinodeTransfer200Response

`func NewGetLinodeTransfer200Response() *GetLinodeTransfer200Response`

NewGetLinodeTransfer200Response instantiates a new GetLinodeTransfer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeTransfer200ResponseWithDefaults

`func NewGetLinodeTransfer200ResponseWithDefaults() *GetLinodeTransfer200Response`

NewGetLinodeTransfer200ResponseWithDefaults instantiates a new GetLinodeTransfer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillable

`func (o *GetLinodeTransfer200Response) GetBillable() int32`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *GetLinodeTransfer200Response) GetBillableOk() (*int32, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *GetLinodeTransfer200Response) SetBillable(v int32)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *GetLinodeTransfer200Response) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### GetQuota

`func (o *GetLinodeTransfer200Response) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetLinodeTransfer200Response) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetLinodeTransfer200Response) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *GetLinodeTransfer200Response) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetUsed

`func (o *GetLinodeTransfer200Response) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *GetLinodeTransfer200Response) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *GetLinodeTransfer200Response) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *GetLinodeTransfer200Response) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


