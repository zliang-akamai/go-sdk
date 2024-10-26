/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the PostUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostUserRequest{}

// PostUserRequest struct for PostUserRequest
type PostUserRequest struct {
	// The email address for the User. Linode sends emails to this address for account management communications. May be used for other communications as configured.
	Email string `json:"email"`
	LastLogin NullableGetUsers200ResponseDataInnerAllOfLastLogin `json:"last_login,omitempty"`
	// The date and time when this User's current password was created.  User passwords are first created during the Account sign-up process, and updated using the [Reset Password](https://login.linode.com/forgot/password) webpage.  `null` if this User has not created a password yet.
	PasswordCreated NullableTime `json:"password_created,omitempty"`
	// If true, the User must be granted access to perform actions or access entities on this Account. Run [List a user's grants](https://techdocs.akamai.com/linode-api/reference/get-user-grants) for details on how to configure grants for a restricted User.
	Restricted *bool `json:"restricted,omitempty"`
	// A list of SSH Key labels added by this User.  Users can add keys with the [Add an SSH key](https://techdocs.akamai.com/linode-api/reference/post-add-ssh-key) operation.  These keys are deployed when this User is included in the `authorized_users` field of the following requests:  - [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance) - [Rebuild a Linode](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance) - [Create a disk](https://techdocs.akamai.com/linode-api/reference/post-add-linode-disk)
	SshKeys []string `json:"ssh_keys,omitempty"`
	// A boolean value indicating if the User has Two Factor Authentication (TFA) enabled. Run the [Create a two factor secret](https://techdocs.akamai.com/linode-api/reference/post-tfa-enable) operation to enable TFA.
	TfaEnabled *bool `json:"tfa_enabled,omitempty"`
	// The User's username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts).
	Username string `json:"username" validate:"regexp=^[a-zA-Z0-9]((?![_-]{2,})[a-zA-Z0-9-_])+[a-zA-Z0-9]$"`
	// The phone number verified for this User Profile with the [Verify a phone number](https://techdocs.akamai.com/linode-api/reference/post-profile-phone-number-verify) operation.  `null` if this User Profile has no verified phone number.
	VerifiedPhoneNumber NullableString `json:"verified_phone_number,omitempty"`
}

type _PostUserRequest PostUserRequest

// NewPostUserRequest instantiates a new PostUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostUserRequest(email string, username string) *PostUserRequest {
	this := PostUserRequest{}
	this.Email = email
	this.Username = username
	return &this
}

// NewPostUserRequestWithDefaults instantiates a new PostUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostUserRequestWithDefaults() *PostUserRequest {
	this := PostUserRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *PostUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PostUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PostUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostUserRequest) GetLastLogin() GetUsers200ResponseDataInnerAllOfLastLogin {
	if o == nil || IsNil(o.LastLogin.Get()) {
		var ret GetUsers200ResponseDataInnerAllOfLastLogin
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostUserRequest) GetLastLoginOk() (*GetUsers200ResponseDataInnerAllOfLastLogin, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *PostUserRequest) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableGetUsers200ResponseDataInnerAllOfLastLogin and assigns it to the LastLogin field.
func (o *PostUserRequest) SetLastLogin(v GetUsers200ResponseDataInnerAllOfLastLogin) {
	o.LastLogin.Set(&v)
}
// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *PostUserRequest) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *PostUserRequest) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetPasswordCreated returns the PasswordCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostUserRequest) GetPasswordCreated() time.Time {
	if o == nil || IsNil(o.PasswordCreated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PasswordCreated.Get()
}

// GetPasswordCreatedOk returns a tuple with the PasswordCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostUserRequest) GetPasswordCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordCreated.Get(), o.PasswordCreated.IsSet()
}

// HasPasswordCreated returns a boolean if a field has been set.
func (o *PostUserRequest) HasPasswordCreated() bool {
	if o != nil && o.PasswordCreated.IsSet() {
		return true
	}

	return false
}

// SetPasswordCreated gets a reference to the given NullableTime and assigns it to the PasswordCreated field.
func (o *PostUserRequest) SetPasswordCreated(v time.Time) {
	o.PasswordCreated.Set(&v)
}
// SetPasswordCreatedNil sets the value for PasswordCreated to be an explicit nil
func (o *PostUserRequest) SetPasswordCreatedNil() {
	o.PasswordCreated.Set(nil)
}

// UnsetPasswordCreated ensures that no value is present for PasswordCreated, not even an explicit nil
func (o *PostUserRequest) UnsetPasswordCreated() {
	o.PasswordCreated.Unset()
}

