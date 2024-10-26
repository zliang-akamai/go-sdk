# GetLinodeIps200ResponseIpv4VpcInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Returns &#x60;true&#x60; if the VPC interface is in use, meaning that the Linode was powered on using the &#x60;config_id&#x60; to which the interface belongs. Otherwise returns &#x60;false&#x60;. | [optional] [readonly] 
**Address** | Pointer to **NullableString** | An IPv4 address configured for this VPC interface. These follow the [RFC 1918](https://datatracker.ietf.org/doc/html/rfc1918) private address format. Displayed as &#x60;null&#x60; if an &#x60;address_range&#x60;. | [optional] [readonly] 
**AddressRange** | Pointer to **NullableString** | A range of IPv4 addresses configured for this VPC interface. Displayed as &#x60;null&#x60; if a single &#x60;address&#x60;. | [optional] [readonly] 
**ConfigId** | Pointer to **int32** | The globally general entity identifier for the Linode configuration profile where the VPC is included. | [optional] [readonly] 
**Gateway** | Pointer to **NullableString** | The default gateway for the VPC subnet that the IP or IP range belongs to. | [optional] [readonly] 
**InterfaceId** | Pointer to **int32** | The globally general API entity identifier for the Linode interface. | [optional] [readonly] 
**LinodeId** | Pointer to **int32** | The identifier for the Linode the VPC interface currently belongs to. | [optional] [readonly] 
**Nat11** | Pointer to **string** | The public IP address used for NAT 1:1 with the VPC. This is empty if NAT 1:1 isn&#39;t used. | [optional] [readonly] 
**Prefix** | Pointer to **NullableInt32** | The number of bits set in the &#x60;subnet_mask&#x60;. | [optional] [readonly] 
**Region** | Pointer to **string** | The region of the VPC. | [optional] [readonly] 
**SubnetId** | Pointer to **int32** | The &#x60;id&#x60; of the VPC Subnet for this interface. | [optional] 
**SubnetMask** | Pointer to **string** | The mask that separates host bits from network bits for the &#x60;address&#x60; or &#x60;address_range&#x60;. | [optional] [readonly] 
**VpcId** | Pointer to **int32** | The unique globally general API entity identifier for the VPC. | [optional] [readonly] 

## Methods

### NewGetLinodeIps200ResponseIpv4VpcInner

`func NewGetLinodeIps200ResponseIpv4VpcInner() *GetLinodeIps200ResponseIpv4VpcInner`

NewGetLinodeIps200ResponseIpv4VpcInner instantiates a new GetLinodeIps200ResponseIpv4VpcInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeIps200ResponseIpv4VpcInnerWithDefaults

`func NewGetLinodeIps200ResponseIpv4VpcInnerWithDefaults() *GetLinodeIps200ResponseIpv4VpcInner`

NewGetLinodeIps200ResponseIpv4VpcInnerWithDefaults instantiates a new GetLinodeIps200ResponseIpv4VpcInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *GetLinodeIps200ResponseIpv4VpcInner) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAddressRange

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetAddressRange() string`

GetAddressRange returns the AddressRange field if non-nil, zero value otherwise.

### GetAddressRangeOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetAddressRangeOk() (*string, bool)`

GetAddressRangeOk returns a tuple with the AddressRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRange

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetAddressRange(v string)`

SetAddressRange sets AddressRange field to given value.

### HasAddressRange

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasAddressRange() bool`

HasAddressRange returns a boolean if a field has been set.

### SetAddressRangeNil

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetAddressRangeNil(b bool)`

 SetAddressRangeNil sets the value for AddressRange to be an explicit nil

### UnsetAddressRange
`func (o *GetLinodeIps200ResponseIpv4VpcInner) UnsetAddressRange()`

UnsetAddressRange ensures that no value is present for AddressRange, not even an explicit nil
### GetConfigId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetGateway

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *GetLinodeIps200ResponseIpv4VpcInner) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetInterfaceId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetLinodeId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetNat11

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetNat11() string`

GetNat11 returns the Nat11 field if non-nil, zero value otherwise.

### GetNat11Ok

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetNat11Ok() (*string, bool)`

GetNat11Ok returns a tuple with the Nat11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNat11

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetNat11(v string)`

SetNat11 sets Nat11 field to given value.

### HasNat11

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasNat11() bool`

HasNat11 returns a boolean if a field has been set.

### GetPrefix

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *GetLinodeIps200ResponseIpv4VpcInner) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetRegion

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSubnetId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetSubnetMask

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetVpcId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetVpcId() int32`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *GetLinodeIps200ResponseIpv4VpcInner) GetVpcIdOk() (*int32, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) SetVpcId(v int32)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *GetLinodeIps200ResponseIpv4VpcInner) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


