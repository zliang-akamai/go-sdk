# PostMutateLinodeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAutoDiskResize** | Pointer to **bool** | Automatically resize disks when resizing a Linode. When resizing down to a smaller plan your Linode&#39;s data must fit within the smaller disk size. | [optional] [default to true]

## Methods

### NewPostMutateLinodeInstanceRequest

`func NewPostMutateLinodeInstanceRequest() *PostMutateLinodeInstanceRequest`

NewPostMutateLinodeInstanceRequest instantiates a new PostMutateLinodeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMutateLinodeInstanceRequestWithDefaults

`func NewPostMutateLinodeInstanceRequestWithDefaults() *PostMutateLinodeInstanceRequest`

NewPostMutateLinodeInstanceRequestWithDefaults instantiates a new PostMutateLinodeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAutoDiskResize

`func (o *PostMutateLinodeInstanceRequest) GetAllowAutoDiskResize() bool`

GetAllowAutoDiskResize returns the AllowAutoDiskResize field if non-nil, zero value otherwise.

### GetAllowAutoDiskResizeOk

`func (o *PostMutateLinodeInstanceRequest) GetAllowAutoDiskResizeOk() (*bool, bool)`

GetAllowAutoDiskResizeOk returns a tuple with the AllowAutoDiskResize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoDiskResize

`func (o *PostMutateLinodeInstanceRequest) SetAllowAutoDiskResize(v bool)`

SetAllowAutoDiskResize sets AllowAutoDiskResize field to given value.

### HasAllowAutoDiskResize

`func (o *PostMutateLinodeInstanceRequest) HasAllowAutoDiskResize() bool`

HasAllowAutoDiskResize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


