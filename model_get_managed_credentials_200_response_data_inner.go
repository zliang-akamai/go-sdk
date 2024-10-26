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

// checks if the GetManagedCredentials200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManagedCredentials200ResponseDataInner{}

// GetManagedCredentials200ResponseDataInner A securely stored Credential that allows Linode's special forces to access a Managed server to respond to Issues.
type GetManagedCredentials200ResponseDataInner struct {
	// This Credential's unique ID.
	Id *int32 `json:"id,omitempty"`
	// The unique label for this Credential. This is for display purposes only.
	Label *string `json:"label,omitempty" validate:"regexp=[a-zA-Z0-9-_ \\\\.]{2,75}"`
	// The date this Credential was last decrypted by a member of Linode special forces.
	LastDecrypted *time.Time `json:"last_decrypted,omitempty"`
}

// NewGetManagedCredentials200ResponseDataInner instantiates a new GetManagedCredentials200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagedCredentials200ResponseDataInner() *GetManagedCredentials200ResponseDataInner {
	this := GetManagedCredentials200ResponseDataInner{}
	return &this
}

// NewGetManagedCredentials200ResponseDataInnerWithDefaults instantiates a new GetManagedCredentials200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagedCredentials200ResponseDataInnerWithDefaults() *GetManagedCredentials200ResponseDataInner {
	this := GetManagedCredentials200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetManagedCredentials200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedCredentials200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetManagedCredentials200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetManagedCredentials200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetManagedCredentials200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedCredentials200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetManagedCredentials200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetManagedCredentials200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetLastDecrypted returns the LastDecrypted field value if set, zero value otherwise.
func (o *GetManagedCredentials200ResponseDataInner) GetLastDecrypted() time.Time {
	if o == nil || IsNil(o.LastDecrypted) {
		var ret time.Time
		return ret
	}
	return *o.LastDecrypted
}

// GetLastDecryptedOk returns a tuple with the LastDecrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedCredentials200ResponseDataInner) GetLastDecryptedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastDecrypted) {
		return nil, false
	}
	return o.LastDecrypted, true
}

// HasLastDecrypted returns a boolean if a field has been set.
func (o *GetManagedCredentials200ResponseDataInner) HasLastDecrypted() bool {
	if o != nil && !IsNil(o.LastDecrypted) {
		return true
	}

	return false
}

// SetLastDecrypted gets a reference to the given time.Time and assigns it to the LastDecrypted field.
func (o *GetManagedCredentials200ResponseDataInner) SetLastDecrypted(v time.Time) {
	o.LastDecrypted = &v
}

func (o GetManagedCredentials200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManagedCredentials200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.LastDecrypted) {
		toSerialize["last_decrypted"] = o.LastDecrypted
	}
	return toSerialize, nil
}

type NullableGetManagedCredentials200ResponseDataInner struct {
	value *GetManagedCredentials200ResponseDataInner
	isSet bool
}

func (v NullableGetManagedCredentials200ResponseDataInner) Get() *GetManagedCredentials200ResponseDataInner {
	return v.value
}

func (v *NullableGetManagedCredentials200ResponseDataInner) Set(val *GetManagedCredentials200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagedCredentials200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagedCredentials200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagedCredentials200ResponseDataInner(val *GetManagedCredentials200ResponseDataInner) *NullableGetManagedCredentials200ResponseDataInner {
	return &NullableGetManagedCredentials200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetManagedCredentials200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagedCredentials200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

