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
	"bytes"
	"fmt"
)

// checks if the PostResizeLinodeInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostResizeLinodeInstanceRequest{}

// PostResizeLinodeInstanceRequest struct for PostResizeLinodeInstanceRequest
type PostResizeLinodeInstanceRequest struct {
	// Automatically resize disks when resizing a Linode. When resizing down to a smaller plan your Linode's data must fit within the smaller disk size.
	AllowAutoDiskResize *bool `json:"allow_auto_disk_resize,omitempty"`
	// Type of migration used in moving to a new host or Linode type.  `warm`: the Linode will not power down until the migration is complete. Warm migrations are not available for DC migrations.  `cold`: the Linode will be powered down and migrated. When the migration is complete, the Linode will be powered on.
	MigrationType *string `json:"migration_type,omitempty"`
	// The ID representing the Linode Type.
	Type string `json:"type"`
}

type _PostResizeLinodeInstanceRequest PostResizeLinodeInstanceRequest

// NewPostResizeLinodeInstanceRequest instantiates a new PostResizeLinodeInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostResizeLinodeInstanceRequest(type_ string) *PostResizeLinodeInstanceRequest {
	this := PostResizeLinodeInstanceRequest{}
	var allowAutoDiskResize bool = true
	this.AllowAutoDiskResize = &allowAutoDiskResize
	var migrationType string = "cold"
	this.MigrationType = &migrationType
	this.Type = type_
	return &this
}

// NewPostResizeLinodeInstanceRequestWithDefaults instantiates a new PostResizeLinodeInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostResizeLinodeInstanceRequestWithDefaults() *PostResizeLinodeInstanceRequest {
	this := PostResizeLinodeInstanceRequest{}
	var allowAutoDiskResize bool = true
	this.AllowAutoDiskResize = &allowAutoDiskResize
	var migrationType string = "cold"
	this.MigrationType = &migrationType
	return &this
}

// GetAllowAutoDiskResize returns the AllowAutoDiskResize field value if set, zero value otherwise.
func (o *PostResizeLinodeInstanceRequest) GetAllowAutoDiskResize() bool {
	if o == nil || IsNil(o.AllowAutoDiskResize) {
		var ret bool
		return ret
	}
	return *o.AllowAutoDiskResize
}

// GetAllowAutoDiskResizeOk returns a tuple with the AllowAutoDiskResize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostResizeLinodeInstanceRequest) GetAllowAutoDiskResizeOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAutoDiskResize) {
		return nil, false
	}
	return o.AllowAutoDiskResize, true
}

// HasAllowAutoDiskResize returns a boolean if a field has been set.
func (o *PostResizeLinodeInstanceRequest) HasAllowAutoDiskResize() bool {
	if o != nil && !IsNil(o.AllowAutoDiskResize) {
		return true
	}

	return false
}

// SetAllowAutoDiskResize gets a reference to the given bool and assigns it to the AllowAutoDiskResize field.
func (o *PostResizeLinodeInstanceRequest) SetAllowAutoDiskResize(v bool) {
	o.AllowAutoDiskResize = &v
}

// GetMigrationType returns the MigrationType field value if set, zero value otherwise.
func (o *PostResizeLinodeInstanceRequest) GetMigrationType() string {
	if o == nil || IsNil(o.MigrationType) {
		var ret string
		return ret
	}
	return *o.MigrationType
}

// GetMigrationTypeOk returns a tuple with the MigrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostResizeLinodeInstanceRequest) GetMigrationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MigrationType) {
		return nil, false
	}
	return o.MigrationType, true
}

// HasMigrationType returns a boolean if a field has been set.
func (o *PostResizeLinodeInstanceRequest) HasMigrationType() bool {
	if o != nil && !IsNil(o.MigrationType) {
		return true
	}

	return false
}

// SetMigrationType gets a reference to the given string and assigns it to the MigrationType field.
func (o *PostResizeLinodeInstanceRequest) SetMigrationType(v string) {
	o.MigrationType = &v
}

// GetType returns the Type field value
func (o *PostResizeLinodeInstanceRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostResizeLinodeInstanceRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostResizeLinodeInstanceRequest) SetType(v string) {
	o.Type = v
}

func (o PostResizeLinodeInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostResizeLinodeInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowAutoDiskResize) {
		toSerialize["allow_auto_disk_resize"] = o.AllowAutoDiskResize
	}
	if !IsNil(o.MigrationType) {
		toSerialize["migration_type"] = o.MigrationType
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *PostResizeLinodeInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varPostResizeLinodeInstanceRequest := _PostResizeLinodeInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostResizeLinodeInstanceRequest)

	if err != nil {
		return err
	}

	*o = PostResizeLinodeInstanceRequest(varPostResizeLinodeInstanceRequest)

	return err
}

type NullablePostResizeLinodeInstanceRequest struct {
	value *PostResizeLinodeInstanceRequest
	isSet bool
}

func (v NullablePostResizeLinodeInstanceRequest) Get() *PostResizeLinodeInstanceRequest {
	return v.value
}

func (v *NullablePostResizeLinodeInstanceRequest) Set(val *PostResizeLinodeInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostResizeLinodeInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostResizeLinodeInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostResizeLinodeInstanceRequest(val *PostResizeLinodeInstanceRequest) *NullablePostResizeLinodeInstanceRequest {
	return &NullablePostResizeLinodeInstanceRequest{value: val, isSet: true}
}

func (v NullablePostResizeLinodeInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostResizeLinodeInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


