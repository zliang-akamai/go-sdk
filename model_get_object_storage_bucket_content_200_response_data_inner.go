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

// checks if the GetObjectStorageBucketContent200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetObjectStorageBucketContent200ResponseDataInner{}

// GetObjectStorageBucketContent200ResponseDataInner An Object in Object Storage, or a \"prefix\" that contains one or more objects when a `delimiter` is used.
type GetObjectStorageBucketContent200ResponseDataInner struct {
	// An MD-5 hash of the object. `null` if this object represents a prefix.
	Etag *string `json:"etag,omitempty"`
	// The date and time this object was last modified. `null` if this object represents a prefix.
	LastModified *time.Time `json:"last_modified,omitempty"`
	// The name of this object or prefix.
	Name *string `json:"name,omitempty"`
	// The owner of this object, as a UUID. `null` if this object represents a prefix.
	Owner *string `json:"owner,omitempty"`
	// The size of this object, in bytes. `null` if this object represents a prefix.
	Size *int32 `json:"size,omitempty"`
}

// NewGetObjectStorageBucketContent200ResponseDataInner instantiates a new GetObjectStorageBucketContent200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetObjectStorageBucketContent200ResponseDataInner() *GetObjectStorageBucketContent200ResponseDataInner {
	this := GetObjectStorageBucketContent200ResponseDataInner{}
	return &this
}

// NewGetObjectStorageBucketContent200ResponseDataInnerWithDefaults instantiates a new GetObjectStorageBucketContent200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetObjectStorageBucketContent200ResponseDataInnerWithDefaults() *GetObjectStorageBucketContent200ResponseDataInner {
	this := GetObjectStorageBucketContent200ResponseDataInner{}
	return &this
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetEtag() string {
	if o == nil || IsNil(o.Etag) {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetEtagOk() (*string, bool) {
	if o == nil || IsNil(o.Etag) {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) HasEtag() bool {
	if o != nil && !IsNil(o.Etag) {
		return true
	}

	return false
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *GetObjectStorageBucketContent200ResponseDataInner) SetEtag(v string) {
	o.Etag = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *GetObjectStorageBucketContent200ResponseDataInner) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetObjectStorageBucketContent200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *GetObjectStorageBucketContent200ResponseDataInner) SetOwner(v string) {
	o.Owner = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GetObjectStorageBucketContent200ResponseDataInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *GetObjectStorageBucketContent200ResponseDataInner) SetSize(v int32) {
	o.Size = &v
}

func (o GetObjectStorageBucketContent200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetObjectStorageBucketContent200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Etag) {
		toSerialize["etag"] = o.Etag
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableGetObjectStorageBucketContent200ResponseDataInner struct {
	value *GetObjectStorageBucketContent200ResponseDataInner
	isSet bool
}

func (v NullableGetObjectStorageBucketContent200ResponseDataInner) Get() *GetObjectStorageBucketContent200ResponseDataInner {
	return v.value
}

func (v *NullableGetObjectStorageBucketContent200ResponseDataInner) Set(val *GetObjectStorageBucketContent200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObjectStorageBucketContent200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObjectStorageBucketContent200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObjectStorageBucketContent200ResponseDataInner(val *GetObjectStorageBucketContent200ResponseDataInner) *NullableGetObjectStorageBucketContent200ResponseDataInner {
	return &NullableGetObjectStorageBucketContent200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetObjectStorageBucketContent200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetObjectStorageBucketContent200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


