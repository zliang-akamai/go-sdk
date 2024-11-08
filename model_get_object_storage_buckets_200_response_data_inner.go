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

// checks if the GetObjectStorageBuckets200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetObjectStorageBuckets200ResponseDataInner{}

// GetObjectStorageBuckets200ResponseDataInner An Object Storage bucket. You should access this through the [S3 API](https://docs.ceph.com/en/latest/radosgw/s3/#api).
type GetObjectStorageBuckets200ResponseDataInner struct {
	// The legacy `clusterId` equivalent for the `regionId` where this bucket lives. The API maintains this for backward compatibility.  > 📘 > > - This value and the `regionId` are interchangeable when used in requests. Best practice is to use the `regionId`. > > - This value is empty for newer regions that don't have a legacy `clusterId`.
	// Deprecated
	Cluster *string `json:"cluster,omitempty"`
	// When this bucket was created.
	Created *time.Time `json:"created,omitempty"`
	// The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public.
	Hostname *string `json:"hostname,omitempty"`
	// The name of this bucket.
	Label *string `json:"label,omitempty"`
	// The number of objects stored in this bucket.
	Objects *int32 `json:"objects,omitempty"`
	// The `id` of the [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where this Object Storage bucket lives.
	Region *string `json:"region,omitempty"`
	// The size of the bucket in bytes.
	Size *int32 `json:"size,omitempty"`
}

// NewGetObjectStorageBuckets200ResponseDataInner instantiates a new GetObjectStorageBuckets200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetObjectStorageBuckets200ResponseDataInner() *GetObjectStorageBuckets200ResponseDataInner {
	this := GetObjectStorageBuckets200ResponseDataInner{}
	return &this
}

// NewGetObjectStorageBuckets200ResponseDataInnerWithDefaults instantiates a new GetObjectStorageBuckets200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetObjectStorageBuckets200ResponseDataInnerWithDefaults() *GetObjectStorageBuckets200ResponseDataInner {
	this := GetObjectStorageBuckets200ResponseDataInner{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
// Deprecated
func (o *GetObjectStorageBuckets200ResponseDataInner) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GetObjectStorageBuckets200ResponseDataInner) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
// Deprecated
func (o *GetObjectStorageBuckets200ResponseDataInner) SetCluster(v string) {
	o.Cluster = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetObjectStorageBuckets200ResponseDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *GetObjectStorageBuckets200ResponseDataInner) SetHostname(v string) {
	o.Hostname = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetObjectStorageBuckets200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetObjects() int32 {
	if o == nil || IsNil(o.Objects) {
		var ret int32
		return ret
	}
	return *o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetObjectsOk() (*int32, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given int32 and assigns it to the Objects field.
func (o *GetObjectStorageBuckets200ResponseDataInner) SetObjects(v int32) {
	o.Objects = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetObjectStorageBuckets200ResponseDataInner) SetRegion(v string) {
	o.Region = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GetObjectStorageBuckets200ResponseDataInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *GetObjectStorageBuckets200ResponseDataInner) SetSize(v int32) {
	o.Size = &v
}

func (o GetObjectStorageBuckets200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetObjectStorageBuckets200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableGetObjectStorageBuckets200ResponseDataInner struct {
	value *GetObjectStorageBuckets200ResponseDataInner
	isSet bool
}

func (v NullableGetObjectStorageBuckets200ResponseDataInner) Get() *GetObjectStorageBuckets200ResponseDataInner {
	return v.value
}

func (v *NullableGetObjectStorageBuckets200ResponseDataInner) Set(val *GetObjectStorageBuckets200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObjectStorageBuckets200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObjectStorageBuckets200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObjectStorageBuckets200ResponseDataInner(val *GetObjectStorageBuckets200ResponseDataInner) *NullableGetObjectStorageBuckets200ResponseDataInner {
	return &NullableGetObjectStorageBuckets200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetObjectStorageBuckets200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetObjectStorageBuckets200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


