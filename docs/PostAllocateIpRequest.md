# PostAllocateIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinodeId** | **int32** | The ID of a Linode you have access to that this address will be allocated to. | 
**Public** | **bool** | Whether to create a public or private IPv4 address. | 
**Type** | **string** | The type of address you are requesting. Only IPv4 addresses may be allocated through this operation. | 

## Methods

### NewPostAllocateIpRequest

`func NewPostAllocateIpRequest(linodeId int32, public bool, type_ string, ) *PostAllocateIpRequest`

NewPostAllocateIpRequest instantiates a new PostAllocateIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAllocateIpRequestWithDefaults

`func NewPostAllocateIpRequestWithDefaults() *PostAllocateIpRequest`

NewPostAllocateIpRequestWithDefaults instantiates a new PostAllocateIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinodeId

`func (o *PostAllocateIpRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostAllocateIpRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostAllocateIpRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.


### GetPublic

`func (o *PostAllocateIpRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *PostAllocateIpRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *PostAllocateIpRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetType

`func (o *PostAllocateIpRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAllocateIpRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAllocateIpRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


