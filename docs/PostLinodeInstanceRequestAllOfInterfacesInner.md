# PostLinodeInstanceRequestAllOfInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Returns &#x60;true&#x60; if the Interface is in use, meaning that Compute Instance has been booted using the Configuration Profile to which the Interface belongs. Otherwise returns &#x60;false&#x60;. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID representing this Interface. | [optional] [readonly] 
**IpRanges** | Pointer to **[]string** | An array of IPv4 CIDR VPC Subnet ranges that are routed to this Interface.  - Array items are only allowed for &#x60;vpc&#x60; type Interfaces. - This must be empty for non-&#x60;vpc&#x60; type Interfaces.  For requests:  - Addresses in submitted ranges must not already be actively assigned. - Submitting values replaces any existing values. - Submitting an empty array removes any existing values. - Omitting this property results in no change to existing values. | [optional] 
**IpamAddress** | Pointer to **NullableString** | This Network Interface&#39;s private IP address in Classless Inter-Domain Routing (CIDR) notation.  For &#x60;vlan&#x60; purpose Interfaces:  - Must be unique among the Linode&#39;s Interfaces to avoid conflicting addresses. - Should be unique among devices attached to the VLAN to avoid conflict. - The Linode is configured to use this address for the associated Interface upon reboot if Network Helper is enabled. If Network Helper is disabled, the address can be enabled with [manual static IP configuration](https://www.linode.com/docs/guides/manual-network-configuration/).  For &#x60;public&#x60; purpose Interfaces:  - In requests, must be an empty string (&#x60;\&quot;\&quot;&#x60;) or &#x60;null&#x60; if included. - In responses, always returns &#x60;null&#x60;.  For &#x60;vpc&#x60; purpose Interfaces:  - In requests, must be an empty string (&#x60;\&quot;\&quot;&#x60;) or &#x60;null&#x60; if included. - In responses, always returns &#x60;null&#x60;. | [optional] 
**Ipv4** | Pointer to [**PostLinodeInstanceRequestAllOfInterfacesInnerIpv4**](PostLinodeInstanceRequestAllOfInterfacesInnerIpv4.md) |  | [optional] 
**Label** | Pointer to **NullableString** | The name of this Interface.  For &#x60;vlan&#x60; purpose Interfaces:  - Required. - Must be unique among the Linode&#39;s Interfaces (a Linode cannot be attached to the same VLAN multiple times). - Can only contain ASCII letters, numbers, and hyphens (&#x60;-&#x60;). You can&#39;t use two consecutive hyphens (&#x60;--&#x60;). - If the VLAN label is new, a VLAN is created. Up to 10 VLANs can be created in each data center region. To view your active VLANs, run the [List VLANs](https://techdocs.akamai.com/linode-api/reference/get-vlans) operation.  For &#x60;public&#x60; purpose Interfaces:  - In requests, must be an empty string (&#x60;\&quot;\&quot;&#x60;) or &#x60;null&#x60; if included. - In responses, always returns &#x60;null&#x60;.  For &#x60;vpc&#x60; purpose Interfaces:  - In requests, must be an empty string (&#x60;\&quot;\&quot;&#x60;) or &#x60;null&#x60; if included. - In responses, always returns &#x60;null&#x60;. | [optional] 
**Primary** | Pointer to **bool** | The primary Interface is configured as the default route to the Linode.  Each Configuration Profile can have up to one &#x60;\&quot;primary\&quot;: true&#x60; Interface at a time.  Must be &#x60;false&#x60; for &#x60;vlan&#x60; type Interfaces.  If no Interface is configured as the primary, the first non-&#x60;vlan&#x60; type Interface in the &#x60;interfaces&#x60; array is automatically treated as the primary Interface. | [optional] 
**Purpose** | **string** | The type of Interface.  - &#x60;public&#x60;   - Only one &#x60;public&#x60; Interface per Linode can be defined.   - The Linode&#39;s default public IPv4 address is assigned to the &#x60;public&#x60; Interface.   - A Linode must have a public Interface in the first/eth0 position to be reachable via the public internet upon boot without additional system configuration. If no &#x60;public&#x60; Interface is configured, the Linode is not directly reachable via the public internet. In this case, access can only be established via [LISH](https://www.linode.com/docs/products/compute/compute-instances/guides/lish/) or other Linodes connected to the same VLAN or VPC.  - &#x60;vlan&#x60;   - Configuring a &#x60;vlan&#x60; purpose Interface attaches this Linode to the VLAN with the specified &#x60;label&#x60;.   - The Linode is configured to use the specified &#x60;ipam_address&#x60;, if any.  - &#x60;vpc&#x60;   - Configuring a &#x60;vpc&#x60; purpose Interface attaches this Linode to the existing VPC Subnet with the specified &#x60;subnet_id&#x60;.   - When the Interface is activated, the Linode is configured to use an IP address from the range in the assigned VPC Subnet. See &#x60;ipv4.vpc&#x60; for more information. | 
**SubnetId** | Pointer to **NullableInt32** | The &#x60;id&#x60; of the VPC Subnet for this Interface.  In requests, this value is used to assign a Linode to a VPC Subnet.  - Required for &#x60;vpc&#x60; type Interfaces. - Returns &#x60;null&#x60; for non-&#x60;vpc&#x60; type Interfaces. - Once a VPC Subnet is assigned to an Interface, it cannot be updated. - The Linode must be rebooted with the Interface&#39;s Configuration Profile to complete assignment to a VPC Subnet. | [optional] 
**VpcId** | Pointer to **NullableInt32** | The &#x60;id&#x60; of the VPC configured for this Interface. Returns &#x60;null&#x60; for non-&#x60;vpc&#x60; type Interfaces. | [optional] [readonly] 