// GetRestricted returns the Restricted field value if set, zero value otherwise.
func (o *PostUserRequest) GetRestricted() bool {
	if o == nil || IsNil(o.Restricted) {
		var ret bool
		return ret
	}
	return *o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostUserRequest) GetRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.Restricted) {
		return nil, false
	}
	return o.Restricted, true
}

// HasRestricted returns a boolean if a field has been set.
func (o *PostUserRequest) HasRestricted() bool {
	if o != nil && !IsNil(o.Restricted) {
		return true
	}

	return false
}

// SetRestricted gets a reference to the given bool and assigns it to the Restricted field.
func (o *PostUserRequest) SetRestricted(v bool) {
	o.Restricted = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *PostUserRequest) GetSshKeys() []string {
	if o == nil || IsNil(o.SshKeys) {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostUserRequest) GetSshKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *PostUserRequest) HasSshKeys() bool {
	if o != nil && !IsNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *PostUserRequest) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetTfaEnabled returns the TfaEnabled field value if set, zero value otherwise.
func (o *PostUserRequest) GetTfaEnabled() bool {
	if o == nil || IsNil(o.TfaEnabled) {
		var ret bool
		return ret
	}
	return *o.TfaEnabled
}

// GetTfaEnabledOk returns a tuple with the TfaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostUserRequest) GetTfaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TfaEnabled) {
		return nil, false
	}
	return o.TfaEnabled, true
}

// HasTfaEnabled returns a boolean if a field has been set.
func (o *PostUserRequest) HasTfaEnabled() bool {
	if o != nil && !IsNil(o.TfaEnabled) {
		return true
	}

	return false
}

// SetTfaEnabled gets a reference to the given bool and assigns it to the TfaEnabled field.
func (o *PostUserRequest) SetTfaEnabled(v bool) {
	o.TfaEnabled = &v
}

// GetUsername returns the Username field value
func (o *PostUserRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *PostUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *PostUserRequest) SetUsername(v string) {
	o.Username = v
}

// GetVerifiedPhoneNumber returns the VerifiedPhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostUserRequest) GetVerifiedPhoneNumber() string {
	if o == nil || IsNil(o.VerifiedPhoneNumber.Get()) {
		var ret string
		return ret
	}
	return *o.VerifiedPhoneNumber.Get()
}

// GetVerifiedPhoneNumberOk returns a tuple with the VerifiedPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostUserRequest) GetVerifiedPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerifiedPhoneNumber.Get(), o.VerifiedPhoneNumber.IsSet()
}

// HasVerifiedPhoneNumber returns a boolean if a field has been set.
func (o *PostUserRequest) HasVerifiedPhoneNumber() bool {
	if o != nil && o.VerifiedPhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetVerifiedPhoneNumber gets a reference to the given NullableString and assigns it to the VerifiedPhoneNumber field.
func (o *PostUserRequest) SetVerifiedPhoneNumber(v string) {
	o.VerifiedPhoneNumber.Set(&v)
}
// SetVerifiedPhoneNumberNil sets the value for VerifiedPhoneNumber to be an explicit nil
func (o *PostUserRequest) SetVerifiedPhoneNumberNil() {
	o.VerifiedPhoneNumber.Set(nil)
}

// UnsetVerifiedPhoneNumber ensures that no value is present for VerifiedPhoneNumber, not even an explicit nil
func (o *PostUserRequest) UnsetVerifiedPhoneNumber() {
	o.VerifiedPhoneNumber.Unset()
}

func (o PostUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if o.LastLogin.IsSet() {
		toSerialize["last_login"] = o.LastLogin.Get()
	}
	if o.PasswordCreated.IsSet() {
		toSerialize["password_created"] = o.PasswordCreated.Get()
	}
	if !IsNil(o.Restricted) {
		toSerialize["restricted"] = o.Restricted
	}
	if !IsNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !IsNil(o.TfaEnabled) {
		toSerialize["tfa_enabled"] = o.TfaEnabled
	}
	toSerialize["username"] = o.Username
	if o.VerifiedPhoneNumber.IsSet() {
		toSerialize["verified_phone_number"] = o.VerifiedPhoneNumber.Get()
	}
	return toSerialize, nil
}

func (o *PostUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"username",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostUserRequest := _PostUserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostUserRequest)

	if err != nil {
		return err
	}

	*o = PostUserRequest(varPostUserRequest)

	return err
}

type NullablePostUserRequest struct {
	value *PostUserRequest
	isSet bool
}

func (v NullablePostUserRequest) Get() *PostUserRequest {
	return v.value
}

func (v *NullablePostUserRequest) Set(val *PostUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostUserRequest(val *PostUserRequest) *NullablePostUserRequest {
	return &NullablePostUserRequest{value: val, isSet: true}
}

func (v NullablePostUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


