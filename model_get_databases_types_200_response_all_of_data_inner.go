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

// checks if the GetDatabasesTypes200ResponseAllOfDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatabasesTypes200ResponseAllOfDataInner{}

// GetDatabasesTypes200ResponseAllOfDataInner Managed Database plan type object.
type GetDatabasesTypes200ResponseAllOfDataInner struct {
	// The compute class category.
	Class *string `json:"class,omitempty"`
	// Whether this Database plan type has been deprecated and is no longer available.
	Deprecated *bool `json:"deprecated,omitempty"`
	// The amount of disk space set aside for Databases of this plan type. The value is represented in megabytes.
	Disk *int32 `json:"disk,omitempty"`
	Engines *GetDatabasesTypes200ResponseAllOfDataInnerEngines `json:"engines,omitempty"`
	// The ID representing the Managed Database node plan type.
	Id *string `json:"id,omitempty"`
	// A human-readable string that describes each plan type. For display purposes only.
	Label *string `json:"label,omitempty"`
	// The amount of RAM allocated to Database created of this plan type. The value is represented in megabytes.
	Memory *int32 `json:"memory,omitempty"`
	// The number of CPUs allocated to databases of this plan type.
	Vcpus *int32 `json:"vcpus,omitempty"`
}

// NewGetDatabasesTypes200ResponseAllOfDataInner instantiates a new GetDatabasesTypes200ResponseAllOfDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatabasesTypes200ResponseAllOfDataInner() *GetDatabasesTypes200ResponseAllOfDataInner {
	this := GetDatabasesTypes200ResponseAllOfDataInner{}
	return &this
}

// NewGetDatabasesTypes200ResponseAllOfDataInnerWithDefaults instantiates a new GetDatabasesTypes200ResponseAllOfDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatabasesTypes200ResponseAllOfDataInnerWithDefaults() *GetDatabasesTypes200ResponseAllOfDataInner {
	this := GetDatabasesTypes200ResponseAllOfDataInner{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetClass(v string) {
	o.Class = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetDisk() int32 {
	if o == nil || IsNil(o.Disk) {
		var ret int32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetDiskOk() (*int32, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given int32 and assigns it to the Disk field.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetDisk(v int32) {
	o.Disk = &v
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetEngines() GetDatabasesTypes200ResponseAllOfDataInnerEngines {
	if o == nil || IsNil(o.Engines) {
		var ret GetDatabasesTypes200ResponseAllOfDataInnerEngines
		return ret
	}
	return *o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetEnginesOk() (*GetDatabasesTypes200ResponseAllOfDataInnerEngines, bool) {
	if o == nil || IsNil(o.Engines) {
		return nil, false
	}
	return o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasEngines() bool {
	if o != nil && !IsNil(o.Engines) {
		return true
	}

	return false
}

// SetEngines gets a reference to the given GetDatabasesTypes200ResponseAllOfDataInnerEngines and assigns it to the Engines field.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetEngines(v GetDatabasesTypes200ResponseAllOfDataInnerEngines) {
	o.Engines = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetMemory(v int32) {
	o.Memory = &v
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetVcpus() int32 {
	if o == nil || IsNil(o.Vcpus) {
		var ret int32
		return ret
	}
	return *o.Vcpus
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetVcpusOk() (*int32, bool) {
	if o == nil || IsNil(o.Vcpus) {
		return nil, false
	}
	return o.Vcpus, true
}

// HasVcpus returns a boolean if a field has been set.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasVcpus() bool {
	if o != nil && !IsNil(o.Vcpus) {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given int32 and assigns it to the Vcpus field.
func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetVcpus(v int32) {
	o.Vcpus = &v
}

func (o GetDatabasesTypes200ResponseAllOfDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatabasesTypes200ResponseAllOfDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.Engines) {
		toSerialize["engines"] = o.Engines
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Vcpus) {
		toSerialize["vcpus"] = o.Vcpus
	}
	return toSerialize, nil
}

type NullableGetDatabasesTypes200ResponseAllOfDataInner struct {
	value *GetDatabasesTypes200ResponseAllOfDataInner
	isSet bool
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInner) Get() *GetDatabasesTypes200ResponseAllOfDataInner {
	return v.value
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInner) Set(val *GetDatabasesTypes200ResponseAllOfDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatabasesTypes200ResponseAllOfDataInner(val *GetDatabasesTypes200ResponseAllOfDataInner) *NullableGetDatabasesTypes200ResponseAllOfDataInner {
	return &NullableGetDatabasesTypes200ResponseAllOfDataInner{value: val, isSet: true}
}

func (v NullableGetDatabasesTypes200ResponseAllOfDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatabasesTypes200ResponseAllOfDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


