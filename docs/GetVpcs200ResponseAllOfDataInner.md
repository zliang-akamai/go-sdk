# GetVpcs200ResponseAllOfDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date-time of VPC creation. | [optional] [readonly] 
**Description** | Pointer to **string** | A written description to help distinguish the VPC. | [optional] [default to ""]
**Id** | Pointer to **int32** | The unique ID of the VPC. | [optional] [readonly] 
**Label** | Pointer to **string** | The VPC&#39;s label, for display purposes only.  - Needs to be unique among the Account&#39;s VPCs. - Can only contain ASCII letters, numbers, and hyphens (&#x60;-&#x60;). You can&#39;t use two consecutive hyphens (&#x60;--&#x60;). | [optional] 
**Region** | Pointer to **string** | The Region for the VPC. | [optional] 
**Subnets** | Pointer to [**[]GetVpcs200ResponseAllOfDataInnerSubnetsInner**](GetVpcs200ResponseAllOfDataInnerSubnetsInner.md) | A list of subnets associated with the VPC. | [optional] 
**Updated** | Pointer to **NullableTime** | The date-time of the most recent VPC update. | [optional] [readonly] 

## Methods

### NewGetVpcs200ResponseAllOfDataInner

`func NewGetVpcs200ResponseAllOfDataInner() *GetVpcs200ResponseAllOfDataInner`

NewGetVpcs200ResponseAllOfDataInner instantiates a new GetVpcs200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVpcs200ResponseAllOfDataInnerWithDefaults

`func NewGetVpcs200ResponseAllOfDataInnerWithDefaults() *GetVpcs200ResponseAllOfDataInner`

NewGetVpcs200ResponseAllOfDataInnerWithDefaults instantiates a new GetVpcs200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetVpcs200ResponseAllOfDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetVpcs200ResponseAllOfDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetVpcs200ResponseAllOfDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetVpcs200ResponseAllOfDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *GetVpcs200ResponseAllOfDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVpcs200ResponseAllOfDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVpcs200ResponseAllOfDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVpcs200ResponseAllOfDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GetVpcs200ResponseAllOfDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVpcs200ResponseAllOfDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVpcs200ResponseAllOfDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetVpcs200ResponseAllOfDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetVpcs200ResponseAllOfDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetVpcs200ResponseAllOfDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetVpcs200ResponseAllOfDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetVpcs200ResponseAllOfDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegion

`func (o *GetVpcs200ResponseAllOfDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetVpcs200ResponseAllOfDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetVpcs200ResponseAllOfDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetVpcs200ResponseAllOfDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSubnets

`func (o *GetVpcs200ResponseAllOfDataInner) GetSubnets() []GetVpcs200ResponseAllOfDataInnerSubnetsInner`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *GetVpcs200ResponseAllOfDataInner) GetSubnetsOk() (*[]GetVpcs200ResponseAllOfDataInnerSubnetsInner, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *GetVpcs200ResponseAllOfDataInner) SetSubnets(v []GetVpcs200ResponseAllOfDataInnerSubnetsInner)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *GetVpcs200ResponseAllOfDataInner) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetUpdated

`func (o *GetVpcs200ResponseAllOfDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetVpcs200ResponseAllOfDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetVpcs200ResponseAllOfDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetVpcs200ResponseAllOfDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *GetVpcs200ResponseAllOfDataInner) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *GetVpcs200ResponseAllOfDataInner) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


