# PostAssignIpsRequestAssignmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The IPv4 address or IPv6 range for this assignment.  - Must be an IPv4 address or an IPv6 range you can access in the Region specified. - IPv6 ranges must include a prefix length of &#x60;/56&#x60; or &#x60;/64&#x60;, for example: &#x60;2001:db8:3c4d:15::/64&#x60;. - Assignment of an IPv6 range to a Linode updates the route target of the range to the assigned Linode&#39;s SLAAC address. - May be a public or private address. | 
**LinodeId** | **int32** | The ID of the Linode to assign this address to. The IP&#39;s previous Linode will lose this address, and must end up with at least one public address and no more than one private address once all assignments have been made. | 

## Methods

### NewPostAssignIpsRequestAssignmentsInner

`func NewPostAssignIpsRequestAssignmentsInner(address string, linodeId int32, ) *PostAssignIpsRequestAssignmentsInner`

NewPostAssignIpsRequestAssignmentsInner instantiates a new PostAssignIpsRequestAssignmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAssignIpsRequestAssignmentsInnerWithDefaults

`func NewPostAssignIpsRequestAssignmentsInnerWithDefaults() *PostAssignIpsRequestAssignmentsInner`

NewPostAssignIpsRequestAssignmentsInnerWithDefaults instantiates a new PostAssignIpsRequestAssignmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PostAssignIpsRequestAssignmentsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PostAssignIpsRequestAssignmentsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PostAssignIpsRequestAssignmentsInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetLinodeId

`func (o *PostAssignIpsRequestAssignmentsInner) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostAssignIpsRequestAssignmentsInner) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostAssignIpsRequestAssignmentsInner) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


