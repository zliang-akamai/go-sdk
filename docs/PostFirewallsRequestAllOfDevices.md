# PostFirewallsRequestAllOfDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Linodes** | Pointer to **[]int32** | An array of Linode IDs. A Firewall Device is created for each ID. | [optional] 
**Nodebalancers** | Pointer to **[]int32** | An array containing a NodeBalancer ID. A Firewall Device is created for the ID.  - A NodeBalancer can have only one Firewall assigned to it. - Firewalls only apply to inbound TCP traffic to NodeBalancers. | [optional] 

## Methods

### NewPostFirewallsRequestAllOfDevices

`func NewPostFirewallsRequestAllOfDevices() *PostFirewallsRequestAllOfDevices`

NewPostFirewallsRequestAllOfDevices instantiates a new PostFirewallsRequestAllOfDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFirewallsRequestAllOfDevicesWithDefaults

`func NewPostFirewallsRequestAllOfDevicesWithDefaults() *PostFirewallsRequestAllOfDevices`

NewPostFirewallsRequestAllOfDevicesWithDefaults instantiates a new PostFirewallsRequestAllOfDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinodes

`func (o *PostFirewallsRequestAllOfDevices) GetLinodes() []int32`

GetLinodes returns the Linodes field if non-nil, zero value otherwise.

### GetLinodesOk

`func (o *PostFirewallsRequestAllOfDevices) GetLinodesOk() (*[]int32, bool)`

GetLinodesOk returns a tuple with the Linodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodes

`func (o *PostFirewallsRequestAllOfDevices) SetLinodes(v []int32)`

SetLinodes sets Linodes field to given value.

### HasLinodes

`func (o *PostFirewallsRequestAllOfDevices) HasLinodes() bool`

HasLinodes returns a boolean if a field has been set.

### GetNodebalancers

`func (o *PostFirewallsRequestAllOfDevices) GetNodebalancers() []int32`

GetNodebalancers returns the Nodebalancers field if non-nil, zero value otherwise.

### GetNodebalancersOk

`func (o *PostFirewallsRequestAllOfDevices) GetNodebalancersOk() (*[]int32, bool)`

GetNodebalancersOk returns a tuple with the Nodebalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancers

`func (o *PostFirewallsRequestAllOfDevices) SetNodebalancers(v []int32)`

SetNodebalancers sets Nodebalancers field to given value.

### HasNodebalancers

`func (o *PostFirewallsRequestAllOfDevices) HasNodebalancers() bool`

HasNodebalancers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


