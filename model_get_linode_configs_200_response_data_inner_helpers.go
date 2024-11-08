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

// checks if the GetLinodeConfigs200ResponseDataInnerHelpers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLinodeConfigs200ResponseDataInnerHelpers{}

// GetLinodeConfigs200ResponseDataInnerHelpers Helpers enabled when booting to this Linode Config.
type GetLinodeConfigs200ResponseDataInnerHelpers struct {
	// Populates the /dev directory early during boot without udev.  Defaults to false.
	DevtmpfsAutomount *bool `json:"devtmpfs_automount,omitempty"`
	// Helps maintain correct inittab/upstart console device.
	Distro *bool `json:"distro,omitempty"`
	// Creates a modules dependency file for the Kernel you run.
	ModulesDep *bool `json:"modules_dep,omitempty"`
	// Automatically configures static networking.
	Network *bool `json:"network,omitempty"`
	// Disables updatedb cron job to avoid disk thrashing.
	UpdatedbDisabled *bool `json:"updatedb_disabled,omitempty"`
}

// NewGetLinodeConfigs200ResponseDataInnerHelpers instantiates a new GetLinodeConfigs200ResponseDataInnerHelpers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLinodeConfigs200ResponseDataInnerHelpers() *GetLinodeConfigs200ResponseDataInnerHelpers {
	this := GetLinodeConfigs200ResponseDataInnerHelpers{}
	return &this
}

// NewGetLinodeConfigs200ResponseDataInnerHelpersWithDefaults instantiates a new GetLinodeConfigs200ResponseDataInnerHelpers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLinodeConfigs200ResponseDataInnerHelpersWithDefaults() *GetLinodeConfigs200ResponseDataInnerHelpers {
	this := GetLinodeConfigs200ResponseDataInnerHelpers{}
	return &this
}

// GetDevtmpfsAutomount returns the DevtmpfsAutomount field value if set, zero value otherwise.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetDevtmpfsAutomount() bool {
	if o == nil || IsNil(o.DevtmpfsAutomount) {
		var ret bool
		return ret
	}
	return *o.DevtmpfsAutomount
}

// GetDevtmpfsAutomountOk returns a tuple with the DevtmpfsAutomount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetDevtmpfsAutomountOk() (*bool, bool) {
	if o == nil || IsNil(o.DevtmpfsAutomount) {
		return nil, false
	}
	return o.DevtmpfsAutomount, true
}

// HasDevtmpfsAutomount returns a boolean if a field has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasDevtmpfsAutomount() bool {
	if o != nil && !IsNil(o.DevtmpfsAutomount) {
		return true
	}

	return false
}

// SetDevtmpfsAutomount gets a reference to the given bool and assigns it to the DevtmpfsAutomount field.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetDevtmpfsAutomount(v bool) {
	o.DevtmpfsAutomount = &v
}

// GetDistro returns the Distro field value if set, zero value otherwise.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetDistro() bool {
	if o == nil || IsNil(o.Distro) {
		var ret bool
		return ret
	}
	return *o.Distro
}

// GetDistroOk returns a tuple with the Distro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetDistroOk() (*bool, bool) {
	if o == nil || IsNil(o.Distro) {
		return nil, false
	}
	return o.Distro, true
}

// HasDistro returns a boolean if a field has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasDistro() bool {
	if o != nil && !IsNil(o.Distro) {
		return true
	}

	return false
}

// SetDistro gets a reference to the given bool and assigns it to the Distro field.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetDistro(v bool) {
	o.Distro = &v
}

// GetModulesDep returns the ModulesDep field value if set, zero value otherwise.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetModulesDep() bool {
	if o == nil || IsNil(o.ModulesDep) {
		var ret bool
		return ret
	}
	return *o.ModulesDep
}

// GetModulesDepOk returns a tuple with the ModulesDep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetModulesDepOk() (*bool, bool) {
	if o == nil || IsNil(o.ModulesDep) {
		return nil, false
	}
	return o.ModulesDep, true
}

// HasModulesDep returns a boolean if a field has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasModulesDep() bool {
	if o != nil && !IsNil(o.ModulesDep) {
		return true
	}

	return false
}

// SetModulesDep gets a reference to the given bool and assigns it to the ModulesDep field.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetModulesDep(v bool) {
	o.ModulesDep = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetNetwork() bool {
	if o == nil || IsNil(o.Network) {
		var ret bool
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given bool and assigns it to the Network field.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetNetwork(v bool) {
	o.Network = &v
}

// GetUpdatedbDisabled returns the UpdatedbDisabled field value if set, zero value otherwise.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetUpdatedbDisabled() bool {
	if o == nil || IsNil(o.UpdatedbDisabled) {
		var ret bool
		return ret
	}
	return *o.UpdatedbDisabled
}

// GetUpdatedbDisabledOk returns a tuple with the UpdatedbDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetUpdatedbDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdatedbDisabled) {
		return nil, false
	}
	return o.UpdatedbDisabled, true
}

// HasUpdatedbDisabled returns a boolean if a field has been set.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasUpdatedbDisabled() bool {
	if o != nil && !IsNil(o.UpdatedbDisabled) {
		return true
	}

	return false
}

// SetUpdatedbDisabled gets a reference to the given bool and assigns it to the UpdatedbDisabled field.
func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetUpdatedbDisabled(v bool) {
	o.UpdatedbDisabled = &v
}

func (o GetLinodeConfigs200ResponseDataInnerHelpers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLinodeConfigs200ResponseDataInnerHelpers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DevtmpfsAutomount) {
		toSerialize["devtmpfs_automount"] = o.DevtmpfsAutomount
	}
	if !IsNil(o.Distro) {
		toSerialize["distro"] = o.Distro
	}
	if !IsNil(o.ModulesDep) {
		toSerialize["modules_dep"] = o.ModulesDep
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.UpdatedbDisabled) {
		toSerialize["updatedb_disabled"] = o.UpdatedbDisabled
	}
	return toSerialize, nil
}

type NullableGetLinodeConfigs200ResponseDataInnerHelpers struct {
	value *GetLinodeConfigs200ResponseDataInnerHelpers
	isSet bool
}

func (v NullableGetLinodeConfigs200ResponseDataInnerHelpers) Get() *GetLinodeConfigs200ResponseDataInnerHelpers {
	return v.value
}

func (v *NullableGetLinodeConfigs200ResponseDataInnerHelpers) Set(val *GetLinodeConfigs200ResponseDataInnerHelpers) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLinodeConfigs200ResponseDataInnerHelpers) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLinodeConfigs200ResponseDataInnerHelpers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLinodeConfigs200ResponseDataInnerHelpers(val *GetLinodeConfigs200ResponseDataInnerHelpers) *NullableGetLinodeConfigs200ResponseDataInnerHelpers {
	return &NullableGetLinodeConfigs200ResponseDataInnerHelpers{value: val, isSet: true}
}

func (v NullableGetLinodeConfigs200ResponseDataInnerHelpers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLinodeConfigs200ResponseDataInnerHelpers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


