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

// checks if the GetUserGrants200ResponseDatabaseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserGrants200ResponseDatabaseInner{}

// GetUserGrants200ResponseDatabaseInner Represents the level of access a restricted User has to a specific resource on the Account.
type GetUserGrants200ResponseDatabaseInner struct {
	// The ID of the entity this grant applies to.
	Id *int32 `json:"id,omitempty"`
	// The current label of the entity this grant applies to, for display purposes.
	Label *string `json:"label,omitempty"`
	// The level of access this User has to this entity.  If null, this User has no access.
	Permissions NullableString `json:"permissions,omitempty"`
}

// NewGetUserGrants200ResponseDatabaseInner instantiates a new GetUserGrants200ResponseDatabaseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserGrants200ResponseDatabaseInner() *GetUserGrants200ResponseDatabaseInner {
	this := GetUserGrants200ResponseDatabaseInner{}
	return &this
}

// NewGetUserGrants200ResponseDatabaseInnerWithDefaults instantiates a new GetUserGrants200ResponseDatabaseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserGrants200ResponseDatabaseInnerWithDefaults() *GetUserGrants200ResponseDatabaseInner {
	this := GetUserGrants200ResponseDatabaseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetUserGrants200ResponseDatabaseInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200ResponseDatabaseInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetUserGrants200ResponseDatabaseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetUserGrants200ResponseDatabaseInner) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetUserGrants200ResponseDatabaseInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200ResponseDatabaseInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetUserGrants200ResponseDatabaseInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetUserGrants200ResponseDatabaseInner) SetLabel(v string) {
	o.Label = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUserGrants200ResponseDatabaseInner) GetPermissions() string {
	if o == nil || IsNil(o.Permissions.Get()) {
		var ret string
		return ret
	}
	return *o.Permissions.Get()
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUserGrants200ResponseDatabaseInner) GetPermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions.Get(), o.Permissions.IsSet()
}

// HasPermissions returns a boolean if a field has been set.
func (o *GetUserGrants200ResponseDatabaseInner) HasPermissions() bool {
	if o != nil && o.Permissions.IsSet() {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given NullableString and assigns it to the Permissions field.
func (o *GetUserGrants200ResponseDatabaseInner) SetPermissions(v string) {
	o.Permissions.Set(&v)
}
// SetPermissionsNil sets the value for Permissions to be an explicit nil
func (o *GetUserGrants200ResponseDatabaseInner) SetPermissionsNil() {
	o.Permissions.Set(nil)
}

// UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
func (o *GetUserGrants200ResponseDatabaseInner) UnsetPermissions() {
	o.Permissions.Unset()
}

func (o GetUserGrants200ResponseDatabaseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserGrants200ResponseDatabaseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Permissions.IsSet() {
		toSerialize["permissions"] = o.Permissions.Get()
	}
	return toSerialize, nil
}

type NullableGetUserGrants200ResponseDatabaseInner struct {
	value *GetUserGrants200ResponseDatabaseInner
	isSet bool
}

func (v NullableGetUserGrants200ResponseDatabaseInner) Get() *GetUserGrants200ResponseDatabaseInner {
	return v.value
}

func (v *NullableGetUserGrants200ResponseDatabaseInner) Set(val *GetUserGrants200ResponseDatabaseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserGrants200ResponseDatabaseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserGrants200ResponseDatabaseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserGrants200ResponseDatabaseInner(val *GetUserGrants200ResponseDatabaseInner) *NullableGetUserGrants200ResponseDatabaseInner {
	return &NullableGetUserGrants200ResponseDatabaseInner{value: val, isSet: true}
}

func (v NullableGetUserGrants200ResponseDatabaseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserGrants200ResponseDatabaseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

