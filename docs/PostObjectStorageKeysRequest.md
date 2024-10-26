# PostObjectStorageKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketAccess** | Pointer to [**[]PostObjectStorageKeysRequestBucketAccessInner**](PostObjectStorageKeysRequestBucketAccessInner.md) | Set up the key to limit access to specific buckets, each with a specific permission level. You can create a limited Object Storage key with access to no buckets. Include an empty &#x60;bucket_access&#x60; array in the request. | [optional] 
**Label** | Pointer to **string** | The label for the Object Storage key, for display purposes only. | [optional] 
**Regions** | Pointer to **[]string** | You can use a key to create new buckets in regions set in this array. But it can&#39;t be used to manage content in those buckets. See [Create an Object Storage key](https://techdocs.akamai.com/linode-api/reference/post-object-storage-keys) for more details. | [optional] 

## Methods

### NewPostObjectStorageKeysRequest

`func NewPostObjectStorageKeysRequest() *PostObjectStorageKeysRequest`

NewPostObjectStorageKeysRequest instantiates a new PostObjectStorageKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostObjectStorageKeysRequestWithDefaults

`func NewPostObjectStorageKeysRequestWithDefaults() *PostObjectStorageKeysRequest`

NewPostObjectStorageKeysRequestWithDefaults instantiates a new PostObjectStorageKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketAccess

`func (o *PostObjectStorageKeysRequest) GetBucketAccess() []PostObjectStorageKeysRequestBucketAccessInner`

GetBucketAccess returns the BucketAccess field if non-nil, zero value otherwise.

### GetBucketAccessOk

`func (o *PostObjectStorageKeysRequest) GetBucketAccessOk() (*[]PostObjectStorageKeysRequestBucketAccessInner, bool)`

GetBucketAccessOk returns a tuple with the BucketAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketAccess

`func (o *PostObjectStorageKeysRequest) SetBucketAccess(v []PostObjectStorageKeysRequestBucketAccessInner)`

SetBucketAccess sets BucketAccess field to given value.

### HasBucketAccess

`func (o *PostObjectStorageKeysRequest) HasBucketAccess() bool`

HasBucketAccess returns a boolean if a field has been set.

### GetLabel

`func (o *PostObjectStorageKeysRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostObjectStorageKeysRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostObjectStorageKeysRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostObjectStorageKeysRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegions

`func (o *PostObjectStorageKeysRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PostObjectStorageKeysRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PostObjectStorageKeysRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PostObjectStorageKeysRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


