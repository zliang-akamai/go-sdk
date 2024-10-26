# GetLinodeIps200ResponseIpv4PublicInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IP address. | [optional] [readonly] 
**Gateway** | Pointer to **NullableString** | The default gateway for this address. | [optional] [readonly] 
**LinodeId** | Pointer to **int32** | The ID of the Linode this address currently belongs to. For IPv4 addresses, this is by default the Linode that this address was assigned to on creation, and these addresses my be moved using the [Assign IPv4s to Linodes](https://techdocs.akamai.com/linode-api/reference/post-assign-ipv4s) operation. For SLAAC and link-local addresses, this value may not be changed. | [optional] [readonly] 
**Prefix** | Pointer to **int32** | The number of bits set in the subnet mask. | [optional] [readonly] 
**Public** | Pointer to **bool** | Whether this is a public or private IP address. | [optional] [readonly] 
**Rdns** | Pointer to **NullableString** | The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set. | [optional] 
**Region** | Pointer to **string** | The Region this IP address resides in. | [optional] [readonly] 
**SubnetMask** | Pointer to **string** | The mask that separates host bits from network bits for this address. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of address this is. | [optional] [readonly] 
**VpcNat11** | Pointer to [**GetLinodeIps200ResponseIpv4PublicInnerVpcNat11**](GetLinodeIps200ResponseIpv4PublicInnerVpcNat11.md) |  | [optional] 

## Methods

### NewGetLinodeIps200ResponseIpv4PublicInner

`func NewGetLinodeIps200ResponseIpv4PublicInner() *GetLinodeIps200ResponseIpv4PublicInner`

NewGetLinodeIps200ResponseIpv4PublicInner instantiates a new GetLinodeIps200ResponseIpv4PublicInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeIps200ResponseIpv4PublicInnerWithDefaults

`func NewGetLinodeIps200ResponseIpv4PublicInnerWithDefaults() *GetLinodeIps200ResponseIpv4PublicInner`

NewGetLinodeIps200ResponseIpv4PublicInnerWithDefaults instantiates a new GetLinodeIps200ResponseIpv4PublicInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *GetLinodeIps200ResponseIpv4PublicInner) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetLinodeId

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetPrefix

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPublic

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRdns

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetRdns() string`

GetRdns returns the Rdns field if non-nil, zero value otherwise.

### GetRdnsOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetRdnsOk() (*string, bool)`

GetRdnsOk returns a tuple with the Rdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdns

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetRdns(v string)`

SetRdns sets Rdns field to given value.

### HasRdns

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasRdns() bool`

HasRdns returns a boolean if a field has been set.

### SetRdnsNil

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetRdnsNil(b bool)`

 SetRdnsNil sets the value for Rdns to be an explicit nil

### UnsetRdns
`func (o *GetLinodeIps200ResponseIpv4PublicInner) UnsetRdns()`

UnsetRdns ensures that no value is present for Rdns, not even an explicit nil
### GetRegion

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSubnetMask

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetType

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVpcNat11

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetVpcNat11() GetLinodeIps200ResponseIpv4PublicInnerVpcNat11`

GetVpcNat11 returns the VpcNat11 field if non-nil, zero value otherwise.

### GetVpcNat11Ok

`func (o *GetLinodeIps200ResponseIpv4PublicInner) GetVpcNat11Ok() (*GetLinodeIps200ResponseIpv4PublicInnerVpcNat11, bool)`

GetVpcNat11Ok returns a tuple with the VpcNat11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNat11

`func (o *GetLinodeIps200ResponseIpv4PublicInner) SetVpcNat11(v GetLinodeIps200ResponseIpv4PublicInnerVpcNat11)`

SetVpcNat11 sets VpcNat11 field to given value.

### HasVpcNat11

`func (o *GetLinodeIps200ResponseIpv4PublicInner) HasVpcNat11() bool`

HasVpcNat11 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


