# PostManagedCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Credential&#39;s unique ID. | [optional] [readonly] 
**Label** | **string** | The unique label for this Credential. This is for display purposes only. | 
**LastDecrypted** | Pointer to **time.Time** | The date this Credential was last decrypted by a member of Linode special forces. | [optional] [readonly] 
**Password** | **string** | The password to use when accessing the Managed Service. | 
**Username** | Pointer to **string** | The username to use when accessing the Managed Service. | [optional] 

## Methods

### NewPostManagedCredentialRequest

`func NewPostManagedCredentialRequest(label string, password string, ) *PostManagedCredentialRequest`

NewPostManagedCredentialRequest instantiates a new PostManagedCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostManagedCredentialRequestWithDefaults

`func NewPostManagedCredentialRequestWithDefaults() *PostManagedCredentialRequest`

NewPostManagedCredentialRequestWithDefaults instantiates a new PostManagedCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostManagedCredentialRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostManagedCredentialRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostManagedCredentialRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostManagedCredentialRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PostManagedCredentialRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostManagedCredentialRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostManagedCredentialRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastDecrypted

`func (o *PostManagedCredentialRequest) GetLastDecrypted() time.Time`

GetLastDecrypted returns the LastDecrypted field if non-nil, zero value otherwise.

### GetLastDecryptedOk

`func (o *PostManagedCredentialRequest) GetLastDecryptedOk() (*time.Time, bool)`

GetLastDecryptedOk returns a tuple with the LastDecrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDecrypted

`func (o *PostManagedCredentialRequest) SetLastDecrypted(v time.Time)`

SetLastDecrypted sets LastDecrypted field to given value.

### HasLastDecrypted

`func (o *PostManagedCredentialRequest) HasLastDecrypted() bool`

HasLastDecrypted returns a boolean if a field has been set.

### GetPassword

`func (o *PostManagedCredentialRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostManagedCredentialRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostManagedCredentialRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *PostManagedCredentialRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PostManagedCredentialRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PostManagedCredentialRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PostManagedCredentialRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


