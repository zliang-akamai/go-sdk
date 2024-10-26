# GetProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **string** | This account&#39;s Cloud Manager authentication type. Authentication types are chosen through Cloud Manager and authorized when logging into your account. These authentication types are either the user&#39;s password (in conjunction with their username), or the name of their identity provider such as GitHub. For example, if a user:  - Has never used Third-Party Authentication, their authentication type will be &#x60;password&#x60;. - Is using Third-Party Authentication, their authentication type will be the name of their Identity Provider (eg. &#x60;github&#x60;). - Has used Third-Party Authentication and has since revoked it, their authentication type will be &#x60;password&#x60;.  __Note__. This functionality may not yet be available in Cloud Manager. See the [Cloud Manager Changelog](https://www.linode.com/docs/products/tools/cloud-manager/release-notes/) for the latest updates. | [optional] [readonly] 
**AuthorizedKeys** | Pointer to **[]string** | The list of SSH Keys authorized to use Lish for your User. This value is ignored if &#x60;lish_auth_method&#x60; is &#x60;disabled&#x60;. | [optional] 
**Email** | Pointer to **string** | Your email address.  This address will be used for communication with Linode as necessary. | [optional] 
**EmailNotifications** | Pointer to **bool** | If true, you will receive email notifications about account activity.  If false, you may still receive business-critical communications through email. | [optional] 
**IpWhitelistEnabled** | Pointer to **bool** | If true, logins for your User will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled. If you disable this setting, you will not be able to re-enable it. | [optional] 
**LishAuthMethod** | Pointer to **string** | The authentication methods that are allowed when connecting to [the Linode Shell (Lish)](https://www.linode.com/docs/guides/lish/).  - &#x60;keys_only&#x60; is the most secure if you intend to use Lish. - &#x60;disabled&#x60; is recommended if you do not intend to use Lish at all. - If this account&#39;s Cloud Manager authentication type is set to a Third-Party Authentication method, &#x60;password_keys&#x60; cannot be used as your Lish authentication method. To view this account&#39;s Cloud Manager &#x60;authentication_type&#x60; field, send a request to the [Get a profile](https://techdocs.akamai.com/linode-api/reference/get-profile) operation. | [optional] 
**Referrals** | Pointer to [**GetProfile200ResponseReferrals**](GetProfile200ResponseReferrals.md) |  | [optional] 
**Restricted** | Pointer to **bool** | If true, your User has restrictions on what can be accessed on your Account. To get details on what entities/actions you can access/perform, run [List grants](https://techdocs.akamai.com/linode-api/reference/get-profile-grants). | [optional] 
**Timezone** | Pointer to **string** | The timezone you prefer to see times in. This is not used by the API directly. It is provided for the benefit of clients such as the Linode Cloud Manager and other clients built on the API. All times returned by the API are in UTC. | [optional] 
**TwoFactorAuth** | Pointer to **bool** | If true, logins from untrusted computers will require Two Factor Authentication.  Run [Create a two factor secret](https://techdocs.akamai.com/linode-api/reference/post-tfa-enable) to enable Two Factor Authentication. | [optional] 
**Uid** | Pointer to **int32** | Your unique ID in our system. This value will never change, and can safely be used to identify your User. | [optional] [readonly] 
**Username** | Pointer to **string** | Your username, used for logging in to our system. | [optional] [readonly] 
**VerifiedPhoneNumber** | Pointer to **NullableString** | The phone number verified for this Profile with the [Verify a phone number](https://techdocs.akamai.com/linode-api/reference/post-profile-phone-number-verify) operation.  &#x60;null&#x60; if this Profile has no verified phone number. | [optional] [readonly] 

## Methods

### NewGetProfile200Response

`func NewGetProfile200Response() *GetProfile200Response`

NewGetProfile200Response instantiates a new GetProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfile200ResponseWithDefaults

`func NewGetProfile200ResponseWithDefaults() *GetProfile200Response`

NewGetProfile200ResponseWithDefaults instantiates a new GetProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *GetProfile200Response) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *GetProfile200Response) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *GetProfile200Response) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *GetProfile200Response) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAuthorizedKeys

`func (o *GetProfile200Response) GetAuthorizedKeys() []string`

GetAuthorizedKeys returns the AuthorizedKeys field if non-nil, zero value otherwise.

### GetAuthorizedKeysOk

`func (o *GetProfile200Response) GetAuthorizedKeysOk() (*[]string, bool)`

GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeys

`func (o *GetProfile200Response) SetAuthorizedKeys(v []string)`

SetAuthorizedKeys sets AuthorizedKeys field to given value.

### HasAuthorizedKeys

`func (o *GetProfile200Response) HasAuthorizedKeys() bool`

HasAuthorizedKeys returns a boolean if a field has been set.

### SetAuthorizedKeysNil

