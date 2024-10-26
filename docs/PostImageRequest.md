# PostImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudInit** | Pointer to **bool** | Whether this image supports [cloud-init](https://www.linode.com/docs/guides/write-files-with-cloud-init/). | [optional] 
**Description** | Pointer to **string** | A detailed description of this image. | [optional] 
**DiskId** | **int32** | The ID of the target Linode disk for this image. | 
**Label** | Pointer to **string** | The short title for this image. If not provided, the disk&#39;s current label is used. | [optional] 
**Tags** | Pointer to **[]string** | Tags used for organizational purposes. A tag can be from 3 to 100 characters long, and an image can have a maximum of 500 total tags. | [optional] 

## Methods

### NewPostImageRequest

`func NewPostImageRequest(diskId int32, ) *PostImageRequest`

NewPostImageRequest instantiates a new PostImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostImageRequestWithDefaults

`func NewPostImageRequestWithDefaults() *PostImageRequest`

NewPostImageRequestWithDefaults instantiates a new PostImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudInit

`func (o *PostImageRequest) GetCloudInit() bool`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *PostImageRequest) GetCloudInitOk() (*bool, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *PostImageRequest) SetCloudInit(v bool)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *PostImageRequest) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.

### GetDescription

`func (o *PostImageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostImageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostImageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostImageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskId

`func (o *PostImageRequest) GetDiskId() int32`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *PostImageRequest) GetDiskIdOk() (*int32, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *PostImageRequest) SetDiskId(v int32)`

SetDiskId sets DiskId field to given value.


### GetLabel

`func (o *PostImageRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostImageRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostImageRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostImageRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTags

`func (o *PostImageRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostImageRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostImageRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostImageRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


