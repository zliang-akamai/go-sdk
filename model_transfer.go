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

// checks if the Transfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transfer{}

// Transfer An object representing your network utilization for the current month, in Gigabytes.  Certain Regions have separate utilization quotas and rates. For Region-specific network utilization data, see `region_transfers`.
type Transfer struct {
	// The amount of your transfer pool that is billable this billing cycle.
	Billable *int32 `json:"billable,omitempty"`
	// The amount of network usage allowed this billing cycle.
	Quota *int32 `json:"quota,omitempty"`
	RegionTransfers []GetTransfer200ResponseRegionTransfersInner `json:"region_transfers,omitempty"`
	// The amount of network usage you have used this billing cycle.
	Used *int32 `json:"used,omitempty"`
}

// NewTransfer instantiates a new Transfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfer() *Transfer {
	this := Transfer{}
	return &this
}

// NewTransferWithDefaults instantiates a new Transfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferWithDefaults() *Transfer {
	this := Transfer{}
	return &this
}

// GetBillable returns the Billable field value if set, zero value otherwise.
func (o *Transfer) GetBillable() int32 {
	if o == nil || IsNil(o.Billable) {
		var ret int32
		return ret
	}
	return *o.Billable
}

// GetBillableOk returns a tuple with the Billable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetBillableOk() (*int32, bool) {
	if o == nil || IsNil(o.Billable) {
		return nil, false
	}
	return o.Billable, true
}

// HasBillable returns a boolean if a field has been set.
func (o *Transfer) HasBillable() bool {
	if o != nil && !IsNil(o.Billable) {
		return true
	}

	return false
}

// SetBillable gets a reference to the given int32 and assigns it to the Billable field.
func (o *Transfer) SetBillable(v int32) {
	o.Billable = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *Transfer) GetQuota() int32 {
	if o == nil || IsNil(o.Quota) {
		var ret int32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetQuotaOk() (*int32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *Transfer) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int32 and assigns it to the Quota field.
func (o *Transfer) SetQuota(v int32) {
	o.Quota = &v
}

// GetRegionTransfers returns the RegionTransfers field value if set, zero value otherwise.
func (o *Transfer) GetRegionTransfers() []GetTransfer200ResponseRegionTransfersInner {
	if o == nil || IsNil(o.RegionTransfers) {
		var ret []GetTransfer200ResponseRegionTransfersInner
		return ret
	}
	return o.RegionTransfers
}

// GetRegionTransfersOk returns a tuple with the RegionTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetRegionTransfersOk() ([]GetTransfer200ResponseRegionTransfersInner, bool) {
	if o == nil || IsNil(o.RegionTransfers) {
		return nil, false
	}
	return o.RegionTransfers, true
}

// HasRegionTransfers returns a boolean if a field has been set.
func (o *Transfer) HasRegionTransfers() bool {
	if o != nil && !IsNil(o.RegionTransfers) {
		return true
	}

	return false
}

// SetRegionTransfers gets a reference to the given []GetTransfer200ResponseRegionTransfersInner and assigns it to the RegionTransfers field.
func (o *Transfer) SetRegionTransfers(v []GetTransfer200ResponseRegionTransfersInner) {
	o.RegionTransfers = v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *Transfer) GetUsed() int32 {
	if o == nil || IsNil(o.Used) {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *Transfer) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *Transfer) SetUsed(v int32) {
	o.Used = &v
}

func (o Transfer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Billable) {
		toSerialize["billable"] = o.Billable
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.RegionTransfers) {
		toSerialize["region_transfers"] = o.RegionTransfers
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return toSerialize, nil
}

type NullableTransfer struct {
	value *Transfer
	isSet bool
}

func (v NullableTransfer) Get() *Transfer {
	return v.value
}

func (v *NullableTransfer) Set(val *Transfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfer(val *Transfer) *NullableTransfer {
	return &NullableTransfer{value: val, isSet: true}
}

func (v NullableTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


