# PostVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **int32** | When creating a Volume attached to a Linode, the ID of the Linode Config to include the new Volume in. This Config must belong to the Linode referenced by &#x60;linode_id&#x60;. Must _not_ be provided if &#x60;linode_id&#x60; is not sent. If a &#x60;linode_id&#x60; is sent without a &#x60;config_id&#x60;, the volume will be attached:    - to the Linode&#39;s only config if it only has one config.   - to the Linode&#39;s last used config, if possible.  If no config can be selected for attachment, an error will be returned. | [optional] 
**Encryption** | Pointer to **string** | Enables encryption on the volume. Full disk encryption ensures the data stored on a block storage volume drive is secure. It protects against unauthorized access by keeping the data encrypted if the volume drive is removed from the data center, decommissioned, or disposed of.  The platform automatically manages the encryption and decryption process for you. You can use an encrypted volume the same way as you use a non-encrypted volume.  You can enable or disable disk encryption only when creating new block storage volumes. After a volume is created, the encryption setting can&#39;t be changed. | [optional] [default to "disabled"]
**Label** | **string** | The Volume&#39;s label, which is also used in the &#x60;filesystem_path&#x60; of the resulting volume. | 
**LinodeId** | Pointer to **int32** | The Linode this volume should be attached to upon creation. If not given, the volume will be created without an attachment. | [optional] 
**Region** | Pointer to **string** | The Region to deploy this Volume in. This is only required if a linode_id is not given. | [optional] 
**Size** | Pointer to **int32** | The initial size of this volume, in GB.  Be aware that volumes may only be resized up after creation. | [optional] [default to 20]
**Tags** | Pointer to **[]string** | An array of Tags applied to this object.  Tags are for organizational purposes only. | [optional] 

## Methods

### NewPostVolumeRequest

`func NewPostVolumeRequest(label string, ) *PostVolumeRequest`

NewPostVolumeRequest instantiates a new PostVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostVolumeRequestWithDefaults

`func NewPostVolumeRequestWithDefaults() *PostVolumeRequest`

NewPostVolumeRequestWithDefaults instantiates a new PostVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *PostVolumeRequest) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *PostVolumeRequest) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *PostVolumeRequest) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *PostVolumeRequest) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetEncryption

`func (o *PostVolumeRequest) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *PostVolumeRequest) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *PostVolumeRequest) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *PostVolumeRequest) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetLabel

`func (o *PostVolumeRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostVolumeRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostVolumeRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLinodeId

`func (o *PostVolumeRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostVolumeRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostVolumeRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *PostVolumeRequest) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetRegion

`func (o *PostVolumeRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostVolumeRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostVolumeRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostVolumeRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSize

`func (o *PostVolumeRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostVolumeRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostVolumeRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PostVolumeRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTags

`func (o *PostVolumeRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostVolumeRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostVolumeRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostVolumeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