`func (o *GetProfile200Response) SetAuthorizedKeysNil(b bool)`

 SetAuthorizedKeysNil sets the value for AuthorizedKeys to be an explicit nil

### UnsetAuthorizedKeys
`func (o *GetProfile200Response) UnsetAuthorizedKeys()`

UnsetAuthorizedKeys ensures that no value is present for AuthorizedKeys, not even an explicit nil
### GetEmail

`func (o *GetProfile200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetProfile200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetProfile200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetProfile200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailNotifications

`func (o *GetProfile200Response) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *GetProfile200Response) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *GetProfile200Response) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *GetProfile200Response) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetIpWhitelistEnabled

`func (o *GetProfile200Response) GetIpWhitelistEnabled() bool`

GetIpWhitelistEnabled returns the IpWhitelistEnabled field if non-nil, zero value otherwise.

### GetIpWhitelistEnabledOk

`func (o *GetProfile200Response) GetIpWhitelistEnabledOk() (*bool, bool)`

GetIpWhitelistEnabledOk returns a tuple with the IpWhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelistEnabled

`func (o *GetProfile200Response) SetIpWhitelistEnabled(v bool)`

SetIpWhitelistEnabled sets IpWhitelistEnabled field to given value.

### HasIpWhitelistEnabled

`func (o *GetProfile200Response) HasIpWhitelistEnabled() bool`

HasIpWhitelistEnabled returns a boolean if a field has been set.

### GetLishAuthMethod

`func (o *GetProfile200Response) GetLishAuthMethod() string`

GetLishAuthMethod returns the LishAuthMethod field if non-nil, zero value otherwise.

### GetLishAuthMethodOk

`func (o *GetProfile200Response) GetLishAuthMethodOk() (*string, bool)`

GetLishAuthMethodOk returns a tuple with the LishAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLishAuthMethod

`func (o *GetProfile200Response) SetLishAuthMethod(v string)`

SetLishAuthMethod sets LishAuthMethod field to given value.

### HasLishAuthMethod

`func (o *GetProfile200Response) HasLishAuthMethod() bool`

HasLishAuthMethod returns a boolean if a field has been set.

### GetReferrals

`func (o *GetProfile200Response) GetReferrals() GetProfile200ResponseReferrals`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *GetProfile200Response) GetReferralsOk() (*GetProfile200ResponseReferrals, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *GetProfile200Response) SetReferrals(v GetProfile200ResponseReferrals)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *GetProfile200Response) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetRestricted

`func (o *GetProfile200Response) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *GetProfile200Response) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *GetProfile200Response) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *GetProfile200Response) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetTimezone

`func (o *GetProfile200Response) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetProfile200Response) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetProfile200Response) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetProfile200Response) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTwoFactorAuth

`func (o *GetProfile200Response) GetTwoFactorAuth() bool`

GetTwoFactorAuth returns the TwoFactorAuth field if non-nil, zero value otherwise.

### GetTwoFactorAuthOk

`func (o *GetProfile200Response) GetTwoFactorAuthOk() (*bool, bool)`

GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuth

`func (o *GetProfile200Response) SetTwoFactorAuth(v bool)`

SetTwoFactorAuth sets TwoFactorAuth field to given value.

### HasTwoFactorAuth

`func (o *GetProfile200Response) HasTwoFactorAuth() bool`

HasTwoFactorAuth returns a boolean if a field has been set.

### GetUid

`func (o *GetProfile200Response) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetProfile200Response) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetProfile200Response) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetProfile200Response) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *GetProfile200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetProfile200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetProfile200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetProfile200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVerifiedPhoneNumber

`func (o *GetProfile200Response) GetVerifiedPhoneNumber() string`

GetVerifiedPhoneNumber returns the VerifiedPhoneNumber field if non-nil, zero value otherwise.

### GetVerifiedPhoneNumberOk

`func (o *GetProfile200Response) GetVerifiedPhoneNumberOk() (*string, bool)`

GetVerifiedPhoneNumberOk returns a tuple with the VerifiedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPhoneNumber

`func (o *GetProfile200Response) SetVerifiedPhoneNumber(v string)`

SetVerifiedPhoneNumber sets VerifiedPhoneNumber field to given value.

### HasVerifiedPhoneNumber

`func (o *GetProfile200Response) HasVerifiedPhoneNumber() bool`

HasVerifiedPhoneNumber returns a boolean if a field has been set.

### SetVerifiedPhoneNumberNil

`func (o *GetProfile200Response) SetVerifiedPhoneNumberNil(b bool)`

 SetVerifiedPhoneNumberNil sets the value for VerifiedPhoneNumber to be an explicit nil

### UnsetVerifiedPhoneNumber
`func (o *GetProfile200Response) UnsetVerifiedPhoneNumber()`

UnsetVerifiedPhoneNumber ensures that no value is present for VerifiedPhoneNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


