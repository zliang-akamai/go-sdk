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

// checks if the GetClients200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetClients200ResponseDataInner{}

// GetClients200ResponseDataInner A third-party application registered to Linode that users may log into with their Linode account through our authentication server at [login.linode.com](https://login.linode.com).  Using an OAuth Client, a third-party developer may be given access to some, or all, of a User's account for the purposes of their application.
type GetClients200ResponseDataInner struct {
	// The OAuth Client ID.  This is used to identify the client, and is a publicly known value (it is not a secret).
	Id *string `json:"id,omitempty"`
	// The name of this application.  This will be presented to users when they are asked to grant it access to their Account.
	Label *string `json:"label,omitempty"`
	// If this is a public or private OAuth Client.  Public clients have a slightly different authentication workflow than private clients.  See the [OAuth spec](https://oauth.net/2/) for more details.
	Public *bool `json:"public,omitempty"`
	// The location a successful log in from [login.linode.com](https://login.linode.com) should be redirected to for this client.  The receiver of this redirect should be ready to accept an OAuth exchange code and finish the OAuth exchange.
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// The OAuth Client secret, used in the OAuth exchange.  This is returned as `<REDACTED>` except when an OAuth Client is created or its secret is reset.  This is a secret, and should not be shared or disclosed publicly.
	Secret *string `json:"secret,omitempty"`
	// The status of this application.  `active` by default.
	Status *string `json:"status,omitempty"`
	// The URL where this client's thumbnail may be viewed, or `null` if this client does not have a thumbnail set.
	ThumbnailUrl NullableString `json:"thumbnail_url,omitempty"`
}

// NewGetClients200ResponseDataInner instantiates a new GetClients200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClients200ResponseDataInner() *GetClients200ResponseDataInner {
	this := GetClients200ResponseDataInner{}
	var public bool = false
	this.Public = &public
	return &this
}

// NewGetClients200ResponseDataInnerWithDefaults instantiates a new GetClients200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClients200ResponseDataInnerWithDefaults() *GetClients200ResponseDataInner {
	this := GetClients200ResponseDataInner{}
	var public bool = false
	this.Public = &public
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetClients200ResponseDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetClients200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetClients200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetClients200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetClients200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetClients200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *GetClients200ResponseDataInner) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseDataInner) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *GetClients200ResponseDataInner) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *GetClients200ResponseDataInner) SetPublic(v bool) {
	o.Public = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *GetClients200ResponseDataInner) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri) {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseDataInner) GetRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *GetClients200ResponseDataInner) HasRedirectUri() bool {
	if o != nil && !IsNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *GetClients200ResponseDataInner) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *GetClients200ResponseDataInner) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseDataInner) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *GetClients200ResponseDataInner) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *GetClients200ResponseDataInner) SetSecret(v string) {
	o.Secret = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetClients200ResponseDataInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetClients200ResponseDataInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetClients200ResponseDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetClients200ResponseDataInner) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetClients200ResponseDataInner) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *GetClients200ResponseDataInner) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *GetClients200ResponseDataInner) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *GetClients200ResponseDataInner) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *GetClients200ResponseDataInner) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

func (o GetClients200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetClients200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.RedirectUri) {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.ThumbnailUrl.IsSet() {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl.Get()
	}
	return toSerialize, nil
}

type NullableGetClients200ResponseDataInner struct {
	value *GetClients200ResponseDataInner
	isSet bool
}

func (v NullableGetClients200ResponseDataInner) Get() *GetClients200ResponseDataInner {
	return v.value
}

func (v *NullableGetClients200ResponseDataInner) Set(val *GetClients200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClients200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClients200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClients200ResponseDataInner(val *GetClients200ResponseDataInner) *NullableGetClients200ResponseDataInner {
	return &NullableGetClients200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetClients200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClients200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


