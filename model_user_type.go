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
)

// checks if the UserType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserType{}

// UserType The type of user on an account. Mostly applies to the use of the parent and child accounts for Akamai partners functionality.
type UserType struct {
	// If the user belongs to a [parent or child account](https://www.linode.com/docs/guides/parent-child-accounts/) relationship, this defines the user type in the respective account. Possible values include:  - `parent`. This is a user in an Akamai partner account. Akamai partners have a contractural relationship with their end customers, to sell Akamai services. This user can either have full access (a parent account admin user) or limited access. Limited users don't have access to manage child accounts, but they can be granted this access by an admin user.  - `child`. This is an Akamai partner's end customer user, in a child account. A child user can have either full or limited access. Full access lets the user manage other child users and the proxy user in a child account.  - `proxy`. This is a user on a child account that gives parent account users access to that child account. A parent account user with the `child_account_access` grant can [Create a proxy user token](https://techdocs.akamai.com/linode-api/reference/post-child-account-token) from the parent account. The parent user can use this token to run API operations from the child account, as if they were a child user.  - `default`. This applies to all regular, non-parent-child account users.
	UserType *string `json:"user_type,omitempty"`
}

// NewUserType instantiates a new UserType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserType() *UserType {
	this := UserType{}
	return &this
}

// NewUserTypeWithDefaults instantiates a new UserType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypeWithDefaults() *UserType {
	this := UserType{}
	return &this
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *UserType) GetUserType() string {
	if o == nil || IsNil(o.UserType) {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserType) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *UserType) SetUserType(v string) {
	o.UserType = &v
}

func (o UserType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserType) {
		toSerialize["user_type"] = o.UserType
	}
	return toSerialize, nil
}

type NullableUserType struct {
	value *UserType
	isSet bool
}

func (v NullableUserType) Get() *UserType {
	return v.value
}

func (v *NullableUserType) Set(val *UserType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserType(val *UserType) *NullableUserType {
	return &NullableUserType{value: val, isSet: true}
}

func (v NullableUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


