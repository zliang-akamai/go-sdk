# PutVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Volume was created. | [optional] [readonly] 
**Encryption** | Pointer to **string** | Displays if encryption is enabled on this volume. | [optional] [readonly] 
**FilesystemPath** | Pointer to **string** | The full filesystem path for the Volume based on the Volume&#39;s label. Path is /dev/disk/by-id/scsi-0Linode_Volume_ + Volume label. | [optional] [readonly] 
**HardwareType** | Pointer to **string** | The storage type of this Volume. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID of this Volume. | [optional] [readonly] 
**Label** | Pointer to **string** | The Volume&#39;s label is for display purposes only. | [optional] 
**LinodeId** | Pointer to **interface{}** |  | [optional] [readonly] 
**LinodeLabel** | Pointer to **NullableString** | If a Volume is attached to a specific Linode, the label of that Linode will be displayed here. | [optional] [readonly] 
**Region** | Pointer to **string** | The unique ID of this Region. | [optional] [readonly] 
**Size** | Pointer to **interface{}** |  | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of the volume.  Can be one of:    - &#x60;creating&#x60;. The volume is being created and is not yet available for use.   - &#x60;active&#x60;. The volume is online and available for use.   - &#x60;resizing&#x60;. The volume is in the process of upgrading its current capacity.   - &#x60;key_rotating&#x60;. The volume is in the process of rotating encryption keys. Requests to resize, delete, or clone a volume fails during encryption key rotation. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | An array of Tags applied to this object.  Tags are for organizational purposes only. | [optional] 
**Updated** | Pointer to **time.Time** | When this Volume was last updated. | [optional] [readonly] 

## Methods

### NewPutVolumeRequest

`func NewPutVolumeRequest() *PutVolumeRequest`

NewPutVolumeRequest instantiates a new PutVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutVolumeRequestWithDefaults

`func NewPutVolumeRequestWithDefaults() *PutVolumeRequest`

NewPutVolumeRequestWithDefaults instantiates a new PutVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PutVolumeRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PutVolumeRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PutVolumeRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PutVolumeRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEncryption

`func (o *PutVolumeRequest) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *PutVolumeRequest) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *PutVolumeRequest) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *PutVolumeRequest) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFilesystemPath

`func (o *PutVolumeRequest) GetFilesystemPath() string`

GetFilesystemPath returns the FilesystemPath field if non-nil, zero value otherwise.

### GetFilesystemPathOk

`func (o *PutVolumeRequest) GetFilesystemPathOk() (*string, bool)`

GetFilesystemPathOk returns a tuple with the FilesystemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemPath

`func (o *PutVolumeRequest) SetFilesystemPath(v string)`

SetFilesystemPath sets FilesystemPath field to given value.

### HasFilesystemPath

`func (o *PutVolumeRequest) HasFilesystemPath() bool`

HasFilesystemPath returns a boolean if a field has been set.

### GetHardwareType

`func (o *PutVolumeRequest) GetHardwareType() string`

GetHardwareType returns the HardwareType field if non-nil, zero value otherwise.

### GetHardwareTypeOk

`func (o *PutVolumeRequest) GetHardwareTypeOk() (*string, bool)`

GetHardwareTypeOk returns a tuple with the HardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareType

`func (o *PutVolumeRequest) SetHardwareType(v string)`

SetHardwareType sets HardwareType field to given value.

### HasHardwareType

`func (o *PutVolumeRequest) HasHardwareType() bool`

HasHardwareType returns a boolean if a field has been set.

### GetId

`func (o *PutVolumeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutVolumeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutVolumeRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PutVolumeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PutVolumeRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutVolumeRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutVolumeRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutVolumeRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLinodeId

`func (o *PutVolumeRequest) GetLinodeId() interface{}`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PutVolumeRequest) GetLinodeIdOk() (*interface{}, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PutVolumeRequest) SetLinodeId(v interface{})`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *PutVolumeRequest) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### SetLinodeIdNil

`func (o *PutVolumeRequest) SetLinodeIdNil(b bool)`

 SetLinodeIdNil sets the value for LinodeId to be an explicit nil

### UnsetLinodeId
`func (o *PutVolumeRequest) UnsetLinodeId()`

UnsetLinodeId ensures that no value is present for LinodeId, not even an explicit nil
### GetLinodeLabel

`func (o *PutVolumeRequest) GetLinodeLabel() string`

GetLinodeLabel returns the LinodeLabel field if non-nil, zero value otherwise.

### GetLinodeLabelOk

`func (o *PutVolumeRequest) GetLinodeLabelOk() (*string, bool)`

GetLinodeLabelOk returns a tuple with the LinodeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeLabel

`func (o *PutVolumeRequest) SetLinodeLabel(v string)`

SetLinodeLabel sets LinodeLabel field to given value.

### HasLinodeLabel

`func (o *PutVolumeRequest) HasLinodeLabel() bool`

HasLinodeLabel returns a boolean if a field has been set.

### SetLinodeLabelNil

`func (o *PutVolumeRequest) SetLinodeLabelNil(b bool)`

 SetLinodeLabelNil sets the value for LinodeLabel to be an explicit nil

### UnsetLinodeLabel
`func (o *PutVolumeRequest) UnsetLinodeLabel()`

UnsetLinodeLabel ensures that no value is present for LinodeLabel, not even an explicit nil
### GetRegion

`func (o *PutVolumeRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PutVolumeRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PutVolumeRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PutVolumeRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSize

`func (o *PutVolumeRequest) GetSize() interface{}`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PutVolumeRequest) GetSizeOk() (*interface{}, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PutVolumeRequest) SetSize(v interface{})`

SetSize sets Size field to given value.

### HasSize

`func (o *PutVolumeRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *PutVolumeRequest) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *PutVolumeRequest) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *PutVolumeRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PutVolumeRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PutVolumeRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PutVolumeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PutVolumeRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PutVolumeRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PutVolumeRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PutVolumeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdated

`func (o *PutVolumeRequest) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PutVolumeRequest) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PutVolumeRequest) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PutVolumeRequest) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


