# PostResizeVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** | The Volume&#39;s size, in GiB. | 

## Methods

### NewPostResizeVolumeRequest

`func NewPostResizeVolumeRequest(size int32, ) *PostResizeVolumeRequest`

NewPostResizeVolumeRequest instantiates a new PostResizeVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostResizeVolumeRequestWithDefaults

`func NewPostResizeVolumeRequestWithDefaults() *PostResizeVolumeRequest`

NewPostResizeVolumeRequestWithDefaults instantiates a new PostResizeVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *PostResizeVolumeRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostResizeVolumeRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostResizeVolumeRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


