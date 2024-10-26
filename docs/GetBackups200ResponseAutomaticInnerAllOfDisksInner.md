# GetBackups200ResponseAutomaticInnerAllOfDisksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filesystem** | Pointer to **string** | The file system of the disk. This can be &#x60;raw&#x60;, which indicates no file system, just a raw binary stream, &#x60;swap&#x60; for a Linux swap area, &#x60;ext3&#x60; or &#x60;ext4&#x60; for the ext3 of ext4 journaling file systems for Linux, respectively, or &#x60;initrd&#x60; for uncompressed initrd, ext2 with a maximum size of 32 MB. | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetBackups200ResponseAutomaticInnerAllOfDisksInner

`func NewGetBackups200ResponseAutomaticInnerAllOfDisksInner() *GetBackups200ResponseAutomaticInnerAllOfDisksInner`

NewGetBackups200ResponseAutomaticInnerAllOfDisksInner instantiates a new GetBackups200ResponseAutomaticInnerAllOfDisksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBackups200ResponseAutomaticInnerAllOfDisksInnerWithDefaults

`func NewGetBackups200ResponseAutomaticInnerAllOfDisksInnerWithDefaults() *GetBackups200ResponseAutomaticInnerAllOfDisksInner`

NewGetBackups200ResponseAutomaticInnerAllOfDisksInnerWithDefaults instantiates a new GetBackups200ResponseAutomaticInnerAllOfDisksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystem

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) GetFilesystem() string`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) GetFilesystemOk() (*string, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) SetFilesystem(v string)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetLabel

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSize

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetBackups200ResponseAutomaticInnerAllOfDisksInner) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


