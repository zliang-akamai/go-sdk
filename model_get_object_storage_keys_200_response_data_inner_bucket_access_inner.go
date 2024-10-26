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

// checks if the GetObjectStorageKeys200ResponseDataInnerBucketAccessInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetObjectStorageKeys200ResponseDataInnerBucketAccessInner{}

// GetObjectStorageKeys200ResponseDataInnerBucketAccessInner struct for GetObjectStorageKeys200ResponseDataInnerBucketAccessInner
type GetObjectStorageKeys200ResponseDataInnerBucketAccessInner struct {
	// The name of the bucket the key can access in the `region`.
	BucketName *string `json:"bucket_name,omitempty"`
	// Identifies the legacy cluster where this key can be used. The key grants access to each specified `bucket_name`, based on the `permissions` set. To support backward compatibility, the API generates this value, based on the `region` set for a new key. See [Create an Object Storage key](https://techdocs.akamai.com/linode-api/reference/post-object-storage-keys) for more information.
	Cluster *string `json:"cluster,omitempty"`
	// The level of access the key grants to the `bucket_name`. Keys with `read_write` access can manage content in the `bucket_name`, while `read_only` can be used to view content. See [Create an Object Storage key(ref:post-object-storage-keys) for more details.
	Permissions *string `json:"permissions,omitempty"`
	// The region where the Object Store `bucket_name` resides.
	Region *string `json:"region,omitempty"`
}

// NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInner instantiates a new GetObjectStorageKeys200ResponseDataInnerBucketAccessInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInner() *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner {
	this := GetObjectStorageKeys200ResponseDataInnerBucketAccessInner{}
	return &this
}

// NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInnerWithDefaults instantiates a new GetObjectStorageKeys200ResponseDataInnerBucketAccessInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInnerWithDefaults() *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner {
	this := GetObjectStorageKeys200ResponseDataInnerBucketAccessInner{}
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) SetBucketName(v string) {
	o.BucketName = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) SetCluster(v string) {
	o.Cluster = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetPermissions() string {
	if o == nil || IsNil(o.Permissions) {
		var ret string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetPermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given string and assigns it to the Permissions field.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) SetPermissions(v string) {
	o.Permissions = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) SetRegion(v string) {
	o.Region = &v
}

func (o GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketName) {
		toSerialize["bucket_name"] = o.BucketName
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner struct {
	value *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner
	isSet bool
}

func (v NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner) Get() *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner {
	return v.value
}

func (v *NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner) Set(val *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner(val *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) *NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner {
	return &NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner{value: val, isSet: true}
}

func (v NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetObjectStorageKeys200ResponseDataInnerBucketAccessInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

