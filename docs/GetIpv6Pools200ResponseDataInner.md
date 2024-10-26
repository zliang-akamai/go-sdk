# GetIpv6Pools200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **int32** | The prefix length of the address. The total number of addresses that can be assigned from this range is calculated as 2&lt;sup&gt;(128 - prefix length)&lt;/sup&gt;. | [optional] 
**Range** | Pointer to **string** | The IPv6 range of addresses in this pool. | [optional] [readonly] 
**Region** | Pointer to **string** | The region for this pool of IPv6 addresses. | [optional] [readonly] 
**RouteTarget** | Pointer to **NullableString** | The last address in this block of IPv6 addresses. | [optional] 

## Methods

### NewGetIpv6Pools200ResponseDataInner

`func NewGetIpv6Pools200ResponseDataInner() *GetIpv6Pools200ResponseDataInner`

NewGetIpv6Pools200ResponseDataInner instantiates a new GetIpv6Pools200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6Pools200ResponseDataInnerWithDefaults

`func NewGetIpv6Pools200ResponseDataInnerWithDefaults() *GetIpv6Pools200ResponseDataInner`

NewGetIpv6Pools200ResponseDataInnerWithDefaults instantiates a new GetIpv6Pools200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *GetIpv6Pools200ResponseDataInner) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetIpv6Pools200ResponseDataInner) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetIpv6Pools200ResponseDataInner) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetIpv6Pools200ResponseDataInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRange

`func (o *GetIpv6Pools200ResponseDataInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetIpv6Pools200ResponseDataInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetIpv6Pools200ResponseDataInner) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GetIpv6Pools200ResponseDataInner) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRegion

`func (o *GetIpv6Pools200ResponseDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetIpv6Pools200ResponseDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetIpv6Pools200ResponseDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetIpv6Pools200ResponseDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRouteTarget

`func (o *GetIpv6Pools200ResponseDataInner) GetRouteTarget() string`

GetRouteTarget returns the RouteTarget field if non-nil, zero value otherwise.

### GetRouteTargetOk

`func (o *GetIpv6Pools200ResponseDataInner) GetRouteTargetOk() (*string, bool)`

GetRouteTargetOk returns a tuple with the RouteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTarget

`func (o *GetIpv6Pools200ResponseDataInner) SetRouteTarget(v string)`

SetRouteTarget sets RouteTarget field to given value.

### HasRouteTarget

`func (o *GetIpv6Pools200ResponseDataInner) HasRouteTarget() bool`

HasRouteTarget returns a boolean if a field has been set.

### SetRouteTargetNil

`func (o *GetIpv6Pools200ResponseDataInner) SetRouteTargetNil(b bool)`

 SetRouteTargetNil sets the value for RouteTarget to be an explicit nil

### UnsetRouteTarget
`func (o *GetIpv6Pools200ResponseDataInner) UnsetRouteTarget()`

UnsetRouteTarget ensures that no value is present for RouteTarget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


