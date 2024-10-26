# PostManagedCredentialUsernamePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The password to use when accessing the Managed Service. | 
**Username** | Pointer to **string** | The username to use when accessing the Managed Service. | [optional] 

## Methods

### NewPostManagedCredentialUsernamePasswordRequest

`func NewPostManagedCredentialUsernamePasswordRequest(password string, ) *PostManagedCredentialUsernamePasswordRequest`

NewPostManagedCredentialUsernamePasswordRequest instantiates a new PostManagedCredentialUsernamePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostManagedCredentialUsernamePasswordRequestWithDefaults

`func NewPostManagedCredentialUsernamePasswordRequestWithDefaults() *PostManagedCredentialUsernamePasswordRequest`

NewPostManagedCredentialUsernamePasswordRequestWithDefaults instantiates a new PostManagedCredentialUsernamePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PostManagedCredentialUsernamePasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostManagedCredentialUsernamePasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostManagedCredentialUsernamePasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *PostManagedCredentialUsernamePasswordRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PostManagedCredentialUsernamePasswordRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PostManagedCredentialUsernamePasswordRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PostManagedCredentialUsernamePasswordRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


