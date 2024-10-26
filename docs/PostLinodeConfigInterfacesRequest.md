# PostLinodeConfigInterfacesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]int32** | An ordered array of existing Configuration Profile Interface &#x60;id&#x60;s.  - All current Interface &#x60;id&#x60;s must be present in the array. - If the Configuration Profile contains Interfaces and is active on the Linode, the Linode must first be shut down prior to running this operation. - Reordering takes effect after rebooting the Linode with this Configuration Profile.  The position in the array determines which of the Linode&#39;s network Interfaces is configured:  - First [0]:  eth0 - Second [1]: eth1 - Third [2]:  eth2 | 

## Methods

### NewPostLinodeConfigInterfacesRequest

`func NewPostLinodeConfigInterfacesRequest(ids []int32, ) *PostLinodeConfigInterfacesRequest`

NewPostLinodeConfigInterfacesRequest instantiates a new PostLinodeConfigInterfacesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLinodeConfigInterfacesRequestWithDefaults

`func NewPostLinodeConfigInterfacesRequestWithDefaults() *PostLinodeConfigInterfacesRequest`

NewPostLinodeConfigInterfacesRequestWithDefaults instantiates a new PostLinodeConfigInterfacesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *PostLinodeConfigInterfacesRequest) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *PostLinodeConfigInterfacesRequest) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *PostLinodeConfigInterfacesRequest) SetIds(v []int32)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


