# GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | Pointer to **[]string** | A list of IPv4 addresses or networks. Addresses must be in IP/mask format. Must not be an empty list.  If &#x60;0.0.0.0/0&#x60; is included in this list, all IPv4 addresses are affected by this rule. | [optional] 
**Ipv6** | Pointer to **[]string** | A list of IPv6 addresses or networks. Addresses must be in IP/mask format and must not include zone_id notation as described in [RFC 4007](https://www.rfc-editor.org/rfc/rfc4007). Must not be an empty list.  If &#x60;::/0&#x60; is included in this list, all IPv6 addresses are affected by this rule. | [optional] 

## Methods

### NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses

`func NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses() *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses`

NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses instantiates a new GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddressesWithDefaults

`func NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddressesWithDefaults() *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses`

NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddressesWithDefaults instantiates a new GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses) GetIpv4() []string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses) GetIpv4Ok() (*[]string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses) SetIpv4(v []string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses) GetIpv6() []string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses) GetIpv6Ok() (*[]string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses) SetIpv6(v []string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


