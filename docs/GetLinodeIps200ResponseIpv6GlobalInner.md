# GetLinodeIps200ResponseIpv6GlobalInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **int32** | The prefix length of the address. The total number of addresses that can be assigned from this range is calculated as 2&lt;sup&gt;(128 - prefix length)&lt;/sup&gt;. | [optional] 
**Range** | Pointer to **string** | The IPv6 address of this range. | [optional] [readonly] 
**Region** | Pointer to **string** | The region for this range of IPv6 addresses. | [optional] [readonly] 
**RouteTarget** | Pointer to **string** | The IPv6 SLAAC address. | [optional] 

## Methods

### NewGetLinodeIps200ResponseIpv6GlobalInner

`func NewGetLinodeIps200ResponseIpv6GlobalInner() *GetLinodeIps200ResponseIpv6GlobalInner`

NewGetLinodeIps200ResponseIpv6GlobalInner instantiates a new GetLinodeIps200ResponseIpv6GlobalInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeIps200ResponseIpv6GlobalInnerWithDefaults

`func NewGetLinodeIps200ResponseIpv6GlobalInnerWithDefaults() *GetLinodeIps200ResponseIpv6GlobalInner`

NewGetLinodeIps200ResponseIpv6GlobalInnerWithDefaults instantiates a new GetLinodeIps200ResponseIpv6GlobalInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRange

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRegion

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRouteTarget

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRouteTarget() string`

GetRouteTarget returns the RouteTarget field if non-nil, zero value otherwise.

### GetRouteTargetOk

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) GetRouteTargetOk() (*string, bool)`

GetRouteTargetOk returns a tuple with the RouteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTarget

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) SetRouteTarget(v string)`

SetRouteTarget sets RouteTarget field to given value.

### HasRouteTarget

`func (o *GetLinodeIps200ResponseIpv6GlobalInner) HasRouteTarget() bool`

HasRouteTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


