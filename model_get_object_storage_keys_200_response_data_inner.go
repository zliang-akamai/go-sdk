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

// checks if the GetObjectStorageKeys200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetObjectStorageKeys200ResponseDataInner{}

// GetObjectStorageKeys200ResponseDataInner A key used to communicate with the Object Storage S3 API.
type GetObjectStorageKeys200ResponseDataInner struct {
	// A unique string chosen by the API to identify this key. Used as a user name to identify this key when making requests to the S3 API.
	AccessKey *string `json:"access_key,omitempty"`
	// Settings that limit access to specific buckets, each with a specific permission level.
	BucketAccess []GetObjectStorageKeys200ResponseDataInnerBucketAccessInner `json:"bucket_access,omitempty"`
	// This Object Storage key's unique ID.
	Id *int32 `json:"id,omitempty"`
	// The label given to this key. For display purposes only.
	Label *string `json:"label,omitempty"`
	// Whether this Object Storage key limits access to specific buckets and permissions. Returns `false` if this key grants full access. Specific limitations are set in `bucket_access`.
	Limited *bool `json:"limited,omitempty"`
	// The key can be used in these regions to create new buckets but it can't be used to manage content in those buckets. See [Create an Object Storage key](https://techdocs.akamai.com/linode-api/reference/post-object-storage-keys) for more details.
	Regions []GetObjectStorageKeys200ResponseDataInnerRegionsInner `json:"regions,omitempty"`
	// This Object Storage key's secret key. Used as a password to validate this key when making requests to the S3 API. This value is only revealed in a response after creating or modifying a key.
	SecretKey *string `json:"secret_key,omitempty"`
}

// NewGetObjectStorageKeys200ResponseDataInner instantiates a new GetObjectStorageKeys200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetObjectStorageKeys200ResponseDataInner() *GetObjectStorageKeys200ResponseDataInner {
	this := GetObjectStorageKeys200ResponseDataInner{}
	return &this
}

// NewGetObjectStorageKeys200ResponseDataInnerWithDefaults instantiates a new GetObjectStorageKeys200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetObjectStorageKeys200ResponseDataInnerWithDefaults() *GetObjectStorageKeys200ResponseDataInner {
	this := GetObjectStorageKeys200ResponseDataInner{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInner) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *GetObjectStorageKeys200ResponseDataInner) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetBucketAccess returns the BucketAccess field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInner) GetBucketAccess() []GetObjectStorageKeys200ResponseDataInnerBucketAccessInner {
	if o == nil || IsNil(o.BucketAccess) {
		var ret []GetObjectStorageKeys200ResponseDataInnerBucketAccessInner
		return ret
	}
	return o.BucketAccess
}

// GetBucketAccessOk returns a tuple with the BucketAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) GetBucketAccessOk() ([]GetObjectStorageKeys200ResponseDataInnerBucketAccessInner, bool) {
	if o == nil || IsNil(o.BucketAccess) {
		return nil, false
	}
	return o.BucketAccess, true
}

// HasBucketAccess returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) HasBucketAccess() bool {
	if o != nil && !IsNil(o.BucketAccess) {
		return true
	}

	return false
}

// SetBucketAccess gets a reference to the given []GetObjectStorageKeys200ResponseDataInnerBucketAccessInner and assigns it to the BucketAccess field.
func (o *GetObjectStorageKeys200ResponseDataInner) SetBucketAccess(v []GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) {
	o.BucketAccess = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetObjectStorageKeys200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetObjectStorageKeys200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetLimited returns the Limited field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInner) GetLimited() bool {
	if o == nil || IsNil(o.Limited) {
		var ret bool
		return ret
	}
	return *o.Limited
}

// GetLimitedOk returns a tuple with the Limited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) GetLimitedOk() (*bool, bool) {
	if o == nil || IsNil(o.Limited) {
		return nil, false
	}
	return o.Limited, true
}

// HasLimited returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) HasLimited() bool {
	if o != nil && !IsNil(o.Limited) {
		return true
	}

	return false
}

// SetLimited gets a reference to the given bool and assigns it to the Limited field.
func (o *GetObjectStorageKeys200ResponseDataInner) SetLimited(v bool) {
	o.Limited = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInner) GetRegions() []GetObjectStorageKeys200ResponseDataInnerRegionsInner {
	if o == nil || IsNil(o.Regions) {
		var ret []GetObjectStorageKeys200ResponseDataInnerRegionsInner
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) GetRegionsOk() ([]GetObjectStorageKeys200ResponseDataInnerRegionsInner, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []GetObjectStorageKeys200ResponseDataInnerRegionsInner and assigns it to the Regions field.
func (o *GetObjectStorageKeys200ResponseDataInner) SetRegions(v []GetObjectStorageKeys200ResponseDataInnerRegionsInner) {
	o.Regions = v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInner) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInner) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *GetObjectStorageKeys200ResponseDataInner) SetSecretKey(v string) {
	o.SecretKey = &v
}

func (o GetObjectStorageKeys200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetObjectStorageKeys200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.BucketAccess) {
		toSerialize["bucket_access"] = o.BucketAccess
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Limited) {
		toSerialize["limited"] = o.Limited
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secret_key"] = o.SecretKey
	}
	return toSerialize, nil
}

type NullableGetObjectStorageKeys200ResponseDataInner struct {
	value *GetObjectStorageKeys200ResponseDataInner
	isSet bool
}

func (v NullableGetObjectStorageKeys200ResponseDataInner) Get() *GetObjectStorageKeys200ResponseDataInner {
	return v.value
}

func (v *NullableGetObjectStorageKeys200ResponseDataInner) Set(val *GetObjectStorageKeys200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObjectStorageKeys200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObjectStorageKeys200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObjectStorageKeys200ResponseDataInner(val *GetObjectStorageKeys200ResponseDataInner) *NullableGetObjectStorageKeys200ResponseDataInner {
	return &NullableGetObjectStorageKeys200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetObjectStorageKeys200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetObjectStorageKeys200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


