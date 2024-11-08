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

// checks if the GetManagedIssues200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManagedIssues200ResponseDataInner{}

// GetManagedIssues200ResponseDataInner An Issue that was detected with a service Linode is managing.
type GetManagedIssues200ResponseDataInner struct {
	// When this Issue was created. Issues are created in response to issues detected with Managed Services, so this is also when the Issue was detected.
	Created *time.Time `json:"created,omitempty"`
	Entity *GetManagedIssues200ResponseDataInnerEntity `json:"entity,omitempty"`
	// This Issue's unique ID.
	Id *int32 `json:"id,omitempty"`
	// An array of Managed Service IDs that were affected by this Issue.
	Services []int32 `json:"services,omitempty"`
}

// NewGetManagedIssues200ResponseDataInner instantiates a new GetManagedIssues200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagedIssues200ResponseDataInner() *GetManagedIssues200ResponseDataInner {
	this := GetManagedIssues200ResponseDataInner{}
	return &this
}

// NewGetManagedIssues200ResponseDataInnerWithDefaults instantiates a new GetManagedIssues200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagedIssues200ResponseDataInnerWithDefaults() *GetManagedIssues200ResponseDataInner {
	this := GetManagedIssues200ResponseDataInner{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetManagedIssues200ResponseDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIssues200ResponseDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetManagedIssues200ResponseDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetManagedIssues200ResponseDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *GetManagedIssues200ResponseDataInner) GetEntity() GetManagedIssues200ResponseDataInnerEntity {
	if o == nil || IsNil(o.Entity) {
		var ret GetManagedIssues200ResponseDataInnerEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIssues200ResponseDataInner) GetEntityOk() (*GetManagedIssues200ResponseDataInnerEntity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *GetManagedIssues200ResponseDataInner) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given GetManagedIssues200ResponseDataInnerEntity and assigns it to the Entity field.
func (o *GetManagedIssues200ResponseDataInner) SetEntity(v GetManagedIssues200ResponseDataInnerEntity) {
	o.Entity = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetManagedIssues200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIssues200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetManagedIssues200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetManagedIssues200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *GetManagedIssues200ResponseDataInner) GetServices() []int32 {
	if o == nil || IsNil(o.Services) {
		var ret []int32
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedIssues200ResponseDataInner) GetServicesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *GetManagedIssues200ResponseDataInner) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []int32 and assigns it to the Services field.
func (o *GetManagedIssues200ResponseDataInner) SetServices(v []int32) {
	o.Services = v
}

func (o GetManagedIssues200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManagedIssues200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	return toSerialize, nil
}

type NullableGetManagedIssues200ResponseDataInner struct {
	value *GetManagedIssues200ResponseDataInner
	isSet bool
}

func (v NullableGetManagedIssues200ResponseDataInner) Get() *GetManagedIssues200ResponseDataInner {
	return v.value
}

func (v *NullableGetManagedIssues200ResponseDataInner) Set(val *GetManagedIssues200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagedIssues200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagedIssues200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagedIssues200ResponseDataInner(val *GetManagedIssues200ResponseDataInner) *NullableGetManagedIssues200ResponseDataInner {
	return &NullableGetManagedIssues200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetManagedIssues200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagedIssues200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


