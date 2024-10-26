# PostUploadImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudInit** | Pointer to **bool** | Whether the uploaded Image supports cloud-init. | [optional] 
**Description** | Pointer to **string** | Description for the uploaded image. | [optional] 
**Label** | **string** | Label for the uploaded image. | 
**Region** | **string** | The region to upload to. Once uploaded, the image can be used in any region. | 
**Tags** | Pointer to **[]string** | Tags you can use to organize images. A tag can be from 3 to 100 characters long, and an image can have a maximum of 500 total tags. | [optional] 

## Methods

### NewPostUploadImageRequest

`func NewPostUploadImageRequest(label string, region string, ) *PostUploadImageRequest`

NewPostUploadImageRequest instantiates a new PostUploadImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUploadImageRequestWithDefaults

`func NewPostUploadImageRequestWithDefaults() *PostUploadImageRequest`

NewPostUploadImageRequestWithDefaults instantiates a new PostUploadImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudInit

`func (o *PostUploadImageRequest) GetCloudInit() bool`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *PostUploadImageRequest) GetCloudInitOk() (*bool, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *PostUploadImageRequest) SetCloudInit(v bool)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *PostUploadImageRequest) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.

### GetDescription

`func (o *PostUploadImageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostUploadImageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostUploadImageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostUploadImageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *PostUploadImageRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostUploadImageRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostUploadImageRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRegion

`func (o *PostUploadImageRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostUploadImageRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostUploadImageRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetTags

`func (o *PostUploadImageRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostUploadImageRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostUploadImageRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostUploadImageRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


