# PostResetDiskPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The new root password for the OS installed on this Disk. The password must meet the complexity strength validation requirements for a strong password. | 

## Methods

### NewPostResetDiskPasswordRequest

`func NewPostResetDiskPasswordRequest(password string, ) *PostResetDiskPasswordRequest`

NewPostResetDiskPasswordRequest instantiates a new PostResetDiskPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostResetDiskPasswordRequestWithDefaults

`func NewPostResetDiskPasswordRequestWithDefaults() *PostResetDiskPasswordRequest`

NewPostResetDiskPasswordRequestWithDefaults instantiates a new PostResetDiskPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PostResetDiskPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostResetDiskPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostResetDiskPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


