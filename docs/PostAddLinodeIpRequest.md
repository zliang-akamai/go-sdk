# PostAddLinodeIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | **bool** | Whether to create a public or private IPv4 address. | 
**Type** | **string** | The type of address you are allocating. Only IPv4 addresses may be allocated through this operation. | 

## Methods

### NewPostAddLinodeIpRequest

`func NewPostAddLinodeIpRequest(public bool, type_ string, ) *PostAddLinodeIpRequest`

NewPostAddLinodeIpRequest instantiates a new PostAddLinodeIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAddLinodeIpRequestWithDefaults

`func NewPostAddLinodeIpRequestWithDefaults() *PostAddLinodeIpRequest`

NewPostAddLinodeIpRequestWithDefaults instantiates a new PostAddLinodeIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *PostAddLinodeIpRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *PostAddLinodeIpRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *PostAddLinodeIpRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetType

`func (o *PostAddLinodeIpRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAddLinodeIpRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAddLinodeIpRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


