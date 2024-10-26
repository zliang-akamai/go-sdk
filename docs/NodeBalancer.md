# NodeBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientConnThrottle** | Pointer to **int32** | Throttle connections per second.  Set to 0 (zero) to disable throttling. | [optional] 
**Created** | Pointer to **time.Time** | When this NodeBalancer was created. | [optional] [readonly] 
**Hostname** | Pointer to **string** | This NodeBalancer&#39;s hostname, beginning with its IP address and ending with _.ip.linodeusercontent.com_. | [optional] [readonly] 
**Id** | Pointer to **int32** | This NodeBalancer&#39;s unique ID. | [optional] [readonly] 
**Ipv4** | Pointer to **string** | This NodeBalancer&#39;s public IPv4 address. | [optional] [readonly] 
**Ipv6** | Pointer to **NullableString** | This NodeBalancer&#39;s public IPv6 address. | [optional] [readonly] 
**Label** | Pointer to **string** | This NodeBalancer&#39;s label. These must be unique on your Account. | [optional] 
**Region** | Pointer to **string** | The Region where this NodeBalancer is located. NodeBalancers only support backends in the same Region. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | An array of Tags applied to this object.  Tags are for organizational purposes only. | [optional] 
**Transfer** | Pointer to [**NodeBalancerTransfer**](NodeBalancerTransfer.md) |  | [optional] 
**Updated** | Pointer to **time.Time** | When this NodeBalancer was last updated. | [optional] [readonly] 

## Methods

### NewNodeBalancer

`func NewNodeBalancer() *NodeBalancer`

NewNodeBalancer instantiates a new NodeBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancerWithDefaults

`func NewNodeBalancerWithDefaults() *NodeBalancer`

NewNodeBalancerWithDefaults instantiates a new NodeBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientConnThrottle

`func (o *NodeBalancer) GetClientConnThrottle() int32`

GetClientConnThrottle returns the ClientConnThrottle field if non-nil, zero value otherwise.

### GetClientConnThrottleOk

`func (o *NodeBalancer) GetClientConnThrottleOk() (*int32, bool)`

GetClientConnThrottleOk returns a tuple with the ClientConnThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnThrottle

`func (o *NodeBalancer) SetClientConnThrottle(v int32)`

SetClientConnThrottle sets ClientConnThrottle field to given value.

### HasClientConnThrottle

`func (o *NodeBalancer) HasClientConnThrottle() bool`

HasClientConnThrottle returns a boolean if a field has been set.

### GetCreated

`func (o *NodeBalancer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NodeBalancer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NodeBalancer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NodeBalancer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetHostname

`func (o *NodeBalancer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeBalancer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeBalancer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NodeBalancer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *NodeBalancer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeBalancer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeBalancer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NodeBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpv4

`func (o *NodeBalancer) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *NodeBalancer) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *NodeBalancer) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *NodeBalancer) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *NodeBalancer) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *NodeBalancer) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *NodeBalancer) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *NodeBalancer) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### SetIpv6Nil

`func (o *NodeBalancer) SetIpv6Nil(b bool)`

 SetIpv6Nil sets the value for Ipv6 to be an explicit nil

### UnsetIpv6
`func (o *NodeBalancer) UnsetIpv6()`

UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
### GetLabel

`func (o *NodeBalancer) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NodeBalancer) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NodeBalancer) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *NodeBalancer) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegion

`func (o *NodeBalancer) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NodeBalancer) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NodeBalancer) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NodeBalancer) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTags

`func (o *NodeBalancer) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NodeBalancer) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NodeBalancer) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NodeBalancer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransfer

`func (o *NodeBalancer) GetTransfer() NodeBalancerTransfer`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *NodeBalancer) GetTransferOk() (*NodeBalancerTransfer, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *NodeBalancer) SetTransfer(v NodeBalancerTransfer)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *NodeBalancer) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetUpdated

`func (o *NodeBalancer) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *NodeBalancer) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *NodeBalancer) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *NodeBalancer) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


