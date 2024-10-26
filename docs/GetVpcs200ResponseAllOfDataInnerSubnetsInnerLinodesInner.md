# GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of a Linode assigned to the VPC Subnet. | [optional] 
**Interfaces** | Pointer to [**[]GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner**](GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner.md) | VPC purpose interfaces with the subnet&#39;s &#x60;subnet_id&#x60; assigned to the Linode. | [optional] 

## Methods

### NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner

`func NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner() *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner`

NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner instantiates a new GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerWithDefaults

`func NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerWithDefaults() *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner`

NewGetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerWithDefaults instantiates a new GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) GetInterfaces() []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) GetInterfacesOk() (*[]GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) SetInterfaces(v []GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetVpcs200ResponseAllOfDataInnerSubnetsInnerLinodesInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


