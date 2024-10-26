# Volume1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Volume was created. | [optional] [readonly] 
**Encryption** | Pointer to **string** | Displays if encryption is enabled on this volume. | [optional] [readonly] 
**FilesystemPath** | Pointer to **string** | The full filesystem path for the Volume based on the Volume&#39;s label. Path is /dev/disk/by-id/scsi-0Linode_Volume_ + Volume label. | [optional] [readonly] 
**HardwareType** | Pointer to **string** | The storage type of this Volume. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID of this Volume. | [optional] [readonly] 
**Label** | Pointer to **string** | The Volume&#39;s label is for display purposes only. | [optional] 
**LinodeId** | Pointer to **NullableInt32** | If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here. | [optional] 
**LinodeLabel** | Pointer to **NullableString** | If a Volume is attached to a specific Linode, the label of that Linode will be displayed here. | [optional] [readonly] 
**Region** | Pointer to **string** | The unique ID of this Region. | [optional] [readonly] 
**Size** | Pointer to **int32** | The Volume&#39;s size, in GiB. | [optional] 
**Status** | Pointer to **string** | The current status of the volume.  Can be one of:    - &#x60;creating&#x60;. The volume is being created and is not yet available for use.   - &#x60;active&#x60;. The volume is online and available for use.   - &#x60;resizing&#x60;. The volume is in the process of upgrading its current capacity.   - &#x60;key_rotating&#x60;. The volume is in the process of rotating encryption keys. Requests to resize, delete, or clone a volume fails during encryption key rotation. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | An array of Tags applied to this object.  Tags are for organizational purposes only. | [optional] 
**Updated** | Pointer to **time.Time** | When this Volume was last updated. | [optional] [readonly] 

## Methods

### NewVolume1

`func NewVolume1() *Volume1`

NewVolume1 instantiates a new Volume1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolume1WithDefaults

`func NewVolume1WithDefaults() *Volume1`

NewVolume1WithDefaults instantiates a new Volume1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Volume1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Volume1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Volume1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Volume1) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEncryption

`func (o *Volume1) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *Volume1) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *Volume1) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *Volume1) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFilesystemPath

`func (o *Volume1) GetFilesystemPath() string`

GetFilesystemPath returns the FilesystemPath field if non-nil, zero value otherwise.

### GetFilesystemPathOk

`func (o *Volume1) GetFilesystemPathOk() (*string, bool)`

GetFilesystemPathOk returns a tuple with the FilesystemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemPath

`func (o *Volume1) SetFilesystemPath(v string)`

SetFilesystemPath sets FilesystemPath field to given value.

### HasFilesystemPath

`func (o *Volume1) HasFilesystemPath() bool`

HasFilesystemPath returns a boolean if a field has been set.

### GetHardwareType

`func (o *Volume1) GetHardwareType() string`

GetHardwareType returns the HardwareType field if non-nil, zero value otherwise.

### GetHardwareTypeOk

`func (o *Volume1) GetHardwareTypeOk() (*string, bool)`

GetHardwareTypeOk returns a tuple with the HardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareType

`func (o *Volume1) SetHardwareType(v string)`

SetHardwareType sets HardwareType field to given value.

### HasHardwareType

`func (o *Volume1) HasHardwareType() bool`

HasHardwareType returns a boolean if a field has been set.

### GetId

`func (o *Volume1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume1) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Volume1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Volume1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Volume1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Volume1) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Volume1) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLinodeId

`func (o *Volume1) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *Volume1) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *Volume1) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *Volume1) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### SetLinodeIdNil

`func (o *Volume1) SetLinodeIdNil(b bool)`

 SetLinodeIdNil sets the value for LinodeId to be an explicit nil

### UnsetLinodeId
`func (o *Volume1) UnsetLinodeId()`

UnsetLinodeId ensures that no value is present for LinodeId, not even an explicit nil
### GetLinodeLabel

`func (o *Volume1) GetLinodeLabel() string`

GetLinodeLabel returns the LinodeLabel field if non-nil, zero value otherwise.

### GetLinodeLabelOk

`func (o *Volume1) GetLinodeLabelOk() (*string, bool)`

GetLinodeLabelOk returns a tuple with the LinodeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeLabel

`func (o *Volume1) SetLinodeLabel(v string)`

SetLinodeLabel sets LinodeLabel field to given value.

### HasLinodeLabel

`func (o *Volume1) HasLinodeLabel() bool`

HasLinodeLabel returns a boolean if a field has been set.

### SetLinodeLabelNil

`func (o *Volume1) SetLinodeLabelNil(b bool)`

 SetLinodeLabelNil sets the value for LinodeLabel to be an explicit nil

### UnsetLinodeLabel
`func (o *Volume1) UnsetLinodeLabel()`

UnsetLinodeLabel ensures that no value is present for LinodeLabel, not even an explicit nil
### GetRegion

`func (o *Volume1) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Volume1) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Volume1) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Volume1) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSize

`func (o *Volume1) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Volume1) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Volume1) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Volume1) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *Volume1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Volume1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Volume1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Volume1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Volume1) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Volume1) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Volume1) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Volume1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdated

`func (o *Volume1) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Volume1) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Volume1) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Volume1) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


