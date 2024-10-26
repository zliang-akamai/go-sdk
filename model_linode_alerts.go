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

// checks if the LinodeAlerts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinodeAlerts{}

// LinodeAlerts struct for LinodeAlerts
type LinodeAlerts struct {
	// The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. Your Linode's total CPU capacity is represented as 100%, multiplied by its number of cores.  For example, a two core Linode's CPU capacity is represented as 200%. If you want to be alerted at 90% of a two core Linode's CPU capacity, set the alert value to `180`.  The default value is 90% multiplied by the number of cores.  If the value is set to `0` (zero), the alert is disabled.
	Cpu *int32 `json:"cpu,omitempty"`
	// The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to `0` (zero), this alert is disabled.
	Io *int32 `json:"io,omitempty"`
	// The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to `0` (zero), the alert is disabled.
	NetworkIn *int32 `json:"network_in,omitempty"`
	// The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to `0` (zero), the alert is disabled.
	NetworkOut *int32 `json:"network_out,omitempty"`
	// The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to `0` (zero), the alert is disabled.
	TransferQuota *int32 `json:"transfer_quota,omitempty"`
}

// NewLinodeAlerts instantiates a new LinodeAlerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinodeAlerts() *LinodeAlerts {
	this := LinodeAlerts{}
	return &this
}

// NewLinodeAlertsWithDefaults instantiates a new LinodeAlerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinodeAlertsWithDefaults() *LinodeAlerts {
	this := LinodeAlerts{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *LinodeAlerts) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeAlerts) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *LinodeAlerts) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *LinodeAlerts) SetCpu(v int32) {
	o.Cpu = &v
}

// GetIo returns the Io field value if set, zero value otherwise.
func (o *LinodeAlerts) GetIo() int32 {
	if o == nil || IsNil(o.Io) {
		var ret int32
		return ret
	}
	return *o.Io
}

// GetIoOk returns a tuple with the Io field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeAlerts) GetIoOk() (*int32, bool) {
	if o == nil || IsNil(o.Io) {
		return nil, false
	}
	return o.Io, true
}

// HasIo returns a boolean if a field has been set.
func (o *LinodeAlerts) HasIo() bool {
	if o != nil && !IsNil(o.Io) {
		return true
	}

	return false
}

// SetIo gets a reference to the given int32 and assigns it to the Io field.
func (o *LinodeAlerts) SetIo(v int32) {
	o.Io = &v
}

// GetNetworkIn returns the NetworkIn field value if set, zero value otherwise.
func (o *LinodeAlerts) GetNetworkIn() int32 {
	if o == nil || IsNil(o.NetworkIn) {
		var ret int32
		return ret
	}
	return *o.NetworkIn
}

// GetNetworkInOk returns a tuple with the NetworkIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeAlerts) GetNetworkInOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkIn) {
		return nil, false
	}
	return o.NetworkIn, true
}

// HasNetworkIn returns a boolean if a field has been set.
func (o *LinodeAlerts) HasNetworkIn() bool {
	if o != nil && !IsNil(o.NetworkIn) {
		return true
	}

	return false
}

// SetNetworkIn gets a reference to the given int32 and assigns it to the NetworkIn field.
func (o *LinodeAlerts) SetNetworkIn(v int32) {
	o.NetworkIn = &v
}

// GetNetworkOut returns the NetworkOut field value if set, zero value otherwise.
func (o *LinodeAlerts) GetNetworkOut() int32 {
	if o == nil || IsNil(o.NetworkOut) {
		var ret int32
		return ret
	}
	return *o.NetworkOut
}

// GetNetworkOutOk returns a tuple with the NetworkOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeAlerts) GetNetworkOutOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkOut) {
		return nil, false
	}
	return o.NetworkOut, true
}

// HasNetworkOut returns a boolean if a field has been set.
func (o *LinodeAlerts) HasNetworkOut() bool {
	if o != nil && !IsNil(o.NetworkOut) {
		return true
	}

	return false
}

// SetNetworkOut gets a reference to the given int32 and assigns it to the NetworkOut field.
func (o *LinodeAlerts) SetNetworkOut(v int32) {
	o.NetworkOut = &v
}

// GetTransferQuota returns the TransferQuota field value if set, zero value otherwise.
func (o *LinodeAlerts) GetTransferQuota() int32 {
	if o == nil || IsNil(o.TransferQuota) {
		var ret int32
		return ret
	}
	return *o.TransferQuota
}

// GetTransferQuotaOk returns a tuple with the TransferQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeAlerts) GetTransferQuotaOk() (*int32, bool) {
	if o == nil || IsNil(o.TransferQuota) {
		return nil, false
	}
	return o.TransferQuota, true
}

// HasTransferQuota returns a boolean if a field has been set.
func (o *LinodeAlerts) HasTransferQuota() bool {
	if o != nil && !IsNil(o.TransferQuota) {
		return true
	}

	return false
}

// SetTransferQuota gets a reference to the given int32 and assigns it to the TransferQuota field.
func (o *LinodeAlerts) SetTransferQuota(v int32) {
	o.TransferQuota = &v
}

func (o LinodeAlerts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinodeAlerts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Io) {
		toSerialize["io"] = o.Io
	}
	if !IsNil(o.NetworkIn) {
		toSerialize["network_in"] = o.NetworkIn
	}
	if !IsNil(o.NetworkOut) {
		toSerialize["network_out"] = o.NetworkOut
	}
	if !IsNil(o.TransferQuota) {
		toSerialize["transfer_quota"] = o.TransferQuota
	}
	return toSerialize, nil
}

type NullableLinodeAlerts struct {
	value *LinodeAlerts
	isSet bool
}

func (v NullableLinodeAlerts) Get() *LinodeAlerts {
	return v.value
}

func (v *NullableLinodeAlerts) Set(val *LinodeAlerts) {
	v.value = val
	v.isSet = true
}

func (v NullableLinodeAlerts) IsSet() bool {
	return v.isSet
}

func (v *NullableLinodeAlerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinodeAlerts(val *LinodeAlerts) *NullableLinodeAlerts {
	return &NullableLinodeAlerts{value: val, isSet: true}
}

func (v NullableLinodeAlerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinodeAlerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


