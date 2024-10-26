# GetIpv6Range200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBgp** | Pointer to **bool** | Whether this IPv6 range is shared. | [optional] [readonly] 
**Linodes** | Pointer to **[]int32** | A list of Linodes targeted by this IPv6 range. Includes Linodes with IP sharing. | [optional] [readonly] 
**Prefix** | Pointer to **int32** | The prefix length of the address. The total number of addresses that can be assigned from this range is calculated as 2&lt;sup&gt;(128 - prefix length)&lt;/sup&gt;. | [optional] 
**Range** | Pointer to **string** | The IPv6 address of this range. | [optional] [readonly] 
**Region** | Pointer to **string** | The region for this range of IPv6 addresses. | [optional] [readonly] 

## Methods

### NewGetIpv6Range200Response

`func NewGetIpv6Range200Response() *GetIpv6Range200Response`

NewGetIpv6Range200Response instantiates a new GetIpv6Range200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6Range200ResponseWithDefaults

`func NewGetIpv6Range200ResponseWithDefaults() *GetIpv6Range200Response`

NewGetIpv6Range200ResponseWithDefaults instantiates a new GetIpv6Range200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBgp

`func (o *GetIpv6Range200Response) GetIsBgp() bool`

GetIsBgp returns the IsBgp field if non-nil, zero value otherwise.

### GetIsBgpOk

`func (o *GetIpv6Range200Response) GetIsBgpOk() (*bool, bool)`

GetIsBgpOk returns a tuple with the IsBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBgp

`func (o *GetIpv6Range200Response) SetIsBgp(v bool)`

SetIsBgp sets IsBgp field to given value.

### HasIsBgp

`func (o *GetIpv6Range200Response) HasIsBgp() bool`

HasIsBgp returns a boolean if a field has been set.

### GetLinodes

`func (o *GetIpv6Range200Response) GetLinodes() []int32`

GetLinodes returns the Linodes field if non-nil, zero value otherwise.

### GetLinodesOk

`func (o *GetIpv6Range200Response) GetLinodesOk() (*[]int32, bool)`

GetLinodesOk returns a tuple with the Linodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodes

`func (o *GetIpv6Range200Response) SetLinodes(v []int32)`

SetLinodes sets Linodes field to given value.

### HasLinodes

`func (o *GetIpv6Range200Response) HasLinodes() bool`

HasLinodes returns a boolean if a field has been set.

### GetPrefix

`func (o *GetIpv6Range200Response) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetIpv6Range200Response) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetIpv6Range200Response) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetIpv6Range200Response) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRange

`func (o *GetIpv6Range200Response) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetIpv6Range200Response) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetIpv6Range200Response) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GetIpv6Range200Response) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRegion

`func (o *GetIpv6Range200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetIpv6Range200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetIpv6Range200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetIpv6Range200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


