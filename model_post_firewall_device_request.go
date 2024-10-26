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

// checks if the PostFirewallDeviceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostFirewallDeviceRequest{}

// PostFirewallDeviceRequest struct for PostFirewallDeviceRequest
type PostFirewallDeviceRequest struct {
	// The entity's ID.
	Id int32 `json:"id"`
	// The entity's label.
	Label *string `json:"label,omitempty"`
	// The entity's type.
	Type string `json:"type"`
	// The API URL path you can use to access this entity.
	Url *string `json:"url,omitempty"`
}

type _PostFirewallDeviceRequest PostFirewallDeviceRequest

// NewPostFirewallDeviceRequest instantiates a new PostFirewallDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostFirewallDeviceRequest(id int32, type_ string) *PostFirewallDeviceRequest {
	this := PostFirewallDeviceRequest{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewPostFirewallDeviceRequestWithDefaults instantiates a new PostFirewallDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostFirewallDeviceRequestWithDefaults() *PostFirewallDeviceRequest {
	this := PostFirewallDeviceRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PostFirewallDeviceRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostFirewallDeviceRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostFirewallDeviceRequest) SetId(v int32) {
	o.Id = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PostFirewallDeviceRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFirewallDeviceRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PostFirewallDeviceRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PostFirewallDeviceRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *PostFirewallDeviceRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostFirewallDeviceRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostFirewallDeviceRequest) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PostFirewallDeviceRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFirewallDeviceRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PostFirewallDeviceRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PostFirewallDeviceRequest) SetUrl(v string) {
	o.Url = &v
}

func (o PostFirewallDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostFirewallDeviceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

func (o *PostFirewallDeviceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPostFirewallDeviceRequest := _PostFirewallDeviceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostFirewallDeviceRequest)

	if err != nil {
		return err
	}

	*o = PostFirewallDeviceRequest(varPostFirewallDeviceRequest)

	return err
}

type NullablePostFirewallDeviceRequest struct {
	value *PostFirewallDeviceRequest
	isSet bool
}

func (v NullablePostFirewallDeviceRequest) Get() *PostFirewallDeviceRequest {
	return v.value
}

func (v *NullablePostFirewallDeviceRequest) Set(val *PostFirewallDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostFirewallDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostFirewallDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostFirewallDeviceRequest(val *PostFirewallDeviceRequest) *NullablePostFirewallDeviceRequest {
	return &NullablePostFirewallDeviceRequest{value: val, isSet: true}
}

func (v NullablePostFirewallDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostFirewallDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

