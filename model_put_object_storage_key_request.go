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

// checks if the PutObjectStorageKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutObjectStorageKeyRequest{}

// PutObjectStorageKeyRequest struct for PutObjectStorageKeyRequest
type PutObjectStorageKeyRequest struct {
	// The label for the Object Storage key. Used for display purposes. Omit this to leave the `label` unchanged.
	Label *string `json:"label,omitempty"`
	// Replace the current list of `regions` set in a specific key. Include an existing region to maintain it or leave it out to remove it. If you include new `regions` in the key, they can't be used to manage content in buckets in that specific region. You can grant these keys this access using [bucket policies](https://www.linode.com/docs/products/storage/object-storage/guides/bucket-policies/). Omit this to leave the existing list unchanged.  > 🚧 > > You can't remove a `region` from a limited key's original `bucket_access` list. If you include the `regions` array in this operation, you need to include all existing `region` entries from the `bucket_access` array. Otherwise, the operation fails with an error. Run [Get an Object Storage key](https://techdocs.akamai.com/linode-api/reference/get-object-storage-key) to review current `region` entries in a limited key.
	Regions []string `json:"regions,omitempty"`
}

// NewPutObjectStorageKeyRequest instantiates a new PutObjectStorageKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutObjectStorageKeyRequest() *PutObjectStorageKeyRequest {
	this := PutObjectStorageKeyRequest{}
	return &this
}

// NewPutObjectStorageKeyRequestWithDefaults instantiates a new PutObjectStorageKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutObjectStorageKeyRequestWithDefaults() *PutObjectStorageKeyRequest {
	this := PutObjectStorageKeyRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PutObjectStorageKeyRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutObjectStorageKeyRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PutObjectStorageKeyRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PutObjectStorageKeyRequest) SetLabel(v string) {
	o.Label = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *PutObjectStorageKeyRequest) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutObjectStorageKeyRequest) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *PutObjectStorageKeyRequest) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *PutObjectStorageKeyRequest) SetRegions(v []string) {
	o.Regions = v
}

func (o PutObjectStorageKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutObjectStorageKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	return toSerialize, nil
}

type NullablePutObjectStorageKeyRequest struct {
	value *PutObjectStorageKeyRequest
	isSet bool
}

func (v NullablePutObjectStorageKeyRequest) Get() *PutObjectStorageKeyRequest {
	return v.value
}

func (v *NullablePutObjectStorageKeyRequest) Set(val *PutObjectStorageKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutObjectStorageKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutObjectStorageKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutObjectStorageKeyRequest(val *PutObjectStorageKeyRequest) *NullablePutObjectStorageKeyRequest {
	return &NullablePutObjectStorageKeyRequest{value: val, isSet: true}
}

func (v NullablePutObjectStorageKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutObjectStorageKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


