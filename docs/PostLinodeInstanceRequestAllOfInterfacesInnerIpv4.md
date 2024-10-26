# PostLinodeInstanceRequestAllOfInterfacesInnerIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nat11** | Pointer to **NullableString** | The 1:1 NAT IPv4 address, used to associate a public IPv4 address with the VPC Subnet IPv4 address assigned to this Interface.  - Only allowed for &#x60;vpc&#x60; type Interfaces. - Returns &#x60;null&#x60; if no 1:1 NAT is set for a &#x60;vpc&#x60; type Interface. - Returns an empty string (&#x60;\&quot;\&quot;&#x60;) for non-&#x60;vpc&#x60; type Interfaces.  For requests:  - Setting this value to &#x60;any&#x60; enables the Linode&#39;s assigned public IPv4 address on this Interface and establishes a 1:1 NAT between the public IPv4 and VPC Subnet IPv4 addresses. - Setting the value to a specific public IPv4 address that is assigned to the Linode enables a 1:1 NAT between that address and the VPC Subnet IPv4 address. - The public IPv4 address can&#39;t be shared with another Linode. - If omitted, set to &#x60;null&#x60;, or set to an empty string (&#x60;\&quot;\&quot;&#x60;), no 1:1 NAT is established.  &gt; 📘 &gt; &gt; When creating a new compute instance, you can&#39;t set this to a specific IPv4 address. When a new compute instance is created, the network establishes a public IPv4 address for it. Since this address doesn&#39;t exist yet, you can&#39;t include a custom IPv4 address to change it. Once your compute instance is created, you can [update your configuration profile interface](https://www.linode.com/docs/api/linode-instances/#configuration-profile-interface-update) to change the &#x60;nat_1_1&#x60; address. | [optional] 
**Vpc** | Pointer to **NullableString** | The VPC Subnet IPv4 address for this Interface.  - Only allowed for &#x60;vpc&#x60; type Interfaces. - Returns an empty string (&#x60;\&quot;\&quot;&#x60;) for non-&#x60;vpc&#x60; type Interfaces.  For requests:  - Must not already be actively assigned as an address or within a range to any Linodes. - Must not be the first two or last two addresses in the Subnet IPv4 Range. - If omitted, a valid address within the Subnet IPv4 range is automatically assigned. | [optional] 

## Methods

### NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4

`func NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4() *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4`

NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4 instantiates a new PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4WithDefaults

`func NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4WithDefaults() *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4`

NewPostLinodeInstanceRequestAllOfInterfacesInnerIpv4WithDefaults instantiates a new PostLinodeInstanceRequestAllOfInterfacesInnerIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNat11

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) GetNat11() string`

GetNat11 returns the Nat11 field if non-nil, zero value otherwise.

### GetNat11Ok

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) GetNat11Ok() (*string, bool)`

GetNat11Ok returns a tuple with the Nat11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNat11

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) SetNat11(v string)`

SetNat11 sets Nat11 field to given value.

### HasNat11

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) HasNat11() bool`

HasNat11 returns a boolean if a field has been set.

### SetNat11Nil

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) SetNat11Nil(b bool)`

 SetNat11Nil sets the value for Nat11 to be an explicit nil

### UnsetNat11
`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) UnsetNat11()`

UnsetNat11 ensures that no value is present for Nat11, not even an explicit nil
### GetVpc

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### SetVpcNil

`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *PostLinodeInstanceRequestAllOfInterfacesInnerIpv4) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


