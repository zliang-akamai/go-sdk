# PostUploadImage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**GetImages200ResponseDataInner**](GetImages200ResponseDataInner.md) |  | [optional] 
**UploadTo** | Pointer to **string** | The URL to upload the Image to. | [optional] 

## Methods

### NewPostUploadImage200Response

`func NewPostUploadImage200Response() *PostUploadImage200Response`

NewPostUploadImage200Response instantiates a new PostUploadImage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUploadImage200ResponseWithDefaults

`func NewPostUploadImage200ResponseWithDefaults() *PostUploadImage200Response`

NewPostUploadImage200ResponseWithDefaults instantiates a new PostUploadImage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *PostUploadImage200Response) GetImage() GetImages200ResponseDataInner`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PostUploadImage200Response) GetImageOk() (*GetImages200ResponseDataInner, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PostUploadImage200Response) SetImage(v GetImages200ResponseDataInner)`

SetImage sets Image field to given value.

### HasImage

`func (o *PostUploadImage200Response) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetUploadTo

`func (o *PostUploadImage200Response) GetUploadTo() string`

GetUploadTo returns the UploadTo field if non-nil, zero value otherwise.

### GetUploadToOk

`func (o *PostUploadImage200Response) GetUploadToOk() (*string, bool)`

GetUploadToOk returns a tuple with the UploadTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTo

`func (o *PostUploadImage200Response) SetUploadTo(v string)`

SetUploadTo sets UploadTo field to given value.

### HasUploadTo

`func (o *PostUploadImage200Response) HasUploadTo() bool`

HasUploadTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


