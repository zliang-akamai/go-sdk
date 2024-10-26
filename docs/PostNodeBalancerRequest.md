# PostNodeBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientConnThrottle** | Pointer to **int32** | Throttle connections per second.  Set to 0 (zero) to disable throttling. | [optional] 
**Configs** | Pointer to [**[]PostNodeBalancerRequestConfigsInner**](PostNodeBalancerRequestConfigsInner.md) | The port Configs to create for this NodeBalancer.  Each Config must have a unique port and at least one Node. | [optional] 
**FirewallId** | Pointer to **int32** | The ID of the Firewall to assign to the NodeBalancer.  - A NodeBalancer can have only one Firewall assigned to it. - Firewalls only apply to inbound TCP traffic to NodeBalancers. | [optional] 
**Label** | Pointer to **string** | This NodeBalancer&#39;s label. These must be unique on your Account. | [optional] 
**Region** | **string** | The ID of the Region to create this NodeBalancer in. | 
**Tags** | Pointer to **[]string** | An array of Tags applied to this object. Tags are for organizational purposes only. | [optional] 

## Methods

### NewPostNodeBalancerRequest

`func NewPostNodeBalancerRequest(region string, ) *PostNodeBalancerRequest`

NewPostNodeBalancerRequest instantiates a new PostNodeBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostNodeBalancerRequestWithDefaults

`func NewPostNodeBalancerRequestWithDefaults() *PostNodeBalancerRequest`

NewPostNodeBalancerRequestWithDefaults instantiates a new PostNodeBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientConnThrottle

`func (o *PostNodeBalancerRequest) GetClientConnThrottle() int32`

GetClientConnThrottle returns the ClientConnThrottle field if non-nil, zero value otherwise.

### GetClientConnThrottleOk

`func (o *PostNodeBalancerRequest) GetClientConnThrottleOk() (*int32, bool)`

GetClientConnThrottleOk returns a tuple with the ClientConnThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnThrottle

`func (o *PostNodeBalancerRequest) SetClientConnThrottle(v int32)`

SetClientConnThrottle sets ClientConnThrottle field to given value.

### HasClientConnThrottle

`func (o *PostNodeBalancerRequest) HasClientConnThrottle() bool`

HasClientConnThrottle returns a boolean if a field has been set.

### GetConfigs

`func (o *PostNodeBalancerRequest) GetConfigs() []PostNodeBalancerRequestConfigsInner`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *PostNodeBalancerRequest) GetConfigsOk() (*[]PostNodeBalancerRequestConfigsInner, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *PostNodeBalancerRequest) SetConfigs(v []PostNodeBalancerRequestConfigsInner)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *PostNodeBalancerRequest) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetFirewallId

`func (o *PostNodeBalancerRequest) GetFirewallId() int32`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *PostNodeBalancerRequest) GetFirewallIdOk() (*int32, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *PostNodeBalancerRequest) SetFirewallId(v int32)`

SetFirewallId sets FirewallId field to given value.

### HasFirewallId

`func (o *PostNodeBalancerRequest) HasFirewallId() bool`

HasFirewallId returns a boolean if a field has been set.

### GetLabel

`func (o *PostNodeBalancerRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostNodeBalancerRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostNodeBalancerRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostNodeBalancerRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegion

`func (o *PostNodeBalancerRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostNodeBalancerRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostNodeBalancerRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetTags

`func (o *PostNodeBalancerRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostNodeBalancerRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostNodeBalancerRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostNodeBalancerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


