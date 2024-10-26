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

// checks if the GetLinodeStats200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLinodeStats200Response{}

// GetLinodeStats200Response CPU, IO, IPv4, and IPv6 statistics. Graph data, if available, is in `[timestamp, reading]` array format. Timestamp is a UNIX timestamp in EST.
type GetLinodeStats200Response struct {
	// Percentage of CPU used.
	Cpu [][]float32 `json:"cpu,omitempty"`
	Io *GetLinodeStats200ResponseIo `json:"io,omitempty"`
	Netv4 *GetLinodeStats200ResponseNetv4 `json:"netv4,omitempty"`
	Netv6 *GetLinodeStats200ResponseNetv6 `json:"netv6,omitempty"`
	// The title for this data set.
	Title *string `json:"title,omitempty"`
}

// NewGetLinodeStats200Response instantiates a new GetLinodeStats200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLinodeStats200Response() *GetLinodeStats200Response {
	this := GetLinodeStats200Response{}
	return &this
}

// NewGetLinodeStats200ResponseWithDefaults instantiates a new GetLinodeStats200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLinodeStats200ResponseWithDefaults() *GetLinodeStats200Response {
	this := GetLinodeStats200Response{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *GetLinodeStats200Response) GetCpu() [][]float32 {
	if o == nil || IsNil(o.Cpu) {
		var ret [][]float32
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeStats200Response) GetCpuOk() ([][]float32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *GetLinodeStats200Response) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given [][]float32 and assigns it to the Cpu field.
func (o *GetLinodeStats200Response) SetCpu(v [][]float32) {
	o.Cpu = v
}

// GetIo returns the Io field value if set, zero value otherwise.
func (o *GetLinodeStats200Response) GetIo() GetLinodeStats200ResponseIo {
	if o == nil || IsNil(o.Io) {
		var ret GetLinodeStats200ResponseIo
		return ret
	}
	return *o.Io
}

// GetIoOk returns a tuple with the Io field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeStats200Response) GetIoOk() (*GetLinodeStats200ResponseIo, bool) {
	if o == nil || IsNil(o.Io) {
		return nil, false
	}
	return o.Io, true
}

// HasIo returns a boolean if a field has been set.
func (o *GetLinodeStats200Response) HasIo() bool {
	if o != nil && !IsNil(o.Io) {
		return true
	}

	return false
}

// SetIo gets a reference to the given GetLinodeStats200ResponseIo and assigns it to the Io field.
func (o *GetLinodeStats200Response) SetIo(v GetLinodeStats200ResponseIo) {
	o.Io = &v
}

// GetNetv4 returns the Netv4 field value if set, zero value otherwise.
func (o *GetLinodeStats200Response) GetNetv4() GetLinodeStats200ResponseNetv4 {
	if o == nil || IsNil(o.Netv4) {
		var ret GetLinodeStats200ResponseNetv4
		return ret
	}
	return *o.Netv4
}

// GetNetv4Ok returns a tuple with the Netv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeStats200Response) GetNetv4Ok() (*GetLinodeStats200ResponseNetv4, bool) {
	if o == nil || IsNil(o.Netv4) {
		return nil, false
	}
	return o.Netv4, true
}

// HasNetv4 returns a boolean if a field has been set.
func (o *GetLinodeStats200Response) HasNetv4() bool {
	if o != nil && !IsNil(o.Netv4) {
		return true
	}

	return false
}

// SetNetv4 gets a reference to the given GetLinodeStats200ResponseNetv4 and assigns it to the Netv4 field.
func (o *GetLinodeStats200Response) SetNetv4(v GetLinodeStats200ResponseNetv4) {
	o.Netv4 = &v
}

// GetNetv6 returns the Netv6 field value if set, zero value otherwise.
func (o *GetLinodeStats200Response) GetNetv6() GetLinodeStats200ResponseNetv6 {
	if o == nil || IsNil(o.Netv6) {
		var ret GetLinodeStats200ResponseNetv6
		return ret
	}
	return *o.Netv6
}

// GetNetv6Ok returns a tuple with the Netv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeStats200Response) GetNetv6Ok() (*GetLinodeStats200ResponseNetv6, bool) {
	if o == nil || IsNil(o.Netv6) {
		return nil, false
	}
	return o.Netv6, true
}

// HasNetv6 returns a boolean if a field has been set.
func (o *GetLinodeStats200Response) HasNetv6() bool {
	if o != nil && !IsNil(o.Netv6) {
		return true
	}

	return false
}

// SetNetv6 gets a reference to the given GetLinodeStats200ResponseNetv6 and assigns it to the Netv6 field.
func (o *GetLinodeStats200Response) SetNetv6(v GetLinodeStats200ResponseNetv6) {
	o.Netv6 = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetLinodeStats200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLinodeStats200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetLinodeStats200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetLinodeStats200Response) SetTitle(v string) {
	o.Title = &v
}

func (o GetLinodeStats200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLinodeStats200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Io) {
		toSerialize["io"] = o.Io
	}
	if !IsNil(o.Netv4) {
		toSerialize["netv4"] = o.Netv4
	}
	if !IsNil(o.Netv6) {
		toSerialize["netv6"] = o.Netv6
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableGetLinodeStats200Response struct {
	value *GetLinodeStats200Response
	isSet bool
}

func (v NullableGetLinodeStats200Response) Get() *GetLinodeStats200Response {
	return v.value
}

func (v *NullableGetLinodeStats200Response) Set(val *GetLinodeStats200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLinodeStats200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLinodeStats200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLinodeStats200Response(val *GetLinodeStats200Response) *NullableGetLinodeStats200Response {
	return &NullableGetLinodeStats200Response{value: val, isSet: true}
}

func (v NullableGetLinodeStats200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLinodeStats200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


