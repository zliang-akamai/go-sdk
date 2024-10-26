# PostIpv6Range200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** | The IPv6 network range, including subnet and prefix length. | [optional] 
**RouteTarget** | Pointer to **string** | The route target IPV6 SLAAC address for this range. Does not include the prefix length. | [optional] 

## Methods

### NewPostIpv6Range200Response

`func NewPostIpv6Range200Response() *PostIpv6Range200Response`

NewPostIpv6Range200Response instantiates a new PostIpv6Range200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostIpv6Range200ResponseWithDefaults

`func NewPostIpv6Range200ResponseWithDefaults() *PostIpv6Range200Response`

NewPostIpv6Range200ResponseWithDefaults instantiates a new PostIpv6Range200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *PostIpv6Range200Response) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *PostIpv6Range200Response) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *PostIpv6Range200Response) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *PostIpv6Range200Response) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRouteTarget

`func (o *PostIpv6Range200Response) GetRouteTarget() string`

GetRouteTarget returns the RouteTarget field if non-nil, zero value otherwise.

### GetRouteTargetOk

`func (o *PostIpv6Range200Response) GetRouteTargetOk() (*string, bool)`

GetRouteTargetOk returns a tuple with the RouteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTarget

`func (o *PostIpv6Range200Response) SetRouteTarget(v string)`

SetRouteTarget sets RouteTarget field to given value.

### HasRouteTarget

`func (o *PostIpv6Range200Response) HasRouteTarget() bool`

HasRouteTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


