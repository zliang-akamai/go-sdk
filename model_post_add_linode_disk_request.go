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
	"bytes"
	"fmt"
)

// checks if the PostAddLinodeDiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAddLinodeDiskRequest{}

// PostAddLinodeDiskRequest struct for PostAddLinodeDiskRequest
type PostAddLinodeDiskRequest struct {
	// A list of public SSH keys that will be automatically appended to the root user's `~/.ssh/authorized_keys` file when deploying from an Image.
	AuthorizedKeys []string `json:"authorized_keys,omitempty"`
	// A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users `~/.ssh/authorized_keys` file automatically when deploying from an Image.
	AuthorizedUsers []string `json:"authorized_users,omitempty"`
	// The file system of the disk. This can be `raw`, which indicates no file system, just a raw binary stream, `swap` for a Linux swap area, `ext3` or `ext4` for the ext3 of ext4 journaling file systems for Linux, respectively, or `initrd` for uncompressed initrd, ext2 with a maximum size of 32 MB.
	Filesystem *string `json:"filesystem,omitempty"`
	// An Image ID to deploy the Linode Disk from.  Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation with authentication to view all available Images. Official Linode Images start with `linode/`, while your Account's Images start with `private/`. Creating a disk from a Private Image requires `read_only` or `read_write` permissions for that Image. Run the [Update a user's grants](https://techdocs.akamai.com/linode-api/reference/put-user-grants) operation to adjust permissions for an Account Image.
	Image *string `json:"image,omitempty"`
	// The name of the disk. This is for display purposes only.
	Label *string `json:"label,omitempty"`
	// This sets the root user's password on a newly created Linode Disk when deploying from an Image.  - __Required__ when creating a Linode Disk from an Image, including when using a StackScript.  - Must meet a password strength score requirement that is calculated internally by the API. If the strength requirement is not met, you will receive a `Password does not meet strength requirement` error.
	RootPass *string `json:"root_pass,omitempty"`
	// The size of the Disk in MB.  Images require a minimum size. Run the [Get an image](https://techdocs.akamai.com/linode-api/reference/get-image) operation to view its size.
	Size int32 `json:"size"`
	// This field is required only if the StackScript being deployed requires input data from the User for successful completion. See [User Defined Fields (UDFs)](https://www.linode.com/docs/products/tools/stackscripts/guides/write-a-custom-script/#declare-user-defined-fields-udfs) for more details.  This field is required to be valid JSON.  Total length cannot exceed 65,535 characters.
	StackscriptData map[string]interface{} `json:"stackscript_data,omitempty"`
	// A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode. A compatible `image` is required to use a StackScript. To get a list of available StackScript and their permitted Images, run [List StackScripts](https://techdocs.akamai.com/linode-api/reference/get-stack-scripts). This field cannot be used when deploying from a Backup or a Private Image.
	StackscriptId *int32 `json:"stackscript_id,omitempty"`
}

type _PostAddLinodeDiskRequest PostAddLinodeDiskRequest

// NewPostAddLinodeDiskRequest instantiates a new PostAddLinodeDiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAddLinodeDiskRequest(size int32) *PostAddLinodeDiskRequest {
	this := PostAddLinodeDiskRequest{}
	this.Size = size
	return &this
}

// NewPostAddLinodeDiskRequestWithDefaults instantiates a new PostAddLinodeDiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAddLinodeDiskRequestWithDefaults() *PostAddLinodeDiskRequest {
	this := PostAddLinodeDiskRequest{}
	return &this
}

// GetAuthorizedKeys returns the AuthorizedKeys field value if set, zero value otherwise.
func (o *PostAddLinodeDiskRequest) GetAuthorizedKeys() []string {
	if o == nil || IsNil(o.AuthorizedKeys) {
		var ret []string
		return ret
	}
	return o.AuthorizedKeys
}

// GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetAuthorizedKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizedKeys) {
		return nil, false
	}
	return o.AuthorizedKeys, true
}

// HasAuthorizedKeys returns a boolean if a field has been set.
func (o *PostAddLinodeDiskRequest) HasAuthorizedKeys() bool {
	if o != nil && !IsNil(o.AuthorizedKeys) {
		return true
	}

	return false
}

// SetAuthorizedKeys gets a reference to the given []string and assigns it to the AuthorizedKeys field.
func (o *PostAddLinodeDiskRequest) SetAuthorizedKeys(v []string) {
	o.AuthorizedKeys = v
}

// GetAuthorizedUsers returns the AuthorizedUsers field value if set, zero value otherwise.
func (o *PostAddLinodeDiskRequest) GetAuthorizedUsers() []string {
	if o == nil || IsNil(o.AuthorizedUsers) {
		var ret []string
		return ret
	}
	return o.AuthorizedUsers
}

// GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetAuthorizedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizedUsers) {
		return nil, false
	}
	return o.AuthorizedUsers, true
}

// HasAuthorizedUsers returns a boolean if a field has been set.
func (o *PostAddLinodeDiskRequest) HasAuthorizedUsers() bool {
	if o != nil && !IsNil(o.AuthorizedUsers) {
		return true
	}

	return false
}

// SetAuthorizedUsers gets a reference to the given []string and assigns it to the AuthorizedUsers field.
func (o *PostAddLinodeDiskRequest) SetAuthorizedUsers(v []string) {
	o.AuthorizedUsers = v
}

// GetFilesystem returns the Filesystem field value if set, zero value otherwise.
func (o *PostAddLinodeDiskRequest) GetFilesystem() string {
	if o == nil || IsNil(o.Filesystem) {
		var ret string
		return ret
	}
	return *o.Filesystem
}

