# GetLinodeIps200ResponseIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Private** | Pointer to [**[]GetLinodeIps200ResponseIpv4PrivateInner**](GetLinodeIps200ResponseIpv4PrivateInner.md) | A list of private IP Address objects belonging to this Linode. | [optional] [readonly] 
**Public** | Pointer to [**[]GetLinodeIps200ResponseIpv4PublicInner**](GetLinodeIps200ResponseIpv4PublicInner.md) | A list of public IP Address objects belonging to this Linode. | [optional] [readonly] 
**Reserved** | Pointer to [**[]GetLinodeIps200ResponseIpv4PublicInner**](GetLinodeIps200ResponseIpv4PublicInner.md) | A list of reserved IP Address objects belonging to this Linode. | [optional] [readonly] 
**Shared** | Pointer to [**[]GetLinodeIps200ResponseIpv4PublicInner**](GetLinodeIps200ResponseIpv4PublicInner.md) | A list of shared IP Address objects assigned to this Linode. | [optional] [readonly] 
**Vpc** | Pointer to [**[]GetLinodeIps200ResponseIpv4VpcInner**](GetLinodeIps200ResponseIpv4VpcInner.md) | A list of Virtual Private Cloud (VPC)-specific addresses or ranges for the Linode. | [optional] [readonly] 

## Methods

### NewGetLinodeIps200ResponseIpv4

`func NewGetLinodeIps200ResponseIpv4() *GetLinodeIps200ResponseIpv4`

NewGetLinodeIps200ResponseIpv4 instantiates a new GetLinodeIps200ResponseIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeIps200ResponseIpv4WithDefaults

`func NewGetLinodeIps200ResponseIpv4WithDefaults() *GetLinodeIps200ResponseIpv4`

NewGetLinodeIps200ResponseIpv4WithDefaults instantiates a new GetLinodeIps200ResponseIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivate

`func (o *GetLinodeIps200ResponseIpv4) GetPrivate() []GetLinodeIps200ResponseIpv4PrivateInner`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GetLinodeIps200ResponseIpv4) GetPrivateOk() (*[]GetLinodeIps200ResponseIpv4PrivateInner, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GetLinodeIps200ResponseIpv4) SetPrivate(v []GetLinodeIps200ResponseIpv4PrivateInner)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GetLinodeIps200ResponseIpv4) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPublic

`func (o *GetLinodeIps200ResponseIpv4) GetPublic() []GetLinodeIps200ResponseIpv4PublicInner`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GetLinodeIps200ResponseIpv4) GetPublicOk() (*[]GetLinodeIps200ResponseIpv4PublicInner, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GetLinodeIps200ResponseIpv4) SetPublic(v []GetLinodeIps200ResponseIpv4PublicInner)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GetLinodeIps200ResponseIpv4) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetReserved

`func (o *GetLinodeIps200ResponseIpv4) GetReserved() []GetLinodeIps200ResponseIpv4PublicInner`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *GetLinodeIps200ResponseIpv4) GetReservedOk() (*[]GetLinodeIps200ResponseIpv4PublicInner, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *GetLinodeIps200ResponseIpv4) SetReserved(v []GetLinodeIps200ResponseIpv4PublicInner)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *GetLinodeIps200ResponseIpv4) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetShared

`func (o *GetLinodeIps200ResponseIpv4) GetShared() []GetLinodeIps200ResponseIpv4PublicInner`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *GetLinodeIps200ResponseIpv4) GetSharedOk() (*[]GetLinodeIps200ResponseIpv4PublicInner, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *GetLinodeIps200ResponseIpv4) SetShared(v []GetLinodeIps200ResponseIpv4PublicInner)`

SetShared sets Shared field to given value.

### HasShared

`func (o *GetLinodeIps200ResponseIpv4) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetVpc

`func (o *GetLinodeIps200ResponseIpv4) GetVpc() []GetLinodeIps200ResponseIpv4VpcInner`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *GetLinodeIps200ResponseIpv4) GetVpcOk() (*[]GetLinodeIps200ResponseIpv4VpcInner, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *GetLinodeIps200ResponseIpv4) SetVpc(v []GetLinodeIps200ResponseIpv4VpcInner)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *GetLinodeIps200ResponseIpv4) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


