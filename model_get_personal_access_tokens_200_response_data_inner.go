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
)

// checks if the GetPersonalAccessTokens200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPersonalAccessTokens200ResponseDataInner{}

// GetPersonalAccessTokens200ResponseDataInner A Personal Access Token is a token generated manually to access the API without going through an OAuth login.  Personal Access Tokens can have scopes just like OAuth tokens do, and are commonly used to give access to command-line tools like the Linode CLI, or when writing your own integrations.
type GetPersonalAccessTokens200ResponseDataInner struct {
	// The date and time this token was created.
	Created *time.Time `json:"created,omitempty"`
	// When this token will expire.  Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated.  Tokens may be created with `null` as their expiry and will never expire unless revoked.
	Expiry *time.Time `json:"expiry,omitempty"`
	// This token's unique ID, which can be used to revoke it.
	Id *int32 `json:"id,omitempty"`
	// This token's label.  This is for display purposes only, but can be used to more easily track what you're using each token for.
	Label *string `json:"label,omitempty"`
	// The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the [Linode CLI](https://github.com/linode/linode-cli), require tokens with access to `*`. Tokens with more restrictive scopes are generally more secure.
	Scopes *string `json:"scopes,omitempty"`
	// The token used to access the API.  When the token is created, the full token is returned here.  Otherwise, only the first 16 characters are returned.
	Token *string `json:"token,omitempty"`
}

// NewGetPersonalAccessTokens200ResponseDataInner instantiates a new GetPersonalAccessTokens200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPersonalAccessTokens200ResponseDataInner() *GetPersonalAccessTokens200ResponseDataInner {
	this := GetPersonalAccessTokens200ResponseDataInner{}
	return &this
}

// NewGetPersonalAccessTokens200ResponseDataInnerWithDefaults instantiates a new GetPersonalAccessTokens200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPersonalAccessTokens200ResponseDataInnerWithDefaults() *GetPersonalAccessTokens200ResponseDataInner {
	this := GetPersonalAccessTokens200ResponseDataInner{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetPersonalAccessTokens200ResponseDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *GetPersonalAccessTokens200ResponseDataInner) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetPersonalAccessTokens200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetPersonalAccessTokens200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetScopes() string {
	if o == nil || IsNil(o.Scopes) {
		var ret string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetScopesOk() (*string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given string and assigns it to the Scopes field.
func (o *GetPersonalAccessTokens200ResponseDataInner) SetScopes(v string) {
	o.Scopes = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetPersonalAccessTokens200ResponseDataInner) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetPersonalAccessTokens200ResponseDataInner) SetToken(v string) {
	o.Token = &v
}

func (o GetPersonalAccessTokens200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPersonalAccessTokens200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableGetPersonalAccessTokens200ResponseDataInner struct {
	value *GetPersonalAccessTokens200ResponseDataInner
	isSet bool
}

func (v NullableGetPersonalAccessTokens200ResponseDataInner) Get() *GetPersonalAccessTokens200ResponseDataInner {
	return v.value
}

func (v *NullableGetPersonalAccessTokens200ResponseDataInner) Set(val *GetPersonalAccessTokens200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPersonalAccessTokens200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPersonalAccessTokens200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPersonalAccessTokens200ResponseDataInner(val *GetPersonalAccessTokens200ResponseDataInner) *NullableGetPersonalAccessTokens200ResponseDataInner {
	return &NullableGetPersonalAccessTokens200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetPersonalAccessTokens200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPersonalAccessTokens200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


