# GetVpcs200ResponseAllOfDataInnerSubnetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date-time of VPC Subnet creation. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID of the VPC Subnet. | [optional] [readonly] 
**Ipv4** | Pointer to **string** | IPv4 range in CIDR canonical form.  - The range must belong to a private address space as defined in [RFC1918](https://datatracker.ietf.org/doc/html/rfc1918). - Allowed prefix lengths: 1-29. - The range must not overlap with 192.168.128.0/17. - The range must not overlap with other Subnets on the same VPC. | [optional] 
**Label** | Pointer to **string** | The VPC Subnet&#39;s label, for display purposes only.  - Must be unique among the VPC&#39;s Subnets. - Can only contain ASCII letters, numbers, and hyphens (&#x60;-&#x60;). You can&#39;t use two consecutive hyphens (&#x60;--&#x60;). | [optional] 
**Linodes** | Pointer to [**[]GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner**](GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner.md) | An array of Linode IDs assigned to the VPC Subnet.  A Linode is assigned to a VPC Subnet if it has a Configuration Profile with a &#x60;vpc&#x60; purpose interface with the subnet&#39;s &#x60;subnet_id&#x60;. Even if the Configuration Profile is not active, meaning the Linode does not have access to the Subnet, the Linode still appears in this array. | [optional] [readonly] 
**Updated** | Pointer to **NullableTime** | The date-time of the most recent VPC Subnet update. | [optional] [readonly] 

## Methods

### NewGetVpcs200ResponseAllOfDataInnerSubnetsInner

`func NewGetVpcs200ResponseAllOfDataInnerSubnetsInner() *GetVpcs200ResponseAllOfDataInnerSubnetsInner`

NewGetVpcs200ResponseAllOfDataInnerSubnetsInner instantiates a new GetVpcs200ResponseAllOfDataInnerSubnetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerWithDefaults

`func NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerWithDefaults() *GetVpcs200ResponseAllOfDataInnerSubnetsInner`

NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerWithDefaults instantiates a new GetVpcs200ResponseAllOfDataInnerSubnetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpv4

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetLabel

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLinodes

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetLinodes() []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner`

GetLinodes returns the Linodes field if non-nil, zero value otherwise.

### GetLinodesOk

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetLinodesOk() (*[]GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner, bool)`

GetLinodesOk returns a tuple with the Linodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodes

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetLinodes(v []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner)`

SetLinodes sets Linodes field to given value.

### HasLinodes

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasLinodes() bool`

HasLinodes returns a boolean if a field has been set.

### GetUpdated

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInner) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


