# GetIps200ResponseAllOfDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IP address. | [optional] [readonly] 
**Gateway** | Pointer to **NullableString** | The default gateway for this address. | [optional] [readonly] 
**LinodeId** | Pointer to **int32** | The ID of the Linode this address currently belongs to. For IPv4 addresses, this defaults to the Linode that this address was assigned to on creation. IPv4 addresses may be moved using the [Assign IPv4s to Linodes](https://techdocs.akamai.com/linode-api/reference/post-assign-ipv4s) operation. For SLAAC and link-local addresses, this value may not be changed. | [optional] [readonly] 
**Prefix** | Pointer to **int32** | The number of bits set in the subnet mask. | [optional] [readonly] 
**Public** | Pointer to **bool** | Whether this is a public or private IP address. | [optional] [readonly] 
**Rdns** | Pointer to **NullableString** | The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set. | [optional] 
**Region** | Pointer to **string** | The Region this IP address resides in. | [optional] [readonly] 
**SubnetMask** | Pointer to **string** | The mask that separates host bits from network bits for this address. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of address this is. | [optional] [readonly] 
**VpcNat11** | Pointer to [**GetIps200ResponseAllOfDataInnerVpcNat11**](GetIps200ResponseAllOfDataInnerVpcNat11.md) |  | [optional] 

## Methods

### NewGetIps200ResponseAllOfDataInner

`func NewGetIps200ResponseAllOfDataInner() *GetIps200ResponseAllOfDataInner`

NewGetIps200ResponseAllOfDataInner instantiates a new GetIps200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIps200ResponseAllOfDataInnerWithDefaults

`func NewGetIps200ResponseAllOfDataInnerWithDefaults() *GetIps200ResponseAllOfDataInner`

NewGetIps200ResponseAllOfDataInnerWithDefaults instantiates a new GetIps200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetIps200ResponseAllOfDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetIps200ResponseAllOfDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetIps200ResponseAllOfDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetIps200ResponseAllOfDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *GetIps200ResponseAllOfDataInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetIps200ResponseAllOfDataInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetIps200ResponseAllOfDataInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetIps200ResponseAllOfDataInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *GetIps200ResponseAllOfDataInner) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *GetIps200ResponseAllOfDataInner) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetLinodeId

`func (o *GetIps200ResponseAllOfDataInner) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *GetIps200ResponseAllOfDataInner) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *GetIps200ResponseAllOfDataInner) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *GetIps200ResponseAllOfDataInner) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetPrefix

`func (o *GetIps200ResponseAllOfDataInner) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetIps200ResponseAllOfDataInner) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetIps200ResponseAllOfDataInner) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetIps200ResponseAllOfDataInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPublic

`func (o *GetIps200ResponseAllOfDataInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GetIps200ResponseAllOfDataInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GetIps200ResponseAllOfDataInner) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GetIps200ResponseAllOfDataInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRdns

`func (o *GetIps200ResponseAllOfDataInner) GetRdns() string`

GetRdns returns the Rdns field if non-nil, zero value otherwise.

### GetRdnsOk

`func (o *GetIps200ResponseAllOfDataInner) GetRdnsOk() (*string, bool)`

GetRdnsOk returns a tuple with the Rdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdns

`func (o *GetIps200ResponseAllOfDataInner) SetRdns(v string)`

SetRdns sets Rdns field to given value.

### HasRdns

`func (o *GetIps200ResponseAllOfDataInner) HasRdns() bool`

HasRdns returns a boolean if a field has been set.

### SetRdnsNil

`func (o *GetIps200ResponseAllOfDataInner) SetRdnsNil(b bool)`

 SetRdnsNil sets the value for Rdns to be an explicit nil

### UnsetRdns
`func (o *GetIps200ResponseAllOfDataInner) UnsetRdns()`

UnsetRdns ensures that no value is present for Rdns, not even an explicit nil
### GetRegion

`func (o *GetIps200ResponseAllOfDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetIps200ResponseAllOfDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetIps200ResponseAllOfDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetIps200ResponseAllOfDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSubnetMask

`func (o *GetIps200ResponseAllOfDataInner) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *GetIps200ResponseAllOfDataInner) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *GetIps200ResponseAllOfDataInner) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *GetIps200ResponseAllOfDataInner) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetType

`func (o *GetIps200ResponseAllOfDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIps200ResponseAllOfDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIps200ResponseAllOfDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIps200ResponseAllOfDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVpcNat11

`func (o *GetIps200ResponseAllOfDataInner) GetVpcNat11() GetIps200ResponseAllOfDataInnerVpcNat11`

GetVpcNat11 returns the VpcNat11 field if non-nil, zero value otherwise.

### GetVpcNat11Ok

`func (o *GetIps200ResponseAllOfDataInner) GetVpcNat11Ok() (*GetIps200ResponseAllOfDataInnerVpcNat11, bool)`

GetVpcNat11Ok returns a tuple with the VpcNat11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNat11

`func (o *GetIps200ResponseAllOfDataInner) SetVpcNat11(v GetIps200ResponseAllOfDataInnerVpcNat11)`

SetVpcNat11 sets VpcNat11 field to given value.

### HasVpcNat11

`func (o *GetIps200ResponseAllOfDataInner) HasVpcNat11() bool`

HasVpcNat11 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


