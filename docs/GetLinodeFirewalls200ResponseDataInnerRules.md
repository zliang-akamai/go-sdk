# GetLinodeFirewalls200ResponseDataInnerRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **string** | The fingerprint is a 32-bit hash. It represents the firewall rules as an 8 character hex string. You can use &#x60;fingerprint&#x60; to compare rule versions. | [optional] [readonly] 
**Inbound** | Pointer to [**[]GetLinodeFirewalls200ResponseDataInnerRulesInboundInner**](GetLinodeFirewalls200ResponseDataInnerRulesInboundInner.md) | The inbound rules for the firewall, as a JSON array. | [optional] 
**InboundPolicy** | Pointer to **string** | The default behavior for inbound traffic. This setting can be overridden by [updating](https://techdocs.akamai.com/linode-api/reference/put-firewall-rules) the &#x60;inbound.action&#x60; property of the Firewall Rule. | [optional] 
**Outbound** | Pointer to [**[]GetLinodeFirewalls200ResponseDataInnerRulesInboundInner**](GetLinodeFirewalls200ResponseDataInnerRulesInboundInner.md) | The outbound rules for the firewall, as a JSON array. | [optional] 
**OutboundPolicy** | Pointer to **string** | The default behavior for outbound traffic. This setting can be overridden by [updating](https://techdocs.akamai.com/linode-api/reference/put-firewall-rules) the &#x60;outbound.action&#x60; property of the Firewall Rule. | [optional] 
**Version** | Pointer to **int32** | The firewall&#39;s rule version. The first version is &#x60;1&#x60;. The version number is incremented when the firewall&#39;s rules change. | [optional] [readonly] 

## Methods

### NewGetLinodeFirewalls200ResponseDataInnerRules

`func NewGetLinodeFirewalls200ResponseDataInnerRules() *GetLinodeFirewalls200ResponseDataInnerRules`

NewGetLinodeFirewalls200ResponseDataInnerRules instantiates a new GetLinodeFirewalls200ResponseDataInnerRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeFirewalls200ResponseDataInnerRulesWithDefaults

`func NewGetLinodeFirewalls200ResponseDataInnerRulesWithDefaults() *GetLinodeFirewalls200ResponseDataInnerRules`

NewGetLinodeFirewalls200ResponseDataInnerRulesWithDefaults instantiates a new GetLinodeFirewalls200ResponseDataInnerRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetInbound

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetInbound() []GetLinodeFirewalls200ResponseDataInnerRulesInboundInner`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetInboundOk() (*[]GetLinodeFirewalls200ResponseDataInnerRulesInboundInner, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) SetInbound(v []GetLinodeFirewalls200ResponseDataInnerRulesInboundInner)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetInboundPolicy

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetInboundPolicy() string`

GetInboundPolicy returns the InboundPolicy field if non-nil, zero value otherwise.

### GetInboundPolicyOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetInboundPolicyOk() (*string, bool)`

GetInboundPolicyOk returns a tuple with the InboundPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundPolicy

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) SetInboundPolicy(v string)`

SetInboundPolicy sets InboundPolicy field to given value.

### HasInboundPolicy

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) HasInboundPolicy() bool`

HasInboundPolicy returns a boolean if a field has been set.

### GetOutbound

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetOutbound() []GetLinodeFirewalls200ResponseDataInnerRulesInboundInner`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetOutboundOk() (*[]GetLinodeFirewalls200ResponseDataInnerRulesInboundInner, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) SetOutbound(v []GetLinodeFirewalls200ResponseDataInnerRulesInboundInner)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### GetOutboundPolicy

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetOutboundPolicy() string`

GetOutboundPolicy returns the OutboundPolicy field if non-nil, zero value otherwise.

### GetOutboundPolicyOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetOutboundPolicyOk() (*string, bool)`

GetOutboundPolicyOk returns a tuple with the OutboundPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPolicy

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) SetOutboundPolicy(v string)`

SetOutboundPolicy sets OutboundPolicy field to given value.

### HasOutboundPolicy

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) HasOutboundPolicy() bool`

HasOutboundPolicy returns a boolean if a field has been set.

### GetVersion

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetLinodeFirewalls200ResponseDataInnerRules) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


