# PostResizeDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** | The desired size, in MB, of the disk. | 

## Methods

### NewPostResizeDiskRequest

`func NewPostResizeDiskRequest(size int32, ) *PostResizeDiskRequest`

NewPostResizeDiskRequest instantiates a new PostResizeDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostResizeDiskRequestWithDefaults

`func NewPostResizeDiskRequestWithDefaults() *PostResizeDiskRequest`

NewPostResizeDiskRequestWithDefaults instantiates a new PostResizeDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *PostResizeDiskRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostResizeDiskRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostResizeDiskRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


