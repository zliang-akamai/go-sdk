# PutFirewallRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **string** | The fingerprint is a 32-bit hash. It represents the firewall rules as an 8 character hex string. You can use &#x60;fingerprint&#x60; to compare rule versions. | [optional] [readonly] 
**Inbound** | Pointer to **interface{}** |  | [optional] 
**InboundPolicy** | Pointer to **string** | The default behavior for inbound traffic. This setting can be overridden by [updating](https://techdocs.akamai.com/linode-api/reference/put-firewall-rules) the &#x60;inbound.action&#x60; property of the Firewall Rule. | [optional] 
**Outbound** | Pointer to **interface{}** |  | [optional] 
**OutboundPolicy** | Pointer to **string** | The default behavior for outbound traffic. This setting can be overridden by [updating](https://techdocs.akamai.com/linode-api/reference/put-firewall-rules) the &#x60;outbound.action&#x60; property of the Firewall Rule. | [optional] 
**Version** | Pointer to **int32** | The firewall&#39;s rule version. The first version is &#x60;1&#x60;. The version number is incremented when the firewall&#39;s rules change. | [optional] [readonly] 

## Methods

### NewPutFirewallRulesRequest

`func NewPutFirewallRulesRequest() *PutFirewallRulesRequest`

NewPutFirewallRulesRequest instantiates a new PutFirewallRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutFirewallRulesRequestWithDefaults

`func NewPutFirewallRulesRequestWithDefaults() *PutFirewallRulesRequest`

NewPutFirewallRulesRequestWithDefaults instantiates a new PutFirewallRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *PutFirewallRulesRequest) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *PutFirewallRulesRequest) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *PutFirewallRulesRequest) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *PutFirewallRulesRequest) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetInbound

`func (o *PutFirewallRulesRequest) GetInbound() interface{}`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *PutFirewallRulesRequest) GetInboundOk() (*interface{}, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *PutFirewallRulesRequest) SetInbound(v interface{})`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *PutFirewallRulesRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### SetInboundNil

`func (o *PutFirewallRulesRequest) SetInboundNil(b bool)`

 SetInboundNil sets the value for Inbound to be an explicit nil

### UnsetInbound
`func (o *PutFirewallRulesRequest) UnsetInbound()`

UnsetInbound ensures that no value is present for Inbound, not even an explicit nil
### GetInboundPolicy

`func (o *PutFirewallRulesRequest) GetInboundPolicy() string`

GetInboundPolicy returns the InboundPolicy field if non-nil, zero value otherwise.

### GetInboundPolicyOk

`func (o *PutFirewallRulesRequest) GetInboundPolicyOk() (*string, bool)`

GetInboundPolicyOk returns a tuple with the InboundPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundPolicy

`func (o *PutFirewallRulesRequest) SetInboundPolicy(v string)`

SetInboundPolicy sets InboundPolicy field to given value.

### HasInboundPolicy

`func (o *PutFirewallRulesRequest) HasInboundPolicy() bool`

HasInboundPolicy returns a boolean if a field has been set.

### GetOutbound

`func (o *PutFirewallRulesRequest) GetOutbound() interface{}`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *PutFirewallRulesRequest) GetOutboundOk() (*interface{}, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *PutFirewallRulesRequest) SetOutbound(v interface{})`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *PutFirewallRulesRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### SetOutboundNil

`func (o *PutFirewallRulesRequest) SetOutboundNil(b bool)`

 SetOutboundNil sets the value for Outbound to be an explicit nil

### UnsetOutbound
`func (o *PutFirewallRulesRequest) UnsetOutbound()`

UnsetOutbound ensures that no value is present for Outbound, not even an explicit nil
### GetOutboundPolicy

`func (o *PutFirewallRulesRequest) GetOutboundPolicy() string`

GetOutboundPolicy returns the OutboundPolicy field if non-nil, zero value otherwise.

### GetOutboundPolicyOk

`func (o *PutFirewallRulesRequest) GetOutboundPolicyOk() (*string, bool)`

GetOutboundPolicyOk returns a tuple with the OutboundPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPolicy

`func (o *PutFirewallRulesRequest) SetOutboundPolicy(v string)`

SetOutboundPolicy sets OutboundPolicy field to given value.

### HasOutboundPolicy

`func (o *PutFirewallRulesRequest) HasOutboundPolicy() bool`

HasOutboundPolicy returns a boolean if a field has been set.

### GetVersion

`func (o *PutFirewallRulesRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PutFirewallRulesRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PutFirewallRulesRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PutFirewallRulesRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


