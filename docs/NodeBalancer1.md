# NodeBalancer1

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

### NewNodeBalancer1

`func NewNodeBalancer1() *NodeBalancer1`

NewNodeBalancer1 instantiates a new NodeBalancer1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancer1WithDefaults

`func NewNodeBalancer1WithDefaults() *NodeBalancer1`

NewNodeBalancer1WithDefaults instantiates a new NodeBalancer1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientConnThrottle

`func (o *NodeBalancer1) GetClientConnThrottle() int32`

GetClientConnThrottle returns the ClientConnThrottle field if non-nil, zero value otherwise.

### GetClientConnThrottleOk

`func (o *NodeBalancer1) GetClientConnThrottleOk() (*int32, bool)`

GetClientConnThrottleOk returns a tuple with the ClientConnThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnThrottle

`func (o *NodeBalancer1) SetClientConnThrottle(v int32)`

SetClientConnThrottle sets ClientConnThrottle field to given value.

### HasClientConnThrottle

`func (o *NodeBalancer1) HasClientConnThrottle() bool`

HasClientConnThrottle returns a boolean if a field has been set.

### GetCreated

`func (o *NodeBalancer1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NodeBalancer1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NodeBalancer1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NodeBalancer1) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetHostname

`func (o *NodeBalancer1) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeBalancer1) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeBalancer1) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NodeBalancer1) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *NodeBalancer1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeBalancer1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeBalancer1) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NodeBalancer1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpv4

`func (o *NodeBalancer1) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *NodeBalancer1) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *NodeBalancer1) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *NodeBalancer1) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *NodeBalancer1) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *NodeBalancer1) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *NodeBalancer1) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *NodeBalancer1) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### SetIpv6Nil

`func (o *NodeBalancer1) SetIpv6Nil(b bool)`

 SetIpv6Nil sets the value for Ipv6 to be an explicit nil

### UnsetIpv6
`func (o *NodeBalancer1) UnsetIpv6()`

UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
### GetLabel

`func (o *NodeBalancer1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NodeBalancer1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NodeBalancer1) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *NodeBalancer1) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegion

`func (o *NodeBalancer1) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NodeBalancer1) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NodeBalancer1) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NodeBalancer1) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTags

`func (o *NodeBalancer1) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NodeBalancer1) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NodeBalancer1) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NodeBalancer1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransfer

`func (o *NodeBalancer1) GetTransfer() NodeBalancerTransfer`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *NodeBalancer1) GetTransferOk() (*NodeBalancerTransfer, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *NodeBalancer1) SetTransfer(v NodeBalancerTransfer)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *NodeBalancer1) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetUpdated

`func (o *NodeBalancer1) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *NodeBalancer1) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *NodeBalancer1) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *NodeBalancer1) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


