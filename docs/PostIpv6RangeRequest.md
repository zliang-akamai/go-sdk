# PostIpv6RangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinodeId** | Pointer to **int32** | The ID of the Linode to assign this range to. The SLAAC address for the provided Linode is used as the range&#39;s &#x60;route_target&#x60;.  - __Required__ if &#x60;route_target&#x60; is omitted from the request.  - Mutually exclusive with &#x60;route_target&#x60;. Submitting values for both properties in a request results in an error. | [optional] 
**PrefixLength** | **int32** | The prefix length of the IPv6 range. | 
**RouteTarget** | Pointer to **string** | The IPv6 SLAAC address to assign this range to.  - __Required__ if &#x60;linode_id&#x60; is omitted from the request.  - Mutually exclusive with &#x60;linode_id&#x60;. Submitting values for both properties in a request results in an error.  - __Note__. Omit the &#x60;/128&#x60; prefix length of the SLAAC address when using this property. | [optional] 

## Methods

### NewPostIpv6RangeRequest

`func NewPostIpv6RangeRequest(prefixLength int32, ) *PostIpv6RangeRequest`

NewPostIpv6RangeRequest instantiates a new PostIpv6RangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostIpv6RangeRequestWithDefaults

`func NewPostIpv6RangeRequestWithDefaults() *PostIpv6RangeRequest`

NewPostIpv6RangeRequestWithDefaults instantiates a new PostIpv6RangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinodeId

`func (o *PostIpv6RangeRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostIpv6RangeRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostIpv6RangeRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *PostIpv6RangeRequest) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetPrefixLength

`func (o *PostIpv6RangeRequest) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *PostIpv6RangeRequest) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *PostIpv6RangeRequest) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetRouteTarget

`func (o *PostIpv6RangeRequest) GetRouteTarget() string`

GetRouteTarget returns the RouteTarget field if non-nil, zero value otherwise.

### GetRouteTargetOk

`func (o *PostIpv6RangeRequest) GetRouteTargetOk() (*string, bool)`

GetRouteTargetOk returns a tuple with the RouteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTarget

`func (o *PostIpv6RangeRequest) SetRouteTarget(v string)`

SetRouteTarget sets RouteTarget field to given value.

### HasRouteTarget

`func (o *PostIpv6RangeRequest) HasRouteTarget() bool`

HasRouteTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


