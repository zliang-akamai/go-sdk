# GetUsers200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address for the User. Linode sends emails to this address for account management communications. May be used for other communications as configured. | [optional] 
**LastLogin** | Pointer to [**NullableGetUsers200ResponseDataInnerAllOfLastLogin**](GetUsers200ResponseDataInnerAllOfLastLogin.md) |  | [optional] 
**PasswordCreated** | Pointer to **NullableTime** | The date and time when this User&#39;s current password was created.  User passwords are first created during the Account sign-up process, and updated using the [Reset Password](https://login.linode.com/forgot/password) webpage.  &#x60;null&#x60; if this User has not created a password yet. | [optional] [readonly] 
**Restricted** | Pointer to **bool** | If true, the User must be granted access to perform actions or access entities on this Account. Run [List a user&#39;s grants](https://techdocs.akamai.com/linode-api/reference/get-user-grants) for details on how to configure grants for a restricted User. | [optional] 
**SshKeys** | Pointer to **[]string** | A list of SSH Key labels added by this User.  Users can add keys with the [Add an SSH key](https://techdocs.akamai.com/linode-api/reference/post-add-ssh-key) operation.  These keys are deployed when this User is included in the &#x60;authorized_users&#x60; field of the following requests:  - [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance) - [Rebuild a Linode](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance) - [Create a disk](https://techdocs.akamai.com/linode-api/reference/post-add-linode-disk) | [optional] [readonly] 
**TfaEnabled** | Pointer to **bool** | A boolean value indicating if the User has Two Factor Authentication (TFA) enabled. Run the [Create a two factor secret](https://techdocs.akamai.com/linode-api/reference/post-tfa-enable) operation to enable TFA. | [optional] [readonly] 
**Username** | Pointer to **string** | The User&#39;s username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts). | [optional] 
**VerifiedPhoneNumber** | Pointer to **NullableString** | The phone number verified for this User Profile with the [Verify a phone number](https://techdocs.akamai.com/linode-api/reference/post-profile-phone-number-verify) operation.  &#x60;null&#x60; if this User Profile has no verified phone number. | [optional] [readonly] 
**UserType** | Pointer to **string** | If the user belongs to a [parent or child account](https://www.linode.com/docs/guides/parent-child-accounts/) relationship, this defines the user type in the respective account. Possible values include:  - &#x60;parent&#x60;. This is a user in an Akamai partner account. Akamai partners have a contractural relationship with their end customers, to sell Akamai services. This user can either have full access (a parent account admin user) or limited access. Limited users don&#39;t have access to manage child accounts, but they can be granted this access by an admin user.  - &#x60;child&#x60;. This is an Akamai partner&#39;s end customer user, in a child account. A child user can have either full or limited access. Full access lets the user manage other child users and the proxy user in a child account.  - &#x60;proxy&#x60;. This is a user on a child account that gives parent account users access to that child account. A parent account user with the &#x60;child_account_access&#x60; grant can [Create a proxy user token](https://techdocs.akamai.com/linode-api/reference/post-child-account-token) from the parent account. The parent user can use this token to run API operations from the child account, as if they were a child user.  - &#x60;default&#x60;. This applies to all regular, non-parent-child account users. | [optional] [readonly] 

## Methods

### NewGetUsers200ResponseDataInner

`func NewGetUsers200ResponseDataInner() *GetUsers200ResponseDataInner`

NewGetUsers200ResponseDataInner instantiates a new GetUsers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsers200ResponseDataInnerWithDefaults

`func NewGetUsers200ResponseDataInnerWithDefaults() *GetUsers200ResponseDataInner`

NewGetUsers200ResponseDataInnerWithDefaults instantiates a new GetUsers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetUsers200ResponseDataInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUsers200ResponseDataInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUsers200ResponseDataInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUsers200ResponseDataInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLastLogin

`func (o *GetUsers200ResponseDataInner) GetLastLogin() GetUsers200ResponseDataInnerAllOfLastLogin`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *GetUsers200ResponseDataInner) GetLastLoginOk() (*GetUsers200ResponseDataInnerAllOfLastLogin, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *GetUsers200ResponseDataInner) SetLastLogin(v GetUsers200ResponseDataInnerAllOfLastLogin)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *GetUsers200ResponseDataInner) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *GetUsers200ResponseDataInner) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *GetUsers200ResponseDataInner) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetPasswordCreated

`func (o *GetUsers200ResponseDataInner) GetPasswordCreated() time.Time`

GetPasswordCreated returns the PasswordCreated field if non-nil, zero value otherwise.

### GetPasswordCreatedOk

`func (o *GetUsers200ResponseDataInner) GetPasswordCreatedOk() (*time.Time, bool)`

GetPasswordCreatedOk returns a tuple with the PasswordCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCreated

`func (o *GetUsers200ResponseDataInner) SetPasswordCreated(v time.Time)`

SetPasswordCreated sets PasswordCreated field to given value.

### HasPasswordCreated

`func (o *GetUsers200ResponseDataInner) HasPasswordCreated() bool`

HasPasswordCreated returns a boolean if a field has been set.

### SetPasswordCreatedNil

`func (o *GetUsers200ResponseDataInner) SetPasswordCreatedNil(b bool)`

 SetPasswordCreatedNil sets the value for PasswordCreated to be an explicit nil

### UnsetPasswordCreated
`func (o *GetUsers200ResponseDataInner) UnsetPasswordCreated()`

UnsetPasswordCreated ensures that no value is present for PasswordCreated, not even an explicit nil
### GetRestricted

`func (o *GetUsers200ResponseDataInner) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *GetUsers200ResponseDataInner) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *GetUsers200ResponseDataInner) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *GetUsers200ResponseDataInner) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetSshKeys

`func (o *GetUsers200ResponseDataInner) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *GetUsers200ResponseDataInner) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *GetUsers200ResponseDataInner) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *GetUsers200ResponseDataInner) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetTfaEnabled

