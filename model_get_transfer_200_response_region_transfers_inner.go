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

// checks if the GetTransfer200ResponseRegionTransfersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransfer200ResponseRegionTransfersInner{}

// GetTransfer200ResponseRegionTransfersInner struct for GetTransfer200ResponseRegionTransfersInner
type GetTransfer200ResponseRegionTransfersInner struct {
	// The amount of your transfer pool that is billable this billing cycle for this Region.
	Billable *int32 `json:"billable,omitempty"`
	// The Region ID for this network utilization data.
	Id *string `json:"id,omitempty"`
	// The amount of network usage allowed this billing cycle for this Region.
	Quota *int32 `json:"quota,omitempty"`
	// The amount of network usage you have used this billing cycle for this Region.
	Used *int32 `json:"used,omitempty"`
}

// NewGetTransfer200ResponseRegionTransfersInner instantiates a new GetTransfer200ResponseRegionTransfersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransfer200ResponseRegionTransfersInner() *GetTransfer200ResponseRegionTransfersInner {
	this := GetTransfer200ResponseRegionTransfersInner{}
	return &this
}

// NewGetTransfer200ResponseRegionTransfersInnerWithDefaults instantiates a new GetTransfer200ResponseRegionTransfersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransfer200ResponseRegionTransfersInnerWithDefaults() *GetTransfer200ResponseRegionTransfersInner {
	this := GetTransfer200ResponseRegionTransfersInner{}
	return &this
}

// GetBillable returns the Billable field value if set, zero value otherwise.
func (o *GetTransfer200ResponseRegionTransfersInner) GetBillable() int32 {
	if o == nil || IsNil(o.Billable) {
		var ret int32
		return ret
	}
	return *o.Billable
}

// GetBillableOk returns a tuple with the Billable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransfer200ResponseRegionTransfersInner) GetBillableOk() (*int32, bool) {
	if o == nil || IsNil(o.Billable) {
		return nil, false
	}
	return o.Billable, true
}

// HasBillable returns a boolean if a field has been set.
func (o *GetTransfer200ResponseRegionTransfersInner) HasBillable() bool {
	if o != nil && !IsNil(o.Billable) {
		return true
	}

	return false
}

// SetBillable gets a reference to the given int32 and assigns it to the Billable field.
func (o *GetTransfer200ResponseRegionTransfersInner) SetBillable(v int32) {
	o.Billable = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetTransfer200ResponseRegionTransfersInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransfer200ResponseRegionTransfersInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetTransfer200ResponseRegionTransfersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetTransfer200ResponseRegionTransfersInner) SetId(v string) {
	o.Id = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *GetTransfer200ResponseRegionTransfersInner) GetQuota() int32 {
	if o == nil || IsNil(o.Quota) {
		var ret int32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransfer200ResponseRegionTransfersInner) GetQuotaOk() (*int32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *GetTransfer200ResponseRegionTransfersInner) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int32 and assigns it to the Quota field.
func (o *GetTransfer200ResponseRegionTransfersInner) SetQuota(v int32) {
	o.Quota = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *GetTransfer200ResponseRegionTransfersInner) GetUsed() int32 {
	if o == nil || IsNil(o.Used) {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransfer200ResponseRegionTransfersInner) GetUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *GetTransfer200ResponseRegionTransfersInner) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *GetTransfer200ResponseRegionTransfersInner) SetUsed(v int32) {
	o.Used = &v
}

func (o GetTransfer200ResponseRegionTransfersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransfer200ResponseRegionTransfersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Billable) {
		toSerialize["billable"] = o.Billable
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return toSerialize, nil
}

type NullableGetTransfer200ResponseRegionTransfersInner struct {
	value *GetTransfer200ResponseRegionTransfersInner
	isSet bool
}

func (v NullableGetTransfer200ResponseRegionTransfersInner) Get() *GetTransfer200ResponseRegionTransfersInner {
	return v.value
}

func (v *NullableGetTransfer200ResponseRegionTransfersInner) Set(val *GetTransfer200ResponseRegionTransfersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransfer200ResponseRegionTransfersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransfer200ResponseRegionTransfersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransfer200ResponseRegionTransfersInner(val *GetTransfer200ResponseRegionTransfersInner) *NullableGetTransfer200ResponseRegionTransfersInner {
	return &NullableGetTransfer200ResponseRegionTransfersInner{value: val, isSet: true}
}

func (v NullableGetTransfer200ResponseRegionTransfersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransfer200ResponseRegionTransfersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


