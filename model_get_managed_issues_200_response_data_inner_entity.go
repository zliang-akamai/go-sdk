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

// checks if the GetManagedIssues200ResponseDataInnerEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManagedIssues200ResponseDataInnerEntity{}

// GetManagedIssues200ResponseDataInnerEntity The ticket this Managed Issue opened.
type GetManagedIssues200ResponseDataInnerEntity struct {
	// This ticket's ID.
	Id *int32 `json:"id,omitempty"`
	// The summary for this Ticket.
	Label *string `json:"label,omitempty"`
	// The type of entity this is. In this case, it is always a Ticket.
	Type *string `json:"type,omitempty"`
	// The relative URL where you can access this Ticket.
	Url *string `json:"url,omitempty"`
}

// NewGetManagedIssues200ResponseDataInnerEntity instantiates a new GetManagedIssues200ResponseDataInnerEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagedIssues200ResponseDataInnerEntity() *GetManagedIssues200ResponseDataInnerEntity {
	this := GetManagedIssues200ResponseDataInnerEntity{}
	return &this
}

// NewGetManagedIssues200ResponseDataInnerEntityWithDefaults instantiates a new GetManagedIssues200ResponseDataInnerEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagedIssues200ResponseDataInnerEntityWithDefaults() *GetManagedIssues200ResponseDataInnerEntity {
	this := GetManagedIssues200ResponseDataInnerEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetManagedIssues200ResponseDataInnerEntity) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIssues200ResponseDataInnerEntity) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetManagedIssues200ResponseDataInnerEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetManagedIssues200ResponseDataInnerEntity) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetManagedIssues200ResponseDataInnerEntity) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIssues200ResponseDataInnerEntity) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetManagedIssues200ResponseDataInnerEntity) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetManagedIssues200ResponseDataInnerEntity) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetManagedIssues200ResponseDataInnerEntity) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIssues200ResponseDataInnerEntity) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetManagedIssues200ResponseDataInnerEntity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetManagedIssues200ResponseDataInnerEntity) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetManagedIssues200ResponseDataInnerEntity) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIssues200ResponseDataInnerEntity) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetManagedIssues200ResponseDataInnerEntity) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetManagedIssues200ResponseDataInnerEntity) SetUrl(v string) {
	o.Url = &v
}

func (o GetManagedIssues200ResponseDataInnerEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManagedIssues200ResponseDataInnerEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableGetManagedIssues200ResponseDataInnerEntity struct {
	value *GetManagedIssues200ResponseDataInnerEntity
	isSet bool
}

func (v NullableGetManagedIssues200ResponseDataInnerEntity) Get() *GetManagedIssues200ResponseDataInnerEntity {
	return v.value
}

func (v *NullableGetManagedIssues200ResponseDataInnerEntity) Set(val *GetManagedIssues200ResponseDataInnerEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagedIssues200ResponseDataInnerEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagedIssues200ResponseDataInnerEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagedIssues200ResponseDataInnerEntity(val *GetManagedIssues200ResponseDataInnerEntity) *NullableGetManagedIssues200ResponseDataInnerEntity {
	return &NullableGetManagedIssues200ResponseDataInnerEntity{value: val, isSet: true}
}

func (v NullableGetManagedIssues200ResponseDataInnerEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagedIssues200ResponseDataInnerEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

