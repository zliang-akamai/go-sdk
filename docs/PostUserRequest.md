# PostUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address for the User. Linode sends emails to this address for account management communications. May be used for other communications as configured. | 
**LastLogin** | Pointer to [**NullableGetUsers200ResponseDataInnerAllOfLastLogin**](GetUsers200ResponseDataInnerAllOfLastLogin.md) |  | [optional] 
**PasswordCreated** | Pointer to **NullableTime** | The date and time when this User&#39;s current password was created.  User passwords are first created during the Account sign-up process, and updated using the [Reset Password](https://login.linode.com/forgot/password) webpage.  &#x60;null&#x60; if this User has not created a password yet. | [optional] [readonly] 
**Restricted** | Pointer to **bool** | If true, the User must be granted access to perform actions or access entities on this Account. Run [List a user&#39;s grants](https://techdocs.akamai.com/linode-api/reference/get-user-grants) for details on how to configure grants for a restricted User. | [optional] 
**SshKeys** | Pointer to **[]string** | A list of SSH Key labels added by this User.  Users can add keys with the [Add an SSH key](https://techdocs.akamai.com/linode-api/reference/post-add-ssh-key) operation.  These keys are deployed when this User is included in the &#x60;authorized_users&#x60; field of the following requests:  - [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance) - [Rebuild a Linode](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance) - [Create a disk](https://techdocs.akamai.com/linode-api/reference/post-add-linode-disk) | [optional] [readonly] 
**TfaEnabled** | Pointer to **bool** | A boolean value indicating if the User has Two Factor Authentication (TFA) enabled. Run the [Create a two factor secret](https://techdocs.akamai.com/linode-api/reference/post-tfa-enable) operation to enable TFA. | [optional] [readonly] 
**Username** | **string** | The User&#39;s username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts). | 
**VerifiedPhoneNumber** | Pointer to **NullableString** | The phone number verified for this User Profile with the [Verify a phone number](https://techdocs.akamai.com/linode-api/reference/post-profile-phone-number-verify) operation.  &#x60;null&#x60; if this User Profile has no verified phone number. | [optional] [readonly] 

## Methods

### NewPostUserRequest

`func NewPostUserRequest(email string, username string, ) *PostUserRequest`

NewPostUserRequest instantiates a new PostUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUserRequestWithDefaults

`func NewPostUserRequestWithDefaults() *PostUserRequest`

NewPostUserRequestWithDefaults instantiates a new PostUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLastLogin

`func (o *PostUserRequest) GetLastLogin() GetUsers200ResponseDataInnerAllOfLastLogin`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PostUserRequest) GetLastLoginOk() (*GetUsers200ResponseDataInnerAllOfLastLogin, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PostUserRequest) SetLastLogin(v GetUsers200ResponseDataInnerAllOfLastLogin)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PostUserRequest) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *PostUserRequest) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *PostUserRequest) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetPasswordCreated

`func (o *PostUserRequest) GetPasswordCreated() time.Time`

GetPasswordCreated returns the PasswordCreated field if non-nil, zero value otherwise.

### GetPasswordCreatedOk

`func (o *PostUserRequest) GetPasswordCreatedOk() (*time.Time, bool)`

GetPasswordCreatedOk returns a tuple with the PasswordCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCreated

`func (o *PostUserRequest) SetPasswordCreated(v time.Time)`

SetPasswordCreated sets PasswordCreated field to given value.

### HasPasswordCreated

`func (o *PostUserRequest) HasPasswordCreated() bool`

HasPasswordCreated returns a boolean if a field has been set.

### SetPasswordCreatedNil

`func (o *PostUserRequest) SetPasswordCreatedNil(b bool)`

 SetPasswordCreatedNil sets the value for PasswordCreated to be an explicit nil

### UnsetPasswordCreated
`func (o *PostUserRequest) UnsetPasswordCreated()`

UnsetPasswordCreated ensures that no value is present for PasswordCreated, not even an explicit nil
### GetRestricted

`func (o *PostUserRequest) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *PostUserRequest) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *PostUserRequest) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *PostUserRequest) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetSshKeys

`func (o *PostUserRequest) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *PostUserRequest) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *PostUserRequest) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *PostUserRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetTfaEnabled

`func (o *PostUserRequest) GetTfaEnabled() bool`

GetTfaEnabled returns the TfaEnabled field if non-nil, zero value otherwise.

### GetTfaEnabledOk

`func (o *PostUserRequest) GetTfaEnabledOk() (*bool, bool)`

GetTfaEnabledOk returns a tuple with the TfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaEnabled

`func (o *PostUserRequest) SetTfaEnabled(v bool)`

SetTfaEnabled sets TfaEnabled field to given value.

### HasTfaEnabled

`func (o *PostUserRequest) HasTfaEnabled() bool`

HasTfaEnabled returns a boolean if a field has been set.

### GetUsername

`func (o *PostUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PostUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PostUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetVerifiedPhoneNumber

`func (o *PostUserRequest) GetVerifiedPhoneNumber() string`

GetVerifiedPhoneNumber returns the VerifiedPhoneNumber field if non-nil, zero value otherwise.

### GetVerifiedPhoneNumberOk

`func (o *PostUserRequest) GetVerifiedPhoneNumberOk() (*string, bool)`

GetVerifiedPhoneNumberOk returns a tuple with the VerifiedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPhoneNumber

`func (o *PostUserRequest) SetVerifiedPhoneNumber(v string)`

SetVerifiedPhoneNumber sets VerifiedPhoneNumber field to given value.

### HasVerifiedPhoneNumber

`func (o *PostUserRequest) HasVerifiedPhoneNumber() bool`

HasVerifiedPhoneNumber returns a boolean if a field has been set.

### SetVerifiedPhoneNumberNil

`func (o *PostUserRequest) SetVerifiedPhoneNumberNil(b bool)`

 SetVerifiedPhoneNumberNil sets the value for VerifiedPhoneNumber to be an explicit nil

### UnsetVerifiedPhoneNumber
`func (o *PostUserRequest) UnsetVerifiedPhoneNumber()`

UnsetVerifiedPhoneNumber ensures that no value is present for VerifiedPhoneNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


