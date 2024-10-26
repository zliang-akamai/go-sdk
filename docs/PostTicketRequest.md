# PostTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseId** | Pointer to **int32** | The ID of the Managed Database this ticket is regarding, if relevant. | [optional] 
**Description** | **string** | The full details of the issue or question. | 
**DomainId** | Pointer to **int32** | The ID of the Domain this ticket is regarding, if relevant. | [optional] 
**FirewallId** | Pointer to **int32** | The ID of the Firewall this ticket is regarding, if relevant. | [optional] 
**LinodeId** | Pointer to **int32** | The ID of the Linode this ticket is regarding, if relevant. | [optional] 
**LkeclusterId** | Pointer to **int32** | The ID of the Kubernetes cluster this ticket is regarding, if relevant. | [optional] 
**LongviewclientId** | Pointer to **int32** | The ID of the Longview client this ticket is regarding, if relevant. | [optional] 
**ManagedIssue** | Pointer to **bool** | Designates if this ticket is related to a [Managed service](https://www.linode.com/products/managed/). If &#x60;true&#x60;, the following constraints will apply:  - No ID attributes (i.e. &#x60;linode_id&#x60;, &#x60;domain_id&#x60;, etc.) should be provided with this request. - Your account must have a managed service [enabled](https://techdocs.akamai.com/linode-api/reference/post-enable-managed-service). | [optional] 
**NodebalancerId** | Pointer to **int32** | The ID of the NodeBalancer this ticket is regarding, if relevant. | [optional] 
**Region** | Pointer to **string** | The [Region](https://techdocs.akamai.com/linode-api/reference/get-regions) ID for the associated VLAN this ticket is regarding.  Only allowed when submitting a VLAN ticket. | [optional] 
**Summary** | **string** | The summary or title for this SupportTicket. | 
**Vlan** | Pointer to **string** | The label of the VLAN this ticket is regarding, if relevant. To view your VLANs, run the [List VLANs](https://techdocs.akamai.com/linode-api/reference/get-vlans)) operation.  Requires a specified &#x60;region&#x60; to identify the VLAN. | [optional] 
**VolumeId** | Pointer to **int32** | The ID of the Volume this ticket is regarding, if relevant. | [optional] 
**VpcId** | Pointer to **int32** | The ID of the VPC this ticket is regarding, if relevant. | [optional] 

## Methods

### NewPostTicketRequest

`func NewPostTicketRequest(description string, summary string, ) *PostTicketRequest`

NewPostTicketRequest instantiates a new PostTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTicketRequestWithDefaults

`func NewPostTicketRequestWithDefaults() *PostTicketRequest`

NewPostTicketRequestWithDefaults instantiates a new PostTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseId

`func (o *PostTicketRequest) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *PostTicketRequest) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *PostTicketRequest) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *PostTicketRequest) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetDescription

`func (o *PostTicketRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostTicketRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostTicketRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDomainId

`func (o *PostTicketRequest) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *PostTicketRequest) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *PostTicketRequest) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *PostTicketRequest) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetFirewallId

`func (o *PostTicketRequest) GetFirewallId() int32`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *PostTicketRequest) GetFirewallIdOk() (*int32, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *PostTicketRequest) SetFirewallId(v int32)`

SetFirewallId sets FirewallId field to given value.

### HasFirewallId

`func (o *PostTicketRequest) HasFirewallId() bool`

HasFirewallId returns a boolean if a field has been set.

### GetLinodeId

`func (o *PostTicketRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostTicketRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostTicketRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *PostTicketRequest) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetLkeclusterId

`func (o *PostTicketRequest) GetLkeclusterId() int32`

GetLkeclusterId returns the LkeclusterId field if non-nil, zero value otherwise.

### GetLkeclusterIdOk

`func (o *PostTicketRequest) GetLkeclusterIdOk() (*int32, bool)`

GetLkeclusterIdOk returns a tuple with the LkeclusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLkeclusterId

`func (o *PostTicketRequest) SetLkeclusterId(v int32)`

SetLkeclusterId sets LkeclusterId field to given value.

### HasLkeclusterId

`func (o *PostTicketRequest) HasLkeclusterId() bool`

HasLkeclusterId returns a boolean if a field has been set.

### GetLongviewclientId

`func (o *PostTicketRequest) GetLongviewclientId() int32`

GetLongviewclientId returns the LongviewclientId field if non-nil, zero value otherwise.

### GetLongviewclientIdOk

`func (o *PostTicketRequest) GetLongviewclientIdOk() (*int32, bool)`

GetLongviewclientIdOk returns a tuple with the LongviewclientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongviewclientId

`func (o *PostTicketRequest) SetLongviewclientId(v int32)`

SetLongviewclientId sets LongviewclientId field to given value.

### HasLongviewclientId

`func (o *PostTicketRequest) HasLongviewclientId() bool`

HasLongviewclientId returns a boolean if a field has been set.

### GetManagedIssue

`func (o *PostTicketRequest) GetManagedIssue() bool`

GetManagedIssue returns the ManagedIssue field if non-nil, zero value otherwise.

### GetManagedIssueOk

`func (o *PostTicketRequest) GetManagedIssueOk() (*bool, bool)`

GetManagedIssueOk returns a tuple with the ManagedIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedIssue

`func (o *PostTicketRequest) SetManagedIssue(v bool)`

SetManagedIssue sets ManagedIssue field to given value.

### HasManagedIssue

`func (o *PostTicketRequest) HasManagedIssue() bool`

HasManagedIssue returns a boolean if a field has been set.

### GetNodebalancerId

`func (o *PostTicketRequest) GetNodebalancerId() int32`

GetNodebalancerId returns the NodebalancerId field if non-nil, zero value otherwise.

### GetNodebalancerIdOk

`func (o *PostTicketRequest) GetNodebalancerIdOk() (*int32, bool)`

GetNodebalancerIdOk returns a tuple with the NodebalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancerId

`func (o *PostTicketRequest) SetNodebalancerId(v int32)`

SetNodebalancerId sets NodebalancerId field to given value.

### HasNodebalancerId

`func (o *PostTicketRequest) HasNodebalancerId() bool`

HasNodebalancerId returns a boolean if a field has been set.

### GetRegion

`func (o *PostTicketRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostTicketRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostTicketRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostTicketRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSummary

`func (o *PostTicketRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PostTicketRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PostTicketRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetVlan

`func (o *PostTicketRequest) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PostTicketRequest) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PostTicketRequest) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PostTicketRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVolumeId

`func (o *PostTicketRequest) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *PostTicketRequest) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *PostTicketRequest) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *PostTicketRequest) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetVpcId

`func (o *PostTicketRequest) GetVpcId() int32`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *PostTicketRequest) GetVpcIdOk() (*int32, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *PostTicketRequest) SetVpcId(v int32)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *PostTicketRequest) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


