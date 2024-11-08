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

// checks if the LinodeBackups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinodeBackups{}

// LinodeBackups Information about this Linode's backups status. For information about available backups, run [List backups](https://techdocs.akamai.com/linode-api/reference/get-backups).
type LinodeBackups struct {
	// Whether Backups for this Linode are available for restoration.  Backups undergoing maintenance are not available for restoration.
	Available *bool `json:"available,omitempty"`
	// If this Linode has the Backup service enabled. To enable backups, run [Enable backups](https://techdocs.akamai.com/linode-api/reference/post-enable-backups).
	Enabled *bool `json:"enabled,omitempty"`
	// The last successful backup time. Displayed as `null` if there was no previous backup.
	LastSuccessful *time.Time `json:"last_successful,omitempty"`
	Schedule *LinodeBackupsSchedule `json:"schedule,omitempty"`
}

// NewLinodeBackups instantiates a new LinodeBackups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinodeBackups() *LinodeBackups {
	this := LinodeBackups{}
	return &this
}

// NewLinodeBackupsWithDefaults instantiates a new LinodeBackups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinodeBackupsWithDefaults() *LinodeBackups {
	this := LinodeBackups{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *LinodeBackups) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeBackups) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *LinodeBackups) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *LinodeBackups) SetAvailable(v bool) {
	o.Available = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LinodeBackups) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeBackups) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LinodeBackups) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LinodeBackups) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLastSuccessful returns the LastSuccessful field value if set, zero value otherwise.
func (o *LinodeBackups) GetLastSuccessful() time.Time {
	if o == nil || IsNil(o.LastSuccessful) {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessful
}

// GetLastSuccessfulOk returns a tuple with the LastSuccessful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeBackups) GetLastSuccessfulOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSuccessful) {
		return nil, false
	}
	return o.LastSuccessful, true
}

// HasLastSuccessful returns a boolean if a field has been set.
func (o *LinodeBackups) HasLastSuccessful() bool {
	if o != nil && !IsNil(o.LastSuccessful) {
		return true
	}

	return false
}

// SetLastSuccessful gets a reference to the given time.Time and assigns it to the LastSuccessful field.
func (o *LinodeBackups) SetLastSuccessful(v time.Time) {
	o.LastSuccessful = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *LinodeBackups) GetSchedule() LinodeBackupsSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret LinodeBackupsSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeBackups) GetScheduleOk() (*LinodeBackupsSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *LinodeBackups) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given LinodeBackupsSchedule and assigns it to the Schedule field.
func (o *LinodeBackups) SetSchedule(v LinodeBackupsSchedule) {
	o.Schedule = &v
}

func (o LinodeBackups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinodeBackups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.LastSuccessful) {
		toSerialize["last_successful"] = o.LastSuccessful
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	return toSerialize, nil
}

type NullableLinodeBackups struct {
	value *LinodeBackups
	isSet bool
}

func (v NullableLinodeBackups) Get() *LinodeBackups {
	return v.value
}

func (v *NullableLinodeBackups) Set(val *LinodeBackups) {
	v.value = val
	v.isSet = true
}

func (v NullableLinodeBackups) IsSet() bool {
	return v.isSet
}

func (v *NullableLinodeBackups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinodeBackups(val *LinodeBackups) *NullableLinodeBackups {
	return &NullableLinodeBackups{value: val, isSet: true}
}

func (v NullableLinodeBackups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinodeBackups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