// GetFilesystemOk returns a tuple with the Filesystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetFilesystemOk() (*string, bool) {
	if o == nil || IsNil(o.Filesystem) {
		return nil, false
	}
	return o.Filesystem, true
}

// HasFilesystem returns a boolean if a field has been set.
func (o *PostAddLinodeDiskRequest) HasFilesystem() bool {
	if o != nil && !IsNil(o.Filesystem) {
		return true
	}

	return false
}

// SetFilesystem gets a reference to the given string and assigns it to the Filesystem field.
func (o *PostAddLinodeDiskRequest) SetFilesystem(v string) {
	o.Filesystem = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *PostAddLinodeDiskRequest) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *PostAddLinodeDiskRequest) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *PostAddLinodeDiskRequest) SetImage(v string) {
	o.Image = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PostAddLinodeDiskRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PostAddLinodeDiskRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PostAddLinodeDiskRequest) SetLabel(v string) {
	o.Label = &v
}

// GetRootPass returns the RootPass field value if set, zero value otherwise.
func (o *PostAddLinodeDiskRequest) GetRootPass() string {
	if o == nil || IsNil(o.RootPass) {
		var ret string
		return ret
	}
	return *o.RootPass
}

// GetRootPassOk returns a tuple with the RootPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetRootPassOk() (*string, bool) {
	if o == nil || IsNil(o.RootPass) {
		return nil, false
	}
	return o.RootPass, true
}

// HasRootPass returns a boolean if a field has been set.
func (o *PostAddLinodeDiskRequest) HasRootPass() bool {
	if o != nil && !IsNil(o.RootPass) {
		return true
	}

	return false
}

// SetRootPass gets a reference to the given string and assigns it to the RootPass field.
func (o *PostAddLinodeDiskRequest) SetRootPass(v string) {
	o.RootPass = &v
}

// GetSize returns the Size field value
func (o *PostAddLinodeDiskRequest) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PostAddLinodeDiskRequest) SetSize(v int32) {
	o.Size = v
}

// GetStackscriptData returns the StackscriptData field value if set, zero value otherwise.
func (o *PostAddLinodeDiskRequest) GetStackscriptData() map[string]interface{} {
	if o == nil || IsNil(o.StackscriptData) {
		var ret map[string]interface{}
		return ret
	}
	return o.StackscriptData
}

// GetStackscriptDataOk returns a tuple with the StackscriptData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetStackscriptDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StackscriptData) {
		return map[string]interface{}{}, false
	}
	return o.StackscriptData, true
}

// HasStackscriptData returns a boolean if a field has been set.
func (o *PostAddLinodeDiskRequest) HasStackscriptData() bool {
	if o != nil && !IsNil(o.StackscriptData) {
		return true
	}

	return false
}

// SetStackscriptData gets a reference to the given map[string]interface{} and assigns it to the StackscriptData field.
func (o *PostAddLinodeDiskRequest) SetStackscriptData(v map[string]interface{}) {
	o.StackscriptData = v
}

// GetStackscriptId returns the StackscriptId field value if set, zero value otherwise.
func (o *PostAddLinodeDiskRequest) GetStackscriptId() int32 {
	if o == nil || IsNil(o.StackscriptId) {
		var ret int32
		return ret
	}
	return *o.StackscriptId
}

// GetStackscriptIdOk returns a tuple with the StackscriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeDiskRequest) GetStackscriptIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StackscriptId) {
		return nil, false
	}
	return o.StackscriptId, true
}

// HasStackscriptId returns a boolean if a field has been set.
func (o *PostAddLinodeDiskRequest) HasStackscriptId() bool {
	if o != nil && !IsNil(o.StackscriptId) {
		return true
	}

	return false
}

// SetStackscriptId gets a reference to the given int32 and assigns it to the StackscriptId field.
func (o *PostAddLinodeDiskRequest) SetStackscriptId(v int32) {
	o.StackscriptId = &v
}

func (o PostAddLinodeDiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAddLinodeDiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizedKeys) {
		toSerialize["authorized_keys"] = o.AuthorizedKeys
	}
	if !IsNil(o.AuthorizedUsers) {
		toSerialize["authorized_users"] = o.AuthorizedUsers
	}
	if !IsNil(o.Filesystem) {
		toSerialize["filesystem"] = o.Filesystem
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.RootPass) {
		toSerialize["root_pass"] = o.RootPass
	}
	toSerialize["size"] = o.Size
	if !IsNil(o.StackscriptData) {
		toSerialize["stackscript_data"] = o.StackscriptData
	}
	if !IsNil(o.StackscriptId) {
		toSerialize["stackscript_id"] = o.StackscriptId
	}
	return toSerialize, nil
}

func (o *PostAddLinodeDiskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostAddLinodeDiskRequest := _PostAddLinodeDiskRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostAddLinodeDiskRequest)

	if err != nil {
		return err
	}

	*o = PostAddLinodeDiskRequest(varPostAddLinodeDiskRequest)

	return err
}

type NullablePostAddLinodeDiskRequest struct {
	value *PostAddLinodeDiskRequest
	isSet bool
}

func (v NullablePostAddLinodeDiskRequest) Get() *PostAddLinodeDiskRequest {
	return v.value
}

func (v *NullablePostAddLinodeDiskRequest) Set(val *PostAddLinodeDiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAddLinodeDiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAddLinodeDiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAddLinodeDiskRequest(val *PostAddLinodeDiskRequest) *NullablePostAddLinodeDiskRequest {
	return &NullablePostAddLinodeDiskRequest{value: val, isSet: true}
}

func (v NullablePostAddLinodeDiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAddLinodeDiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


