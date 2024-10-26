# GetLinodeFirewalls200ResponseDataInnerRulesInboundInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall&#39;s &#x60;inbound_policy&#x60; if this is an inbound rule, or the &#x60;outbound_policy&#x60; if this is an outbound rule. | [optional] 
**Addresses** | Pointer to [**GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses**](GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses.md) |  | [optional] 
**Description** | Pointer to **string** | Used to describe this rule. For display purposes only. | [optional] 
**Label** | Pointer to **string** | Used to identify this rule. For display purposes only. | [optional] 
**Ports** | Pointer to **NullableString** | A string representing the port or ports affected by this rule:  - The string may be a single port, a range of ports, or a comma-separated list of single ports and port ranges. A space is permitted following each comma. - A range of ports is inclusive of the start and end values for the range. The end value of the range must be greater than the start value. - Ports must be within 1 and 65535, and may not contain any leading zeroes. For example, port &#x60;080&#x60; is not allowed. - The ports string can have up to 15 _pieces_, where a single port is treated as one piece, and a port range is treated as two pieces. For example, the string \&quot;22-24, 80, 443\&quot; has four pieces. - If no ports are configured, all ports are affected. - Only allowed for the TCP and UDP protocols. Ports are not allowed for the ICMP and IPENCAP protocols. | [optional] 
**Protocol** | Pointer to **string** | The type of network traffic affected by this rule. | [optional] 

## Methods

### NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInner

`func NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInner() *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner`

NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInner instantiates a new GetLinodeFirewalls200ResponseDataInnerRulesInboundInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerWithDefaults

`func NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerWithDefaults() *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner`

NewGetLinodeFirewalls200ResponseDataInnerRulesInboundInnerWithDefaults instantiates a new GetLinodeFirewalls200ResponseDataInnerRulesInboundInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAddresses

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetAddresses() GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetAddressesOk() (*GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) SetAddresses(v GetLinodeFirewalls200ResponseDataInnerRulesInboundInnerAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetDescription

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPorts

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetProtocol

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetLinodeFirewalls200ResponseDataInnerRulesInboundInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


