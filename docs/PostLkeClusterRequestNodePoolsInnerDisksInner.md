# PostLkeClusterRequestNodePoolsInnerDisksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | The size of this custom disk partition in MB. The size of this disk partition can&#39;t exceed the capacity of the node&#39;s plan type. | [optional] 
**Type** | Pointer to **string** | This custom disk partition&#39;s filesystem type. | [optional] 

## Methods

### NewPostLkeClusterRequestNodePoolsInnerDisksInner

`func NewPostLkeClusterRequestNodePoolsInnerDisksInner() *PostLkeClusterRequestNodePoolsInnerDisksInner`

NewPostLkeClusterRequestNodePoolsInnerDisksInner instantiates a new PostLkeClusterRequestNodePoolsInnerDisksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLkeClusterRequestNodePoolsInnerDisksInnerWithDefaults

`func NewPostLkeClusterRequestNodePoolsInnerDisksInnerWithDefaults() *PostLkeClusterRequestNodePoolsInnerDisksInner`

NewPostLkeClusterRequestNodePoolsInnerDisksInnerWithDefaults instantiates a new PostLkeClusterRequestNodePoolsInnerDisksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *PostLkeClusterRequestNodePoolsInnerDisksInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostLkeClusterRequestNodePoolsInnerDisksInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostLkeClusterRequestNodePoolsInnerDisksInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PostLkeClusterRequestNodePoolsInnerDisksInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *PostLkeClusterRequestNodePoolsInnerDisksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostLkeClusterRequestNodePoolsInnerDisksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostLkeClusterRequestNodePoolsInnerDisksInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostLkeClusterRequestNodePoolsInnerDisksInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


