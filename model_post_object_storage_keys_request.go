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

// checks if the PostObjectStorageKeysRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostObjectStorageKeysRequest{}

// PostObjectStorageKeysRequest struct for PostObjectStorageKeysRequest
type PostObjectStorageKeysRequest struct {
	// Set up the key to limit access to specific buckets, each with a specific permission level. You can create a limited Object Storage key with access to no buckets. Include an empty `bucket_access` array in the request.
	BucketAccess []PostObjectStorageKeysRequestBucketAccessInner `json:"bucket_access,omitempty"`
	// The label for the Object Storage key, for display purposes only.
	Label *string `json:"label,omitempty"`
	// You can use a key to create new buckets in regions set in this array. But it can't be used to manage content in those buckets. See [Create an Object Storage key](https://techdocs.akamai.com/linode-api/reference/post-object-storage-keys) for more details.
	Regions []string `json:"regions,omitempty"`
}

// NewPostObjectStorageKeysRequest instantiates a new PostObjectStorageKeysRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostObjectStorageKeysRequest() *PostObjectStorageKeysRequest {
	this := PostObjectStorageKeysRequest{}
	return &this
}

// NewPostObjectStorageKeysRequestWithDefaults instantiates a new PostObjectStorageKeysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostObjectStorageKeysRequestWithDefaults() *PostObjectStorageKeysRequest {
	this := PostObjectStorageKeysRequest{}
	return &this
}

// GetBucketAccess returns the BucketAccess field value if set, zero value otherwise.
func (o *PostObjectStorageKeysRequest) GetBucketAccess() []PostObjectStorageKeysRequestBucketAccessInner {
	if o == nil || IsNil(o.BucketAccess) {
		var ret []PostObjectStorageKeysRequestBucketAccessInner
		return ret
	}
	return o.BucketAccess
}

// GetBucketAccessOk returns a tuple with the BucketAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostObjectStorageKeysRequest) GetBucketAccessOk() ([]PostObjectStorageKeysRequestBucketAccessInner, bool) {
	if o == nil || IsNil(o.BucketAccess) {
		return nil, false
	}
	return o.BucketAccess, true
}

// HasBucketAccess returns a boolean if a field has been set.
func (o *PostObjectStorageKeysRequest) HasBucketAccess() bool {
	if o != nil && !IsNil(o.BucketAccess) {
		return true
	}

	return false
}

// SetBucketAccess gets a reference to the given []PostObjectStorageKeysRequestBucketAccessInner and assigns it to the BucketAccess field.
func (o *PostObjectStorageKeysRequest) SetBucketAccess(v []PostObjectStorageKeysRequestBucketAccessInner) {
	o.BucketAccess = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PostObjectStorageKeysRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostObjectStorageKeysRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PostObjectStorageKeysRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PostObjectStorageKeysRequest) SetLabel(v string) {
	o.Label = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *PostObjectStorageKeysRequest) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostObjectStorageKeysRequest) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *PostObjectStorageKeysRequest) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *PostObjectStorageKeysRequest) SetRegions(v []string) {
	o.Regions = v
}

func (o PostObjectStorageKeysRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostObjectStorageKeysRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketAccess) {
		toSerialize["bucket_access"] = o.BucketAccess
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	return toSerialize, nil
}

type NullablePostObjectStorageKeysRequest struct {
	value *PostObjectStorageKeysRequest
	isSet bool
}

func (v NullablePostObjectStorageKeysRequest) Get() *PostObjectStorageKeysRequest {
	return v.value
}

func (v *NullablePostObjectStorageKeysRequest) Set(val *PostObjectStorageKeysRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostObjectStorageKeysRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostObjectStorageKeysRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostObjectStorageKeysRequest(val *PostObjectStorageKeysRequest) *NullablePostObjectStorageKeysRequest {
	return &NullablePostObjectStorageKeysRequest{value: val, isSet: true}
}

func (v NullablePostObjectStorageKeysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostObjectStorageKeysRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


