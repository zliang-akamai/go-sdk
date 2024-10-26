# PostShareIpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | **[]string** | A list of secondary Linode IPs to share with the primary Linode.  - Can include both IPv4 addresses and IPv6 ranges (omit /56 and /64 prefix lengths) - Can include both private and public IPv4 addresses. - You must have access to all of these addresses and they must be in the same Region as the primary Linode. - Enter an empty array to remove all shared IP addresses. | 
**LinodeId** | **int32** | The ID of the primary Linode that the addresses will be shared with. | 

## Methods

### NewPostShareIpsRequest

`func NewPostShareIpsRequest(ips []string, linodeId int32, ) *PostShareIpsRequest`

NewPostShareIpsRequest instantiates a new PostShareIpsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostShareIpsRequestWithDefaults

`func NewPostShareIpsRequestWithDefaults() *PostShareIpsRequest`

NewPostShareIpsRequestWithDefaults instantiates a new PostShareIpsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *PostShareIpsRequest) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *PostShareIpsRequest) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *PostShareIpsRequest) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetLinodeId

`func (o *PostShareIpsRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostShareIpsRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostShareIpsRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


