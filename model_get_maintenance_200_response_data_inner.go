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

// checks if the GetMaintenance200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMaintenance200ResponseDataInner{}

// GetMaintenance200ResponseDataInner Information about maintenance affecting an entity.
type GetMaintenance200ResponseDataInner struct {
	Entity *GetMaintenance200ResponseDataInnerEntity `json:"entity,omitempty"`
	// The reason maintenance is being performed.
	Reason *string `json:"reason,omitempty"`
	// The maintenance status.  Maintenance progresses in the following sequence: pending, started, then completed.
	Status *string `json:"status,omitempty"`
	// The type of maintenance.
	Type *string `json:"type,omitempty"`
	// When the maintenance will begin.  [Filterable](https://techdocs.akamai.com/linode-api/reference/filtering-and-sorting) with the following parameters:  - A single value in date-time string format (`%Y-%m-%dT%H:%M:%S`), which returns only matches to that value.  - A dictionary containing pairs of inequality operator string keys (`+or`, `+gt`, `+gte`, `+lt`, `+lte`, or `+neq`) and single date-time string format values (`%Y-%m-%dT%H:%M:%S`). `+or` accepts an array of values that may consist of single date-time strings or dictionaries of inequality operator pairs.
	When *time.Time `json:"when,omitempty"`
}

// NewGetMaintenance200ResponseDataInner instantiates a new GetMaintenance200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMaintenance200ResponseDataInner() *GetMaintenance200ResponseDataInner {
	this := GetMaintenance200ResponseDataInner{}
	return &this
}

// NewGetMaintenance200ResponseDataInnerWithDefaults instantiates a new GetMaintenance200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMaintenance200ResponseDataInnerWithDefaults() *GetMaintenance200ResponseDataInner {
	this := GetMaintenance200ResponseDataInner{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *GetMaintenance200ResponseDataInner) GetEntity() GetMaintenance200ResponseDataInnerEntity {
	if o == nil || IsNil(o.Entity) {
		var ret GetMaintenance200ResponseDataInnerEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMaintenance200ResponseDataInner) GetEntityOk() (*GetMaintenance200ResponseDataInnerEntity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *GetMaintenance200ResponseDataInner) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given GetMaintenance200ResponseDataInnerEntity and assigns it to the Entity field.
func (o *GetMaintenance200ResponseDataInner) SetEntity(v GetMaintenance200ResponseDataInnerEntity) {
	o.Entity = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *GetMaintenance200ResponseDataInner) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMaintenance200ResponseDataInner) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GetMaintenance200ResponseDataInner) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *GetMaintenance200ResponseDataInner) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetMaintenance200ResponseDataInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMaintenance200ResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetMaintenance200ResponseDataInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetMaintenance200ResponseDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetMaintenance200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMaintenance200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetMaintenance200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetMaintenance200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *GetMaintenance200ResponseDataInner) GetWhen() time.Time {
	if o == nil || IsNil(o.When) {
		var ret time.Time
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMaintenance200ResponseDataInner) GetWhenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.When) {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *GetMaintenance200ResponseDataInner) HasWhen() bool {
	if o != nil && !IsNil(o.When) {
		return true
	}

	return false
}

// SetWhen gets a reference to the given time.Time and assigns it to the When field.
func (o *GetMaintenance200ResponseDataInner) SetWhen(v time.Time) {
	o.When = &v
}

func (o GetMaintenance200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMaintenance200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.When) {
		toSerialize["when"] = o.When
	}
	return toSerialize, nil
}

type NullableGetMaintenance200ResponseDataInner struct {
	value *GetMaintenance200ResponseDataInner
	isSet bool
}

func (v NullableGetMaintenance200ResponseDataInner) Get() *GetMaintenance200ResponseDataInner {
	return v.value
}

func (v *NullableGetMaintenance200ResponseDataInner) Set(val *GetMaintenance200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMaintenance200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMaintenance200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMaintenance200ResponseDataInner(val *GetMaintenance200ResponseDataInner) *NullableGetMaintenance200ResponseDataInner {
	return &NullableGetMaintenance200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetMaintenance200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMaintenance200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


