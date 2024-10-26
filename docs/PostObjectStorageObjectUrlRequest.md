# PostObjectStorageObjectUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | The expected &#x60;Content-type&#x60; header of the request this signed URL will be valid for.  If provided, the &#x60;Content-type&#x60; header _must_ be sent with the request when this URL is used, and _must_ be the same as it was when the signed URL was created. Required for all methods _except_ &#x60;GET&#x60; or &#x60;DELETE&#x60;. | [optional] 
**ExpiresIn** | Pointer to **int32** | How long this signed URL will be valid for, in seconds.  If omitted, the URL will be valid for 3600 seconds (1 hour). | [optional] [default to 3600]
**Method** | **string** | The HTTP method allowed to be used with the pre-signed URL. | [default to "GET"]
**Name** | **string** | The name of the object that will be accessed with the pre-signed URL. This object need not exist, and no error will be returned if it doesn&#39;t. This behavior is useful for generating pre-signed URLs to upload new objects to by setting the &#x60;method&#x60; to &#x60;PUT&#x60;. | 

## Methods

### NewPostObjectStorageObjectUrlRequest

`func NewPostObjectStorageObjectUrlRequest(method string, name string, ) *PostObjectStorageObjectUrlRequest`

NewPostObjectStorageObjectUrlRequest instantiates a new PostObjectStorageObjectUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostObjectStorageObjectUrlRequestWithDefaults

`func NewPostObjectStorageObjectUrlRequestWithDefaults() *PostObjectStorageObjectUrlRequest`

NewPostObjectStorageObjectUrlRequestWithDefaults instantiates a new PostObjectStorageObjectUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *PostObjectStorageObjectUrlRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PostObjectStorageObjectUrlRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PostObjectStorageObjectUrlRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PostObjectStorageObjectUrlRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExpiresIn

`func (o *PostObjectStorageObjectUrlRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *PostObjectStorageObjectUrlRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *PostObjectStorageObjectUrlRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *PostObjectStorageObjectUrlRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetMethod

`func (o *PostObjectStorageObjectUrlRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PostObjectStorageObjectUrlRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PostObjectStorageObjectUrlRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetName

`func (o *PostObjectStorageObjectUrlRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostObjectStorageObjectUrlRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostObjectStorageObjectUrlRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