## Methods

### NewPostLinodeInstanceRequestAllOfInterfacesInner

`func NewPostLinodeInstanceRequestAllOfInterfacesInner(purpose string, ) *PostLinodeInstanceRequestAllOfInterfacesInner`

NewPostLinodeInstanceRequestAllOfInterfacesInner instantiates a new PostLinodeInstanceRequestAllOfInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLinodeInstanceRequestAllOfInterfacesInnerWithDefaults

`func NewPostLinodeInstanceRequestAllOfInterfacesInnerWithDefaults() *PostLinodeInstanceRequestAllOfInterfacesInner`

NewPostLinodeInstanceRequestAllOfInterfacesInnerWithDefaults instantiates a new PostLinodeInstanceRequestAllOfInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpRanges

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRangesNil

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpRangesNil(b bool)`

 SetIpRangesNil sets the value for IpRanges to be an explicit nil

### UnsetIpRanges
`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetIpRanges()`

UnsetIpRanges ensures that no value is present for IpRanges, not even an explicit nil
### GetIpamAddress

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpamAddress() string`

GetIpamAddress returns the IpamAddress field if non-nil, zero value otherwise.

### GetIpamAddressOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpamAddressOk() (*string, bool)`

GetIpamAddressOk returns a tuple with the IpamAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamAddress

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpamAddress(v string)`

SetIpamAddress sets IpamAddress field to given value.

### HasIpamAddress

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasIpamAddress() bool`

HasIpamAddress returns a boolean if a field has been set.

### SetIpamAddressNil

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpamAddressNil(b bool)`

 SetIpamAddressNil sets the value for IpamAddress to be an explicit nil

### UnsetIpamAddress
`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetIpamAddress()`

UnsetIpamAddress ensures that no value is present for IpamAddress, not even an explicit nil
### GetIpv4

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpv4() PostLinodeInstanceRequestAllOfInterfacesInnerIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetIpv4Ok() (*PostLinodeInstanceRequestAllOfInterfacesInnerIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetIpv4(v PostLinodeInstanceRequestAllOfInterfacesInnerIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetLabel

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetPrimary

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPurpose

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetSubnetId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetVpcId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetVpcId() int32`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) GetVpcIdOk() (*int32, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetVpcId(v int32)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *PostLinodeInstanceRequestAllOfInterfacesInner) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