`func (o *GetUsers200ResponseDataInner) GetTfaEnabled() bool`

GetTfaEnabled returns the TfaEnabled field if non-nil, zero value otherwise.

### GetTfaEnabledOk

`func (o *GetUsers200ResponseDataInner) GetTfaEnabledOk() (*bool, bool)`

GetTfaEnabledOk returns a tuple with the TfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaEnabled

`func (o *GetUsers200ResponseDataInner) SetTfaEnabled(v bool)`

SetTfaEnabled sets TfaEnabled field to given value.

### HasTfaEnabled

`func (o *GetUsers200ResponseDataInner) HasTfaEnabled() bool`

HasTfaEnabled returns a boolean if a field has been set.

### GetUsername

`func (o *GetUsers200ResponseDataInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetUsers200ResponseDataInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetUsers200ResponseDataInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetUsers200ResponseDataInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVerifiedPhoneNumber

`func (o *GetUsers200ResponseDataInner) GetVerifiedPhoneNumber() string`

GetVerifiedPhoneNumber returns the VerifiedPhoneNumber field if non-nil, zero value otherwise.

### GetVerifiedPhoneNumberOk

`func (o *GetUsers200ResponseDataInner) GetVerifiedPhoneNumberOk() (*string, bool)`

GetVerifiedPhoneNumberOk returns a tuple with the VerifiedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPhoneNumber

`func (o *GetUsers200ResponseDataInner) SetVerifiedPhoneNumber(v string)`

SetVerifiedPhoneNumber sets VerifiedPhoneNumber field to given value.

### HasVerifiedPhoneNumber

`func (o *GetUsers200ResponseDataInner) HasVerifiedPhoneNumber() bool`

HasVerifiedPhoneNumber returns a boolean if a field has been set.

### SetVerifiedPhoneNumberNil

`func (o *GetUsers200ResponseDataInner) SetVerifiedPhoneNumberNil(b bool)`

 SetVerifiedPhoneNumberNil sets the value for VerifiedPhoneNumber to be an explicit nil

### UnsetVerifiedPhoneNumber
`func (o *GetUsers200ResponseDataInner) UnsetVerifiedPhoneNumber()`

UnsetVerifiedPhoneNumber ensures that no value is present for VerifiedPhoneNumber, not even an explicit nil
### GetUserType

`func (o *GetUsers200ResponseDataInner) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *GetUsers200ResponseDataInner) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *GetUsers200ResponseDataInner) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *GetUsers200ResponseDataInner) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


