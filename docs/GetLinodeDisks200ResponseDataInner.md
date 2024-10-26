# GetLinodeDisks200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Disk was created. | [optional] [readonly] 
**DiskEncryption** | Pointer to **string** | Displays if encryption is enabled on this disk. This setting is based on the &#x60;disk_encryption&#x60; setting of the Linode. | [optional] [readonly] [default to "enabled"]
**Filesystem** | Pointer to **string** | The file system of the disk. This can be &#x60;raw&#x60;, which indicates no file system, just a raw binary stream, &#x60;swap&#x60; for a Linux swap area, &#x60;ext3&#x60; or &#x60;ext4&#x60; for the ext3 of ext4 journaling file systems for Linux, respectively, or &#x60;initrd&#x60; for uncompressed initrd, ext2 with a maximum size of 32 MB. | [optional] 
**Id** | Pointer to **int32** | This disk&#39;s ID. You need this value to run other operations that interact with the disk. | [optional] [readonly] 
**Label** | Pointer to **string** | The name of the disk. This is for display purposes only. | [optional] 
**Size** | Pointer to **int32** | The size of the disk in MB. | [optional] 
**Status** | Pointer to **string** | The current state of the disk. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this disk was last updated. | [optional] [readonly] 

## Methods

### NewGetLinodeDisks200ResponseDataInner

`func NewGetLinodeDisks200ResponseDataInner() *GetLinodeDisks200ResponseDataInner`

NewGetLinodeDisks200ResponseDataInner instantiates a new GetLinodeDisks200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeDisks200ResponseDataInnerWithDefaults

`func NewGetLinodeDisks200ResponseDataInnerWithDefaults() *GetLinodeDisks200ResponseDataInner`

NewGetLinodeDisks200ResponseDataInnerWithDefaults instantiates a new GetLinodeDisks200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetLinodeDisks200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetLinodeDisks200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetLinodeDisks200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetLinodeDisks200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *GetLinodeDisks200ResponseDataInner) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *GetLinodeDisks200ResponseDataInner) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *GetLinodeDisks200ResponseDataInner) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *GetLinodeDisks200ResponseDataInner) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetFilesystem

`func (o *GetLinodeDisks200ResponseDataInner) GetFilesystem() string`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *GetLinodeDisks200ResponseDataInner) GetFilesystemOk() (*string, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *GetLinodeDisks200ResponseDataInner) SetFilesystem(v string)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *GetLinodeDisks200ResponseDataInner) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetId

`func (o *GetLinodeDisks200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLinodeDisks200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLinodeDisks200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetLinodeDisks200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetLinodeDisks200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetLinodeDisks200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetLinodeDisks200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetLinodeDisks200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSize

`func (o *GetLinodeDisks200ResponseDataInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetLinodeDisks200ResponseDataInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetLinodeDisks200ResponseDataInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetLinodeDisks200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *GetLinodeDisks200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLinodeDisks200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLinodeDisks200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLinodeDisks200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *GetLinodeDisks200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetLinodeDisks200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetLinodeDisks200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetLinodeDisks200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


